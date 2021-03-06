// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteVMReader is a Reader for the DeleteVM structure.
type DeleteVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMOK creates a DeleteVMOK with default headers values
func NewDeleteVMOK() *DeleteVMOK {
	return &DeleteVMOK{}
}

/* DeleteVMOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVMOK struct {
	Payload []*DeleteVMOKBodyItems0
}

func (o *DeleteVMOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmOK  %+v", 200, o.Payload)
}
func (o *DeleteVMOK) GetPayload() []*DeleteVMOKBodyItems0 {
	return o.Payload
}

func (o *DeleteVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMBadRequest creates a DeleteVMBadRequest with default headers values
func NewDeleteVMBadRequest() *DeleteVMBadRequest {
	return &DeleteVMBadRequest{}
}

/* DeleteVMBadRequest describes a response with status code 400, with default header values.

DeleteVMBadRequest delete Vm bad request
*/
type DeleteVMBadRequest struct {
	Payload string
}

func (o *DeleteVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteVMOKBodyItems0 delete VM o k body items0
swagger:model DeleteVMOKBodyItems0
*/
type DeleteVMOKBodyItems0 struct {

	// data
	// Required: true
	Data interface{} `json:"data"`

	// task id
	// Required: true
	TaskID *string `json:"task_id"`
}

// Validate validates this delete VM o k body items0
func (o *DeleteVMOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteVMOKBodyItems0) validateData(formats strfmt.Registry) error {

	if o.Data == nil {
		return errors.Required("data", "body", nil)
	}

	return nil
}

func (o *DeleteVMOKBodyItems0) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("task_id", "body", o.TaskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete VM o k body items0 based on context it is used
func (o *DeleteVMOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteVMOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteVMOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res DeleteVMOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
