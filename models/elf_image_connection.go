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

// ElfImageConnection elf image connection
//
// swagger:model ElfImageConnection
type ElfImageConnection struct {

	// aggregate
	// Required: true
	Aggregate *ElfImageConnectionAggregate `json:"aggregate"`
}

// Validate validates this elf image connection
func (m *ElfImageConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImageConnection) validateAggregate(formats strfmt.Registry) error {

	if err := validate.Required("aggregate", "body", m.Aggregate); err != nil {
		return err
	}

	if m.Aggregate != nil {
		if err := m.Aggregate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this elf image connection based on the context it is used
func (m *ElfImageConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImageConnection) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregate != nil {
		if err := m.Aggregate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageConnection) UnmarshalBinary(b []byte) error {
	var res ElfImageConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfImageConnectionAggregate elf image connection aggregate
//
// swagger:model ElfImageConnectionAggregate
type ElfImageConnectionAggregate struct {

	// count
	// Required: true
	Count *float64 `json:"count"`
}

// Validate validates this elf image connection aggregate
func (m *ElfImageConnectionAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImageConnectionAggregate) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("aggregate"+"."+"count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf image connection aggregate based on context it is used
func (m *ElfImageConnectionAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageConnectionAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageConnectionAggregate) UnmarshalBinary(b []byte) error {
	var res ElfImageConnectionAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
