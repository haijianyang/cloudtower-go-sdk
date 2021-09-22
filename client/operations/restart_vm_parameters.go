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

// NewRestartVMParams creates a new RestartVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartVMParams() *RestartVMParams {
	return &RestartVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartVMParamsWithTimeout creates a new RestartVMParams object
// with the ability to set a timeout on a request.
func NewRestartVMParamsWithTimeout(timeout time.Duration) *RestartVMParams {
	return &RestartVMParams{
		timeout: timeout,
	}
}

// NewRestartVMParamsWithContext creates a new RestartVMParams object
// with the ability to set a context for a request.
func NewRestartVMParamsWithContext(ctx context.Context) *RestartVMParams {
	return &RestartVMParams{
		Context: ctx,
	}
}

// NewRestartVMParamsWithHTTPClient creates a new RestartVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartVMParamsWithHTTPClient(client *http.Client) *RestartVMParams {
	return &RestartVMParams{
		HTTPClient: client,
	}
}

/* RestartVMParams contains all the parameters to send to the API endpoint
   for the restart Vm operation.

   Typically these are written to a http.Request.
*/
type RestartVMParams struct {

	// RequestBody.
	RequestBody *models.VMOperateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restart Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartVMParams) WithDefaults() *RestartVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restart Vm params
func (o *RestartVMParams) WithTimeout(timeout time.Duration) *RestartVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart Vm params
func (o *RestartVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart Vm params
func (o *RestartVMParams) WithContext(ctx context.Context) *RestartVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart Vm params
func (o *RestartVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart Vm params
func (o *RestartVMParams) WithHTTPClient(client *http.Client) *RestartVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart Vm params
func (o *RestartVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the restart Vm params
func (o *RestartVMParams) WithRequestBody(requestBody *models.VMOperateParams) *RestartVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the restart Vm params
func (o *RestartVMParams) SetRequestBody(requestBody *models.VMOperateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RestartVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
