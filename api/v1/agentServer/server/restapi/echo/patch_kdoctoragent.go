// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchKdoctoragentHandlerFunc turns a function with the right signature into a patch kdoctoragent handler
type PatchKdoctoragentHandlerFunc func(PatchKdoctoragentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchKdoctoragentHandlerFunc) Handle(params PatchKdoctoragentParams) middleware.Responder {
	return fn(params)
}

// PatchKdoctoragentHandler interface for that can handle valid patch kdoctoragent params
type PatchKdoctoragentHandler interface {
	Handle(PatchKdoctoragentParams) middleware.Responder
}

// NewPatchKdoctoragent creates a new http.Handler for the patch kdoctoragent operation
func NewPatchKdoctoragent(ctx *middleware.Context, handler PatchKdoctoragentHandler) *PatchKdoctoragent {
	return &PatchKdoctoragent{Context: ctx, Handler: handler}
}

/*
	PatchKdoctoragent swagger:route PATCH /kdoctoragent echo patchKdoctoragent

echo http request

echo http request
*/
type PatchKdoctoragent struct {
	Context *middleware.Context
	Handler PatchKdoctoragentHandler
}

func (o *PatchKdoctoragent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchKdoctoragentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
