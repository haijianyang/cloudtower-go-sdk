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

// NewUpdateUsbDeviceParams creates a new UpdateUsbDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUsbDeviceParams() *UpdateUsbDeviceParams {
	return &UpdateUsbDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUsbDeviceParamsWithTimeout creates a new UpdateUsbDeviceParams object
// with the ability to set a timeout on a request.
func NewUpdateUsbDeviceParamsWithTimeout(timeout time.Duration) *UpdateUsbDeviceParams {
	return &UpdateUsbDeviceParams{
		timeout: timeout,
	}
}

// NewUpdateUsbDeviceParamsWithContext creates a new UpdateUsbDeviceParams object
// with the ability to set a context for a request.
func NewUpdateUsbDeviceParamsWithContext(ctx context.Context) *UpdateUsbDeviceParams {
	return &UpdateUsbDeviceParams{
		Context: ctx,
	}
}

// NewUpdateUsbDeviceParamsWithHTTPClient creates a new UpdateUsbDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUsbDeviceParamsWithHTTPClient(client *http.Client) *UpdateUsbDeviceParams {
	return &UpdateUsbDeviceParams{
		HTTPClient: client,
	}
}

/* UpdateUsbDeviceParams contains all the parameters to send to the API endpoint
   for the update usb device operation.

   Typically these are written to a http.Request.
*/
type UpdateUsbDeviceParams struct {

	// RequestBody.
	RequestBody *models.UsbDeviceMountParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update usb device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsbDeviceParams) WithDefaults() *UpdateUsbDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update usb device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsbDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update usb device params
func (o *UpdateUsbDeviceParams) WithTimeout(timeout time.Duration) *UpdateUsbDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update usb device params
func (o *UpdateUsbDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update usb device params
func (o *UpdateUsbDeviceParams) WithContext(ctx context.Context) *UpdateUsbDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update usb device params
func (o *UpdateUsbDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update usb device params
func (o *UpdateUsbDeviceParams) WithHTTPClient(client *http.Client) *UpdateUsbDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update usb device params
func (o *UpdateUsbDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update usb device params
func (o *UpdateUsbDeviceParams) WithRequestBody(requestBody *models.UsbDeviceMountParams) *UpdateUsbDeviceParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update usb device params
func (o *UpdateUsbDeviceParams) SetRequestBody(requestBody *models.UsbDeviceMountParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUsbDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
