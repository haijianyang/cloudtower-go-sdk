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

// NfsExportWhereInput nfs export where input
//
// swagger:model NfsExportWhereInput
type NfsExportWhereInput struct {

	// a n d
	AND []*NfsExportWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*NfsExportWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*NfsExportWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster interface{} `json:"cluster,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot interface{} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// export inode id
	ExportInodeID *string `json:"export_inode_id,omitempty"`

	// export inode id contains
	ExportInodeIDContains *string `json:"export_inode_id_contains,omitempty"`

	// export inode id ends with
	ExportInodeIDEndsWith *string `json:"export_inode_id_ends_with,omitempty"`

	// export inode id gt
	ExportInodeIDGt *string `json:"export_inode_id_gt,omitempty"`

	// export inode id gte
	ExportInodeIDGte *string `json:"export_inode_id_gte,omitempty"`

	// export inode id in
	ExportInodeIDIn []string `json:"export_inode_id_in,omitempty"`

	// export inode id lt
	ExportInodeIDLt *string `json:"export_inode_id_lt,omitempty"`

	// export inode id lte
	ExportInodeIDLte *string `json:"export_inode_id_lte,omitempty"`

	// export inode id not
	ExportInodeIDNot *string `json:"export_inode_id_not,omitempty"`

	// export inode id not contains
	ExportInodeIDNotContains *string `json:"export_inode_id_not_contains,omitempty"`

	// export inode id not ends with
	ExportInodeIDNotEndsWith *string `json:"export_inode_id_not_ends_with,omitempty"`

	// export inode id not in
	ExportInodeIDNotIn []string `json:"export_inode_id_not_in,omitempty"`

	// export inode id not starts with
	ExportInodeIDNotStartsWith *string `json:"export_inode_id_not_starts_with,omitempty"`

	// export inode id starts with
	ExportInodeIDStartsWith *string `json:"export_inode_id_starts_with,omitempty"`

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

	// inodes every
	InodesEvery interface{} `json:"inodes_every,omitempty"`

	// inodes none
	InodesNone interface{} `json:"inodes_none,omitempty"`

	// inodes some
	InodesSome interface{} `json:"inodes_some,omitempty"`

	// internal
	Internal *bool `json:"internal,omitempty"`

	// internal not
	InternalNot *bool `json:"internal_not,omitempty"`

	// ip whitelist
	IPWhitelist *string `json:"ip_whitelist,omitempty"`

	// ip whitelist contains
	IPWhitelistContains *string `json:"ip_whitelist_contains,omitempty"`

	// ip whitelist ends with
	IPWhitelistEndsWith *string `json:"ip_whitelist_ends_with,omitempty"`

	// ip whitelist gt
	IPWhitelistGt *string `json:"ip_whitelist_gt,omitempty"`

	// ip whitelist gte
	IPWhitelistGte *string `json:"ip_whitelist_gte,omitempty"`

	// ip whitelist in
	IPWhitelistIn []string `json:"ip_whitelist_in,omitempty"`

	// ip whitelist lt
	IPWhitelistLt *string `json:"ip_whitelist_lt,omitempty"`

	// ip whitelist lte
	IPWhitelistLte *string `json:"ip_whitelist_lte,omitempty"`

	// ip whitelist not
	IPWhitelistNot *string `json:"ip_whitelist_not,omitempty"`

	// ip whitelist not contains
	IPWhitelistNotContains *string `json:"ip_whitelist_not_contains,omitempty"`

	// ip whitelist not ends with
	IPWhitelistNotEndsWith *string `json:"ip_whitelist_not_ends_with,omitempty"`

	// ip whitelist not in
	IPWhitelistNotIn []string `json:"ip_whitelist_not_in,omitempty"`

	// ip whitelist not starts with
	IPWhitelistNotStartsWith *string `json:"ip_whitelist_not_starts_with,omitempty"`

	// ip whitelist starts with
	IPWhitelistStartsWith *string `json:"ip_whitelist_starts_with,omitempty"`

	// labels every
	LabelsEvery interface{} `json:"labels_every,omitempty"`

	// labels none
	LabelsNone interface{} `json:"labels_none,omitempty"`

	// labels some
	LabelsSome interface{} `json:"labels_some,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// replica num
	ReplicaNum *float64 `json:"replica_num,omitempty"`

	// replica num gt
	ReplicaNumGt *float64 `json:"replica_num_gt,omitempty"`

	// replica num gte
	ReplicaNumGte *float64 `json:"replica_num_gte,omitempty"`

	// replica num in
	ReplicaNumIn []float64 `json:"replica_num_in,omitempty"`

	// replica num lt
	ReplicaNumLt *float64 `json:"replica_num_lt,omitempty"`

	// replica num lte
	ReplicaNumLte *float64 `json:"replica_num_lte,omitempty"`

	// replica num not
	ReplicaNumNot *float64 `json:"replica_num_not,omitempty"`

	// replica num not in
	ReplicaNumNotIn []float64 `json:"replica_num_not_in,omitempty"`

	// thin provision
	ThinProvision *bool `json:"thin_provision,omitempty"`

	// thin provision not
	ThinProvisionNot *bool `json:"thin_provision_not,omitempty"`
}

// Validate validates this nfs export where input
func (m *NfsExportWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsExportWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *NfsExportWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *NfsExportWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *NfsExportWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *NfsExportWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this nfs export where input based on the context it is used
func (m *NfsExportWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsExportWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NfsExportWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NfsExportWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NfsExportWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *NfsExportWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NfsExportWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NfsExportWhereInput) UnmarshalBinary(b []byte) error {
	var res NfsExportWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
