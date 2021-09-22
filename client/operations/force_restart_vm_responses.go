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

// ForceRestartVMReader is a Reader for the ForceRestartVM structure.
type ForceRestartVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForceRestartVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForceRestartVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewForceRestartVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewForceRestartVMOK creates a ForceRestartVMOK with default headers values
func NewForceRestartVMOK() *ForceRestartVMOK {
	return &ForceRestartVMOK{}
}

/* ForceRestartVMOK describes a response with status code 200, with default header values.

Ok
*/
type ForceRestartVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *ForceRestartVMOK) Error() string {
	return fmt.Sprintf("[POST /force-restart-vm][%d] forceRestartVmOK  %+v", 200, o.Payload)
}
func (o *ForceRestartVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *ForceRestartVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForceRestartVMBadRequest creates a ForceRestartVMBadRequest with default headers values
func NewForceRestartVMBadRequest() *ForceRestartVMBadRequest {
	return &ForceRestartVMBadRequest{}
}

/* ForceRestartVMBadRequest describes a response with status code 400, with default header values.

ForceRestartVMBadRequest force restart Vm bad request
*/
type ForceRestartVMBadRequest struct {
	Payload string
}

func (o *ForceRestartVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /force-restart-vm][%d] forceRestartVmBadRequest  %+v", 400, o.Payload)
}
func (o *ForceRestartVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ForceRestartVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
