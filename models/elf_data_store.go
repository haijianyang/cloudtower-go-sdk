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

// ElfDataStore elf data store
//
// swagger:model ElfDataStore
type ElfDataStore struct {

	// cluster
	// Required: true
	Cluster *ElfDataStoreCluster `json:"cluster"`

	// description
	// Required: true
	Description *string `json:"description"`

	// external use
	// Required: true
	ExternalUse *bool `json:"external_use"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// ip whitelist
	// Required: true
	IPWhitelist *string `json:"ip_whitelist"`

	// iscsi target
	IscsiTarget *ElfDataStoreIscsiTarget `json:"iscsi_target,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nfs export
	NfsExport *ElfDataStoreNfsExport `json:"nfs_export,omitempty"`

	// nvmf subsystem
	NvmfSubsystem *ElfDataStoreNvmfSubsystem `json:"nvmf_subsystem,omitempty"`

	// replica num
	// Required: true
	ReplicaNum *float64 `json:"replica_num"`

	// thin provision
	// Required: true
	ThinProvision *bool `json:"thin_provision"`

	// type
	// Required: true
	Type *ElfDataStoreType `json:"type"`
}

// Validate validates this elf data store
func (m *ElfDataStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThinProvision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfDataStore) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateExternalUse(formats strfmt.Registry) error {

	if err := validate.Required("external_use", "body", m.ExternalUse); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateIPWhitelist(formats strfmt.Registry) error {

	if err := validate.Required("ip_whitelist", "body", m.IPWhitelist); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateIscsiTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiTarget) { // not required
		return nil
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

func (m *ElfDataStore) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateNfsExport(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsExport) { // not required
		return nil
	}

	if m.NfsExport != nil {
		if err := m.NfsExport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs_export")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) validateNvmfSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.NvmfSubsystem) { // not required
		return nil
	}

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) validateReplicaNum(formats strfmt.Registry) error {

	if err := validate.Required("replica_num", "body", m.ReplicaNum); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateThinProvision(formats strfmt.Registry) error {

	if err := validate.Required("thin_provision", "body", m.ThinProvision); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this elf data store based on the context it is used
func (m *ElfDataStore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsExport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfDataStore) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) contextValidateIscsiTarget(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ElfDataStore) contextValidateNfsExport(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsExport != nil {
		if err := m.NfsExport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs_export")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) contextValidateNvmfSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElfDataStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfDataStore) UnmarshalBinary(b []byte) error {
	var res ElfDataStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfDataStoreCluster elf data store cluster
//
// swagger:model ElfDataStoreCluster
type ElfDataStoreCluster struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this elf data store cluster
func (m *ElfDataStoreCluster) Validate(formats strfmt.Registry) error {
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

func (m *ElfDataStoreCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStoreCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf data store cluster based on context it is used
func (m *ElfDataStoreCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfDataStoreCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfDataStoreCluster) UnmarshalBinary(b []byte) error {
	var res ElfDataStoreCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfDataStoreIscsiTarget elf data store iscsi target
//
// swagger:model ElfDataStoreIscsiTarget
type ElfDataStoreIscsiTarget struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this elf data store iscsi target
func (m *ElfDataStoreIscsiTarget) Validate(formats strfmt.Registry) error {
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

func (m *ElfDataStoreIscsiTarget) validateID(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_target"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStoreIscsiTarget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_target"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf data store iscsi target based on context it is used
func (m *ElfDataStoreIscsiTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfDataStoreIscsiTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfDataStoreIscsiTarget) UnmarshalBinary(b []byte) error {
	var res ElfDataStoreIscsiTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfDataStoreNfsExport elf data store nfs export
//
// swagger:model ElfDataStoreNfsExport
type ElfDataStoreNfsExport struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this elf data store nfs export
func (m *ElfDataStoreNfsExport) Validate(formats strfmt.Registry) error {
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

func (m *ElfDataStoreNfsExport) validateID(formats strfmt.Registry) error {

	if err := validate.Required("nfs_export"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStoreNfsExport) validateName(formats strfmt.Registry) error {

	if err := validate.Required("nfs_export"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf data store nfs export based on context it is used
func (m *ElfDataStoreNfsExport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfDataStoreNfsExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfDataStoreNfsExport) UnmarshalBinary(b []byte) error {
	var res ElfDataStoreNfsExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfDataStoreNvmfSubsystem elf data store nvmf subsystem
//
// swagger:model ElfDataStoreNvmfSubsystem
type ElfDataStoreNvmfSubsystem struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this elf data store nvmf subsystem
func (m *ElfDataStoreNvmfSubsystem) Validate(formats strfmt.Registry) error {
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

func (m *ElfDataStoreNvmfSubsystem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStoreNvmfSubsystem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf data store nvmf subsystem based on context it is used
func (m *ElfDataStoreNvmfSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfDataStoreNvmfSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfDataStoreNvmfSubsystem) UnmarshalBinary(b []byte) error {
	var res ElfDataStoreNvmfSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
