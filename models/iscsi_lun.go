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

// IscsiLun iscsi lun
//
// swagger:model IscsiLun
type IscsiLun struct {

	// allowed initiators
	// Required: true
	AllowedInitiators *string `json:"allowed_initiators"`

	// assigned size
	// Required: true
	AssignedSize *float64 `json:"assigned_size"`

	// bps
	// Required: true
	Bps *float64 `json:"bps"`

	// bps max
	// Required: true
	BpsMax *float64 `json:"bps_max"`

	// bps max length
	// Required: true
	BpsMaxLength *float64 `json:"bps_max_length"`

	// bps rd
	// Required: true
	BpsRd *float64 `json:"bps_rd"`

	// bps rd max
	// Required: true
	BpsRdMax *float64 `json:"bps_rd_max"`

	// bps rd max length
	// Required: true
	BpsRdMaxLength *float64 `json:"bps_rd_max_length"`

	// bps wr
	// Required: true
	BpsWr *float64 `json:"bps_wr"`

	// bps wr max
	// Required: true
	BpsWrMax *float64 `json:"bps_wr_max"`

	// bps wr max length
	// Required: true
	BpsWrMaxLength *float64 `json:"bps_wr_max_length"`

	// consistency group
	ConsistencyGroup *IscsiLunConsistencyGroup `json:"consistency_group,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// io size
	// Required: true
	IoSize *float64 `json:"io_size"`

	// iops
	// Required: true
	Iops *float64 `json:"iops"`

	// iops max
	// Required: true
	IopsMax *float64 `json:"iops_max"`

	// iops max length
	// Required: true
	IopsMaxLength *float64 `json:"iops_max_length"`

	// iops rd
	// Required: true
	IopsRd *float64 `json:"iops_rd"`

	// iops rd max
	// Required: true
	IopsRdMax *float64 `json:"iops_rd_max"`

	// iops rd max length
	// Required: true
	IopsRdMaxLength *float64 `json:"iops_rd_max_length"`

	// iops wr
	// Required: true
	IopsWr *float64 `json:"iops_wr"`

	// iops wr max
	// Required: true
	IopsWrMax *float64 `json:"iops_wr_max"`

	// iops wr max length
	// Required: true
	IopsWrMaxLength *float64 `json:"iops_wr_max_length"`

	// iscsi target
	// Required: true
	IscsiTarget *IscsiLunIscsiTarget `json:"iscsi_target"`

	// labels
	Labels []*IscsiLunLabelsItems0 `json:"labels,omitempty"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// lun id
	// Required: true
	LunID *float64 `json:"lun_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// replica num
	// Required: true
	ReplicaNum *float64 `json:"replica_num"`

	// shared size
	// Required: true
	SharedSize *float64 `json:"shared_size"`

	// snapshot num
	// Required: true
	SnapshotNum *float64 `json:"snapshot_num"`

	// stripe num
	// Required: true
	StripeNum *float64 `json:"stripe_num"`

	// stripe size
	// Required: true
	StripeSize *float64 `json:"stripe_size"`

	// thin provision
	// Required: true
	ThinProvision *bool `json:"thin_provision"`

	// unique size
	// Required: true
	UniqueSize *float64 `json:"unique_size"`

	// zbs volume id
	// Required: true
	ZbsVolumeID *string `json:"zbs_volume_id"`
}

// Validate validates this iscsi lun
func (m *IscsiLun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedInitiators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRdMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRdMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWrMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWrMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsRd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsRdMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsRdMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsWr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsWrMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsWrMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLunID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThinProvision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniqueSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZbsVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLun) validateAllowedInitiators(formats strfmt.Registry) error {

	if err := validate.Required("allowed_initiators", "body", m.AllowedInitiators); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateAssignedSize(formats strfmt.Registry) error {

	if err := validate.Required("assigned_size", "body", m.AssignedSize); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBps(formats strfmt.Registry) error {

	if err := validate.Required("bps", "body", m.Bps); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsMax(formats strfmt.Registry) error {

	if err := validate.Required("bps_max", "body", m.BpsMax); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("bps_max_length", "body", m.BpsMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsRd(formats strfmt.Registry) error {

	if err := validate.Required("bps_rd", "body", m.BpsRd); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsRdMax(formats strfmt.Registry) error {

	if err := validate.Required("bps_rd_max", "body", m.BpsRdMax); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsRdMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("bps_rd_max_length", "body", m.BpsRdMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsWr(formats strfmt.Registry) error {

	if err := validate.Required("bps_wr", "body", m.BpsWr); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsWrMax(formats strfmt.Registry) error {

	if err := validate.Required("bps_wr_max", "body", m.BpsWrMax); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateBpsWrMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("bps_wr_max_length", "body", m.BpsWrMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateConsistencyGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroup) { // not required
		return nil
	}

	if m.ConsistencyGroup != nil {
		if err := m.ConsistencyGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLun) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIoSize(formats strfmt.Registry) error {

	if err := validate.Required("io_size", "body", m.IoSize); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIops(formats strfmt.Registry) error {

	if err := validate.Required("iops", "body", m.Iops); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsMax(formats strfmt.Registry) error {

	if err := validate.Required("iops_max", "body", m.IopsMax); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("iops_max_length", "body", m.IopsMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsRd(formats strfmt.Registry) error {

	if err := validate.Required("iops_rd", "body", m.IopsRd); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsRdMax(formats strfmt.Registry) error {

	if err := validate.Required("iops_rd_max", "body", m.IopsRdMax); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsRdMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("iops_rd_max_length", "body", m.IopsRdMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsWr(formats strfmt.Registry) error {

	if err := validate.Required("iops_wr", "body", m.IopsWr); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsWrMax(formats strfmt.Registry) error {

	if err := validate.Required("iops_wr_max", "body", m.IopsWrMax); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIopsWrMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("iops_wr_max_length", "body", m.IopsWrMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateIscsiTarget(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_target", "body", m.IscsiTarget); err != nil {
		return err
	}

	if m.IscsiTarget != nil {
		if err := m.IscsiTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi_target")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLun) validateLabels(formats strfmt.Registry) error {
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

func (m *IscsiLun) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateLunID(formats strfmt.Registry) error {

	if err := validate.Required("lun_id", "body", m.LunID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateReplicaNum(formats strfmt.Registry) error {

	if err := validate.Required("replica_num", "body", m.ReplicaNum); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateSharedSize(formats strfmt.Registry) error {

	if err := validate.Required("shared_size", "body", m.SharedSize); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateSnapshotNum(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_num", "body", m.SnapshotNum); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateStripeNum(formats strfmt.Registry) error {

	if err := validate.Required("stripe_num", "body", m.StripeNum); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateStripeSize(formats strfmt.Registry) error {

	if err := validate.Required("stripe_size", "body", m.StripeSize); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateThinProvision(formats strfmt.Registry) error {

	if err := validate.Required("thin_provision", "body", m.ThinProvision); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateUniqueSize(formats strfmt.Registry) error {

	if err := validate.Required("unique_size", "body", m.UniqueSize); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLun) validateZbsVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("zbs_volume_id", "body", m.ZbsVolumeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this iscsi lun based on the context it is used
func (m *IscsiLun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsistencyGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLun) contextValidateConsistencyGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsistencyGroup != nil {
		if err := m.ConsistencyGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLun) contextValidateIscsiTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.IscsiTarget != nil {
		if err := m.IscsiTarget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi_target")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiLun) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IscsiLun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLun) UnmarshalBinary(b []byte) error {
	var res IscsiLun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiLunConsistencyGroup iscsi lun consistency group
//
// swagger:model IscsiLunConsistencyGroup
type IscsiLunConsistencyGroup struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this iscsi lun consistency group
func (m *IscsiLunConsistencyGroup) Validate(formats strfmt.Registry) error {
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

func (m *IscsiLunConsistencyGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("consistency_group"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLunConsistencyGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("consistency_group"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi lun consistency group based on context it is used
func (m *IscsiLunConsistencyGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiLunConsistencyGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLunConsistencyGroup) UnmarshalBinary(b []byte) error {
	var res IscsiLunConsistencyGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiLunIscsiTarget iscsi lun iscsi target
//
// swagger:model IscsiLunIscsiTarget
type IscsiLunIscsiTarget struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this iscsi lun iscsi target
func (m *IscsiLunIscsiTarget) Validate(formats strfmt.Registry) error {
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

func (m *IscsiLunIscsiTarget) validateID(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_target"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLunIscsiTarget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_target"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi lun iscsi target based on context it is used
func (m *IscsiLunIscsiTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiLunIscsiTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLunIscsiTarget) UnmarshalBinary(b []byte) error {
	var res IscsiLunIscsiTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiLunLabelsItems0 iscsi lun labels items0
//
// swagger:model IscsiLunLabelsItems0
type IscsiLunLabelsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this iscsi lun labels items0
func (m *IscsiLunLabelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLunLabelsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi lun labels items0 based on context it is used
func (m *IscsiLunLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiLunLabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLunLabelsItems0) UnmarshalBinary(b []byte) error {
	var res IscsiLunLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
