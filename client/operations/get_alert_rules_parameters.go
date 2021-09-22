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

// NewGetAlertRulesParams creates a new GetAlertRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertRulesParams() *GetAlertRulesParams {
	return &GetAlertRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertRulesParamsWithTimeout creates a new GetAlertRulesParams object
// with the ability to set a timeout on a request.
func NewGetAlertRulesParamsWithTimeout(timeout time.Duration) *GetAlertRulesParams {
	return &GetAlertRulesParams{
		timeout: timeout,
	}
}

// NewGetAlertRulesParamsWithContext creates a new GetAlertRulesParams object
// with the ability to set a context for a request.
func NewGetAlertRulesParamsWithContext(ctx context.Context) *GetAlertRulesParams {
	return &GetAlertRulesParams{
		Context: ctx,
	}
}

// NewGetAlertRulesParamsWithHTTPClient creates a new GetAlertRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertRulesParamsWithHTTPClient(client *http.Client) *GetAlertRulesParams {
	return &GetAlertRulesParams{
		HTTPClient: client,
	}
}

/* GetAlertRulesParams contains all the parameters to send to the API endpoint
   for the get alert rules operation.

   Typically these are written to a http.Request.
*/
type GetAlertRulesParams struct {

	// RequestBody.
	RequestBody *models.GetAlertRulesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRulesParams) WithDefaults() *GetAlertRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alert rules params
func (o *GetAlertRulesParams) WithTimeout(timeout time.Duration) *GetAlertRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert rules params
func (o *GetAlertRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert rules params
func (o *GetAlertRulesParams) WithContext(ctx context.Context) *GetAlertRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert rules params
func (o *GetAlertRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert rules params
func (o *GetAlertRulesParams) WithHTTPClient(client *http.Client) *GetAlertRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert rules params
func (o *GetAlertRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get alert rules params
func (o *GetAlertRulesParams) WithRequestBody(requestBody *models.GetAlertRulesRequestBody) *GetAlertRulesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get alert rules params
func (o *GetAlertRulesParams) SetRequestBody(requestBody *models.GetAlertRulesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
