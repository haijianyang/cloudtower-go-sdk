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

// GetVMEntityFilterResultsReader is a Reader for the GetVMEntityFilterResults structure.
type GetVMEntityFilterResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMEntityFilterResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMEntityFilterResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMEntityFilterResultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMEntityFilterResultsOK creates a GetVMEntityFilterResultsOK with default headers values
func NewGetVMEntityFilterResultsOK() *GetVMEntityFilterResultsOK {
	return &GetVMEntityFilterResultsOK{}
}

/* GetVMEntityFilterResultsOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMEntityFilterResultsOK struct {
	Payload []*models.VMEntityFilterResult
}

func (o *GetVMEntityFilterResultsOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-entity-filter-results][%d] getVmEntityFilterResultsOK  %+v", 200, o.Payload)
}
func (o *GetVMEntityFilterResultsOK) GetPayload() []*models.VMEntityFilterResult {
	return o.Payload
}

func (o *GetVMEntityFilterResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMEntityFilterResultsBadRequest creates a GetVMEntityFilterResultsBadRequest with default headers values
func NewGetVMEntityFilterResultsBadRequest() *GetVMEntityFilterResultsBadRequest {
	return &GetVMEntityFilterResultsBadRequest{}
}

/* GetVMEntityFilterResultsBadRequest describes a response with status code 400, with default header values.

GetVMEntityFilterResultsBadRequest get Vm entity filter results bad request
*/
type GetVMEntityFilterResultsBadRequest struct {
	Payload string
}

func (o *GetVMEntityFilterResultsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-entity-filter-results][%d] getVmEntityFilterResultsBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMEntityFilterResultsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVMEntityFilterResultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}