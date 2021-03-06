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
)

// IscsiConnectionWhereInput iscsi connection where input
//
// swagger:model IscsiConnectionWhereInput
type IscsiConnectionWhereInput struct {

	// a n d
	AND []*IscsiConnectionWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*IscsiConnectionWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*IscsiConnectionWhereInput `json:"OR,omitempty"`

	// client port
	ClientPort *float64 `json:"client_port,omitempty"`

	// client port gt
	ClientPortGt *float64 `json:"client_port_gt,omitempty"`

	// client port gte
	ClientPortGte *float64 `json:"client_port_gte,omitempty"`

	// client port in
	ClientPortIn []float64 `json:"client_port_in,omitempty"`

	// client port lt
	ClientPortLt *float64 `json:"client_port_lt,omitempty"`

	// client port lte
	ClientPortLte *float64 `json:"client_port_lte,omitempty"`

	// client port not
	ClientPortNot *float64 `json:"client_port_not,omitempty"`

	// client port not in
	ClientPortNotIn []float64 `json:"client_port_not_in,omitempty"`

	// cluster
	Cluster interface{} `json:"cluster,omitempty"`

	// host
	Host interface{} `json:"host,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// initiator ip
	InitiatorIP *string `json:"initiator_ip,omitempty"`

	// initiator ip contains
	InitiatorIPContains *string `json:"initiator_ip_contains,omitempty"`

	// initiator ip ends with
	InitiatorIPEndsWith *string `json:"initiator_ip_ends_with,omitempty"`

	// initiator ip gt
	InitiatorIPGt *string `json:"initiator_ip_gt,omitempty"`

	// initiator ip gte
	InitiatorIPGte *string `json:"initiator_ip_gte,omitempty"`

	// initiator ip in
	InitiatorIPIn []string `json:"initiator_ip_in,omitempty"`

	// initiator ip lt
	InitiatorIPLt *string `json:"initiator_ip_lt,omitempty"`

	// initiator ip lte
	InitiatorIPLte *string `json:"initiator_ip_lte,omitempty"`

	// initiator ip not
	InitiatorIPNot *string `json:"initiator_ip_not,omitempty"`

	// initiator ip not contains
	InitiatorIPNotContains *string `json:"initiator_ip_not_contains,omitempty"`

	// initiator ip not ends with
	InitiatorIPNotEndsWith *string `json:"initiator_ip_not_ends_with,omitempty"`

	// initiator ip not in
	InitiatorIPNotIn []string `json:"initiator_ip_not_in,omitempty"`

	// initiator ip not starts with
	InitiatorIPNotStartsWith *string `json:"initiator_ip_not_starts_with,omitempty"`

	// initiator ip starts with
	InitiatorIPStartsWith *string `json:"initiator_ip_starts_with,omitempty"`

	// iscsi target
	IscsiTarget interface{} `json:"iscsi_target,omitempty"`

	// nvmf subsystem
	NvmfSubsystem interface{} `json:"nvmf_subsystem,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// type in
	TypeIn []StoreConnectionType `json:"type_in,omitempty"`

	// type not
	TypeNot interface{} `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []StoreConnectionType `json:"type_not_in,omitempty"`
}

// Validate validates this iscsi connection where input
func (m *IscsiConnectionWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) validateTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this iscsi connection where input based on the context it is used
func (m *IscsiConnectionWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IscsiConnectionWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiConnectionWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnectionWhereInput) UnmarshalBinary(b []byte) error {
	var res IscsiConnectionWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
