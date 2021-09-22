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

// ElfImage elf image
//
// swagger:model ElfImage
type ElfImage struct {

	// cluster
	Cluster *ElfImageCluster `json:"cluster,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// labels
	Labels []*ElfImageLabelsItems0 `json:"labels,omitempty"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// path
	// Required: true
	Path *string `json:"path"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// upload task
	UploadTask *ElfImageUploadTask `json:"upload_task,omitempty"`

	// vm disks
	VMDisks []*ElfImageVMDisksItems0 `json:"vm_disks,omitempty"`

	// vm snapshots
	VMSnapshots []*ElfImageVMSnapshotsItems0 `json:"vm_snapshots,omitempty"`

	// vm templates
	VMTemplates []*ElfImageVMTemplatesItems0 `json:"vm_templates,omitempty"`
}

// Validate validates this elf image
func (m *ElfImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadTask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImage) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
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

func (m *ElfImage) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ElfImage) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfImage) validateLabels(formats strfmt.Registry) error {
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

func (m *ElfImage) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ElfImage) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *ElfImage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ElfImage) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *ElfImage) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *ElfImage) validateUploadTask(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadTask) { // not required
		return nil
	}

	if m.UploadTask != nil {
		if err := m.UploadTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upload_task")
			}
			return err
		}
	}

	return nil
}

func (m *ElfImage) validateVMDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.VMDisks); i++ {
		if swag.IsZero(m.VMDisks[i]) { // not required
			continue
		}

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElfImage) validateVMSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.VMSnapshots); i++ {
		if swag.IsZero(m.VMSnapshots[i]) { // not required
			continue
		}

		if m.VMSnapshots[i] != nil {
			if err := m.VMSnapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElfImage) validateVMTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.VMTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.VMTemplates); i++ {
		if swag.IsZero(m.VMTemplates[i]) { // not required
			continue
		}

		if m.VMTemplates[i] != nil {
			if err := m.VMTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this elf image based on the context it is used
func (m *ElfImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUploadTask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImage) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ElfImage) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ElfImage) contextValidateUploadTask(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadTask != nil {
		if err := m.UploadTask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upload_task")
			}
			return err
		}
	}

	return nil
}

func (m *ElfImage) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMDisks); i++ {

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElfImage) contextValidateVMSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMSnapshots); i++ {

		if m.VMSnapshots[i] != nil {
			if err := m.VMSnapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElfImage) contextValidateVMTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMTemplates); i++ {

		if m.VMTemplates[i] != nil {
			if err := m.VMTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElfImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImage) UnmarshalBinary(b []byte) error {
	var res ElfImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfImageCluster elf image cluster
//
// swagger:model ElfImageCluster
type ElfImageCluster struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this elf image cluster
func (m *ElfImageCluster) Validate(formats strfmt.Registry) error {
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

func (m *ElfImageCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfImageCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf image cluster based on context it is used
func (m *ElfImageCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageCluster) UnmarshalBinary(b []byte) error {
	var res ElfImageCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfImageLabelsItems0 elf image labels items0
//
// swagger:model ElfImageLabelsItems0
type ElfImageLabelsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this elf image labels items0
func (m *ElfImageLabelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImageLabelsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf image labels items0 based on context it is used
func (m *ElfImageLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageLabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageLabelsItems0) UnmarshalBinary(b []byte) error {
	var res ElfImageLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfImageUploadTask elf image upload task
//
// swagger:model ElfImageUploadTask
type ElfImageUploadTask struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this elf image upload task
func (m *ElfImageUploadTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImageUploadTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("upload_task"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf image upload task based on context it is used
func (m *ElfImageUploadTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageUploadTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageUploadTask) UnmarshalBinary(b []byte) error {
	var res ElfImageUploadTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfImageVMDisksItems0 elf image VM disks items0
//
// swagger:model ElfImageVMDisksItems0
type ElfImageVMDisksItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this elf image VM disks items0
func (m *ElfImageVMDisksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfImageVMDisksItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf image VM disks items0 based on context it is used
func (m *ElfImageVMDisksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageVMDisksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageVMDisksItems0) UnmarshalBinary(b []byte) error {
	var res ElfImageVMDisksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfImageVMSnapshotsItems0 elf image VM snapshots items0
//
// swagger:model ElfImageVMSnapshotsItems0
type ElfImageVMSnapshotsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this elf image VM snapshots items0
func (m *ElfImageVMSnapshotsItems0) Validate(formats strfmt.Registry) error {
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

func (m *ElfImageVMSnapshotsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfImageVMSnapshotsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf image VM snapshots items0 based on context it is used
func (m *ElfImageVMSnapshotsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageVMSnapshotsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageVMSnapshotsItems0) UnmarshalBinary(b []byte) error {
	var res ElfImageVMSnapshotsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ElfImageVMTemplatesItems0 elf image VM templates items0
//
// swagger:model ElfImageVMTemplatesItems0
type ElfImageVMTemplatesItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this elf image VM templates items0
func (m *ElfImageVMTemplatesItems0) Validate(formats strfmt.Registry) error {
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

func (m *ElfImageVMTemplatesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfImageVMTemplatesItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elf image VM templates items0 based on context it is used
func (m *ElfImageVMTemplatesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElfImageVMTemplatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfImageVMTemplatesItems0) UnmarshalBinary(b []byte) error {
	var res ElfImageVMTemplatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}