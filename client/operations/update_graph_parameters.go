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

// NewUpdateGraphParams creates a new UpdateGraphParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGraphParams() *UpdateGraphParams {
	return &UpdateGraphParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGraphParamsWithTimeout creates a new UpdateGraphParams object
// with the ability to set a timeout on a request.
func NewUpdateGraphParamsWithTimeout(timeout time.Duration) *UpdateGraphParams {
	return &UpdateGraphParams{
		timeout: timeout,
	}
}

// NewUpdateGraphParamsWithContext creates a new UpdateGraphParams object
// with the ability to set a context for a request.
func NewUpdateGraphParamsWithContext(ctx context.Context) *UpdateGraphParams {
	return &UpdateGraphParams{
		Context: ctx,
	}
}

// NewUpdateGraphParamsWithHTTPClient creates a new UpdateGraphParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGraphParamsWithHTTPClient(client *http.Client) *UpdateGraphParams {
	return &UpdateGraphParams{
		HTTPClient: client,
	}
}

/* UpdateGraphParams contains all the parameters to send to the API endpoint
   for the update graph operation.

   Typically these are written to a http.Request.
*/
type UpdateGraphParams struct {

	// RequestBody.
	RequestBody *models.GraphUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGraphParams) WithDefaults() *UpdateGraphParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGraphParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update graph params
func (o *UpdateGraphParams) WithTimeout(timeout time.Duration) *UpdateGraphParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update graph params
func (o *UpdateGraphParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update graph params
func (o *UpdateGraphParams) WithContext(ctx context.Context) *UpdateGraphParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update graph params
func (o *UpdateGraphParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update graph params
func (o *UpdateGraphParams) WithHTTPClient(client *http.Client) *UpdateGraphParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update graph params
func (o *UpdateGraphParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update graph params
func (o *UpdateGraphParams) WithRequestBody(requestBody *models.GraphUpdationParams) *UpdateGraphParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update graph params
func (o *UpdateGraphParams) SetRequestBody(requestBody *models.GraphUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGraphParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
