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

// UploadTaskConnection upload task connection
//
// swagger:model UploadTaskConnection
type UploadTaskConnection struct {

	// aggregate
	// Required: true
	Aggregate *UploadTaskConnectionAggregate `json:"aggregate"`
}

// Validate validates this upload task connection
func (m *UploadTaskConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadTaskConnection) validateAggregate(formats strfmt.Registry) error {

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

// ContextValidate validate this upload task connection based on the context it is used
func (m *UploadTaskConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadTaskConnection) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UploadTaskConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadTaskConnection) UnmarshalBinary(b []byte) error {
	var res UploadTaskConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UploadTaskConnectionAggregate upload task connection aggregate
//
// swagger:model UploadTaskConnectionAggregate
type UploadTaskConnectionAggregate struct {

	// count
	// Required: true
	Count *float64 `json:"count"`
}

// Validate validates this upload task connection aggregate
func (m *UploadTaskConnectionAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadTaskConnectionAggregate) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("aggregate"+"."+"count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this upload task connection aggregate based on context it is used
func (m *UploadTaskConnectionAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UploadTaskConnectionAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadTaskConnectionAggregate) UnmarshalBinary(b []byte) error {
	var res UploadTaskConnectionAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
