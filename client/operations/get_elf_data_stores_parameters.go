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

// NewGetElfDataStoresParams creates a new GetElfDataStoresParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetElfDataStoresParams() *GetElfDataStoresParams {
	return &GetElfDataStoresParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetElfDataStoresParamsWithTimeout creates a new GetElfDataStoresParams object
// with the ability to set a timeout on a request.
func NewGetElfDataStoresParamsWithTimeout(timeout time.Duration) *GetElfDataStoresParams {
	return &GetElfDataStoresParams{
		timeout: timeout,
	}
}

// NewGetElfDataStoresParamsWithContext creates a new GetElfDataStoresParams object
// with the ability to set a context for a request.
func NewGetElfDataStoresParamsWithContext(ctx context.Context) *GetElfDataStoresParams {
	return &GetElfDataStoresParams{
		Context: ctx,
	}
}

// NewGetElfDataStoresParamsWithHTTPClient creates a new GetElfDataStoresParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetElfDataStoresParamsWithHTTPClient(client *http.Client) *GetElfDataStoresParams {
	return &GetElfDataStoresParams{
		HTTPClient: client,
	}
}

/* GetElfDataStoresParams contains all the parameters to send to the API endpoint
   for the get elf data stores operation.

   Typically these are written to a http.Request.
*/
type GetElfDataStoresParams struct {

	// RequestBody.
	RequestBody *models.GetElfDataStoresRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get elf data stores params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetElfDataStoresParams) WithDefaults() *GetElfDataStoresParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get elf data stores params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetElfDataStoresParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get elf data stores params
func (o *GetElfDataStoresParams) WithTimeout(timeout time.Duration) *GetElfDataStoresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get elf data stores params
func (o *GetElfDataStoresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get elf data stores params
func (o *GetElfDataStoresParams) WithContext(ctx context.Context) *GetElfDataStoresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get elf data stores params
func (o *GetElfDataStoresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get elf data stores params
func (o *GetElfDataStoresParams) WithHTTPClient(client *http.Client) *GetElfDataStoresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get elf data stores params
func (o *GetElfDataStoresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get elf data stores params
func (o *GetElfDataStoresParams) WithRequestBody(requestBody *models.GetElfDataStoresRequestBody) *GetElfDataStoresParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get elf data stores params
func (o *GetElfDataStoresParams) SetRequestBody(requestBody *models.GetElfDataStoresRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetElfDataStoresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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