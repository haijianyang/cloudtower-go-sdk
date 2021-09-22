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

// NewGetEntityFiltersParams creates a new GetEntityFiltersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEntityFiltersParams() *GetEntityFiltersParams {
	return &GetEntityFiltersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEntityFiltersParamsWithTimeout creates a new GetEntityFiltersParams object
// with the ability to set a timeout on a request.
func NewGetEntityFiltersParamsWithTimeout(timeout time.Duration) *GetEntityFiltersParams {
	return &GetEntityFiltersParams{
		timeout: timeout,
	}
}

// NewGetEntityFiltersParamsWithContext creates a new GetEntityFiltersParams object
// with the ability to set a context for a request.
func NewGetEntityFiltersParamsWithContext(ctx context.Context) *GetEntityFiltersParams {
	return &GetEntityFiltersParams{
		Context: ctx,
	}
}

// NewGetEntityFiltersParamsWithHTTPClient creates a new GetEntityFiltersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEntityFiltersParamsWithHTTPClient(client *http.Client) *GetEntityFiltersParams {
	return &GetEntityFiltersParams{
		HTTPClient: client,
	}
}

/* GetEntityFiltersParams contains all the parameters to send to the API endpoint
   for the get entity filters operation.

   Typically these are written to a http.Request.
*/
type GetEntityFiltersParams struct {

	// RequestBody.
	RequestBody *models.GetEntityFiltersRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get entity filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEntityFiltersParams) WithDefaults() *GetEntityFiltersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get entity filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEntityFiltersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get entity filters params
func (o *GetEntityFiltersParams) WithTimeout(timeout time.Duration) *GetEntityFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get entity filters params
func (o *GetEntityFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get entity filters params
func (o *GetEntityFiltersParams) WithContext(ctx context.Context) *GetEntityFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get entity filters params
func (o *GetEntityFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get entity filters params
func (o *GetEntityFiltersParams) WithHTTPClient(client *http.Client) *GetEntityFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get entity filters params
func (o *GetEntityFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get entity filters params
func (o *GetEntityFiltersParams) WithRequestBody(requestBody *models.GetEntityFiltersRequestBody) *GetEntityFiltersParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get entity filters params
func (o *GetEntityFiltersParams) SetRequestBody(requestBody *models.GetEntityFiltersRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetEntityFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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