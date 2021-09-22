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

// GetNfsExportsConnectionReader is a Reader for the GetNfsExportsConnection structure.
type GetNfsExportsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsExportsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNfsExportsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNfsExportsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNfsExportsConnectionOK creates a GetNfsExportsConnectionOK with default headers values
func NewGetNfsExportsConnectionOK() *GetNfsExportsConnectionOK {
	return &GetNfsExportsConnectionOK{}
}

/* GetNfsExportsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNfsExportsConnectionOK struct {
	Payload *models.NfsExportConnection
}

func (o *GetNfsExportsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports-connection][%d] getNfsExportsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNfsExportsConnectionOK) GetPayload() *models.NfsExportConnection {
	return o.Payload
}

func (o *GetNfsExportsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsExportConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsExportsConnectionBadRequest creates a GetNfsExportsConnectionBadRequest with default headers values
func NewGetNfsExportsConnectionBadRequest() *GetNfsExportsConnectionBadRequest {
	return &GetNfsExportsConnectionBadRequest{}
}

/* GetNfsExportsConnectionBadRequest describes a response with status code 400, with default header values.

GetNfsExportsConnectionBadRequest get nfs exports connection bad request
*/
type GetNfsExportsConnectionBadRequest struct {
	Payload string
}

func (o *GetNfsExportsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports-connection][%d] getNfsExportsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNfsExportsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNfsExportsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
