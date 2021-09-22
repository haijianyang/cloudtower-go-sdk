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

// GetVcenterAccountsConnectionReader is a Reader for the GetVcenterAccountsConnection structure.
type GetVcenterAccountsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVcenterAccountsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVcenterAccountsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVcenterAccountsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVcenterAccountsConnectionOK creates a GetVcenterAccountsConnectionOK with default headers values
func NewGetVcenterAccountsConnectionOK() *GetVcenterAccountsConnectionOK {
	return &GetVcenterAccountsConnectionOK{}
}

/* GetVcenterAccountsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVcenterAccountsConnectionOK struct {
	Payload *models.VcenterAccountConnection
}

func (o *GetVcenterAccountsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vcenter-accounts-connection][%d] getVcenterAccountsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVcenterAccountsConnectionOK) GetPayload() *models.VcenterAccountConnection {
	return o.Payload
}

func (o *GetVcenterAccountsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcenterAccountConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVcenterAccountsConnectionBadRequest creates a GetVcenterAccountsConnectionBadRequest with default headers values
func NewGetVcenterAccountsConnectionBadRequest() *GetVcenterAccountsConnectionBadRequest {
	return &GetVcenterAccountsConnectionBadRequest{}
}

/* GetVcenterAccountsConnectionBadRequest describes a response with status code 400, with default header values.

GetVcenterAccountsConnectionBadRequest get vcenter accounts connection bad request
*/
type GetVcenterAccountsConnectionBadRequest struct {
	Payload string
}

func (o *GetVcenterAccountsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vcenter-accounts-connection][%d] getVcenterAccountsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVcenterAccountsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVcenterAccountsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
