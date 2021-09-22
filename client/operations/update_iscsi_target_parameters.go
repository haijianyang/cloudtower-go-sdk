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

// NewUpdateIscsiTargetParams creates a new UpdateIscsiTargetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIscsiTargetParams() *UpdateIscsiTargetParams {
	return &UpdateIscsiTargetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIscsiTargetParamsWithTimeout creates a new UpdateIscsiTargetParams object
// with the ability to set a timeout on a request.
func NewUpdateIscsiTargetParamsWithTimeout(timeout time.Duration) *UpdateIscsiTargetParams {
	return &UpdateIscsiTargetParams{
		timeout: timeout,
	}
}

// NewUpdateIscsiTargetParamsWithContext creates a new UpdateIscsiTargetParams object
// with the ability to set a context for a request.
func NewUpdateIscsiTargetParamsWithContext(ctx context.Context) *UpdateIscsiTargetParams {
	return &UpdateIscsiTargetParams{
		Context: ctx,
	}
}

// NewUpdateIscsiTargetParamsWithHTTPClient creates a new UpdateIscsiTargetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIscsiTargetParamsWithHTTPClient(client *http.Client) *UpdateIscsiTargetParams {
	return &UpdateIscsiTargetParams{
		HTTPClient: client,
	}
}

/* UpdateIscsiTargetParams contains all the parameters to send to the API endpoint
   for the update iscsi target operation.

   Typically these are written to a http.Request.
*/
type UpdateIscsiTargetParams struct {

	// RequestBody.
	RequestBody *models.IscsiTargetUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update iscsi target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIscsiTargetParams) WithDefaults() *UpdateIscsiTargetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update iscsi target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIscsiTargetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update iscsi target params
func (o *UpdateIscsiTargetParams) WithTimeout(timeout time.Duration) *UpdateIscsiTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update iscsi target params
func (o *UpdateIscsiTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update iscsi target params
func (o *UpdateIscsiTargetParams) WithContext(ctx context.Context) *UpdateIscsiTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update iscsi target params
func (o *UpdateIscsiTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update iscsi target params
func (o *UpdateIscsiTargetParams) WithHTTPClient(client *http.Client) *UpdateIscsiTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update iscsi target params
func (o *UpdateIscsiTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update iscsi target params
func (o *UpdateIscsiTargetParams) WithRequestBody(requestBody *models.IscsiTargetUpdationParams) *UpdateIscsiTargetParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update iscsi target params
func (o *UpdateIscsiTargetParams) SetRequestBody(requestBody *models.IscsiTargetUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIscsiTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
