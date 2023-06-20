// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package scheduler

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kdoctor-io/kdoctor/pkg/k8s/apis/kdoctor.io/v1beta1"
	"github.com/kdoctor-io/kdoctor/pkg/types"
)

const (
	kdoctor = "kdoctor"

	containerArgTaskKind = "--task-kind"
	containerArgTaskName = "--task-name"
	uniqueMatchLabelKey  = "app.kubernetes.io/name"
)

// appRuntimeName generates a unique name for task runtime.
// notice: different kind tasks could use the same CR object name, so we need add their kind to generate name.
func appRuntimeName(taskKind, taskName string) string {
	appName := fmt.Sprintf("%s-%s-%s", kdoctor, strings.ToLower(taskKind), taskName)
	if len(appName) > 63 {
		appName = appName[:63]
	}

	return appName
}

func CreateTaskRuntimeIfNotExist(ctx context.Context, clientSet client.Client, taskKind string, task metav1.Object, agentSpec v1beta1.AgentSpec, log *zap.Logger) (v1beta1.TaskResource, error) {
	appName := appRuntimeName(taskKind, task.GetName())
	resource := v1beta1.TaskResource{
		RuntimeName:   appName,
		RuntimeType:   "",
		ServiceName:   appName,
		RuntimeStatus: v1beta1.RuntimeCreating,
	}

	var app client.Object
	if agentSpec.Kind == types.KindDeployment {
		app = &appsv1.Deployment{}
		resource.RuntimeType = types.KindDeployment
	} else {
		app = &appsv1.DaemonSet{}
		resource.RuntimeType = types.KindDaemonSet
	}

	var needCreate bool
	objectKey := client.ObjectKey{
		// reuse kdoctor-controller namespace
		Namespace: types.ControllerConfig.PodNamespace,
		Name:      appName,
	}

	log.Sugar().Infof("try to get task %s/%s cooresponding runtime %s", taskKind, task.GetName(), appName)
	err := clientSet.Get(ctx, objectKey, app)
	if nil != err {
		if errors.IsNotFound(err) {
			log.Sugar().Infof("task  %s/%s cooresponding runtime %s not found, try to create one", taskKind, task.GetName(), appName)
			needCreate = true
		} else {
			return resource, err
		}
	}

	if needCreate {
		if agentSpec.Kind == types.KindDeployment {
			app = generateDeployment(taskKind, task.GetName(), agentSpec)
		} else {
			app = generateDaemonSet(taskKind, task.GetName(), agentSpec)
		}

		app.SetName(appName)

		err := controllerruntime.SetControllerReference(task, app, clientSet.Scheme())
		if nil != err {
			return resource, fmt.Errorf("failed to set task %s/%s corresponding runtime %s controllerReference, error: %v", taskKind, task.GetName(), appName, err)
		}

		log.Sugar().Infof("try to create task  %s/%s cooresponding runtime %s", taskKind, task.GetName(), app)
		err = clientSet.Create(ctx, app)
		if nil != err {
			return resource, err
		}
	}

	return resource, nil
}

func generateDaemonSet(taskKind, taskName string, agentSpec v1beta1.AgentSpec) *appsv1.DaemonSet {
	daemonset := types.DaemonsetTempl.DeepCopy()

	// add unique selector
	uniqueValue := appRuntimeName(taskKind, taskName)
	var selector metav1.LabelSelector
	if daemonset.Spec.Selector != nil {
		daemonset.Spec.Selector.DeepCopyInto(&selector)
	}
	selector.MatchLabels[uniqueMatchLabelKey] = uniqueValue
	daemonset.Spec.Selector = &selector

	// replace AgentSpec properties
	if len(agentSpec.Annotation) != 0 {
		daemonset.SetAnnotations(appendAnnotation(daemonset.Annotations, agentSpec.Annotation))
	}

	// assemble
	podTemplateSpec := generatePodTemplateSpec(taskKind, taskName, uniqueValue, agentSpec)
	daemonset.Spec.Template = podTemplateSpec

	return daemonset
}

func generateDeployment(taskKind, taskName string, agentSpec v1beta1.AgentSpec) *appsv1.Deployment {
	deployment := types.DeploymentTempl.DeepCopy()

	// add unique selector
	uniqueValue := appRuntimeName(taskKind, taskName)
	var selector metav1.LabelSelector
	if deployment.Spec.Selector != nil {
		deployment.Spec.Selector.DeepCopyInto(&selector)
	}
	selector.MatchLabels[uniqueMatchLabelKey] = uniqueValue
	deployment.Spec.Selector = &selector

	// replace AgentSpec properties
	if len(agentSpec.Annotation) != 0 {
		deployment.SetAnnotations(appendAnnotation(deployment.Annotations, agentSpec.Annotation))
	}
	if agentSpec.DeploymentReplicas != nil {
		deployment.Spec.Replicas = agentSpec.DeploymentReplicas
	}

	// assemble
	podTemplateSpec := generatePodTemplateSpec(taskKind, taskName, uniqueValue, agentSpec)
	deployment.Spec.Template = podTemplateSpec

	return deployment
}

func generatePodTemplateSpec(taskKind, taskName string, uniqueLabelVal string, agentSpec v1beta1.AgentSpec) corev1.PodTemplateSpec {
	pod := types.PodTempl.DeepCopy()

	// 1. add container start parameters
	for index := range pod.Spec.Containers {
		tmpArgs := pod.Spec.Containers[index].Args
		tmpArgs = append(tmpArgs,
			fmt.Sprintf("%s=%s", containerArgTaskKind, taskKind),
			fmt.Sprintf("%s=%s", containerArgTaskName, taskName))
		pod.Spec.Containers[index].Args = tmpArgs
	}

	// 2. add unique selector
	{
		podLabels := make(map[string]string)
		if len(pod.Labels) != 0 {
			podLabels = pod.GetLabels()
		}
		podLabels[uniqueMatchLabelKey] = uniqueLabelVal
		pod.SetLabels(podLabels)
	}

	// 3. replace AgentSpec properties
	{
		if len(agentSpec.Annotation) != 0 {
			pod.SetAnnotations(appendAnnotation(pod.Annotations, agentSpec.Annotation))
		}

		if agentSpec.Affinity != nil {
			pod.Spec.Affinity = agentSpec.Affinity
		}

		if pod.Spec.HostNetwork != agentSpec.HostNetwork {
			if agentSpec.HostNetwork {
				pod.Spec.HostNetwork = true
				pod.Spec.DNSPolicy = corev1.DNSClusterFirstWithHostNet
			} else {
				pod.Spec.HostNetwork = false
				pod.Spec.DNSPolicy = corev1.DNSClusterFirst
			}
		}

		for index := range pod.Spec.Containers {
			if len(agentSpec.Env) != 0 {
				pod.Spec.Containers[index].Env = append(pod.Spec.Containers[index].Env, agentSpec.Env...)
			}

			if agentSpec.Resources != nil {
				pod.Spec.Containers[index].Resources = *agentSpec.Resources
			}
		}
	}

	podTemplateSpec := corev1.PodTemplateSpec{
		ObjectMeta: pod.ObjectMeta,
		Spec:       pod.Spec,
	}

	return podTemplateSpec
}

// TODO: SetName
func generateService(uniqueLabelVal string, agentSpec v1beta1.AgentSpec, ipFamily corev1.IPFamily) *corev1.Service {
	service := types.ServiceTempl.DeepCopy()

	// add unique selector
	selector := map[string]string{}
	if service.Spec.Selector != nil {
		selector = service.Spec.Selector
	}
	selector[uniqueMatchLabelKey] = uniqueLabelVal
	service.Spec.Selector = selector

	// replace AgentSpec properties
	if len(agentSpec.Annotation) != 0 {
		service.SetAnnotations(appendAnnotation(service.Annotations, agentSpec.Annotation))
	}

	// set IP Family
	service.Spec.IPFamilies = []corev1.IPFamily{ipFamily}

	return service
}

func appendAnnotation(origin, addition map[string]string) map[string]string {
	if origin == nil {
		origin = make(map[string]string)
	}

	for key, val := range addition {
		origin[key] = val
	}

	return origin
}
