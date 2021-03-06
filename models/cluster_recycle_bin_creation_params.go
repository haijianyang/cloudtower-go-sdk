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

// ClusterRecycleBinCreationParams cluster recycle bin creation params
//
// swagger:model ClusterRecycleBinCreationParams
type ClusterRecycleBinCreationParams struct {

	// data
	// Required: true
	Data *ClusterRecycleBinCreationParamsData `json:"data"`

	// where
	// Required: true
	Where *ClusterWhereInput `json:"where"`
}

// Validate validates this cluster recycle bin creation params
func (m *ClusterRecycleBinCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterRecycleBinCreationParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterRecycleBinCreationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster recycle bin creation params based on the context it is used
func (m *ClusterRecycleBinCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterRecycleBinCreationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterRecycleBinCreationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterRecycleBinCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterRecycleBinCreationParams) UnmarshalBinary(b []byte) error {
	var res ClusterRecycleBinCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterRecycleBinCreationParamsData cluster recycle bin creation params data
//
// swagger:model ClusterRecycleBinCreationParamsData
type ClusterRecycleBinCreationParamsData struct {

	// enaled
	// Required: true
	Enaled *bool `json:"enaled"`

	// retain
	// Required: true
	Retain *float64 `json:"retain"`
}

// Validate validates this cluster recycle bin creation params data
func (m *ClusterRecycleBinCreationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnaled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterRecycleBinCreationParamsData) validateEnaled(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"enaled", "body", m.Enaled); err != nil {
		return err
	}

	return nil
}

func (m *ClusterRecycleBinCreationParamsData) validateRetain(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"retain", "body", m.Retain); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster recycle bin creation params data based on context it is used
func (m *ClusterRecycleBinCreationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterRecycleBinCreationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterRecycleBinCreationParamsData) UnmarshalBinary(b []byte) error {
	var res ClusterRecycleBinCreationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
