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

// NewDeleteVMTemplateFromVMParams creates a new DeleteVMTemplateFromVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVMTemplateFromVMParams() *DeleteVMTemplateFromVMParams {
	return &DeleteVMTemplateFromVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVMTemplateFromVMParamsWithTimeout creates a new DeleteVMTemplateFromVMParams object
// with the ability to set a timeout on a request.
func NewDeleteVMTemplateFromVMParamsWithTimeout(timeout time.Duration) *DeleteVMTemplateFromVMParams {
	return &DeleteVMTemplateFromVMParams{
		timeout: timeout,
	}
}

// NewDeleteVMTemplateFromVMParamsWithContext creates a new DeleteVMTemplateFromVMParams object
// with the ability to set a context for a request.
func NewDeleteVMTemplateFromVMParamsWithContext(ctx context.Context) *DeleteVMTemplateFromVMParams {
	return &DeleteVMTemplateFromVMParams{
		Context: ctx,
	}
}

// NewDeleteVMTemplateFromVMParamsWithHTTPClient creates a new DeleteVMTemplateFromVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVMTemplateFromVMParamsWithHTTPClient(client *http.Client) *DeleteVMTemplateFromVMParams {
	return &DeleteVMTemplateFromVMParams{
		HTTPClient: client,
	}
}

/* DeleteVMTemplateFromVMParams contains all the parameters to send to the API endpoint
   for the delete Vm template from Vm operation.

   Typically these are written to a http.Request.
*/
type DeleteVMTemplateFromVMParams struct {

	// RequestBody.
	RequestBody *models.VMTemplateDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Vm template from Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMTemplateFromVMParams) WithDefaults() *DeleteVMTemplateFromVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Vm template from Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMTemplateFromVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) WithTimeout(timeout time.Duration) *DeleteVMTemplateFromVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) WithContext(ctx context.Context) *DeleteVMTemplateFromVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) WithHTTPClient(client *http.Client) *DeleteVMTemplateFromVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) WithRequestBody(requestBody *models.VMTemplateDeletionParams) *DeleteVMTemplateFromVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete Vm template from Vm params
func (o *DeleteVMTemplateFromVMParams) SetRequestBody(requestBody *models.VMTemplateDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVMTemplateFromVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
