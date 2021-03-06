// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Nic nic
//
// swagger:model Nic
type Nic struct {

	// driver
	Driver *string `json:"driver,omitempty"`

	// driver state
	DriverState interface{} `json:"driver_state,omitempty"`

	// gateway ip
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// host
	// Required: true
	Host *NicHost `json:"host"`

	// ibdev
	Ibdev *string `json:"ibdev,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// is sriov
	IsSriov *bool `json:"is_sriov,omitempty"`

	// labels
	Labels []*NicLabelsItems0 `json:"labels,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mac address
	// Required: true
	MacAddress *string `json:"mac_address"`

	// max vf num
	MaxVfNum *float64 `json:"max_vf_num,omitempty"`

	// model
	Model *string `json:"model,omitempty"`

	// mtu
	// Required: true
	Mtu *float64 `json:"mtu"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nic uuid
	NicUUID *string `json:"nic_uuid,omitempty"`

	// physical
	// Required: true
	Physical *bool `json:"physical"`

	// rdma enabled
	RdmaEnabled *bool `json:"rdma_enabled,omitempty"`

	// running
	// Required: true
	Running *bool `json:"running"`

	// speed
	Speed *float64 `json:"speed,omitempty"`

	// subnet mask
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// total vf num
	TotalVfNum *float64 `json:"total_vf_num,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// up
	// Required: true
	Up *bool `json:"up"`

	// used vf num
	UsedVfNum *float64 `json:"used_vf_num,omitempty"`

	// vds
	Vds *NicVds `json:"vds,omitempty"`
}

// Validate validates this nic
func (m *Nic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Nic) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Nic) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateMacAddress(formats strfmt.Registry) error {

	if err := validate.Required("mac_address", "body", m.MacAddress); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateMtu(formats strfmt.Registry) error {

	if err := validate.Required("mtu", "body", m.Mtu); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validatePhysical(formats strfmt.Registry) error {

	if err := validate.Required("physical", "body", m.Physical); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateRunning(formats strfmt.Registry) error {

	if err := validate.Required("running", "body", m.Running); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateUp(formats strfmt.Registry) error {

	if err := validate.Required("up", "body", m.Up); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateVds(formats strfmt.Registry) error {
	if swag.IsZero(m.Vds) { // not required
		return nil
	}

	if m.Vds != nil {
		if err := m.Vds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nic based on the context it is used
func (m *Nic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Nic) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Nic) contextValidateVds(ctx context.Context, formats strfmt.Registry) error {

	if m.Vds != nil {
		if err := m.Vds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Nic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Nic) UnmarshalBinary(b []byte) error {
	var res Nic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NicHost nic host
//
// swagger:model NicHost
type NicHost struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this nic host
func (m *NicHost) Validate(formats strfmt.Registry) error {
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

func (m *NicHost) validateID(formats strfmt.Registry) error {

	if err := validate.Required("host"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NicHost) validateName(formats strfmt.Registry) error {

	if err := validate.Required("host"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nic host based on context it is used
func (m *NicHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NicHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NicHost) UnmarshalBinary(b []byte) error {
	var res NicHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NicLabelsItems0 nic labels items0
//
// swagger:model NicLabelsItems0
type NicLabelsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this nic labels items0
func (m *NicLabelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NicLabelsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nic labels items0 based on context it is used
func (m *NicLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NicLabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NicLabelsItems0) UnmarshalBinary(b []byte) error {
	var res NicLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NicVds nic vds
//
// swagger:model NicVds
type NicVds struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this nic vds
func (m *NicVds) Validate(formats strfmt.Registry) error {
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

func (m *NicVds) validateID(formats strfmt.Registry) error {

	if err := validate.Required("vds"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NicVds) validateName(formats strfmt.Registry) error {

	if err := validate.Required("vds"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nic vds based on context it is used
func (m *NicVds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NicVds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NicVds) UnmarshalBinary(b []byte) error {
	var res NicVds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
