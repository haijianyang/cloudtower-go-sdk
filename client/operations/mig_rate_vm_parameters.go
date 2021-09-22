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

// NewMigRateVMParams creates a new MigRateVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMigRateVMParams() *MigRateVMParams {
	return &MigRateVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMigRateVMParamsWithTimeout creates a new MigRateVMParams object
// with the ability to set a timeout on a request.
func NewMigRateVMParamsWithTimeout(timeout time.Duration) *MigRateVMParams {
	return &MigRateVMParams{
		timeout: timeout,
	}
}

// NewMigRateVMParamsWithContext creates a new MigRateVMParams object
// with the ability to set a context for a request.
func NewMigRateVMParamsWithContext(ctx context.Context) *MigRateVMParams {
	return &MigRateVMParams{
		Context: ctx,
	}
}

// NewMigRateVMParamsWithHTTPClient creates a new MigRateVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewMigRateVMParamsWithHTTPClient(client *http.Client) *MigRateVMParams {
	return &MigRateVMParams{
		HTTPClient: client,
	}
}

/* MigRateVMParams contains all the parameters to send to the API endpoint
   for the mig rate Vm operation.

   Typically these are written to a http.Request.
*/
type MigRateVMParams struct {

	// RequestBody.
	RequestBody *models.VMMigrateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mig rate Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigRateVMParams) WithDefaults() *MigRateVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mig rate Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigRateVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mig rate Vm params
func (o *MigRateVMParams) WithTimeout(timeout time.Duration) *MigRateVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mig rate Vm params
func (o *MigRateVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mig rate Vm params
func (o *MigRateVMParams) WithContext(ctx context.Context) *MigRateVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mig rate Vm params
func (o *MigRateVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mig rate Vm params
func (o *MigRateVMParams) WithHTTPClient(client *http.Client) *MigRateVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mig rate Vm params
func (o *MigRateVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the mig rate Vm params
func (o *MigRateVMParams) WithRequestBody(requestBody *models.VMMigrateParams) *MigRateVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the mig rate Vm params
func (o *MigRateVMParams) SetRequestBody(requestBody *models.VMMigrateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *MigRateVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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