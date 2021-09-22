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

// NewGetConsistencyGroupSnapshotsConnectionParams creates a new GetConsistencyGroupSnapshotsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConsistencyGroupSnapshotsConnectionParams() *GetConsistencyGroupSnapshotsConnectionParams {
	return &GetConsistencyGroupSnapshotsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConsistencyGroupSnapshotsConnectionParamsWithTimeout creates a new GetConsistencyGroupSnapshotsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetConsistencyGroupSnapshotsConnectionParamsWithTimeout(timeout time.Duration) *GetConsistencyGroupSnapshotsConnectionParams {
	return &GetConsistencyGroupSnapshotsConnectionParams{
		timeout: timeout,
	}
}

// NewGetConsistencyGroupSnapshotsConnectionParamsWithContext creates a new GetConsistencyGroupSnapshotsConnectionParams object
// with the ability to set a context for a request.
func NewGetConsistencyGroupSnapshotsConnectionParamsWithContext(ctx context.Context) *GetConsistencyGroupSnapshotsConnectionParams {
	return &GetConsistencyGroupSnapshotsConnectionParams{
		Context: ctx,
	}
}

// NewGetConsistencyGroupSnapshotsConnectionParamsWithHTTPClient creates a new GetConsistencyGroupSnapshotsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConsistencyGroupSnapshotsConnectionParamsWithHTTPClient(client *http.Client) *GetConsistencyGroupSnapshotsConnectionParams {
	return &GetConsistencyGroupSnapshotsConnectionParams{
		HTTPClient: client,
	}
}

/* GetConsistencyGroupSnapshotsConnectionParams contains all the parameters to send to the API endpoint
   for the get consistency group snapshots connection operation.

   Typically these are written to a http.Request.
*/
type GetConsistencyGroupSnapshotsConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetConsistencyGroupSnapshotsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get consistency group snapshots connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsistencyGroupSnapshotsConnectionParams) WithDefaults() *GetConsistencyGroupSnapshotsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get consistency group snapshots connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConsistencyGroupSnapshotsConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) WithTimeout(timeout time.Duration) *GetConsistencyGroupSnapshotsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) WithContext(ctx context.Context) *GetConsistencyGroupSnapshotsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) WithHTTPClient(client *http.Client) *GetConsistencyGroupSnapshotsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) WithRequestBody(requestBody *models.GetConsistencyGroupSnapshotsConnectionRequestBody) *GetConsistencyGroupSnapshotsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get consistency group snapshots connection params
func (o *GetConsistencyGroupSnapshotsConnectionParams) SetRequestBody(requestBody *models.GetConsistencyGroupSnapshotsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetConsistencyGroupSnapshotsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
