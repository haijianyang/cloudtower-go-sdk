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

// GraphConnection graph connection
//
// swagger:model GraphConnection
type GraphConnection struct {

	// aggregate
	// Required: true
	Aggregate *GraphConnectionAggregate `json:"aggregate"`
}

// Validate validates this graph connection
func (m *GraphConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphConnection) validateAggregate(formats strfmt.Registry) error {

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

// ContextValidate validate this graph connection based on the context it is used
func (m *GraphConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphConnection) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GraphConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphConnection) UnmarshalBinary(b []byte) error {
	var res GraphConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GraphConnectionAggregate graph connection aggregate
//
// swagger:model GraphConnectionAggregate
type GraphConnectionAggregate struct {

	// count
	// Required: true
	Count *float64 `json:"count"`
}

// Validate validates this graph connection aggregate
func (m *GraphConnectionAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphConnectionAggregate) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("aggregate"+"."+"count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this graph connection aggregate based on context it is used
func (m *GraphConnectionAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GraphConnectionAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphConnectionAggregate) UnmarshalBinary(b []byte) error {
	var res GraphConnectionAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
