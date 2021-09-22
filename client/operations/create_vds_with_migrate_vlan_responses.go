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

// CreateVdsWithMigrateVlanReader is a Reader for the CreateVdsWithMigrateVlan structure.
type CreateVdsWithMigrateVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVdsWithMigrateVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVdsWithMigrateVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVdsWithMigrateVlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVdsWithMigrateVlanOK creates a CreateVdsWithMigrateVlanOK with default headers values
func NewCreateVdsWithMigrateVlanOK() *CreateVdsWithMigrateVlanOK {
	return &CreateVdsWithMigrateVlanOK{}
}

/* CreateVdsWithMigrateVlanOK describes a response with status code 200, with default header values.

Ok
*/
type CreateVdsWithMigrateVlanOK struct {
	Payload []*models.WithTaskVds
}

func (o *CreateVdsWithMigrateVlanOK) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-migrate-vlan][%d] createVdsWithMigrateVlanOK  %+v", 200, o.Payload)
}
func (o *CreateVdsWithMigrateVlanOK) GetPayload() []*models.WithTaskVds {
	return o.Payload
}

func (o *CreateVdsWithMigrateVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVdsWithMigrateVlanBadRequest creates a CreateVdsWithMigrateVlanBadRequest with default headers values
func NewCreateVdsWithMigrateVlanBadRequest() *CreateVdsWithMigrateVlanBadRequest {
	return &CreateVdsWithMigrateVlanBadRequest{}
}

/* CreateVdsWithMigrateVlanBadRequest describes a response with status code 400, with default header values.

CreateVdsWithMigrateVlanBadRequest create vds with migrate vlan bad request
*/
type CreateVdsWithMigrateVlanBadRequest struct {
	Payload string
}

func (o *CreateVdsWithMigrateVlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-migrate-vlan][%d] createVdsWithMigrateVlanBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVdsWithMigrateVlanBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateVdsWithMigrateVlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
