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

// NewDeleteGraphParams creates a new DeleteGraphParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGraphParams() *DeleteGraphParams {
	return &DeleteGraphParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGraphParamsWithTimeout creates a new DeleteGraphParams object
// with the ability to set a timeout on a request.
func NewDeleteGraphParamsWithTimeout(timeout time.Duration) *DeleteGraphParams {
	return &DeleteGraphParams{
		timeout: timeout,
	}
}

// NewDeleteGraphParamsWithContext creates a new DeleteGraphParams object
// with the ability to set a context for a request.
func NewDeleteGraphParamsWithContext(ctx context.Context) *DeleteGraphParams {
	return &DeleteGraphParams{
		Context: ctx,
	}
}

// NewDeleteGraphParamsWithHTTPClient creates a new DeleteGraphParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGraphParamsWithHTTPClient(client *http.Client) *DeleteGraphParams {
	return &DeleteGraphParams{
		HTTPClient: client,
	}
}

/* DeleteGraphParams contains all the parameters to send to the API endpoint
   for the delete graph operation.

   Typically these are written to a http.Request.
*/
type DeleteGraphParams struct {

	// RequestBody.
	RequestBody *models.GraphDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGraphParams) WithDefaults() *DeleteGraphParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete graph params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGraphParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete graph params
func (o *DeleteGraphParams) WithTimeout(timeout time.Duration) *DeleteGraphParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete graph params
func (o *DeleteGraphParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete graph params
func (o *DeleteGraphParams) WithContext(ctx context.Context) *DeleteGraphParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete graph params
func (o *DeleteGraphParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete graph params
func (o *DeleteGraphParams) WithHTTPClient(client *http.Client) *DeleteGraphParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete graph params
func (o *DeleteGraphParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete graph params
func (o *DeleteGraphParams) WithRequestBody(requestBody *models.GraphDeletionParams) *DeleteGraphParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete graph params
func (o *DeleteGraphParams) SetRequestBody(requestBody *models.GraphDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGraphParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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