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

// NewGetNfsExportsParams creates a new GetNfsExportsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNfsExportsParams() *GetNfsExportsParams {
	return &GetNfsExportsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsExportsParamsWithTimeout creates a new GetNfsExportsParams object
// with the ability to set a timeout on a request.
func NewGetNfsExportsParamsWithTimeout(timeout time.Duration) *GetNfsExportsParams {
	return &GetNfsExportsParams{
		timeout: timeout,
	}
}

// NewGetNfsExportsParamsWithContext creates a new GetNfsExportsParams object
// with the ability to set a context for a request.
func NewGetNfsExportsParamsWithContext(ctx context.Context) *GetNfsExportsParams {
	return &GetNfsExportsParams{
		Context: ctx,
	}
}

// NewGetNfsExportsParamsWithHTTPClient creates a new GetNfsExportsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNfsExportsParamsWithHTTPClient(client *http.Client) *GetNfsExportsParams {
	return &GetNfsExportsParams{
		HTTPClient: client,
	}
}

/* GetNfsExportsParams contains all the parameters to send to the API endpoint
   for the get nfs exports operation.

   Typically these are written to a http.Request.
*/
type GetNfsExportsParams struct {

	// RequestBody.
	RequestBody *models.GetNfsExportsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nfs exports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsExportsParams) WithDefaults() *GetNfsExportsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nfs exports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsExportsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nfs exports params
func (o *GetNfsExportsParams) WithTimeout(timeout time.Duration) *GetNfsExportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs exports params
func (o *GetNfsExportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs exports params
func (o *GetNfsExportsParams) WithContext(ctx context.Context) *GetNfsExportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs exports params
func (o *GetNfsExportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs exports params
func (o *GetNfsExportsParams) WithHTTPClient(client *http.Client) *GetNfsExportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs exports params
func (o *GetNfsExportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nfs exports params
func (o *GetNfsExportsParams) WithRequestBody(requestBody *models.GetNfsExportsRequestBody) *GetNfsExportsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nfs exports params
func (o *GetNfsExportsParams) SetRequestBody(requestBody *models.GetNfsExportsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsExportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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