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

// NewGetClustersParams creates a new GetClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClustersParams() *GetClustersParams {
	return &GetClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersParamsWithTimeout creates a new GetClustersParams object
// with the ability to set a timeout on a request.
func NewGetClustersParamsWithTimeout(timeout time.Duration) *GetClustersParams {
	return &GetClustersParams{
		timeout: timeout,
	}
}

// NewGetClustersParamsWithContext creates a new GetClustersParams object
// with the ability to set a context for a request.
func NewGetClustersParamsWithContext(ctx context.Context) *GetClustersParams {
	return &GetClustersParams{
		Context: ctx,
	}
}

// NewGetClustersParamsWithHTTPClient creates a new GetClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClustersParamsWithHTTPClient(client *http.Client) *GetClustersParams {
	return &GetClustersParams{
		HTTPClient: client,
	}
}

/* GetClustersParams contains all the parameters to send to the API endpoint
   for the get clusters operation.

   Typically these are written to a http.Request.
*/
type GetClustersParams struct {

	// RequestBody.
	RequestBody *models.GetClustersRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersParams) WithDefaults() *GetClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) WithTimeout(timeout time.Duration) *GetClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters params
func (o *GetClustersParams) WithContext(ctx context.Context) *GetClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters params
func (o *GetClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) WithHTTPClient(client *http.Client) *GetClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get clusters params
func (o *GetClustersParams) WithRequestBody(requestBody *models.GetClustersRequestBody) *GetClustersParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get clusters params
func (o *GetClustersParams) SetRequestBody(requestBody *models.GetClustersRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
