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

// NewUpdateConsistencyGroupSnapshotParams creates a new UpdateConsistencyGroupSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConsistencyGroupSnapshotParams() *UpdateConsistencyGroupSnapshotParams {
	return &UpdateConsistencyGroupSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConsistencyGroupSnapshotParamsWithTimeout creates a new UpdateConsistencyGroupSnapshotParams object
// with the ability to set a timeout on a request.
func NewUpdateConsistencyGroupSnapshotParamsWithTimeout(timeout time.Duration) *UpdateConsistencyGroupSnapshotParams {
	return &UpdateConsistencyGroupSnapshotParams{
		timeout: timeout,
	}
}

// NewUpdateConsistencyGroupSnapshotParamsWithContext creates a new UpdateConsistencyGroupSnapshotParams object
// with the ability to set a context for a request.
func NewUpdateConsistencyGroupSnapshotParamsWithContext(ctx context.Context) *UpdateConsistencyGroupSnapshotParams {
	return &UpdateConsistencyGroupSnapshotParams{
		Context: ctx,
	}
}

// NewUpdateConsistencyGroupSnapshotParamsWithHTTPClient creates a new UpdateConsistencyGroupSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConsistencyGroupSnapshotParamsWithHTTPClient(client *http.Client) *UpdateConsistencyGroupSnapshotParams {
	return &UpdateConsistencyGroupSnapshotParams{
		HTTPClient: client,
	}
}

/* UpdateConsistencyGroupSnapshotParams contains all the parameters to send to the API endpoint
   for the update consistency group snapshot operation.

   Typically these are written to a http.Request.
*/
type UpdateConsistencyGroupSnapshotParams struct {

	// RequestBody.
	RequestBody *models.ConsistencyGroupSnapshotUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update consistency group snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConsistencyGroupSnapshotParams) WithDefaults() *UpdateConsistencyGroupSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update consistency group snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConsistencyGroupSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) WithTimeout(timeout time.Duration) *UpdateConsistencyGroupSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) WithContext(ctx context.Context) *UpdateConsistencyGroupSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) WithHTTPClient(client *http.Client) *UpdateConsistencyGroupSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) WithRequestBody(requestBody *models.ConsistencyGroupSnapshotUpdationParams) *UpdateConsistencyGroupSnapshotParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update consistency group snapshot params
func (o *UpdateConsistencyGroupSnapshotParams) SetRequestBody(requestBody *models.ConsistencyGroupSnapshotUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConsistencyGroupSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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