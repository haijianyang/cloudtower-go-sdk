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

// NewCreateGraphParams creates a new CreateGraphParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGraphParams() *CreateGraphParams {
	return &CreateGraphParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGraphParamsWithTimeout creates a new CreateGraphParams object
// with the ability to set a timeout on a request.
func NewCreateGraphParamsWithTimeout(timeout time.Duration) *CreateGraphParams {
	return &CreateGraphParams{
		timeout: timeout,
	}
}

// NewCreateGraphParamsWithContext creates a new CreateGraphParams object
// with the ability to set a context for a request.
func NewCreateGraphParamsWithContext(ctx context.Context) *CreateGraphParams {
	return &CreateGraphParams{
		Context: ctx,
	}
}

// NewCreateGraphParamsWithHTTPClient creates a new CreateGraphParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGraphParamsWithHTTPClient(client *http.Client) *CreateGraphParams {
	return &CreateGraphParams{
		HTTPClient: client,
	}
}

/* CreateGraphParams contains all the parameters to send to the API endpoint
   for the create graph operation.

   Typically these are written to a http.Request.
*/
type CreateGraphParams struct {

	// RequestBody.
	RequestBody []*models.GraphCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGraphParams) WithDefaults() *CreateGraphParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGraphParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create graph params
func (o *CreateGraphParams) WithTimeout(timeout time.Duration) *CreateGraphParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create graph params
func (o *CreateGraphParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create graph params
func (o *CreateGraphParams) WithContext(ctx context.Context) *CreateGraphParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create graph params
func (o *CreateGraphParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create graph params
func (o *CreateGraphParams) WithHTTPClient(client *http.Client) *CreateGraphParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create graph params
func (o *CreateGraphParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the create graph params
func (o *CreateGraphParams) WithRequestBody(requestBody []*models.GraphCreationParams) *CreateGraphParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create graph params
func (o *CreateGraphParams) SetRequestBody(requestBody []*models.GraphCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGraphParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
