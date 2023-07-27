// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package runtime

import (
	"context"

	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type runtimeDaemonSet struct {
	client    client.Client
	Namespace string
	Name      string

	log *zap.Logger
}

func NewDaemonSetRuntime(c client.Client, namespace, name string, log *zap.Logger) TaskRuntime {
	rd := &runtimeDaemonSet{
		client:    c,
		Namespace: namespace,
		Name:      name,
		log:       log,
	}

	return rd
}

func (rd *runtimeDaemonSet) IsReady(ctx context.Context) bool {
	var daemonset appsv1.DaemonSet

	err := rd.client.Get(ctx, types.NamespacedName{
		Namespace: rd.Namespace,
		Name:      rd.Name,
	}, &daemonset)
	if nil != err {
		rd.log.Sugar().Errorf("failed to get deployment %s/%s, error: %v", rd.Namespace, rd.Name, err)
		return false
	}

	if daemonset.Status.DesiredNumberScheduled == daemonset.Status.NumberReady {
		return true
	}

	return false
}

func (rd *runtimeDaemonSet) Delete(ctx context.Context) error {
	var daemonset appsv1.DaemonSet

	err := rd.client.Get(ctx, types.NamespacedName{
		Namespace: rd.Namespace,
		Name:      rd.Name,
	}, &daemonset)
	if nil != err {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}

	if daemonset.DeletionTimestamp != nil {
		return nil
	}

	err = rd.client.Delete(ctx, &daemonset)
	if nil != err {
		return err
	}

	return nil
}
