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

// NewGetIpmisParams creates a new GetIpmisParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIpmisParams() *GetIpmisParams {
	return &GetIpmisParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIpmisParamsWithTimeout creates a new GetIpmisParams object
// with the ability to set a timeout on a request.
func NewGetIpmisParamsWithTimeout(timeout time.Duration) *GetIpmisParams {
	return &GetIpmisParams{
		timeout: timeout,
	}
}

// NewGetIpmisParamsWithContext creates a new GetIpmisParams object
// with the ability to set a context for a request.
func NewGetIpmisParamsWithContext(ctx context.Context) *GetIpmisParams {
	return &GetIpmisParams{
		Context: ctx,
	}
}

// NewGetIpmisParamsWithHTTPClient creates a new GetIpmisParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIpmisParamsWithHTTPClient(client *http.Client) *GetIpmisParams {
	return &GetIpmisParams{
		HTTPClient: client,
	}
}

/* GetIpmisParams contains all the parameters to send to the API endpoint
   for the get ipmis operation.

   Typically these are written to a http.Request.
*/
type GetIpmisParams struct {

	// RequestBody.
	RequestBody *models.GetIpmisRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ipmis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpmisParams) WithDefaults() *GetIpmisParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ipmis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpmisParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ipmis params
func (o *GetIpmisParams) WithTimeout(timeout time.Duration) *GetIpmisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ipmis params
func (o *GetIpmisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ipmis params
func (o *GetIpmisParams) WithContext(ctx context.Context) *GetIpmisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ipmis params
func (o *GetIpmisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ipmis params
func (o *GetIpmisParams) WithHTTPClient(client *http.Client) *GetIpmisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ipmis params
func (o *GetIpmisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get ipmis params
func (o *GetIpmisParams) WithRequestBody(requestBody *models.GetIpmisRequestBody) *GetIpmisParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get ipmis params
func (o *GetIpmisParams) SetRequestBody(requestBody *models.GetIpmisRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetIpmisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
