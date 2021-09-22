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

// GetConsistencyGroupsReader is a Reader for the GetConsistencyGroups structure.
type GetConsistencyGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsistencyGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsistencyGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConsistencyGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConsistencyGroupsOK creates a GetConsistencyGroupsOK with default headers values
func NewGetConsistencyGroupsOK() *GetConsistencyGroupsOK {
	return &GetConsistencyGroupsOK{}
}

/* GetConsistencyGroupsOK describes a response with status code 200, with default header values.

Ok
*/
type GetConsistencyGroupsOK struct {
	Payload []*models.ConsistencyGroup
}

func (o *GetConsistencyGroupsOK) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups][%d] getConsistencyGroupsOK  %+v", 200, o.Payload)
}
func (o *GetConsistencyGroupsOK) GetPayload() []*models.ConsistencyGroup {
	return o.Payload
}

func (o *GetConsistencyGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsistencyGroupsBadRequest creates a GetConsistencyGroupsBadRequest with default headers values
func NewGetConsistencyGroupsBadRequest() *GetConsistencyGroupsBadRequest {
	return &GetConsistencyGroupsBadRequest{}
}

/* GetConsistencyGroupsBadRequest describes a response with status code 400, with default header values.

GetConsistencyGroupsBadRequest get consistency groups bad request
*/
type GetConsistencyGroupsBadRequest struct {
	Payload string
}

func (o *GetConsistencyGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups][%d] getConsistencyGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *GetConsistencyGroupsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetConsistencyGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}