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

// NewGetUploadTasksParams creates a new GetUploadTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUploadTasksParams() *GetUploadTasksParams {
	return &GetUploadTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUploadTasksParamsWithTimeout creates a new GetUploadTasksParams object
// with the ability to set a timeout on a request.
func NewGetUploadTasksParamsWithTimeout(timeout time.Duration) *GetUploadTasksParams {
	return &GetUploadTasksParams{
		timeout: timeout,
	}
}

// NewGetUploadTasksParamsWithContext creates a new GetUploadTasksParams object
// with the ability to set a context for a request.
func NewGetUploadTasksParamsWithContext(ctx context.Context) *GetUploadTasksParams {
	return &GetUploadTasksParams{
		Context: ctx,
	}
}

// NewGetUploadTasksParamsWithHTTPClient creates a new GetUploadTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUploadTasksParamsWithHTTPClient(client *http.Client) *GetUploadTasksParams {
	return &GetUploadTasksParams{
		HTTPClient: client,
	}
}

/* GetUploadTasksParams contains all the parameters to send to the API endpoint
   for the get upload tasks operation.

   Typically these are written to a http.Request.
*/
type GetUploadTasksParams struct {

	// RequestBody.
	RequestBody *models.GetUploadTasksRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upload tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUploadTasksParams) WithDefaults() *GetUploadTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upload tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUploadTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upload tasks params
func (o *GetUploadTasksParams) WithTimeout(timeout time.Duration) *GetUploadTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upload tasks params
func (o *GetUploadTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upload tasks params
func (o *GetUploadTasksParams) WithContext(ctx context.Context) *GetUploadTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upload tasks params
func (o *GetUploadTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upload tasks params
func (o *GetUploadTasksParams) WithHTTPClient(client *http.Client) *GetUploadTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upload tasks params
func (o *GetUploadTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get upload tasks params
func (o *GetUploadTasksParams) WithRequestBody(requestBody *models.GetUploadTasksRequestBody) *GetUploadTasksParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get upload tasks params
func (o *GetUploadTasksParams) SetRequestBody(requestBody *models.GetUploadTasksRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUploadTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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