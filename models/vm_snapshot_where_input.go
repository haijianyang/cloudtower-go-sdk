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

// VMSnapshotWhereInput Vm snapshot where input
//
// swagger:model VmSnapshotWhereInput
type VMSnapshotWhereInput struct {

	// a n d
	AND []*VMSnapshotWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VMSnapshotWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VMSnapshotWhereInput `json:"OR,omitempty"`

	// clock offset
	ClockOffset interface{} `json:"clock_offset,omitempty"`

	// clock offset in
	ClockOffsetIn []VMClockOffset `json:"clock_offset_in,omitempty"`

	// clock offset not
	ClockOffsetNot interface{} `json:"clock_offset_not,omitempty"`

	// clock offset not in
	ClockOffsetNotIn []VMClockOffset `json:"clock_offset_not_in,omitempty"`

	// cluster
	Cluster interface{} `json:"cluster,omitempty"`

	// consistent type
	ConsistentType interface{} `json:"consistent_type,omitempty"`

	// consistent type in
	ConsistentTypeIn []ConsistentType `json:"consistent_type_in,omitempty"`

	// consistent type not
	ConsistentTypeNot interface{} `json:"consistent_type_not,omitempty"`

	// consistent type not in
	ConsistentTypeNotIn []ConsistentType `json:"consistent_type_not_in,omitempty"`

	// cpu model
	CPUModel *string `json:"cpu_model,omitempty"`

	// cpu model contains
	CPUModelContains *string `json:"cpu_model_contains,omitempty"`

	// cpu model ends with
	CPUModelEndsWith *string `json:"cpu_model_ends_with,omitempty"`

	// cpu model gt
	CPUModelGt *string `json:"cpu_model_gt,omitempty"`

	// cpu model gte
	CPUModelGte *string `json:"cpu_model_gte,omitempty"`

	// cpu model in
	CPUModelIn []string `json:"cpu_model_in,omitempty"`

	// cpu model lt
	CPUModelLt *string `json:"cpu_model_lt,omitempty"`

	// cpu model lte
	CPUModelLte *string `json:"cpu_model_lte,omitempty"`

	// cpu model not
	CPUModelNot *string `json:"cpu_model_not,omitempty"`

	// cpu model not contains
	CPUModelNotContains *string `json:"cpu_model_not_contains,omitempty"`

	// cpu model not ends with
	CPUModelNotEndsWith *string `json:"cpu_model_not_ends_with,omitempty"`

	// cpu model not in
	CPUModelNotIn []string `json:"cpu_model_not_in,omitempty"`

	// cpu model not starts with
	CPUModelNotStartsWith *string `json:"cpu_model_not_starts_with,omitempty"`

	// cpu model starts with
	CPUModelStartsWith *string `json:"cpu_model_starts_with,omitempty"`

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

	// firmware
	Firmware interface{} `json:"firmware,omitempty"`

	// firmware in
	FirmwareIn []VMFirmware `json:"firmware_in,omitempty"`

	// firmware not
	FirmwareNot interface{} `json:"firmware_not,omitempty"`

	// firmware not in
	FirmwareNotIn []VMFirmware `json:"firmware_not_in,omitempty"`

	// ha
	Ha *bool `json:"ha,omitempty"`

	// ha not
	HaNot *bool `json:"ha_not,omitempty"`

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

	// io policy
	IoPolicy interface{} `json:"io_policy,omitempty"`

	// io policy in
	IoPolicyIn []VMDiskIoPolicy `json:"io_policy_in,omitempty"`

	// io policy not
	IoPolicyNot interface{} `json:"io_policy_not,omitempty"`

	// io policy not in
	IoPolicyNotIn []VMDiskIoPolicy `json:"io_policy_not_in,omitempty"`

	// labels every
	LabelsEvery interface{} `json:"labels_every,omitempty"`

	// labels none
	LabelsNone interface{} `json:"labels_none,omitempty"`

	// labels some
	LabelsSome interface{} `json:"labels_some,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local created at gt
	LocalCreatedAtGt *string `json:"local_created_at_gt,omitempty"`

	// local created at gte
	LocalCreatedAtGte *string `json:"local_created_at_gte,omitempty"`

	// local created at in
	LocalCreatedAtIn []string `json:"local_created_at_in,omitempty"`

	// local created at lt
	LocalCreatedAtLt *string `json:"local_created_at_lt,omitempty"`

	// local created at lte
	LocalCreatedAtLte *string `json:"local_created_at_lte,omitempty"`

	// local created at not
	LocalCreatedAtNot *string `json:"local_created_at_not,omitempty"`

	// local created at not in
	LocalCreatedAtNotIn []string `json:"local_created_at_not_in,omitempty"`

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

	// max bandwidth
	MaxBandwidth *float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth gt
	MaxBandwidthGt *float64 `json:"max_bandwidth_gt,omitempty"`

	// max bandwidth gte
	MaxBandwidthGte *float64 `json:"max_bandwidth_gte,omitempty"`

	// max bandwidth in
	MaxBandwidthIn []float64 `json:"max_bandwidth_in,omitempty"`

	// max bandwidth lt
	MaxBandwidthLt *float64 `json:"max_bandwidth_lt,omitempty"`

	// max bandwidth lte
	MaxBandwidthLte *float64 `json:"max_bandwidth_lte,omitempty"`

	// max bandwidth not
	MaxBandwidthNot *float64 `json:"max_bandwidth_not,omitempty"`

	// max bandwidth not in
	MaxBandwidthNotIn []float64 `json:"max_bandwidth_not_in,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy interface{} `json:"max_bandwidth_policy,omitempty"`

	// max bandwidth policy in
	MaxBandwidthPolicyIn []VMDiskIoRestrictType `json:"max_bandwidth_policy_in,omitempty"`

	// max bandwidth policy not
	MaxBandwidthPolicyNot interface{} `json:"max_bandwidth_policy_not,omitempty"`

	// max bandwidth policy not in
	MaxBandwidthPolicyNotIn []VMDiskIoRestrictType `json:"max_bandwidth_policy_not_in,omitempty"`

	// max iops
	MaxIops *float64 `json:"max_iops,omitempty"`

	// max iops gt
	MaxIopsGt *float64 `json:"max_iops_gt,omitempty"`

	// max iops gte
	MaxIopsGte *float64 `json:"max_iops_gte,omitempty"`

	// max iops in
	MaxIopsIn []float64 `json:"max_iops_in,omitempty"`

	// max iops lt
	MaxIopsLt *float64 `json:"max_iops_lt,omitempty"`

	// max iops lte
	MaxIopsLte *float64 `json:"max_iops_lte,omitempty"`

	// max iops not
	MaxIopsNot *float64 `json:"max_iops_not,omitempty"`

	// max iops not in
	MaxIopsNotIn []float64 `json:"max_iops_not_in,omitempty"`

	// max iops policy
	MaxIopsPolicy interface{} `json:"max_iops_policy,omitempty"`

	// max iops policy in
	MaxIopsPolicyIn []VMDiskIoRestrictType `json:"max_iops_policy_in,omitempty"`

	// max iops policy not
	MaxIopsPolicyNot interface{} `json:"max_iops_policy_not,omitempty"`

	// max iops policy not in
	MaxIopsPolicyNotIn []VMDiskIoRestrictType `json:"max_iops_policy_not_in,omitempty"`

	// memory
	Memory *float64 `json:"memory,omitempty"`

	// memory gt
	MemoryGt *float64 `json:"memory_gt,omitempty"`

	// memory gte
	MemoryGte *float64 `json:"memory_gte,omitempty"`

	// memory in
	MemoryIn []float64 `json:"memory_in,omitempty"`

	// memory lt
	MemoryLt *float64 `json:"memory_lt,omitempty"`

	// memory lte
	MemoryLte *float64 `json:"memory_lte,omitempty"`

	// memory not
	MemoryNot *float64 `json:"memory_not,omitempty"`

	// memory not in
	MemoryNotIn []float64 `json:"memory_not_in,omitempty"`

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

	// size
	Size *float64 `json:"size,omitempty"`

	// size gt
	SizeGt *float64 `json:"size_gt,omitempty"`

	// size gte
	SizeGte *float64 `json:"size_gte,omitempty"`

	// size in
	SizeIn []float64 `json:"size_in,omitempty"`

	// size lt
	SizeLt *float64 `json:"size_lt,omitempty"`

	// size lte
	SizeLte *float64 `json:"size_lte,omitempty"`

	// size not
	SizeNot *float64 `json:"size_not,omitempty"`

	// size not in
	SizeNotIn []float64 `json:"size_not_in,omitempty"`

	// snapshot group
	SnapshotGroup interface{} `json:"snapshot_group,omitempty"`

	// vcpu
	Vcpu *float64 `json:"vcpu,omitempty"`

	// vcpu gt
	VcpuGt *float64 `json:"vcpu_gt,omitempty"`

	// vcpu gte
	VcpuGte *float64 `json:"vcpu_gte,omitempty"`

	// vcpu in
	VcpuIn []float64 `json:"vcpu_in,omitempty"`

	// vcpu lt
	VcpuLt *float64 `json:"vcpu_lt,omitempty"`

	// vcpu lte
	VcpuLte *float64 `json:"vcpu_lte,omitempty"`

	// vcpu not
	VcpuNot *float64 `json:"vcpu_not,omitempty"`

	// vcpu not in
	VcpuNotIn []float64 `json:"vcpu_not_in,omitempty"`

	// vm
	VM interface{} `json:"vm,omitempty"`

	// win opt
	WinOpt *bool `json:"win_opt,omitempty"`

	// win opt not
	WinOptNot *bool `json:"win_opt_not,omitempty"`
}

// Validate validates this Vm snapshot where input
func (m *VMSnapshotWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateClockOffsetIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClockOffsetNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistentTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistentTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmwareIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmwareNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicyIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicyNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicyIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicyNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicyIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicyNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMSnapshotWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *VMSnapshotWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *VMSnapshotWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *VMSnapshotWhereInput) validateClockOffsetIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ClockOffsetIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ClockOffsetIn); i++ {

		if err := m.ClockOffsetIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateClockOffsetNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ClockOffsetNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ClockOffsetNotIn); i++ {

		if err := m.ClockOffsetNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateConsistentTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistentTypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsistentTypeIn); i++ {

		if err := m.ConsistentTypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistent_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateConsistentTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistentTypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ConsistentTypeNotIn); i++ {

		if err := m.ConsistentTypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistent_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *VMSnapshotWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

func (m *VMSnapshotWhereInput) validateFirmwareIn(formats strfmt.Registry) error {
	if swag.IsZero(m.FirmwareIn) { // not required
		return nil
	}

	for i := 0; i < len(m.FirmwareIn); i++ {

		if err := m.FirmwareIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateFirmwareNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.FirmwareNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.FirmwareNotIn); i++ {

		if err := m.FirmwareNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateIoPolicyIn(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicyIn) { // not required
		return nil
	}

	for i := 0; i < len(m.IoPolicyIn); i++ {

		if err := m.IoPolicyIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateIoPolicyNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicyNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.IoPolicyNotIn); i++ {

		if err := m.IoPolicyNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateMaxBandwidthPolicyIn(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicyIn) { // not required
		return nil
	}

	for i := 0; i < len(m.MaxBandwidthPolicyIn); i++ {

		if err := m.MaxBandwidthPolicyIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateMaxBandwidthPolicyNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicyNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.MaxBandwidthPolicyNotIn); i++ {

		if err := m.MaxBandwidthPolicyNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateMaxIopsPolicyIn(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicyIn) { // not required
		return nil
	}

	for i := 0; i < len(m.MaxIopsPolicyIn); i++ {

		if err := m.MaxIopsPolicyIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) validateMaxIopsPolicyNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicyNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.MaxIopsPolicyNotIn); i++ {

		if err := m.MaxIopsPolicyNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this Vm snapshot where input based on the context it is used
func (m *VMSnapshotWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateClockOffsetIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClockOffsetNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistentTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistentTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmwareIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmwareNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicyIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicyNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicyIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicyNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicyIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicyNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMSnapshotWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMSnapshotWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMSnapshotWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMSnapshotWhereInput) contextValidateClockOffsetIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClockOffsetIn); i++ {

		if err := m.ClockOffsetIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateClockOffsetNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClockOffsetNotIn); i++ {

		if err := m.ClockOffsetNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateConsistentTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsistentTypeIn); i++ {

		if err := m.ConsistentTypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistent_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateConsistentTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConsistentTypeNotIn); i++ {

		if err := m.ConsistentTypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistent_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMSnapshotWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMSnapshotWhereInput) contextValidateFirmwareIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FirmwareIn); i++ {

		if err := m.FirmwareIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateFirmwareNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FirmwareNotIn); i++ {

		if err := m.FirmwareNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateIoPolicyIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IoPolicyIn); i++ {

		if err := m.IoPolicyIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateIoPolicyNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IoPolicyNotIn); i++ {

		if err := m.IoPolicyNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateMaxBandwidthPolicyIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MaxBandwidthPolicyIn); i++ {

		if err := m.MaxBandwidthPolicyIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateMaxBandwidthPolicyNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MaxBandwidthPolicyNotIn); i++ {

		if err := m.MaxBandwidthPolicyNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateMaxIopsPolicyIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MaxIopsPolicyIn); i++ {

		if err := m.MaxIopsPolicyIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMSnapshotWhereInput) contextValidateMaxIopsPolicyNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MaxIopsPolicyNotIn); i++ {

		if err := m.MaxIopsPolicyNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMSnapshotWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMSnapshotWhereInput) UnmarshalBinary(b []byte) error {
	var res VMSnapshotWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
