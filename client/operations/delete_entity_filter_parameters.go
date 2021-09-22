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

// NewDeleteEntityFilterParams creates a new DeleteEntityFilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEntityFilterParams() *DeleteEntityFilterParams {
	return &DeleteEntityFilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEntityFilterParamsWithTimeout creates a new DeleteEntityFilterParams object
// with the ability to set a timeout on a request.
func NewDeleteEntityFilterParamsWithTimeout(timeout time.Duration) *DeleteEntityFilterParams {
	return &DeleteEntityFilterParams{
		timeout: timeout,
	}
}

// NewDeleteEntityFilterParamsWithContext creates a new DeleteEntityFilterParams object
// with the ability to set a context for a request.
func NewDeleteEntityFilterParamsWithContext(ctx context.Context) *DeleteEntityFilterParams {
	return &DeleteEntityFilterParams{
		Context: ctx,
	}
}

// NewDeleteEntityFilterParamsWithHTTPClient creates a new DeleteEntityFilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEntityFilterParamsWithHTTPClient(client *http.Client) *DeleteEntityFilterParams {
	return &DeleteEntityFilterParams{
		HTTPClient: client,
	}
}

/* DeleteEntityFilterParams contains all the parameters to send to the API endpoint
   for the delete entity filter operation.

   Typically these are written to a http.Request.
*/
type DeleteEntityFilterParams struct {

	// RequestBody.
	RequestBody *models.EntityFilterDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete entity filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEntityFilterParams) WithDefaults() *DeleteEntityFilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete entity filter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEntityFilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete entity filter params
func (o *DeleteEntityFilterParams) WithTimeout(timeout time.Duration) *DeleteEntityFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete entity filter params
func (o *DeleteEntityFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete entity filter params
func (o *DeleteEntityFilterParams) WithContext(ctx context.Context) *DeleteEntityFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete entity filter params
func (o *DeleteEntityFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete entity filter params
func (o *DeleteEntityFilterParams) WithHTTPClient(client *http.Client) *DeleteEntityFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete entity filter params
func (o *DeleteEntityFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete entity filter params
func (o *DeleteEntityFilterParams) WithRequestBody(requestBody *models.EntityFilterDeletionParams) *DeleteEntityFilterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete entity filter params
func (o *DeleteEntityFilterParams) SetRequestBody(requestBody *models.EntityFilterDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEntityFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
