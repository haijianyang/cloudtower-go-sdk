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

// NewCreateVMFromTemplateParams creates a new CreateVMFromTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVMFromTemplateParams() *CreateVMFromTemplateParams {
	return &CreateVMFromTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMFromTemplateParamsWithTimeout creates a new CreateVMFromTemplateParams object
// with the ability to set a timeout on a request.
func NewCreateVMFromTemplateParamsWithTimeout(timeout time.Duration) *CreateVMFromTemplateParams {
	return &CreateVMFromTemplateParams{
		timeout: timeout,
	}
}

// NewCreateVMFromTemplateParamsWithContext creates a new CreateVMFromTemplateParams object
// with the ability to set a context for a request.
func NewCreateVMFromTemplateParamsWithContext(ctx context.Context) *CreateVMFromTemplateParams {
	return &CreateVMFromTemplateParams{
		Context: ctx,
	}
}

// NewCreateVMFromTemplateParamsWithHTTPClient creates a new CreateVMFromTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVMFromTemplateParamsWithHTTPClient(client *http.Client) *CreateVMFromTemplateParams {
	return &CreateVMFromTemplateParams{
		HTTPClient: client,
	}
}

/* CreateVMFromTemplateParams contains all the parameters to send to the API endpoint
   for the create Vm from template operation.

   Typically these are written to a http.Request.
*/
type CreateVMFromTemplateParams struct {

	// RequestBody.
	RequestBody []*models.VMCreateVMFromTemplateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create Vm from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMFromTemplateParams) WithDefaults() *CreateVMFromTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Vm from template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMFromTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create Vm from template params
func (o *CreateVMFromTemplateParams) WithTimeout(timeout time.Duration) *CreateVMFromTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Vm from template params
func (o *CreateVMFromTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Vm from template params
func (o *CreateVMFromTemplateParams) WithContext(ctx context.Context) *CreateVMFromTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Vm from template params
func (o *CreateVMFromTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Vm from template params
func (o *CreateVMFromTemplateParams) WithHTTPClient(client *http.Client) *CreateVMFromTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Vm from template params
func (o *CreateVMFromTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the create Vm from template params
func (o *CreateVMFromTemplateParams) WithRequestBody(requestBody []*models.VMCreateVMFromTemplateParams) *CreateVMFromTemplateParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create Vm from template params
func (o *CreateVMFromTemplateParams) SetRequestBody(requestBody []*models.VMCreateVMFromTemplateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMFromTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
