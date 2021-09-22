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

// NewGetGlobalSettingsesParams creates a new GetGlobalSettingsesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalSettingsesParams() *GetGlobalSettingsesParams {
	return &GetGlobalSettingsesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalSettingsesParamsWithTimeout creates a new GetGlobalSettingsesParams object
// with the ability to set a timeout on a request.
func NewGetGlobalSettingsesParamsWithTimeout(timeout time.Duration) *GetGlobalSettingsesParams {
	return &GetGlobalSettingsesParams{
		timeout: timeout,
	}
}

// NewGetGlobalSettingsesParamsWithContext creates a new GetGlobalSettingsesParams object
// with the ability to set a context for a request.
func NewGetGlobalSettingsesParamsWithContext(ctx context.Context) *GetGlobalSettingsesParams {
	return &GetGlobalSettingsesParams{
		Context: ctx,
	}
}

// NewGetGlobalSettingsesParamsWithHTTPClient creates a new GetGlobalSettingsesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalSettingsesParamsWithHTTPClient(client *http.Client) *GetGlobalSettingsesParams {
	return &GetGlobalSettingsesParams{
		HTTPClient: client,
	}
}

/* GetGlobalSettingsesParams contains all the parameters to send to the API endpoint
   for the get global settingses operation.

   Typically these are written to a http.Request.
*/
type GetGlobalSettingsesParams struct {

	// RequestBody.
	RequestBody *models.GetGlobalSettingsesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global settingses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalSettingsesParams) WithDefaults() *GetGlobalSettingsesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global settingses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalSettingsesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get global settingses params
func (o *GetGlobalSettingsesParams) WithTimeout(timeout time.Duration) *GetGlobalSettingsesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global settingses params
func (o *GetGlobalSettingsesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global settingses params
func (o *GetGlobalSettingsesParams) WithContext(ctx context.Context) *GetGlobalSettingsesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global settingses params
func (o *GetGlobalSettingsesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global settingses params
func (o *GetGlobalSettingsesParams) WithHTTPClient(client *http.Client) *GetGlobalSettingsesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global settingses params
func (o *GetGlobalSettingsesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get global settingses params
func (o *GetGlobalSettingsesParams) WithRequestBody(requestBody *models.GetGlobalSettingsesRequestBody) *GetGlobalSettingsesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get global settingses params
func (o *GetGlobalSettingsesParams) SetRequestBody(requestBody *models.GetGlobalSettingsesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalSettingsesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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