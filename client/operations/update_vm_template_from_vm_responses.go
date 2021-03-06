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

// UpdateVMTemplateFromVMReader is a Reader for the UpdateVMTemplateFromVM structure.
type UpdateVMTemplateFromVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMTemplateFromVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMTemplateFromVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVMTemplateFromVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMTemplateFromVMOK creates a UpdateVMTemplateFromVMOK with default headers values
func NewUpdateVMTemplateFromVMOK() *UpdateVMTemplateFromVMOK {
	return &UpdateVMTemplateFromVMOK{}
}

/* UpdateVMTemplateFromVMOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateVMTemplateFromVMOK struct {
	Payload []*models.WithTaskVMTemplate
}

func (o *UpdateVMTemplateFromVMOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-template][%d] updateVmTemplateFromVmOK  %+v", 200, o.Payload)
}
func (o *UpdateVMTemplateFromVMOK) GetPayload() []*models.WithTaskVMTemplate {
	return o.Payload
}

func (o *UpdateVMTemplateFromVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMTemplateFromVMBadRequest creates a UpdateVMTemplateFromVMBadRequest with default headers values
func NewUpdateVMTemplateFromVMBadRequest() *UpdateVMTemplateFromVMBadRequest {
	return &UpdateVMTemplateFromVMBadRequest{}
}

/* UpdateVMTemplateFromVMBadRequest describes a response with status code 400, with default header values.

UpdateVMTemplateFromVMBadRequest update Vm template from Vm bad request
*/
type UpdateVMTemplateFromVMBadRequest struct {
	Payload string
}

func (o *UpdateVMTemplateFromVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-template][%d] updateVmTemplateFromVmBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMTemplateFromVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateVMTemplateFromVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
