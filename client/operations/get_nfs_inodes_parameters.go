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

// NewGetNfsInodesParams creates a new GetNfsInodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNfsInodesParams() *GetNfsInodesParams {
	return &GetNfsInodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsInodesParamsWithTimeout creates a new GetNfsInodesParams object
// with the ability to set a timeout on a request.
func NewGetNfsInodesParamsWithTimeout(timeout time.Duration) *GetNfsInodesParams {
	return &GetNfsInodesParams{
		timeout: timeout,
	}
}

// NewGetNfsInodesParamsWithContext creates a new GetNfsInodesParams object
// with the ability to set a context for a request.
func NewGetNfsInodesParamsWithContext(ctx context.Context) *GetNfsInodesParams {
	return &GetNfsInodesParams{
		Context: ctx,
	}
}

// NewGetNfsInodesParamsWithHTTPClient creates a new GetNfsInodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNfsInodesParamsWithHTTPClient(client *http.Client) *GetNfsInodesParams {
	return &GetNfsInodesParams{
		HTTPClient: client,
	}
}

/* GetNfsInodesParams contains all the parameters to send to the API endpoint
   for the get nfs inodes operation.

   Typically these are written to a http.Request.
*/
type GetNfsInodesParams struct {

	// RequestBody.
	RequestBody *models.GetNfsInodesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nfs inodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsInodesParams) WithDefaults() *GetNfsInodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nfs inodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsInodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nfs inodes params
func (o *GetNfsInodesParams) WithTimeout(timeout time.Duration) *GetNfsInodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs inodes params
func (o *GetNfsInodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs inodes params
func (o *GetNfsInodesParams) WithContext(ctx context.Context) *GetNfsInodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs inodes params
func (o *GetNfsInodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs inodes params
func (o *GetNfsInodesParams) WithHTTPClient(client *http.Client) *GetNfsInodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs inodes params
func (o *GetNfsInodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nfs inodes params
func (o *GetNfsInodesParams) WithRequestBody(requestBody *models.GetNfsInodesRequestBody) *GetNfsInodesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nfs inodes params
func (o *GetNfsInodesParams) SetRequestBody(requestBody *models.GetNfsInodesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsInodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
