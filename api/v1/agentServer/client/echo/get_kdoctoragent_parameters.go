// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetKdoctoragentParams creates a new GetKdoctoragentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKdoctoragentParams() *GetKdoctoragentParams {
	return &GetKdoctoragentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKdoctoragentParamsWithTimeout creates a new GetKdoctoragentParams object
// with the ability to set a timeout on a request.
func NewGetKdoctoragentParamsWithTimeout(timeout time.Duration) *GetKdoctoragentParams {
	return &GetKdoctoragentParams{
		timeout: timeout,
	}
}

// NewGetKdoctoragentParamsWithContext creates a new GetKdoctoragentParams object
// with the ability to set a context for a request.
func NewGetKdoctoragentParamsWithContext(ctx context.Context) *GetKdoctoragentParams {
	return &GetKdoctoragentParams{
		Context: ctx,
	}
}

// NewGetKdoctoragentParamsWithHTTPClient creates a new GetKdoctoragentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKdoctoragentParamsWithHTTPClient(client *http.Client) *GetKdoctoragentParams {
	return &GetKdoctoragentParams{
		HTTPClient: client,
	}
}

/*
GetKdoctoragentParams contains all the parameters to send to the API endpoint

	for the get kdoctoragent operation.

	Typically these are written to a http.Request.
*/
type GetKdoctoragentParams struct {

	/* Delay.

	   delay some second return response
	*/
	Delay *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kdoctoragent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKdoctoragentParams) WithDefaults() *GetKdoctoragentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kdoctoragent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKdoctoragentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kdoctoragent params
func (o *GetKdoctoragentParams) WithTimeout(timeout time.Duration) *GetKdoctoragentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kdoctoragent params
func (o *GetKdoctoragentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kdoctoragent params
func (o *GetKdoctoragentParams) WithContext(ctx context.Context) *GetKdoctoragentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kdoctoragent params
func (o *GetKdoctoragentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kdoctoragent params
func (o *GetKdoctoragentParams) WithHTTPClient(client *http.Client) *GetKdoctoragentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kdoctoragent params
func (o *GetKdoctoragentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelay adds the delay to the get kdoctoragent params
func (o *GetKdoctoragentParams) WithDelay(delay *int64) *GetKdoctoragentParams {
	o.SetDelay(delay)
	return o
}

// SetDelay adds the delay to the get kdoctoragent params
func (o *GetKdoctoragentParams) SetDelay(delay *int64) {
	o.Delay = delay
}

// WriteToRequest writes these params to a swagger request
func (o *GetKdoctoragentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Delay != nil {

		// query param delay
		var qrDelay int64

		if o.Delay != nil {
			qrDelay = *o.Delay
		}
		qDelay := swag.FormatInt64(qrDelay)
		if qDelay != "" {

			if err := r.SetQueryParam("delay", qDelay); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
