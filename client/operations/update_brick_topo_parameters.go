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

// NewUpdateBrickTopoParams creates a new UpdateBrickTopoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBrickTopoParams() *UpdateBrickTopoParams {
	return &UpdateBrickTopoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBrickTopoParamsWithTimeout creates a new UpdateBrickTopoParams object
// with the ability to set a timeout on a request.
func NewUpdateBrickTopoParamsWithTimeout(timeout time.Duration) *UpdateBrickTopoParams {
	return &UpdateBrickTopoParams{
		timeout: timeout,
	}
}

// NewUpdateBrickTopoParamsWithContext creates a new UpdateBrickTopoParams object
// with the ability to set a context for a request.
func NewUpdateBrickTopoParamsWithContext(ctx context.Context) *UpdateBrickTopoParams {
	return &UpdateBrickTopoParams{
		Context: ctx,
	}
}

// NewUpdateBrickTopoParamsWithHTTPClient creates a new UpdateBrickTopoParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBrickTopoParamsWithHTTPClient(client *http.Client) *UpdateBrickTopoParams {
	return &UpdateBrickTopoParams{
		HTTPClient: client,
	}
}

/* UpdateBrickTopoParams contains all the parameters to send to the API endpoint
   for the update brick topo operation.

   Typically these are written to a http.Request.
*/
type UpdateBrickTopoParams struct {

	// RequestBody.
	RequestBody *models.BrickTopoUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update brick topo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBrickTopoParams) WithDefaults() *UpdateBrickTopoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update brick topo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBrickTopoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update brick topo params
func (o *UpdateBrickTopoParams) WithTimeout(timeout time.Duration) *UpdateBrickTopoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update brick topo params
func (o *UpdateBrickTopoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update brick topo params
func (o *UpdateBrickTopoParams) WithContext(ctx context.Context) *UpdateBrickTopoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update brick topo params
func (o *UpdateBrickTopoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update brick topo params
func (o *UpdateBrickTopoParams) WithHTTPClient(client *http.Client) *UpdateBrickTopoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update brick topo params
func (o *UpdateBrickTopoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update brick topo params
func (o *UpdateBrickTopoParams) WithRequestBody(requestBody *models.BrickTopoUpdationParams) *UpdateBrickTopoParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update brick topo params
func (o *UpdateBrickTopoParams) SetRequestBody(requestBody *models.BrickTopoUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBrickTopoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
