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

// CreateBrickTopoReader is a Reader for the CreateBrickTopo structure.
type CreateBrickTopoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBrickTopoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBrickTopoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBrickTopoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBrickTopoOK creates a CreateBrickTopoOK with default headers values
func NewCreateBrickTopoOK() *CreateBrickTopoOK {
	return &CreateBrickTopoOK{}
}

/* CreateBrickTopoOK describes a response with status code 200, with default header values.

Ok
*/
type CreateBrickTopoOK struct {
	Payload []*models.WithTaskBrickTopo
}

func (o *CreateBrickTopoOK) Error() string {
	return fmt.Sprintf("[POST /create-brick-topo][%d] createBrickTopoOK  %+v", 200, o.Payload)
}
func (o *CreateBrickTopoOK) GetPayload() []*models.WithTaskBrickTopo {
	return o.Payload
}

func (o *CreateBrickTopoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBrickTopoBadRequest creates a CreateBrickTopoBadRequest with default headers values
func NewCreateBrickTopoBadRequest() *CreateBrickTopoBadRequest {
	return &CreateBrickTopoBadRequest{}
}

/* CreateBrickTopoBadRequest describes a response with status code 400, with default header values.

CreateBrickTopoBadRequest create brick topo bad request
*/
type CreateBrickTopoBadRequest struct {
	Payload string
}

func (o *CreateBrickTopoBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-brick-topo][%d] createBrickTopoBadRequest  %+v", 400, o.Payload)
}
func (o *CreateBrickTopoBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateBrickTopoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
