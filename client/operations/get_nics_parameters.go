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

// NewGetNicsParams creates a new GetNicsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNicsParams() *GetNicsParams {
	return &GetNicsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNicsParamsWithTimeout creates a new GetNicsParams object
// with the ability to set a timeout on a request.
func NewGetNicsParamsWithTimeout(timeout time.Duration) *GetNicsParams {
	return &GetNicsParams{
		timeout: timeout,
	}
}

// NewGetNicsParamsWithContext creates a new GetNicsParams object
// with the ability to set a context for a request.
func NewGetNicsParamsWithContext(ctx context.Context) *GetNicsParams {
	return &GetNicsParams{
		Context: ctx,
	}
}

// NewGetNicsParamsWithHTTPClient creates a new GetNicsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNicsParamsWithHTTPClient(client *http.Client) *GetNicsParams {
	return &GetNicsParams{
		HTTPClient: client,
	}
}

/* GetNicsParams contains all the parameters to send to the API endpoint
   for the get nics operation.

   Typically these are written to a http.Request.
*/
type GetNicsParams struct {

	// RequestBody.
	RequestBody *models.GetNicsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNicsParams) WithDefaults() *GetNicsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNicsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nics params
func (o *GetNicsParams) WithTimeout(timeout time.Duration) *GetNicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nics params
func (o *GetNicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nics params
func (o *GetNicsParams) WithContext(ctx context.Context) *GetNicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nics params
func (o *GetNicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nics params
func (o *GetNicsParams) WithHTTPClient(client *http.Client) *GetNicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nics params
func (o *GetNicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nics params
func (o *GetNicsParams) WithRequestBody(requestBody *models.GetNicsRequestBody) *GetNicsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nics params
func (o *GetNicsParams) SetRequestBody(requestBody *models.GetNicsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
