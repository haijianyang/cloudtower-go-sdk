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

// RemoveVMCdRomReader is a Reader for the RemoveVMCdRom structure.
type RemoveVMCdRomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveVMCdRomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveVMCdRomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveVMCdRomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveVMCdRomOK creates a RemoveVMCdRomOK with default headers values
func NewRemoveVMCdRomOK() *RemoveVMCdRomOK {
	return &RemoveVMCdRomOK{}
}

/* RemoveVMCdRomOK describes a response with status code 200, with default header values.

Ok
*/
type RemoveVMCdRomOK struct {
	Payload []*models.WithTaskVM
}

func (o *RemoveVMCdRomOK) Error() string {
	return fmt.Sprintf("[POST /remove-vm-cd-rom][%d] removeVmCdRomOK  %+v", 200, o.Payload)
}
func (o *RemoveVMCdRomOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *RemoveVMCdRomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveVMCdRomBadRequest creates a RemoveVMCdRomBadRequest with default headers values
func NewRemoveVMCdRomBadRequest() *RemoveVMCdRomBadRequest {
	return &RemoveVMCdRomBadRequest{}
}

/* RemoveVMCdRomBadRequest describes a response with status code 400, with default header values.

RemoveVMCdRomBadRequest remove Vm cd rom bad request
*/
type RemoveVMCdRomBadRequest struct {
	Payload string
}

func (o *RemoveVMCdRomBadRequest) Error() string {
	return fmt.Sprintf("[POST /remove-vm-cd-rom][%d] removeVmCdRomBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveVMCdRomBadRequest) GetPayload() string {
	return o.Payload
}

func (o *RemoveVMCdRomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
