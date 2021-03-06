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

// NewDeleteIscsiLunParams creates a new DeleteIscsiLunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIscsiLunParams() *DeleteIscsiLunParams {
	return &DeleteIscsiLunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIscsiLunParamsWithTimeout creates a new DeleteIscsiLunParams object
// with the ability to set a timeout on a request.
func NewDeleteIscsiLunParamsWithTimeout(timeout time.Duration) *DeleteIscsiLunParams {
	return &DeleteIscsiLunParams{
		timeout: timeout,
	}
}

// NewDeleteIscsiLunParamsWithContext creates a new DeleteIscsiLunParams object
// with the ability to set a context for a request.
func NewDeleteIscsiLunParamsWithContext(ctx context.Context) *DeleteIscsiLunParams {
	return &DeleteIscsiLunParams{
		Context: ctx,
	}
}

// NewDeleteIscsiLunParamsWithHTTPClient creates a new DeleteIscsiLunParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIscsiLunParamsWithHTTPClient(client *http.Client) *DeleteIscsiLunParams {
	return &DeleteIscsiLunParams{
		HTTPClient: client,
	}
}

/* DeleteIscsiLunParams contains all the parameters to send to the API endpoint
   for the delete iscsi lun operation.

   Typically these are written to a http.Request.
*/
type DeleteIscsiLunParams struct {

	// RequestBody.
	RequestBody *models.IscsiLunDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete iscsi lun params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIscsiLunParams) WithDefaults() *DeleteIscsiLunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete iscsi lun params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIscsiLunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete iscsi lun params
func (o *DeleteIscsiLunParams) WithTimeout(timeout time.Duration) *DeleteIscsiLunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iscsi lun params
func (o *DeleteIscsiLunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iscsi lun params
func (o *DeleteIscsiLunParams) WithContext(ctx context.Context) *DeleteIscsiLunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iscsi lun params
func (o *DeleteIscsiLunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iscsi lun params
func (o *DeleteIscsiLunParams) WithHTTPClient(client *http.Client) *DeleteIscsiLunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iscsi lun params
func (o *DeleteIscsiLunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete iscsi lun params
func (o *DeleteIscsiLunParams) WithRequestBody(requestBody *models.IscsiLunDeletionParams) *DeleteIscsiLunParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete iscsi lun params
func (o *DeleteIscsiLunParams) SetRequestBody(requestBody *models.IscsiLunDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIscsiLunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
