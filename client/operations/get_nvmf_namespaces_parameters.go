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

// NewGetNvmfNamespacesParams creates a new GetNvmfNamespacesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNvmfNamespacesParams() *GetNvmfNamespacesParams {
	return &GetNvmfNamespacesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNvmfNamespacesParamsWithTimeout creates a new GetNvmfNamespacesParams object
// with the ability to set a timeout on a request.
func NewGetNvmfNamespacesParamsWithTimeout(timeout time.Duration) *GetNvmfNamespacesParams {
	return &GetNvmfNamespacesParams{
		timeout: timeout,
	}
}

// NewGetNvmfNamespacesParamsWithContext creates a new GetNvmfNamespacesParams object
// with the ability to set a context for a request.
func NewGetNvmfNamespacesParamsWithContext(ctx context.Context) *GetNvmfNamespacesParams {
	return &GetNvmfNamespacesParams{
		Context: ctx,
	}
}

// NewGetNvmfNamespacesParamsWithHTTPClient creates a new GetNvmfNamespacesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNvmfNamespacesParamsWithHTTPClient(client *http.Client) *GetNvmfNamespacesParams {
	return &GetNvmfNamespacesParams{
		HTTPClient: client,
	}
}

/* GetNvmfNamespacesParams contains all the parameters to send to the API endpoint
   for the get nvmf namespaces operation.

   Typically these are written to a http.Request.
*/
type GetNvmfNamespacesParams struct {

	// RequestBody.
	RequestBody *models.GetNvmfNamespacesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nvmf namespaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespacesParams) WithDefaults() *GetNvmfNamespacesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nvmf namespaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespacesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) WithTimeout(timeout time.Duration) *GetNvmfNamespacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) WithContext(ctx context.Context) *GetNvmfNamespacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) WithHTTPClient(client *http.Client) *GetNvmfNamespacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) WithRequestBody(requestBody *models.GetNvmfNamespacesRequestBody) *GetNvmfNamespacesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nvmf namespaces params
func (o *GetNvmfNamespacesParams) SetRequestBody(requestBody *models.GetNvmfNamespacesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNvmfNamespacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
