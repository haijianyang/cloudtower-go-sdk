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

// NewGetAlertRulesConnectionParams creates a new GetAlertRulesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertRulesConnectionParams() *GetAlertRulesConnectionParams {
	return &GetAlertRulesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertRulesConnectionParamsWithTimeout creates a new GetAlertRulesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetAlertRulesConnectionParamsWithTimeout(timeout time.Duration) *GetAlertRulesConnectionParams {
	return &GetAlertRulesConnectionParams{
		timeout: timeout,
	}
}

// NewGetAlertRulesConnectionParamsWithContext creates a new GetAlertRulesConnectionParams object
// with the ability to set a context for a request.
func NewGetAlertRulesConnectionParamsWithContext(ctx context.Context) *GetAlertRulesConnectionParams {
	return &GetAlertRulesConnectionParams{
		Context: ctx,
	}
}

// NewGetAlertRulesConnectionParamsWithHTTPClient creates a new GetAlertRulesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertRulesConnectionParamsWithHTTPClient(client *http.Client) *GetAlertRulesConnectionParams {
	return &GetAlertRulesConnectionParams{
		HTTPClient: client,
	}
}

/* GetAlertRulesConnectionParams contains all the parameters to send to the API endpoint
   for the get alert rules connection operation.

   Typically these are written to a http.Request.
*/
type GetAlertRulesConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetAlertRulesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert rules connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRulesConnectionParams) WithDefaults() *GetAlertRulesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert rules connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertRulesConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) WithTimeout(timeout time.Duration) *GetAlertRulesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) WithContext(ctx context.Context) *GetAlertRulesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) WithHTTPClient(client *http.Client) *GetAlertRulesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) WithRequestBody(requestBody *models.GetAlertRulesConnectionRequestBody) *GetAlertRulesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get alert rules connection params
func (o *GetAlertRulesConnectionParams) SetRequestBody(requestBody *models.GetAlertRulesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertRulesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
