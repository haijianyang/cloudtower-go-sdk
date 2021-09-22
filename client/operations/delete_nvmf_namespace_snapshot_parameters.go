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

// NewDeleteNvmfNamespaceSnapshotParams creates a new DeleteNvmfNamespaceSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNvmfNamespaceSnapshotParams() *DeleteNvmfNamespaceSnapshotParams {
	return &DeleteNvmfNamespaceSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNvmfNamespaceSnapshotParamsWithTimeout creates a new DeleteNvmfNamespaceSnapshotParams object
// with the ability to set a timeout on a request.
func NewDeleteNvmfNamespaceSnapshotParamsWithTimeout(timeout time.Duration) *DeleteNvmfNamespaceSnapshotParams {
	return &DeleteNvmfNamespaceSnapshotParams{
		timeout: timeout,
	}
}

// NewDeleteNvmfNamespaceSnapshotParamsWithContext creates a new DeleteNvmfNamespaceSnapshotParams object
// with the ability to set a context for a request.
func NewDeleteNvmfNamespaceSnapshotParamsWithContext(ctx context.Context) *DeleteNvmfNamespaceSnapshotParams {
	return &DeleteNvmfNamespaceSnapshotParams{
		Context: ctx,
	}
}

// NewDeleteNvmfNamespaceSnapshotParamsWithHTTPClient creates a new DeleteNvmfNamespaceSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNvmfNamespaceSnapshotParamsWithHTTPClient(client *http.Client) *DeleteNvmfNamespaceSnapshotParams {
	return &DeleteNvmfNamespaceSnapshotParams{
		HTTPClient: client,
	}
}

/* DeleteNvmfNamespaceSnapshotParams contains all the parameters to send to the API endpoint
   for the delete nvmf namespace snapshot operation.

   Typically these are written to a http.Request.
*/
type DeleteNvmfNamespaceSnapshotParams struct {

	// RequestBody.
	RequestBody *models.NvmfNamespaceSnapshotDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete nvmf namespace snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNvmfNamespaceSnapshotParams) WithDefaults() *DeleteNvmfNamespaceSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete nvmf namespace snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNvmfNamespaceSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) WithTimeout(timeout time.Duration) *DeleteNvmfNamespaceSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) WithContext(ctx context.Context) *DeleteNvmfNamespaceSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) WithHTTPClient(client *http.Client) *DeleteNvmfNamespaceSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) WithRequestBody(requestBody *models.NvmfNamespaceSnapshotDeletionParams) *DeleteNvmfNamespaceSnapshotParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete nvmf namespace snapshot params
func (o *DeleteNvmfNamespaceSnapshotParams) SetRequestBody(requestBody *models.NvmfNamespaceSnapshotDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNvmfNamespaceSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
