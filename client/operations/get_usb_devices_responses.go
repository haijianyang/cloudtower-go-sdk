// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haijianyang/cloudtower-go-sdk/models"
)

// GetUsbDevicesReader is a Reader for the GetUsbDevices structure.
type GetUsbDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsbDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsbDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsbDevicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsbDevicesOK creates a GetUsbDevicesOK with default headers values
func NewGetUsbDevicesOK() *GetUsbDevicesOK {
	return &GetUsbDevicesOK{}
}

/* GetUsbDevicesOK describes a response with status code 200, with default header values.

Ok
*/
type GetUsbDevicesOK struct {
	Payload []*models.UsbDevice
}

func (o *GetUsbDevicesOK) Error() string {
	return fmt.Sprintf("[POST /get-usb-devices][%d] getUsbDevicesOK  %+v", 200, o.Payload)
}
func (o *GetUsbDevicesOK) GetPayload() []*models.UsbDevice {
	return o.Payload
}

func (o *GetUsbDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsbDevicesBadRequest creates a GetUsbDevicesBadRequest with default headers values
func NewGetUsbDevicesBadRequest() *GetUsbDevicesBadRequest {
	return &GetUsbDevicesBadRequest{}
}

/* GetUsbDevicesBadRequest describes a response with status code 400, with default header values.

GetUsbDevicesBadRequest get usb devices bad request
*/
type GetUsbDevicesBadRequest struct {
	Payload string
}

func (o *GetUsbDevicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-usb-devices][%d] getUsbDevicesBadRequest  %+v", 400, o.Payload)
}
func (o *GetUsbDevicesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetUsbDevicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
