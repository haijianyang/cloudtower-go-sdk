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

// GetVMPlacementGroupsReader is a Reader for the GetVMPlacementGroups structure.
type GetVMPlacementGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMPlacementGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMPlacementGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMPlacementGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMPlacementGroupsOK creates a GetVMPlacementGroupsOK with default headers values
func NewGetVMPlacementGroupsOK() *GetVMPlacementGroupsOK {
	return &GetVMPlacementGroupsOK{}
}

/* GetVMPlacementGroupsOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMPlacementGroupsOK struct {
	Payload []*models.VMPlacementGroup
}

func (o *GetVMPlacementGroupsOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups][%d] getVmPlacementGroupsOK  %+v", 200, o.Payload)
}
func (o *GetVMPlacementGroupsOK) GetPayload() []*models.VMPlacementGroup {
	return o.Payload
}

func (o *GetVMPlacementGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMPlacementGroupsBadRequest creates a GetVMPlacementGroupsBadRequest with default headers values
func NewGetVMPlacementGroupsBadRequest() *GetVMPlacementGroupsBadRequest {
	return &GetVMPlacementGroupsBadRequest{}
}

/* GetVMPlacementGroupsBadRequest describes a response with status code 400, with default header values.

GetVMPlacementGroupsBadRequest get Vm placement groups bad request
*/
type GetVMPlacementGroupsBadRequest struct {
	Payload string
}

func (o *GetVMPlacementGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups][%d] getVmPlacementGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMPlacementGroupsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVMPlacementGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
