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

// GetGlobalSettingsesReader is a Reader for the GetGlobalSettingses structure.
type GetGlobalSettingsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalSettingsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalSettingsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGlobalSettingsesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGlobalSettingsesOK creates a GetGlobalSettingsesOK with default headers values
func NewGetGlobalSettingsesOK() *GetGlobalSettingsesOK {
	return &GetGlobalSettingsesOK{}
}

/* GetGlobalSettingsesOK describes a response with status code 200, with default header values.

Ok
*/
type GetGlobalSettingsesOK struct {
	Payload []*models.GlobalSettings
}

func (o *GetGlobalSettingsesOK) Error() string {
	return fmt.Sprintf("[POST /get-global-settingses][%d] getGlobalSettingsesOK  %+v", 200, o.Payload)
}
func (o *GetGlobalSettingsesOK) GetPayload() []*models.GlobalSettings {
	return o.Payload
}

func (o *GetGlobalSettingsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalSettingsesBadRequest creates a GetGlobalSettingsesBadRequest with default headers values
func NewGetGlobalSettingsesBadRequest() *GetGlobalSettingsesBadRequest {
	return &GetGlobalSettingsesBadRequest{}
}

/* GetGlobalSettingsesBadRequest describes a response with status code 400, with default header values.

GetGlobalSettingsesBadRequest get global settingses bad request
*/
type GetGlobalSettingsesBadRequest struct {
	Payload string
}

func (o *GetGlobalSettingsesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-global-settingses][%d] getGlobalSettingsesBadRequest  %+v", 400, o.Payload)
}
func (o *GetGlobalSettingsesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetGlobalSettingsesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
