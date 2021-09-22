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

// NewInstallVmtoolsParams creates a new InstallVmtoolsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstallVmtoolsParams() *InstallVmtoolsParams {
	return &InstallVmtoolsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstallVmtoolsParamsWithTimeout creates a new InstallVmtoolsParams object
// with the ability to set a timeout on a request.
func NewInstallVmtoolsParamsWithTimeout(timeout time.Duration) *InstallVmtoolsParams {
	return &InstallVmtoolsParams{
		timeout: timeout,
	}
}

// NewInstallVmtoolsParamsWithContext creates a new InstallVmtoolsParams object
// with the ability to set a context for a request.
func NewInstallVmtoolsParamsWithContext(ctx context.Context) *InstallVmtoolsParams {
	return &InstallVmtoolsParams{
		Context: ctx,
	}
}

// NewInstallVmtoolsParamsWithHTTPClient creates a new InstallVmtoolsParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstallVmtoolsParamsWithHTTPClient(client *http.Client) *InstallVmtoolsParams {
	return &InstallVmtoolsParams{
		HTTPClient: client,
	}
}

/* InstallVmtoolsParams contains all the parameters to send to the API endpoint
   for the install vmtools operation.

   Typically these are written to a http.Request.
*/
type InstallVmtoolsParams struct {

	// RequestBody.
	RequestBody *models.InstallVmtoolsParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the install vmtools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallVmtoolsParams) WithDefaults() *InstallVmtoolsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the install vmtools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallVmtoolsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the install vmtools params
func (o *InstallVmtoolsParams) WithTimeout(timeout time.Duration) *InstallVmtoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install vmtools params
func (o *InstallVmtoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install vmtools params
func (o *InstallVmtoolsParams) WithContext(ctx context.Context) *InstallVmtoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install vmtools params
func (o *InstallVmtoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install vmtools params
func (o *InstallVmtoolsParams) WithHTTPClient(client *http.Client) *InstallVmtoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install vmtools params
func (o *InstallVmtoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the install vmtools params
func (o *InstallVmtoolsParams) WithRequestBody(requestBody *models.InstallVmtoolsParams) *InstallVmtoolsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the install vmtools params
func (o *InstallVmtoolsParams) SetRequestBody(requestBody *models.InstallVmtoolsParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *InstallVmtoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
