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

// NewGetClusterSettingsesParams creates a new GetClusterSettingsesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterSettingsesParams() *GetClusterSettingsesParams {
	return &GetClusterSettingsesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterSettingsesParamsWithTimeout creates a new GetClusterSettingsesParams object
// with the ability to set a timeout on a request.
func NewGetClusterSettingsesParamsWithTimeout(timeout time.Duration) *GetClusterSettingsesParams {
	return &GetClusterSettingsesParams{
		timeout: timeout,
	}
}

// NewGetClusterSettingsesParamsWithContext creates a new GetClusterSettingsesParams object
// with the ability to set a context for a request.
func NewGetClusterSettingsesParamsWithContext(ctx context.Context) *GetClusterSettingsesParams {
	return &GetClusterSettingsesParams{
		Context: ctx,
	}
}

// NewGetClusterSettingsesParamsWithHTTPClient creates a new GetClusterSettingsesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterSettingsesParamsWithHTTPClient(client *http.Client) *GetClusterSettingsesParams {
	return &GetClusterSettingsesParams{
		HTTPClient: client,
	}
}

/* GetClusterSettingsesParams contains all the parameters to send to the API endpoint
   for the get cluster settingses operation.

   Typically these are written to a http.Request.
*/
type GetClusterSettingsesParams struct {

	// RequestBody.
	RequestBody *models.GetClusterSettingsesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster settingses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSettingsesParams) WithDefaults() *GetClusterSettingsesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster settingses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSettingsesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster settingses params
func (o *GetClusterSettingsesParams) WithTimeout(timeout time.Duration) *GetClusterSettingsesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster settingses params
func (o *GetClusterSettingsesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster settingses params
func (o *GetClusterSettingsesParams) WithContext(ctx context.Context) *GetClusterSettingsesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster settingses params
func (o *GetClusterSettingsesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster settingses params
func (o *GetClusterSettingsesParams) WithHTTPClient(client *http.Client) *GetClusterSettingsesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster settingses params
func (o *GetClusterSettingsesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get cluster settingses params
func (o *GetClusterSettingsesParams) WithRequestBody(requestBody *models.GetClusterSettingsesRequestBody) *GetClusterSettingsesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get cluster settingses params
func (o *GetClusterSettingsesParams) SetRequestBody(requestBody *models.GetClusterSettingsesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterSettingsesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
