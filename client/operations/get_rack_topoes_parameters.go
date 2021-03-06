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

// NewGetRackTopoesParams creates a new GetRackTopoesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRackTopoesParams() *GetRackTopoesParams {
	return &GetRackTopoesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRackTopoesParamsWithTimeout creates a new GetRackTopoesParams object
// with the ability to set a timeout on a request.
func NewGetRackTopoesParamsWithTimeout(timeout time.Duration) *GetRackTopoesParams {
	return &GetRackTopoesParams{
		timeout: timeout,
	}
}

// NewGetRackTopoesParamsWithContext creates a new GetRackTopoesParams object
// with the ability to set a context for a request.
func NewGetRackTopoesParamsWithContext(ctx context.Context) *GetRackTopoesParams {
	return &GetRackTopoesParams{
		Context: ctx,
	}
}

// NewGetRackTopoesParamsWithHTTPClient creates a new GetRackTopoesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRackTopoesParamsWithHTTPClient(client *http.Client) *GetRackTopoesParams {
	return &GetRackTopoesParams{
		HTTPClient: client,
	}
}

/* GetRackTopoesParams contains all the parameters to send to the API endpoint
   for the get rack topoes operation.

   Typically these are written to a http.Request.
*/
type GetRackTopoesParams struct {

	// RequestBody.
	RequestBody *models.GetRackTopoesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get rack topoes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRackTopoesParams) WithDefaults() *GetRackTopoesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get rack topoes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRackTopoesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get rack topoes params
func (o *GetRackTopoesParams) WithTimeout(timeout time.Duration) *GetRackTopoesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rack topoes params
func (o *GetRackTopoesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rack topoes params
func (o *GetRackTopoesParams) WithContext(ctx context.Context) *GetRackTopoesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rack topoes params
func (o *GetRackTopoesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rack topoes params
func (o *GetRackTopoesParams) WithHTTPClient(client *http.Client) *GetRackTopoesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rack topoes params
func (o *GetRackTopoesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get rack topoes params
func (o *GetRackTopoesParams) WithRequestBody(requestBody *models.GetRackTopoesRequestBody) *GetRackTopoesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get rack topoes params
func (o *GetRackTopoesParams) SetRequestBody(requestBody *models.GetRackTopoesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetRackTopoesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
