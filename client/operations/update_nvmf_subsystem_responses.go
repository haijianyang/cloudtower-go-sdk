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

// UpdateNvmfSubsystemReader is a Reader for the UpdateNvmfSubsystem structure.
type UpdateNvmfSubsystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNvmfSubsystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNvmfSubsystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNvmfSubsystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNvmfSubsystemOK creates a UpdateNvmfSubsystemOK with default headers values
func NewUpdateNvmfSubsystemOK() *UpdateNvmfSubsystemOK {
	return &UpdateNvmfSubsystemOK{}
}

/* UpdateNvmfSubsystemOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateNvmfSubsystemOK struct {
	Payload []*models.WithTaskNvmfSubsystem
}

func (o *UpdateNvmfSubsystemOK) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-subsystem][%d] updateNvmfSubsystemOK  %+v", 200, o.Payload)
}
func (o *UpdateNvmfSubsystemOK) GetPayload() []*models.WithTaskNvmfSubsystem {
	return o.Payload
}

func (o *UpdateNvmfSubsystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNvmfSubsystemBadRequest creates a UpdateNvmfSubsystemBadRequest with default headers values
func NewUpdateNvmfSubsystemBadRequest() *UpdateNvmfSubsystemBadRequest {
	return &UpdateNvmfSubsystemBadRequest{}
}

/* UpdateNvmfSubsystemBadRequest describes a response with status code 400, with default header values.

UpdateNvmfSubsystemBadRequest update nvmf subsystem bad request
*/
type UpdateNvmfSubsystemBadRequest struct {
	Payload string
}

func (o *UpdateNvmfSubsystemBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-subsystem][%d] updateNvmfSubsystemBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateNvmfSubsystemBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateNvmfSubsystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}