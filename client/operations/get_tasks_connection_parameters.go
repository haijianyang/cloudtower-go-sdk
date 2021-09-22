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

// NewGetTasksConnectionParams creates a new GetTasksConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksConnectionParams() *GetTasksConnectionParams {
	return &GetTasksConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksConnectionParamsWithTimeout creates a new GetTasksConnectionParams object
// with the ability to set a timeout on a request.
func NewGetTasksConnectionParamsWithTimeout(timeout time.Duration) *GetTasksConnectionParams {
	return &GetTasksConnectionParams{
		timeout: timeout,
	}
}

// NewGetTasksConnectionParamsWithContext creates a new GetTasksConnectionParams object
// with the ability to set a context for a request.
func NewGetTasksConnectionParamsWithContext(ctx context.Context) *GetTasksConnectionParams {
	return &GetTasksConnectionParams{
		Context: ctx,
	}
}

// NewGetTasksConnectionParamsWithHTTPClient creates a new GetTasksConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksConnectionParamsWithHTTPClient(client *http.Client) *GetTasksConnectionParams {
	return &GetTasksConnectionParams{
		HTTPClient: client,
	}
}

/* GetTasksConnectionParams contains all the parameters to send to the API endpoint
   for the get tasks connection operation.

   Typically these are written to a http.Request.
*/
type GetTasksConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetTasksConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksConnectionParams) WithDefaults() *GetTasksConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tasks connection params
func (o *GetTasksConnectionParams) WithTimeout(timeout time.Duration) *GetTasksConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks connection params
func (o *GetTasksConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks connection params
func (o *GetTasksConnectionParams) WithContext(ctx context.Context) *GetTasksConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks connection params
func (o *GetTasksConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks connection params
func (o *GetTasksConnectionParams) WithHTTPClient(client *http.Client) *GetTasksConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks connection params
func (o *GetTasksConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get tasks connection params
func (o *GetTasksConnectionParams) WithRequestBody(requestBody *models.GetTasksConnectionRequestBody) *GetTasksConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get tasks connection params
func (o *GetTasksConnectionParams) SetRequestBody(requestBody *models.GetTasksConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
