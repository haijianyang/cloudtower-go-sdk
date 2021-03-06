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

// SnmpTransportUpdationParams snmp transport updation params
//
// swagger:model SnmpTransportUpdationParams
type SnmpTransportUpdationParams struct {

	// data
	// Required: true
	Data *SnmpTransportUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *SnmpTransportWhereInput `json:"where"`
}

// Validate validates this snmp transport updation params
func (m *SnmpTransportUpdationParams) Validate(formats strfmt.Registry) error {
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

func (m *SnmpTransportUpdationParams) validateData(formats strfmt.Registry) error {

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

func (m *SnmpTransportUpdationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this snmp transport updation params based on the context it is used
func (m *SnmpTransportUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *SnmpTransportUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnmpTransportUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnmpTransportUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpTransportUpdationParams) UnmarshalBinary(b []byte) error {
	var res SnmpTransportUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnmpTransportUpdationParamsData snmp transport updation params data
//
// swagger:model SnmpTransportUpdationParamsData
type SnmpTransportUpdationParamsData struct {

	// auth pass phrase
	AuthPassPhrase string `json:"auth_pass_phrase,omitempty"`

	// auth protocol
	AuthProtocol SnmpAuthProtocol `json:"auth_protocol,omitempty"`

	// community
	Community string `json:"community,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// port
	Port float64 `json:"port,omitempty"`

	// privacy pass phrase
	PrivacyPassPhrase string `json:"privacy_pass_phrase,omitempty"`

	// privacy protocol
	PrivacyProtocol SnmpPrivacyProtocol `json:"privacy_protocol,omitempty"`

	// protocol
	Protocol SnmpProtocol `json:"protocol,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// version
	Version SnmpVersion `json:"version,omitempty"`
}

// Validate validates this snmp transport updation params data
func (m *SnmpTransportUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
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

func (m *SnmpTransportUpdationParamsData) validateAuthProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthProtocol) { // not required
		return nil
	}

	if err := m.AuthProtocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "auth_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTransportUpdationParamsData) validatePrivacyProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivacyProtocol) { // not required
		return nil
	}

	if err := m.PrivacyProtocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "privacy_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTransportUpdationParamsData) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTransportUpdationParamsData) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := m.Version.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "version")
		}
		return err
	}

	return nil
}

// ContextValidate validate this snmp transport updation params data based on the context it is used
func (m *SnmpTransportUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivacyProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpTransportUpdationParamsData) contextValidateAuthProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthProtocol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "auth_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTransportUpdationParamsData) contextValidatePrivacyProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PrivacyProtocol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "privacy_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTransportUpdationParamsData) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTransportUpdationParamsData) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Version.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "version")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpTransportUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpTransportUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res SnmpTransportUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
