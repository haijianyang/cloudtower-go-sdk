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

// UpdateCustomizeAlertRuleReader is a Reader for the UpdateCustomizeAlertRule structure.
type UpdateCustomizeAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomizeAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomizeAlertRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomizeAlertRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCustomizeAlertRuleOK creates a UpdateCustomizeAlertRuleOK with default headers values
func NewUpdateCustomizeAlertRuleOK() *UpdateCustomizeAlertRuleOK {
	return &UpdateCustomizeAlertRuleOK{}
}

/* UpdateCustomizeAlertRuleOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateCustomizeAlertRuleOK struct {
	Payload []*models.WithTaskGlobalAlertRule
}

func (o *UpdateCustomizeAlertRuleOK) Error() string {
	return fmt.Sprintf("[POST /update-customize-alert-rule][%d] updateCustomizeAlertRuleOK  %+v", 200, o.Payload)
}
func (o *UpdateCustomizeAlertRuleOK) GetPayload() []*models.WithTaskGlobalAlertRule {
	return o.Payload
}

func (o *UpdateCustomizeAlertRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomizeAlertRuleBadRequest creates a UpdateCustomizeAlertRuleBadRequest with default headers values
func NewUpdateCustomizeAlertRuleBadRequest() *UpdateCustomizeAlertRuleBadRequest {
	return &UpdateCustomizeAlertRuleBadRequest{}
}

/* UpdateCustomizeAlertRuleBadRequest describes a response with status code 400, with default header values.

UpdateCustomizeAlertRuleBadRequest update customize alert rule bad request
*/
type UpdateCustomizeAlertRuleBadRequest struct {
	Payload string
}

func (o *UpdateCustomizeAlertRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-customize-alert-rule][%d] updateCustomizeAlertRuleBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCustomizeAlertRuleBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateCustomizeAlertRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
