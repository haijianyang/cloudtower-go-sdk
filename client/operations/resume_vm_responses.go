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

// ResumeVMReader is a Reader for the ResumeVM structure.
type ResumeVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResumeVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResumeVMOK creates a ResumeVMOK with default headers values
func NewResumeVMOK() *ResumeVMOK {
	return &ResumeVMOK{}
}

/* ResumeVMOK describes a response with status code 200, with default header values.

Ok
*/
type ResumeVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *ResumeVMOK) Error() string {
	return fmt.Sprintf("[POST /resume-vm][%d] resumeVmOK  %+v", 200, o.Payload)
}
func (o *ResumeVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *ResumeVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeVMBadRequest creates a ResumeVMBadRequest with default headers values
func NewResumeVMBadRequest() *ResumeVMBadRequest {
	return &ResumeVMBadRequest{}
}

/* ResumeVMBadRequest describes a response with status code 400, with default header values.

ResumeVMBadRequest resume Vm bad request
*/
type ResumeVMBadRequest struct {
	Payload string
}

func (o *ResumeVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /resume-vm][%d] resumeVmBadRequest  %+v", 400, o.Payload)
}
func (o *ResumeVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ResumeVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}