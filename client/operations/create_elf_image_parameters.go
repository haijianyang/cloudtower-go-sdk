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
)

// NewCreateElfImageParams creates a new CreateElfImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateElfImageParams() *CreateElfImageParams {
	return &CreateElfImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateElfImageParamsWithTimeout creates a new CreateElfImageParams object
// with the ability to set a timeout on a request.
func NewCreateElfImageParamsWithTimeout(timeout time.Duration) *CreateElfImageParams {
	return &CreateElfImageParams{
		timeout: timeout,
	}
}

// NewCreateElfImageParamsWithContext creates a new CreateElfImageParams object
// with the ability to set a context for a request.
func NewCreateElfImageParamsWithContext(ctx context.Context) *CreateElfImageParams {
	return &CreateElfImageParams{
		Context: ctx,
	}
}

// NewCreateElfImageParamsWithHTTPClient creates a new CreateElfImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateElfImageParamsWithHTTPClient(client *http.Client) *CreateElfImageParams {
	return &CreateElfImageParams{
		HTTPClient: client,
	}
}

/* CreateElfImageParams contains all the parameters to send to the API endpoint
   for the create elf image operation.

   Typically these are written to a http.Request.
*/
type CreateElfImageParams struct {

	// ClusterID.
	ClusterID string

	// Description.
	Description string

	// File.
	File runtime.NamedReadCloser

	// Name.
	Name string

	// Size.
	Size string

	// UploadTaskID.
	UploadTaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create elf image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateElfImageParams) WithDefaults() *CreateElfImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create elf image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateElfImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create elf image params
func (o *CreateElfImageParams) WithTimeout(timeout time.Duration) *CreateElfImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create elf image params
func (o *CreateElfImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create elf image params
func (o *CreateElfImageParams) WithContext(ctx context.Context) *CreateElfImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create elf image params
func (o *CreateElfImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create elf image params
func (o *CreateElfImageParams) WithHTTPClient(client *http.Client) *CreateElfImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create elf image params
func (o *CreateElfImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the create elf image params
func (o *CreateElfImageParams) WithClusterID(clusterID string) *CreateElfImageParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create elf image params
func (o *CreateElfImageParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDescription adds the description to the create elf image params
func (o *CreateElfImageParams) WithDescription(description string) *CreateElfImageParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the create elf image params
func (o *CreateElfImageParams) SetDescription(description string) {
	o.Description = description
}

// WithFile adds the file to the create elf image params
func (o *CreateElfImageParams) WithFile(file runtime.NamedReadCloser) *CreateElfImageParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the create elf image params
func (o *CreateElfImageParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithName adds the name to the create elf image params
func (o *CreateElfImageParams) WithName(name string) *CreateElfImageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create elf image params
func (o *CreateElfImageParams) SetName(name string) {
	o.Name = name
}

// WithSize adds the size to the create elf image params
func (o *CreateElfImageParams) WithSize(size string) *CreateElfImageParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the create elf image params
func (o *CreateElfImageParams) SetSize(size string) {
	o.Size = size
}

// WithUploadTaskID adds the uploadTaskID to the create elf image params
func (o *CreateElfImageParams) WithUploadTaskID(uploadTaskID string) *CreateElfImageParams {
	o.SetUploadTaskID(uploadTaskID)
	return o
}

// SetUploadTaskID adds the uploadTaskId to the create elf image params
func (o *CreateElfImageParams) SetUploadTaskID(uploadTaskID string) {
	o.UploadTaskID = uploadTaskID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateElfImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param cluster_id
	frClusterID := o.ClusterID
	fClusterID := frClusterID
	if fClusterID != "" {
		if err := r.SetFormParam("cluster_id", fClusterID); err != nil {
			return err
		}
	}

	// form param description
	frDescription := o.Description
	fDescription := frDescription
	if fDescription != "" {
		if err := r.SetFormParam("description", fDescription); err != nil {
			return err
		}
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	// form param size
	frSize := o.Size
	fSize := frSize
	if fSize != "" {
		if err := r.SetFormParam("size", fSize); err != nil {
			return err
		}
	}

	// form param upload_task_id
	frUploadTaskID := o.UploadTaskID
	fUploadTaskID := frUploadTaskID
	if fUploadTaskID != "" {
		if err := r.SetFormParam("upload_task_id", fUploadTaskID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}