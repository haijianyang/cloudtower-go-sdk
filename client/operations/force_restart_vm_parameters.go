// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/haijianyang/cloudtower-go-sdk/models"
)

// NewForceRestartVMParams creates a new ForceRestartVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForceRestartVMParams() *ForceRestartVMParams {
	return &ForceRestartVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForceRestartVMParamsWithTimeout creates a new ForceRestartVMParams object
// with the ability to set a timeout on a request.
func NewForceRestartVMParamsWithTimeout(timeout time.Duration) *ForceRestartVMParams {
	return &ForceRestartVMParams{
		timeout: timeout,
	}
}

// NewForceRestartVMParamsWithContext creates a new ForceRestartVMParams object
// with the ability to set a context for a request.
func NewForceRestartVMParamsWithContext(ctx context.Context) *ForceRestartVMParams {
	return &ForceRestartVMParams{
		Context: ctx,
	}
}

// NewForceRestartVMParamsWithHTTPClient creates a new ForceRestartVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewForceRestartVMParamsWithHTTPClient(client *http.Client) *ForceRestartVMParams {
	return &ForceRestartVMParams{
		HTTPClient: client,
	}
}

/* ForceRestartVMParams contains all the parameters to send to the API endpoint
   for the force restart Vm operation.

   Typically these are written to a http.Request.
*/
type ForceRestartVMParams struct {

	// RequestBody.
	RequestBody *models.VMOperateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the force restart Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForceRestartVMParams) WithDefaults() *ForceRestartVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the force restart Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForceRestartVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the force restart Vm params
func (o *ForceRestartVMParams) WithTimeout(timeout time.Duration) *ForceRestartVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the force restart Vm params
func (o *ForceRestartVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the force restart Vm params
func (o *ForceRestartVMParams) WithContext(ctx context.Context) *ForceRestartVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the force restart Vm params
func (o *ForceRestartVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the force restart Vm params
func (o *ForceRestartVMParams) WithHTTPClient(client *http.Client) *ForceRestartVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the force restart Vm params
func (o *ForceRestartVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the force restart Vm params
func (o *ForceRestartVMParams) WithRequestBody(requestBody *models.VMOperateParams) *ForceRestartVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the force restart Vm params
func (o *ForceRestartVMParams) SetRequestBody(requestBody *models.VMOperateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *ForceRestartVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
