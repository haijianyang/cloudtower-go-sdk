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

// AddVMDiskReader is a Reader for the AddVMDisk structure.
type AddVMDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddVMDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddVMDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddVMDiskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddVMDiskOK creates a AddVMDiskOK with default headers values
func NewAddVMDiskOK() *AddVMDiskOK {
	return &AddVMDiskOK{}
}

/* AddVMDiskOK describes a response with status code 200, with default header values.

Ok
*/
type AddVMDiskOK struct {
	Payload []*models.WithTaskVM
}

func (o *AddVMDiskOK) Error() string {
	return fmt.Sprintf("[POST /add-vm-disk][%d] addVmDiskOK  %+v", 200, o.Payload)
}
func (o *AddVMDiskOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *AddVMDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddVMDiskBadRequest creates a AddVMDiskBadRequest with default headers values
func NewAddVMDiskBadRequest() *AddVMDiskBadRequest {
	return &AddVMDiskBadRequest{}
}

/* AddVMDiskBadRequest describes a response with status code 400, with default header values.

AddVMDiskBadRequest add Vm disk bad request
*/
type AddVMDiskBadRequest struct {
	Payload string
}

func (o *AddVMDiskBadRequest) Error() string {
	return fmt.Sprintf("[POST /add-vm-disk][%d] addVmDiskBadRequest  %+v", 400, o.Payload)
}
func (o *AddVMDiskBadRequest) GetPayload() string {
	return o.Payload
}

func (o *AddVMDiskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
