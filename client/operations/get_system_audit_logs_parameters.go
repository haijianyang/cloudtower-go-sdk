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

// NewGetSystemAuditLogsParams creates a new GetSystemAuditLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemAuditLogsParams() *GetSystemAuditLogsParams {
	return &GetSystemAuditLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemAuditLogsParamsWithTimeout creates a new GetSystemAuditLogsParams object
// with the ability to set a timeout on a request.
func NewGetSystemAuditLogsParamsWithTimeout(timeout time.Duration) *GetSystemAuditLogsParams {
	return &GetSystemAuditLogsParams{
		timeout: timeout,
	}
}

// NewGetSystemAuditLogsParamsWithContext creates a new GetSystemAuditLogsParams object
// with the ability to set a context for a request.
func NewGetSystemAuditLogsParamsWithContext(ctx context.Context) *GetSystemAuditLogsParams {
	return &GetSystemAuditLogsParams{
		Context: ctx,
	}
}

// NewGetSystemAuditLogsParamsWithHTTPClient creates a new GetSystemAuditLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemAuditLogsParamsWithHTTPClient(client *http.Client) *GetSystemAuditLogsParams {
	return &GetSystemAuditLogsParams{
		HTTPClient: client,
	}
}

/* GetSystemAuditLogsParams contains all the parameters to send to the API endpoint
   for the get system audit logs operation.

   Typically these are written to a http.Request.
*/
type GetSystemAuditLogsParams struct {

	// RequestBody.
	RequestBody *models.GetSystemAuditLogsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system audit logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemAuditLogsParams) WithDefaults() *GetSystemAuditLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system audit logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemAuditLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system audit logs params
func (o *GetSystemAuditLogsParams) WithTimeout(timeout time.Duration) *GetSystemAuditLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system audit logs params
func (o *GetSystemAuditLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system audit logs params
func (o *GetSystemAuditLogsParams) WithContext(ctx context.Context) *GetSystemAuditLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system audit logs params
func (o *GetSystemAuditLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system audit logs params
func (o *GetSystemAuditLogsParams) WithHTTPClient(client *http.Client) *GetSystemAuditLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system audit logs params
func (o *GetSystemAuditLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get system audit logs params
func (o *GetSystemAuditLogsParams) WithRequestBody(requestBody *models.GetSystemAuditLogsRequestBody) *GetSystemAuditLogsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get system audit logs params
func (o *GetSystemAuditLogsParams) SetRequestBody(requestBody *models.GetSystemAuditLogsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemAuditLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
