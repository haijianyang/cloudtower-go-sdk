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

// NewUpdateClusterLicenseParams creates a new UpdateClusterLicenseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateClusterLicenseParams() *UpdateClusterLicenseParams {
	return &UpdateClusterLicenseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterLicenseParamsWithTimeout creates a new UpdateClusterLicenseParams object
// with the ability to set a timeout on a request.
func NewUpdateClusterLicenseParamsWithTimeout(timeout time.Duration) *UpdateClusterLicenseParams {
	return &UpdateClusterLicenseParams{
		timeout: timeout,
	}
}

// NewUpdateClusterLicenseParamsWithContext creates a new UpdateClusterLicenseParams object
// with the ability to set a context for a request.
func NewUpdateClusterLicenseParamsWithContext(ctx context.Context) *UpdateClusterLicenseParams {
	return &UpdateClusterLicenseParams{
		Context: ctx,
	}
}

// NewUpdateClusterLicenseParamsWithHTTPClient creates a new UpdateClusterLicenseParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateClusterLicenseParamsWithHTTPClient(client *http.Client) *UpdateClusterLicenseParams {
	return &UpdateClusterLicenseParams{
		HTTPClient: client,
	}
}

/* UpdateClusterLicenseParams contains all the parameters to send to the API endpoint
   for the update cluster license operation.

   Typically these are written to a http.Request.
*/
type UpdateClusterLicenseParams struct {

	// RequestBody.
	RequestBody *models.ClusterLicenseUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update cluster license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClusterLicenseParams) WithDefaults() *UpdateClusterLicenseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update cluster license params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClusterLicenseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update cluster license params
func (o *UpdateClusterLicenseParams) WithTimeout(timeout time.Duration) *UpdateClusterLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster license params
func (o *UpdateClusterLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster license params
func (o *UpdateClusterLicenseParams) WithContext(ctx context.Context) *UpdateClusterLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster license params
func (o *UpdateClusterLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster license params
func (o *UpdateClusterLicenseParams) WithHTTPClient(client *http.Client) *UpdateClusterLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster license params
func (o *UpdateClusterLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update cluster license params
func (o *UpdateClusterLicenseParams) WithRequestBody(requestBody *models.ClusterLicenseUpdationParams) *UpdateClusterLicenseParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update cluster license params
func (o *UpdateClusterLicenseParams) SetRequestBody(requestBody *models.ClusterLicenseUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
