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

// NewGetVMPlacementGroupsConnectionParams creates a new GetVMPlacementGroupsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMPlacementGroupsConnectionParams() *GetVMPlacementGroupsConnectionParams {
	return &GetVMPlacementGroupsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMPlacementGroupsConnectionParamsWithTimeout creates a new GetVMPlacementGroupsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVMPlacementGroupsConnectionParamsWithTimeout(timeout time.Duration) *GetVMPlacementGroupsConnectionParams {
	return &GetVMPlacementGroupsConnectionParams{
		timeout: timeout,
	}
}

// NewGetVMPlacementGroupsConnectionParamsWithContext creates a new GetVMPlacementGroupsConnectionParams object
// with the ability to set a context for a request.
func NewGetVMPlacementGroupsConnectionParamsWithContext(ctx context.Context) *GetVMPlacementGroupsConnectionParams {
	return &GetVMPlacementGroupsConnectionParams{
		Context: ctx,
	}
}

// NewGetVMPlacementGroupsConnectionParamsWithHTTPClient creates a new GetVMPlacementGroupsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMPlacementGroupsConnectionParamsWithHTTPClient(client *http.Client) *GetVMPlacementGroupsConnectionParams {
	return &GetVMPlacementGroupsConnectionParams{
		HTTPClient: client,
	}
}

/* GetVMPlacementGroupsConnectionParams contains all the parameters to send to the API endpoint
   for the get Vm placement groups connection operation.

   Typically these are written to a http.Request.
*/
type GetVMPlacementGroupsConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetVMPlacementGroupsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm placement groups connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMPlacementGroupsConnectionParams) WithDefaults() *GetVMPlacementGroupsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm placement groups connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMPlacementGroupsConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) WithTimeout(timeout time.Duration) *GetVMPlacementGroupsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) WithContext(ctx context.Context) *GetVMPlacementGroupsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) WithHTTPClient(client *http.Client) *GetVMPlacementGroupsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) WithRequestBody(requestBody *models.GetVMPlacementGroupsConnectionRequestBody) *GetVMPlacementGroupsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm placement groups connection params
func (o *GetVMPlacementGroupsConnectionParams) SetRequestBody(requestBody *models.GetVMPlacementGroupsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMPlacementGroupsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
