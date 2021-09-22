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

// DeleteVMVolumeFromVMReader is a Reader for the DeleteVMVolumeFromVM structure.
type DeleteVMVolumeFromVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMVolumeFromVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMVolumeFromVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMVolumeFromVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMVolumeFromVMOK creates a DeleteVMVolumeFromVMOK with default headers values
func NewDeleteVMVolumeFromVMOK() *DeleteVMVolumeFromVMOK {
	return &DeleteVMVolumeFromVMOK{}
}

/* DeleteVMVolumeFromVMOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVMVolumeFromVMOK struct {
	Payload []*models.WithTaskDeleteVMVolume
}

func (o *DeleteVMVolumeFromVMOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume][%d] deleteVmVolumeFromVmOK  %+v", 200, o.Payload)
}
func (o *DeleteVMVolumeFromVMOK) GetPayload() []*models.WithTaskDeleteVMVolume {
	return o.Payload
}

func (o *DeleteVMVolumeFromVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMVolumeFromVMBadRequest creates a DeleteVMVolumeFromVMBadRequest with default headers values
func NewDeleteVMVolumeFromVMBadRequest() *DeleteVMVolumeFromVMBadRequest {
	return &DeleteVMVolumeFromVMBadRequest{}
}

/* DeleteVMVolumeFromVMBadRequest describes a response with status code 400, with default header values.

DeleteVMVolumeFromVMBadRequest delete Vm volume from Vm bad request
*/
type DeleteVMVolumeFromVMBadRequest struct {
	Payload string
}

func (o *DeleteVMVolumeFromVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume][%d] deleteVmVolumeFromVmBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMVolumeFromVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteVMVolumeFromVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}