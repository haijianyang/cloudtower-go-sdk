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

// NewGetSnmpTrapReceiversParams creates a new GetSnmpTrapReceiversParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSnmpTrapReceiversParams() *GetSnmpTrapReceiversParams {
	return &GetSnmpTrapReceiversParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnmpTrapReceiversParamsWithTimeout creates a new GetSnmpTrapReceiversParams object
// with the ability to set a timeout on a request.
func NewGetSnmpTrapReceiversParamsWithTimeout(timeout time.Duration) *GetSnmpTrapReceiversParams {
	return &GetSnmpTrapReceiversParams{
		timeout: timeout,
	}
}

// NewGetSnmpTrapReceiversParamsWithContext creates a new GetSnmpTrapReceiversParams object
// with the ability to set a context for a request.
func NewGetSnmpTrapReceiversParamsWithContext(ctx context.Context) *GetSnmpTrapReceiversParams {
	return &GetSnmpTrapReceiversParams{
		Context: ctx,
	}
}

// NewGetSnmpTrapReceiversParamsWithHTTPClient creates a new GetSnmpTrapReceiversParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSnmpTrapReceiversParamsWithHTTPClient(client *http.Client) *GetSnmpTrapReceiversParams {
	return &GetSnmpTrapReceiversParams{
		HTTPClient: client,
	}
}

/* GetSnmpTrapReceiversParams contains all the parameters to send to the API endpoint
   for the get snmp trap receivers operation.

   Typically these are written to a http.Request.
*/
type GetSnmpTrapReceiversParams struct {

	// RequestBody.
	RequestBody *models.GetSnmpTrapReceiversRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get snmp trap receivers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnmpTrapReceiversParams) WithDefaults() *GetSnmpTrapReceiversParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get snmp trap receivers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnmpTrapReceiversParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) WithTimeout(timeout time.Duration) *GetSnmpTrapReceiversParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) WithContext(ctx context.Context) *GetSnmpTrapReceiversParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) WithHTTPClient(client *http.Client) *GetSnmpTrapReceiversParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) WithRequestBody(requestBody *models.GetSnmpTrapReceiversRequestBody) *GetSnmpTrapReceiversParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get snmp trap receivers params
func (o *GetSnmpTrapReceiversParams) SetRequestBody(requestBody *models.GetSnmpTrapReceiversRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnmpTrapReceiversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
