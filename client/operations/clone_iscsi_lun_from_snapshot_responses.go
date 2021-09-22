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

// CloneIscsiLunFromSnapshotReader is a Reader for the CloneIscsiLunFromSnapshot structure.
type CloneIscsiLunFromSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneIscsiLunFromSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneIscsiLunFromSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloneIscsiLunFromSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneIscsiLunFromSnapshotOK creates a CloneIscsiLunFromSnapshotOK with default headers values
func NewCloneIscsiLunFromSnapshotOK() *CloneIscsiLunFromSnapshotOK {
	return &CloneIscsiLunFromSnapshotOK{}
}

/* CloneIscsiLunFromSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type CloneIscsiLunFromSnapshotOK struct {
	Payload []*models.WithTaskIscsiLun
}

func (o *CloneIscsiLunFromSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /clone-iscsi-lun-from-snapshot][%d] cloneIscsiLunFromSnapshotOK  %+v", 200, o.Payload)
}
func (o *CloneIscsiLunFromSnapshotOK) GetPayload() []*models.WithTaskIscsiLun {
	return o.Payload
}

func (o *CloneIscsiLunFromSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneIscsiLunFromSnapshotBadRequest creates a CloneIscsiLunFromSnapshotBadRequest with default headers values
func NewCloneIscsiLunFromSnapshotBadRequest() *CloneIscsiLunFromSnapshotBadRequest {
	return &CloneIscsiLunFromSnapshotBadRequest{}
}

/* CloneIscsiLunFromSnapshotBadRequest describes a response with status code 400, with default header values.

CloneIscsiLunFromSnapshotBadRequest clone iscsi lun from snapshot bad request
*/
type CloneIscsiLunFromSnapshotBadRequest struct {
	Payload string
}

func (o *CloneIscsiLunFromSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /clone-iscsi-lun-from-snapshot][%d] cloneIscsiLunFromSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CloneIscsiLunFromSnapshotBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CloneIscsiLunFromSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
