// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConvertVMTemplateToVMParams convert Vm template to Vm params
//
// swagger:model ConvertVmTemplateToVmParams
type ConvertVMTemplateToVMParams struct {

	// converted from template id
	// Required: true
	ConvertedFromTemplateID *string `json:"converted_from_template_id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this convert Vm template to Vm params
func (m *ConvertVMTemplateToVMParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConvertedFromTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConvertVMTemplateToVMParams) validateConvertedFromTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("converted_from_template_id", "body", m.ConvertedFromTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *ConvertVMTemplateToVMParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this convert Vm template to Vm params based on context it is used
func (m *ConvertVMTemplateToVMParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConvertVMTemplateToVMParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConvertVMTemplateToVMParams) UnmarshalBinary(b []byte) error {
	var res ConvertVMTemplateToVMParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
