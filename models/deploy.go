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

// Deploy deploy
//
// swagger:model Deploy
type Deploy struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// license
	License *DeployLicense `json:"license,omitempty"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this deploy
func (m *Deploy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deploy) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Deploy) validateLicense(formats strfmt.Registry) error {
	if swag.IsZero(m.License) { // not required
		return nil
	}

	if m.License != nil {
		if err := m.License.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license")
			}
			return err
		}
	}

	return nil
}

func (m *Deploy) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this deploy based on the context it is used
func (m *Deploy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLicense(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Deploy) contextValidateLicense(ctx context.Context, formats strfmt.Registry) error {

	if m.License != nil {
		if err := m.License.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Deploy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Deploy) UnmarshalBinary(b []byte) error {
	var res Deploy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeployLicense deploy license
//
// swagger:model DeployLicense
type DeployLicense struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this deploy license
func (m *DeployLicense) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeployLicense) validateID(formats strfmt.Registry) error {

	if err := validate.Required("license"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this deploy license based on context it is used
func (m *DeployLicense) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeployLicense) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeployLicense) UnmarshalBinary(b []byte) error {
	var res DeployLicense
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
