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

// LicenseWhereInput license where input
//
// swagger:model LicenseWhereInput
type LicenseWhereInput struct {

	// a n d
	AND []*LicenseWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*LicenseWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*LicenseWhereInput `json:"OR,omitempty"`

	// expire date
	ExpireDate *string `json:"expire_date,omitempty"`

	// expire date gt
	ExpireDateGt *string `json:"expire_date_gt,omitempty"`

	// expire date gte
	ExpireDateGte *string `json:"expire_date_gte,omitempty"`

	// expire date in
	ExpireDateIn []string `json:"expire_date_in,omitempty"`

	// expire date lt
	ExpireDateLt *string `json:"expire_date_lt,omitempty"`

	// expire date lte
	ExpireDateLte *string `json:"expire_date_lte,omitempty"`

	// expire date not
	ExpireDateNot *string `json:"expire_date_not,omitempty"`

	// expire date not in
	ExpireDateNotIn []string `json:"expire_date_not_in,omitempty"`

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

	// license serial
	LicenseSerial *string `json:"license_serial,omitempty"`

	// license serial contains
	LicenseSerialContains *string `json:"license_serial_contains,omitempty"`

	// license serial ends with
	LicenseSerialEndsWith *string `json:"license_serial_ends_with,omitempty"`

	// license serial gt
	LicenseSerialGt *string `json:"license_serial_gt,omitempty"`

	// license serial gte
	LicenseSerialGte *string `json:"license_serial_gte,omitempty"`

	// license serial in
	LicenseSerialIn []string `json:"license_serial_in,omitempty"`

	// license serial lt
	LicenseSerialLt *string `json:"license_serial_lt,omitempty"`

	// license serial lte
	LicenseSerialLte *string `json:"license_serial_lte,omitempty"`

	// license serial not
	LicenseSerialNot *string `json:"license_serial_not,omitempty"`

	// license serial not contains
	LicenseSerialNotContains *string `json:"license_serial_not_contains,omitempty"`

	// license serial not ends with
	LicenseSerialNotEndsWith *string `json:"license_serial_not_ends_with,omitempty"`

	// license serial not in
	LicenseSerialNotIn []string `json:"license_serial_not_in,omitempty"`

	// license serial not starts with
	LicenseSerialNotStartsWith *string `json:"license_serial_not_starts_with,omitempty"`

	// license serial starts with
	LicenseSerialStartsWith *string `json:"license_serial_starts_with,omitempty"`

	// maintenance end date
	MaintenanceEndDate *string `json:"maintenance_end_date,omitempty"`

	// maintenance end date gt
	MaintenanceEndDateGt *string `json:"maintenance_end_date_gt,omitempty"`

	// maintenance end date gte
	MaintenanceEndDateGte *string `json:"maintenance_end_date_gte,omitempty"`

	// maintenance end date in
	MaintenanceEndDateIn []string `json:"maintenance_end_date_in,omitempty"`

	// maintenance end date lt
	MaintenanceEndDateLt *string `json:"maintenance_end_date_lt,omitempty"`

	// maintenance end date lte
	MaintenanceEndDateLte *string `json:"maintenance_end_date_lte,omitempty"`

	// maintenance end date not
	MaintenanceEndDateNot *string `json:"maintenance_end_date_not,omitempty"`

	// maintenance end date not in
	MaintenanceEndDateNotIn []string `json:"maintenance_end_date_not_in,omitempty"`

	// maintenance start date
	MaintenanceStartDate *string `json:"maintenance_start_date,omitempty"`

	// maintenance start date gt
	MaintenanceStartDateGt *string `json:"maintenance_start_date_gt,omitempty"`

	// maintenance start date gte
	MaintenanceStartDateGte *string `json:"maintenance_start_date_gte,omitempty"`

	// maintenance start date in
	MaintenanceStartDateIn []string `json:"maintenance_start_date_in,omitempty"`

	// maintenance start date lt
	MaintenanceStartDateLt *string `json:"maintenance_start_date_lt,omitempty"`

	// maintenance start date lte
	MaintenanceStartDateLte *string `json:"maintenance_start_date_lte,omitempty"`

	// maintenance start date not
	MaintenanceStartDateNot *string `json:"maintenance_start_date_not,omitempty"`

	// maintenance start date not in
	MaintenanceStartDateNotIn []string `json:"maintenance_start_date_not_in,omitempty"`

	// max chunk num
	MaxChunkNum *float64 `json:"max_chunk_num,omitempty"`

	// max chunk num gt
	MaxChunkNumGt *float64 `json:"max_chunk_num_gt,omitempty"`

	// max chunk num gte
	MaxChunkNumGte *float64 `json:"max_chunk_num_gte,omitempty"`

	// max chunk num in
	MaxChunkNumIn []float64 `json:"max_chunk_num_in,omitempty"`

	// max chunk num lt
	MaxChunkNumLt *float64 `json:"max_chunk_num_lt,omitempty"`

	// max chunk num lte
	MaxChunkNumLte *float64 `json:"max_chunk_num_lte,omitempty"`

	// max chunk num not
	MaxChunkNumNot *float64 `json:"max_chunk_num_not,omitempty"`

	// max chunk num not in
	MaxChunkNumNotIn []float64 `json:"max_chunk_num_not_in,omitempty"`

	// max cluster num
	MaxClusterNum *float64 `json:"max_cluster_num,omitempty"`

	// max cluster num gt
	MaxClusterNumGt *float64 `json:"max_cluster_num_gt,omitempty"`

	// max cluster num gte
	MaxClusterNumGte *float64 `json:"max_cluster_num_gte,omitempty"`

	// max cluster num in
	MaxClusterNumIn []float64 `json:"max_cluster_num_in,omitempty"`

	// max cluster num lt
	MaxClusterNumLt *float64 `json:"max_cluster_num_lt,omitempty"`

	// max cluster num lte
	MaxClusterNumLte *float64 `json:"max_cluster_num_lte,omitempty"`

	// max cluster num not
	MaxClusterNumNot *float64 `json:"max_cluster_num_not,omitempty"`

	// max cluster num not in
	MaxClusterNumNotIn []float64 `json:"max_cluster_num_not_in,omitempty"`

	// sign date
	SignDate *string `json:"sign_date,omitempty"`

	// sign date gt
	SignDateGt *string `json:"sign_date_gt,omitempty"`

	// sign date gte
	SignDateGte *string `json:"sign_date_gte,omitempty"`

	// sign date in
	SignDateIn []string `json:"sign_date_in,omitempty"`

	// sign date lt
	SignDateLt *string `json:"sign_date_lt,omitempty"`

	// sign date lte
	SignDateLte *string `json:"sign_date_lte,omitempty"`

	// sign date not
	SignDateNot *string `json:"sign_date_not,omitempty"`

	// sign date not in
	SignDateNotIn []string `json:"sign_date_not_in,omitempty"`

	// software edition
	SoftwareEdition interface{} `json:"software_edition,omitempty"`

	// software edition in
	SoftwareEditionIn []SoftwareEdition `json:"software_edition_in,omitempty"`

	// software edition not
	SoftwareEditionNot interface{} `json:"software_edition_not,omitempty"`

	// software edition not in
	SoftwareEditionNotIn []SoftwareEdition `json:"software_edition_not_in,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// type in
	TypeIn []LicenseType `json:"type_in,omitempty"`

	// type not
	TypeNot interface{} `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []LicenseType `json:"type_not_in,omitempty"`
}

// Validate validates this license where input
func (m *LicenseWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateSoftwareEditionIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEditionNotIn(formats); err != nil {
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

func (m *LicenseWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *LicenseWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *LicenseWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *LicenseWhereInput) validateSoftwareEditionIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareEditionIn); i++ {

		if err := m.SoftwareEditionIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *LicenseWhereInput) validateSoftwareEditionNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareEditionNotIn); i++ {

		if err := m.SoftwareEditionNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *LicenseWhereInput) validateTypeIn(formats strfmt.Registry) error {
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

func (m *LicenseWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
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

// ContextValidate validate this license where input based on the context it is used
func (m *LicenseWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateSoftwareEditionIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEditionNotIn(ctx, formats); err != nil {
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

func (m *LicenseWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LicenseWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LicenseWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LicenseWhereInput) contextValidateSoftwareEditionIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SoftwareEditionIn); i++ {

		if err := m.SoftwareEditionIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *LicenseWhereInput) contextValidateSoftwareEditionNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SoftwareEditionNotIn); i++ {

		if err := m.SoftwareEditionNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *LicenseWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LicenseWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LicenseWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseWhereInput) UnmarshalBinary(b []byte) error {
	var res LicenseWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
