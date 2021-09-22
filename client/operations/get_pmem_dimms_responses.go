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

// GetPmemDimmsReader is a Reader for the GetPmemDimms structure.
type GetPmemDimmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPmemDimmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPmemDimmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPmemDimmsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPmemDimmsOK creates a GetPmemDimmsOK with default headers values
func NewGetPmemDimmsOK() *GetPmemDimmsOK {
	return &GetPmemDimmsOK{}
}

/* GetPmemDimmsOK describes a response with status code 200, with default header values.

Ok
*/
type GetPmemDimmsOK struct {
	Payload []*models.PmemDimm
}

func (o *GetPmemDimmsOK) Error() string {
	return fmt.Sprintf("[POST /get-pmem-dimms][%d] getPmemDimmsOK  %+v", 200, o.Payload)
}
func (o *GetPmemDimmsOK) GetPayload() []*models.PmemDimm {
	return o.Payload
}

func (o *GetPmemDimmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPmemDimmsBadRequest creates a GetPmemDimmsBadRequest with default headers values
func NewGetPmemDimmsBadRequest() *GetPmemDimmsBadRequest {
	return &GetPmemDimmsBadRequest{}
}

/* GetPmemDimmsBadRequest describes a response with status code 400, with default header values.

GetPmemDimmsBadRequest get pmem dimms bad request
*/
type GetPmemDimmsBadRequest struct {
	Payload string
}

func (o *GetPmemDimmsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-pmem-dimms][%d] getPmemDimmsBadRequest  %+v", 400, o.Payload)
}
func (o *GetPmemDimmsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetPmemDimmsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
