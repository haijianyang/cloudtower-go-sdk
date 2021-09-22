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

// NewGetSvtImagesConnectionParams creates a new GetSvtImagesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSvtImagesConnectionParams() *GetSvtImagesConnectionParams {
	return &GetSvtImagesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSvtImagesConnectionParamsWithTimeout creates a new GetSvtImagesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetSvtImagesConnectionParamsWithTimeout(timeout time.Duration) *GetSvtImagesConnectionParams {
	return &GetSvtImagesConnectionParams{
		timeout: timeout,
	}
}

// NewGetSvtImagesConnectionParamsWithContext creates a new GetSvtImagesConnectionParams object
// with the ability to set a context for a request.
func NewGetSvtImagesConnectionParamsWithContext(ctx context.Context) *GetSvtImagesConnectionParams {
	return &GetSvtImagesConnectionParams{
		Context: ctx,
	}
}

// NewGetSvtImagesConnectionParamsWithHTTPClient creates a new GetSvtImagesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSvtImagesConnectionParamsWithHTTPClient(client *http.Client) *GetSvtImagesConnectionParams {
	return &GetSvtImagesConnectionParams{
		HTTPClient: client,
	}
}

/* GetSvtImagesConnectionParams contains all the parameters to send to the API endpoint
   for the get svt images connection operation.

   Typically these are written to a http.Request.
*/
type GetSvtImagesConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetSvtImagesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get svt images connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSvtImagesConnectionParams) WithDefaults() *GetSvtImagesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get svt images connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSvtImagesConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get svt images connection params
func (o *GetSvtImagesConnectionParams) WithTimeout(timeout time.Duration) *GetSvtImagesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get svt images connection params
func (o *GetSvtImagesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get svt images connection params
func (o *GetSvtImagesConnectionParams) WithContext(ctx context.Context) *GetSvtImagesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get svt images connection params
func (o *GetSvtImagesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get svt images connection params
func (o *GetSvtImagesConnectionParams) WithHTTPClient(client *http.Client) *GetSvtImagesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get svt images connection params
func (o *GetSvtImagesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get svt images connection params
func (o *GetSvtImagesConnectionParams) WithRequestBody(requestBody *models.GetSvtImagesConnectionRequestBody) *GetSvtImagesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get svt images connection params
func (o *GetSvtImagesConnectionParams) SetRequestBody(requestBody *models.GetSvtImagesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetSvtImagesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
