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

// MigrateTransmitterConnection migrate transmitter connection
//
// swagger:model MigrateTransmitterConnection
type MigrateTransmitterConnection struct {

	// aggregate
	// Required: true
	Aggregate *MigrateTransmitterConnectionAggregate `json:"aggregate"`
}

// Validate validates this migrate transmitter connection
func (m *MigrateTransmitterConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateTransmitterConnection) validateAggregate(formats strfmt.Registry) error {

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

// ContextValidate validate this migrate transmitter connection based on the context it is used
func (m *MigrateTransmitterConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateTransmitterConnection) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MigrateTransmitterConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateTransmitterConnection) UnmarshalBinary(b []byte) error {
	var res MigrateTransmitterConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MigrateTransmitterConnectionAggregate migrate transmitter connection aggregate
//
// swagger:model MigrateTransmitterConnectionAggregate
type MigrateTransmitterConnectionAggregate struct {

	// count
	// Required: true
	Count *float64 `json:"count"`
}

// Validate validates this migrate transmitter connection aggregate
func (m *MigrateTransmitterConnectionAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateTransmitterConnectionAggregate) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("aggregate"+"."+"count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this migrate transmitter connection aggregate based on context it is used
func (m *MigrateTransmitterConnectionAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MigrateTransmitterConnectionAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateTransmitterConnectionAggregate) UnmarshalBinary(b []byte) error {
	var res MigrateTransmitterConnectionAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
