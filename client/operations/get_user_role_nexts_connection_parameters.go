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

// NewGetUserRoleNextsConnectionParams creates a new GetUserRoleNextsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserRoleNextsConnectionParams() *GetUserRoleNextsConnectionParams {
	return &GetUserRoleNextsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRoleNextsConnectionParamsWithTimeout creates a new GetUserRoleNextsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetUserRoleNextsConnectionParamsWithTimeout(timeout time.Duration) *GetUserRoleNextsConnectionParams {
	return &GetUserRoleNextsConnectionParams{
		timeout: timeout,
	}
}

// NewGetUserRoleNextsConnectionParamsWithContext creates a new GetUserRoleNextsConnectionParams object
// with the ability to set a context for a request.
func NewGetUserRoleNextsConnectionParamsWithContext(ctx context.Context) *GetUserRoleNextsConnectionParams {
	return &GetUserRoleNextsConnectionParams{
		Context: ctx,
	}
}

// NewGetUserRoleNextsConnectionParamsWithHTTPClient creates a new GetUserRoleNextsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserRoleNextsConnectionParamsWithHTTPClient(client *http.Client) *GetUserRoleNextsConnectionParams {
	return &GetUserRoleNextsConnectionParams{
		HTTPClient: client,
	}
}

/* GetUserRoleNextsConnectionParams contains all the parameters to send to the API endpoint
   for the get user role nexts connection operation.

   Typically these are written to a http.Request.
*/
type GetUserRoleNextsConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetUserRoleNextsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user role nexts connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserRoleNextsConnectionParams) WithDefaults() *GetUserRoleNextsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user role nexts connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserRoleNextsConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) WithTimeout(timeout time.Duration) *GetUserRoleNextsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) WithContext(ctx context.Context) *GetUserRoleNextsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) WithHTTPClient(client *http.Client) *GetUserRoleNextsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) WithRequestBody(requestBody *models.GetUserRoleNextsConnectionRequestBody) *GetUserRoleNextsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get user role nexts connection params
func (o *GetUserRoleNextsConnectionParams) SetRequestBody(requestBody *models.GetUserRoleNextsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRoleNextsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
