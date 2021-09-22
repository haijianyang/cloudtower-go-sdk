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

// NewGetSnapshotGroupsParams creates a new GetSnapshotGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSnapshotGroupsParams() *GetSnapshotGroupsParams {
	return &GetSnapshotGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotGroupsParamsWithTimeout creates a new GetSnapshotGroupsParams object
// with the ability to set a timeout on a request.
func NewGetSnapshotGroupsParamsWithTimeout(timeout time.Duration) *GetSnapshotGroupsParams {
	return &GetSnapshotGroupsParams{
		timeout: timeout,
	}
}

// NewGetSnapshotGroupsParamsWithContext creates a new GetSnapshotGroupsParams object
// with the ability to set a context for a request.
func NewGetSnapshotGroupsParamsWithContext(ctx context.Context) *GetSnapshotGroupsParams {
	return &GetSnapshotGroupsParams{
		Context: ctx,
	}
}

// NewGetSnapshotGroupsParamsWithHTTPClient creates a new GetSnapshotGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSnapshotGroupsParamsWithHTTPClient(client *http.Client) *GetSnapshotGroupsParams {
	return &GetSnapshotGroupsParams{
		HTTPClient: client,
	}
}

/* GetSnapshotGroupsParams contains all the parameters to send to the API endpoint
   for the get snapshot groups operation.

   Typically these are written to a http.Request.
*/
type GetSnapshotGroupsParams struct {

	// RequestBody.
	RequestBody *models.GetSnapshotGroupsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get snapshot groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotGroupsParams) WithDefaults() *GetSnapshotGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get snapshot groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get snapshot groups params
func (o *GetSnapshotGroupsParams) WithTimeout(timeout time.Duration) *GetSnapshotGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot groups params
func (o *GetSnapshotGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot groups params
func (o *GetSnapshotGroupsParams) WithContext(ctx context.Context) *GetSnapshotGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot groups params
func (o *GetSnapshotGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot groups params
func (o *GetSnapshotGroupsParams) WithHTTPClient(client *http.Client) *GetSnapshotGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot groups params
func (o *GetSnapshotGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get snapshot groups params
func (o *GetSnapshotGroupsParams) WithRequestBody(requestBody *models.GetSnapshotGroupsRequestBody) *GetSnapshotGroupsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get snapshot groups params
func (o *GetSnapshotGroupsParams) SetRequestBody(requestBody *models.GetSnapshotGroupsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
