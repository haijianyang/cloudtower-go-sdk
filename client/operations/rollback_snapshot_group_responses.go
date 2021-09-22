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

// RollbackSnapshotGroupReader is a Reader for the RollbackSnapshotGroup structure.
type RollbackSnapshotGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RollbackSnapshotGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRollbackSnapshotGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRollbackSnapshotGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRollbackSnapshotGroupOK creates a RollbackSnapshotGroupOK with default headers values
func NewRollbackSnapshotGroupOK() *RollbackSnapshotGroupOK {
	return &RollbackSnapshotGroupOK{}
}

/* RollbackSnapshotGroupOK describes a response with status code 200, with default header values.

Ok
*/
type RollbackSnapshotGroupOK struct {
	Payload []*models.WithTaskSnapshotGroup
}

func (o *RollbackSnapshotGroupOK) Error() string {
	return fmt.Sprintf("[POST /rollback-snapshot-group][%d] rollbackSnapshotGroupOK  %+v", 200, o.Payload)
}
func (o *RollbackSnapshotGroupOK) GetPayload() []*models.WithTaskSnapshotGroup {
	return o.Payload
}

func (o *RollbackSnapshotGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRollbackSnapshotGroupBadRequest creates a RollbackSnapshotGroupBadRequest with default headers values
func NewRollbackSnapshotGroupBadRequest() *RollbackSnapshotGroupBadRequest {
	return &RollbackSnapshotGroupBadRequest{}
}

/* RollbackSnapshotGroupBadRequest describes a response with status code 400, with default header values.

RollbackSnapshotGroupBadRequest rollback snapshot group bad request
*/
type RollbackSnapshotGroupBadRequest struct {
	Payload string
}

func (o *RollbackSnapshotGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /rollback-snapshot-group][%d] rollbackSnapshotGroupBadRequest  %+v", 400, o.Payload)
}
func (o *RollbackSnapshotGroupBadRequest) GetPayload() string {
	return o.Payload
}

func (o *RollbackSnapshotGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
