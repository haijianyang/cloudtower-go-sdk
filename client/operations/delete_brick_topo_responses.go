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

// DeleteBrickTopoReader is a Reader for the DeleteBrickTopo structure.
type DeleteBrickTopoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBrickTopoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBrickTopoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteBrickTopoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteBrickTopoOK creates a DeleteBrickTopoOK with default headers values
func NewDeleteBrickTopoOK() *DeleteBrickTopoOK {
	return &DeleteBrickTopoOK{}
}

/* DeleteBrickTopoOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteBrickTopoOK struct {
	Payload []*models.WithTaskDeleteBrickTopo
}

func (o *DeleteBrickTopoOK) Error() string {
	return fmt.Sprintf("[POST /delete-brick-topo][%d] deleteBrickTopoOK  %+v", 200, o.Payload)
}
func (o *DeleteBrickTopoOK) GetPayload() []*models.WithTaskDeleteBrickTopo {
	return o.Payload
}

func (o *DeleteBrickTopoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBrickTopoBadRequest creates a DeleteBrickTopoBadRequest with default headers values
func NewDeleteBrickTopoBadRequest() *DeleteBrickTopoBadRequest {
	return &DeleteBrickTopoBadRequest{}
}

/* DeleteBrickTopoBadRequest describes a response with status code 400, with default header values.

DeleteBrickTopoBadRequest delete brick topo bad request
*/
type DeleteBrickTopoBadRequest struct {
	Payload string
}

func (o *DeleteBrickTopoBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-brick-topo][%d] deleteBrickTopoBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteBrickTopoBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteBrickTopoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
