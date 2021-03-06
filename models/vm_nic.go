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

// VMNic Vm nic
//
// swagger:model VmNic
type VMNic struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// gateway
	Gateway *string `json:"gateway,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// interface id
	InterfaceID *string `json:"interface_id,omitempty"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mac address
	MacAddress *string `json:"mac_address,omitempty"`

	// mirror
	Mirror *bool `json:"mirror,omitempty"`

	// model
	Model interface{} `json:"model,omitempty"`

	// nic
	Nic *VMNicNic `json:"nic,omitempty"`

	// order
	Order *float64 `json:"order,omitempty"`

	// subnet mask
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// vlan
	Vlan *VMNicVlan `json:"vlan,omitempty"`

	// vm
	// Required: true
	VM *VMNicVM `json:"vm"`
}

// Validate validates this Vm nic
func (m *VMNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMNic) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMNic) validateNic(formats strfmt.Registry) error {
	if swag.IsZero(m.Nic) { // not required
		return nil
	}

	if m.Nic != nil {
		if err := m.Nic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nic")
			}
			return err
		}
	}

	return nil
}

func (m *VMNic) validateVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlan) { // not required
		return nil
	}

	if m.Vlan != nil {
		if err := m.Vlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

func (m *VMNic) validateVM(formats strfmt.Registry) error {

	if err := validate.Required("vm", "body", m.VM); err != nil {
		return err
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm nic based on the context it is used
func (m *VMNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNic) contextValidateNic(ctx context.Context, formats strfmt.Registry) error {

	if m.Nic != nil {
		if err := m.Nic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nic")
			}
			return err
		}
	}

	return nil
}

func (m *VMNic) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlan != nil {
		if err := m.Vlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

func (m *VMNic) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNic) UnmarshalBinary(b []byte) error {
	var res VMNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMNicNic VM nic nic
//
// swagger:model VMNicNic
type VMNicNic struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this VM nic nic
func (m *VMNicNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *VMNicNic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("nic"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMNicNic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("nic"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM nic nic based on context it is used
func (m *VMNicNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMNicNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNicNic) UnmarshalBinary(b []byte) error {
	var res VMNicNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMNicVM VM nic VM
//
// swagger:model VMNicVM
type VMNicVM struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this VM nic VM
func (m *VMNicVM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *VMNicVM) validateID(formats strfmt.Registry) error {

	if err := validate.Required("vm"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMNicVM) validateName(formats strfmt.Registry) error {

	if err := validate.Required("vm"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM nic VM based on context it is used
func (m *VMNicVM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMNicVM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNicVM) UnmarshalBinary(b []byte) error {
	var res VMNicVM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMNicVlan VM nic vlan
//
// swagger:model VMNicVlan
type VMNicVlan struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this VM nic vlan
func (m *VMNicVlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *VMNicVlan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMNicVlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM nic vlan based on context it is used
func (m *VMNicVlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMNicVlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNicVlan) UnmarshalBinary(b []byte) error {
	var res VMNicVlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
