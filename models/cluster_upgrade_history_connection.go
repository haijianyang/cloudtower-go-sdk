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

// ClusterUpgradeHistoryConnection cluster upgrade history connection
//
// swagger:model ClusterUpgradeHistoryConnection
type ClusterUpgradeHistoryConnection struct {

	// aggregate
	// Required: true
	Aggregate *ClusterUpgradeHistoryConnectionAggregate `json:"aggregate"`
}

// Validate validates this cluster upgrade history connection
func (m *ClusterUpgradeHistoryConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpgradeHistoryConnection) validateAggregate(formats strfmt.Registry) error {

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

// ContextValidate validate this cluster upgrade history connection based on the context it is used
func (m *ClusterUpgradeHistoryConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpgradeHistoryConnection) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClusterUpgradeHistoryConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpgradeHistoryConnection) UnmarshalBinary(b []byte) error {
	var res ClusterUpgradeHistoryConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterUpgradeHistoryConnectionAggregate cluster upgrade history connection aggregate
//
// swagger:model ClusterUpgradeHistoryConnectionAggregate
type ClusterUpgradeHistoryConnectionAggregate struct {

	// count
	// Required: true
	Count *float64 `json:"count"`
}

// Validate validates this cluster upgrade history connection aggregate
func (m *ClusterUpgradeHistoryConnectionAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpgradeHistoryConnectionAggregate) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("aggregate"+"."+"count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster upgrade history connection aggregate based on context it is used
func (m *ClusterUpgradeHistoryConnectionAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUpgradeHistoryConnectionAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpgradeHistoryConnectionAggregate) UnmarshalBinary(b []byte) error {
	var res ClusterUpgradeHistoryConnectionAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
