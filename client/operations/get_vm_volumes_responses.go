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

// GetVMVolumesReader is a Reader for the GetVMVolumes structure.
type GetVMVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMVolumesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMVolumesOK creates a GetVMVolumesOK with default headers values
func NewGetVMVolumesOK() *GetVMVolumesOK {
	return &GetVMVolumesOK{}
}

/* GetVMVolumesOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMVolumesOK struct {
	Payload []*models.VMVolume
}

func (o *GetVMVolumesOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-volumes][%d] getVmVolumesOK  %+v", 200, o.Payload)
}
func (o *GetVMVolumesOK) GetPayload() []*models.VMVolume {
	return o.Payload
}

func (o *GetVMVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMVolumesBadRequest creates a GetVMVolumesBadRequest with default headers values
func NewGetVMVolumesBadRequest() *GetVMVolumesBadRequest {
	return &GetVMVolumesBadRequest{}
}

/* GetVMVolumesBadRequest describes a response with status code 400, with default header values.

GetVMVolumesBadRequest get Vm volumes bad request
*/
type GetVMVolumesBadRequest struct {
	Payload string
}

func (o *GetVMVolumesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-volumes][%d] getVmVolumesBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMVolumesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVMVolumesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
