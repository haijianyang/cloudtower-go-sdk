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

// NewUpdateVMParams creates a new UpdateVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMParams() *UpdateVMParams {
	return &UpdateVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMParamsWithTimeout creates a new UpdateVMParams object
// with the ability to set a timeout on a request.
func NewUpdateVMParamsWithTimeout(timeout time.Duration) *UpdateVMParams {
	return &UpdateVMParams{
		timeout: timeout,
	}
}

// NewUpdateVMParamsWithContext creates a new UpdateVMParams object
// with the ability to set a context for a request.
func NewUpdateVMParamsWithContext(ctx context.Context) *UpdateVMParams {
	return &UpdateVMParams{
		Context: ctx,
	}
}

// NewUpdateVMParamsWithHTTPClient creates a new UpdateVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMParamsWithHTTPClient(client *http.Client) *UpdateVMParams {
	return &UpdateVMParams{
		HTTPClient: client,
	}
}

/* UpdateVMParams contains all the parameters to send to the API endpoint
   for the update Vm operation.

   Typically these are written to a http.Request.
*/
type UpdateVMParams struct {

	// RequestBody.
	RequestBody *models.VMUpdateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMParams) WithDefaults() *UpdateVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Vm params
func (o *UpdateVMParams) WithTimeout(timeout time.Duration) *UpdateVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm params
func (o *UpdateVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm params
func (o *UpdateVMParams) WithContext(ctx context.Context) *UpdateVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm params
func (o *UpdateVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm params
func (o *UpdateVMParams) WithHTTPClient(client *http.Client) *UpdateVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm params
func (o *UpdateVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update Vm params
func (o *UpdateVMParams) WithRequestBody(requestBody *models.VMUpdateParams) *UpdateVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm params
func (o *UpdateVMParams) SetRequestBody(requestBody *models.VMUpdateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
