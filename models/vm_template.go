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

// VMTemplate Vm template
//
// swagger:model VmTemplate
type VMTemplate struct {

	// clock offset
	// Required: true
	ClockOffset *VMClockOffset `json:"clock_offset"`

	// cloud init supported
	// Required: true
	CloudInitSupported *bool `json:"cloud_init_supported"`

	// cluster
	// Required: true
	Cluster *VMTemplateCluster `json:"cluster"`

	// cpu
	// Required: true
	CPU *VMTemplateCPU `json:"cpu"`

	// cpu model
	// Required: true
	CPUModel *string `json:"cpu_model"`

	// description
	// Required: true
	Description *string `json:"description"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// firmware
	// Required: true
	Firmware *VMFirmware `json:"firmware"`

	// ha
	// Required: true
	Ha *bool `json:"ha"`

	// id
	// Required: true
	ID *string `json:"id"`

	// io policy
	IoPolicy interface{} `json:"io_policy,omitempty"`

	// labels
	Labels []*VMTemplateLabelsItems0 `json:"labels,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// max bandwidth
	MaxBandwidth *float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy interface{} `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *float64 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy interface{} `json:"max_iops_policy,omitempty"`

	// memory
	// Required: true
	Memory *float64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// vcpu
	// Required: true
	Vcpu *float64 `json:"vcpu"`

	// video type
	VideoType *string `json:"video_type,omitempty"`

	// vm disks
	VMDisks []*VMTemplateVMDisksItems0 `json:"vm_disks,omitempty"`

	// vm nics
	VMNics []*VMTemplateVMNicsItems0 `json:"vm_nics,omitempty"`

	// win opt
	// Required: true
	WinOpt *bool `json:"win_opt"`
}

// Validate validates this Vm template
func (m *VMTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClockOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudInitSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHa(formats); err != nil {
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

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWinOpt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplate) validateClockOffset(formats strfmt.Registry) error {

	if err := validate.Required("clock_offset", "body", m.ClockOffset); err != nil {
		return err
	}

	if err := validate.Required("clock_offset", "body", m.ClockOffset); err != nil {
		return err
	}

	if m.ClockOffset != nil {
		if err := m.ClockOffset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCloudInitSupported(formats strfmt.Registry) error {

	if err := validate.Required("cloud_init_supported", "body", m.CloudInitSupported); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateCluster(formats strfmt.Registry) error {

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

func (m *VMTemplate) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCPUModel(formats strfmt.Registry) error {

	if err := validate.Required("cpu_model", "body", m.CPUModel); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateFirmware(formats strfmt.Registry) error {

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateHa(formats strfmt.Registry) error {

	if err := validate.Required("ha", "body", m.Ha); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateLabels(formats strfmt.Registry) error {
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

func (m *VMTemplate) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateVcpu(formats strfmt.Registry) error {

	if err := validate.Required("vcpu", "body", m.Vcpu); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateVMDisks(formats strfmt.Registry) error {
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

func (m *VMTemplate) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	for i := 0; i < len(m.VMNics); i++ {
		if swag.IsZero(m.VMNics[i]) { // not required
			continue
		}

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) validateWinOpt(formats strfmt.Registry) error {

	if err := validate.Required("win_opt", "body", m.WinOpt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Vm template based on the context it is used
func (m *VMTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClockOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplate) contextValidateClockOffset(ctx context.Context, formats strfmt.Registry) error {

	if m.ClockOffset != nil {
		if err := m.ClockOffset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMTemplate) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {
		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if m.Firmware != nil {
		if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMTemplate) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMTemplate) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMNics); i++ {

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplate) UnmarshalBinary(b []byte) error {
	var res VMTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMTemplateCPU VM template CPU
//
// swagger:model VMTemplateCPU
type VMTemplateCPU struct {

	// cores
	// Required: true
	Cores *float64 `json:"cores"`

	// sockets
	// Required: true
	Sockets *float64 `json:"sockets"`
}

// Validate validates this VM template CPU
func (m *VMTemplateCPU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSockets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateCPU) validateCores(formats strfmt.Registry) error {

	if err := validate.Required("cpu"+"."+"cores", "body", m.Cores); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateCPU) validateSockets(formats strfmt.Registry) error {

	if err := validate.Required("cpu"+"."+"sockets", "body", m.Sockets); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM template CPU based on context it is used
func (m *VMTemplateCPU) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplateCPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateCPU) UnmarshalBinary(b []byte) error {
	var res VMTemplateCPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMTemplateCluster VM template cluster
//
// swagger:model VMTemplateCluster
type VMTemplateCluster struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this VM template cluster
func (m *VMTemplateCluster) Validate(formats strfmt.Registry) error {
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

func (m *VMTemplateCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM template cluster based on context it is used
func (m *VMTemplateCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplateCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateCluster) UnmarshalBinary(b []byte) error {
	var res VMTemplateCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMTemplateLabelsItems0 VM template labels items0
//
// swagger:model VMTemplateLabelsItems0
type VMTemplateLabelsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this VM template labels items0
func (m *VMTemplateLabelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateLabelsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM template labels items0 based on context it is used
func (m *VMTemplateLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplateLabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateLabelsItems0) UnmarshalBinary(b []byte) error {
	var res VMTemplateLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMTemplateVMDisksItems0 VM template VM disks items0
//
// swagger:model VMTemplateVMDisksItems0
type VMTemplateVMDisksItems0 struct {

	// boot
	// Required: true
	Boot *float64 `json:"boot"`

	// bus
	// Required: true
	Bus *Bus `json:"bus"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// disk name
	DiskName *string `json:"disk_name,omitempty"`

	// elf image local id
	// Required: true
	ElfImageLocalID *string `json:"elf_image_local_id"`

	// image name
	ImageName *string `json:"image_name,omitempty"`

	// index
	// Required: true
	Index *float64 `json:"index"`

	// key
	Key *float64 `json:"key,omitempty"`

	// max bandwidth
	MaxBandwidth *float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy interface{} `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *float64 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy interface{} `json:"max_iops_policy,omitempty"`

	// path
	// Required: true
	Path *string `json:"path"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// snapshot local id
	SnapshotLocalID *string `json:"snapshot_local_id,omitempty"`

	// storage policy uuid
	// Required: true
	StoragePolicyUUID *string `json:"storage_policy_uuid"`

	// svt image local id
	// Required: true
	SvtImageLocalID *string `json:"svt_image_local_id"`

	// type
	// Required: true
	Type *VMDiskType `json:"type"`

	// vm volume local id
	// Required: true
	VMVolumeLocalID *string `json:"vm_volume_local_id"`
}

// Validate validates this VM template VM disks items0
func (m *VMTemplateVMDisksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfImageLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicyUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvtImageLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVolumeLocalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateVMDisksItems0) validateBoot(formats strfmt.Registry) error {

	if err := validate.Required("boot", "body", m.Boot); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validateBus(formats strfmt.Registry) error {

	if err := validate.Required("bus", "body", m.Bus); err != nil {
		return err
	}

	if err := validate.Required("bus", "body", m.Bus); err != nil {
		return err
	}

	if m.Bus != nil {
		if err := m.Bus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bus")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validateElfImageLocalID(formats strfmt.Registry) error {

	if err := validate.Required("elf_image_local_id", "body", m.ElfImageLocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validateStoragePolicyUUID(formats strfmt.Registry) error {

	if err := validate.Required("storage_policy_uuid", "body", m.StoragePolicyUUID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validateSvtImageLocalID(formats strfmt.Registry) error {

	if err := validate.Required("svt_image_local_id", "body", m.SvtImageLocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) validateType(formats strfmt.Registry) error {

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

func (m *VMTemplateVMDisksItems0) validateVMVolumeLocalID(formats strfmt.Registry) error {

	if err := validate.Required("vm_volume_local_id", "body", m.VMVolumeLocalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this VM template VM disks items0 based on the context it is used
func (m *VMTemplateVMDisksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
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

func (m *VMTemplateVMDisksItems0) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

	if m.Bus != nil {
		if err := m.Bus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bus")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplateVMDisksItems0) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMTemplateVMDisksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateVMDisksItems0) UnmarshalBinary(b []byte) error {
	var res VMTemplateVMDisksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMTemplateVMNicsItems0 VM template VM nics items0
//
// swagger:model VMTemplateVMNicsItems0
type VMTemplateVMNicsItems0 struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// index
	// Required: true
	Index *float64 `json:"index"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// mac address
	MacAddress *string `json:"mac_address,omitempty"`

	// mirror
	Mirror *bool `json:"mirror,omitempty"`

	// model
	Model interface{} `json:"model,omitempty"`

	// vlan
	// Required: true
	Vlan *VMTemplateVMNicsItems0Vlan `json:"vlan"`
}

// Validate validates this VM template VM nics items0
func (m *VMTemplateVMNicsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateVMNicsItems0) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMNicsItems0) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("vlan", "body", m.Vlan); err != nil {
		return err
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

// ContextValidate validate this VM template VM nics items0 based on the context it is used
func (m *VMTemplateVMNicsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateVMNicsItems0) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *VMTemplateVMNicsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateVMNicsItems0) UnmarshalBinary(b []byte) error {
	var res VMTemplateVMNicsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMTemplateVMNicsItems0Vlan VM template VM nics items0 vlan
//
// swagger:model VMTemplateVMNicsItems0Vlan
type VMTemplateVMNicsItems0Vlan struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// vds ovs
	// Required: true
	VdsOvs *string `json:"vds_ovs"`

	// vlan id
	// Required: true
	VlanID *float64 `json:"vlan_id"`

	// vlan local id
	// Required: true
	VlanLocalID *string `json:"vlan_local_id"`
}

// Validate validates this VM template VM nics items0 vlan
func (m *VMTemplateVMNicsItems0Vlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdsOvs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanLocalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateVMNicsItems0Vlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMNicsItems0Vlan) validateVdsOvs(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"vds_ovs", "body", m.VdsOvs); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMNicsItems0Vlan) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateVMNicsItems0Vlan) validateVlanLocalID(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"vlan_local_id", "body", m.VlanLocalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM template VM nics items0 vlan based on context it is used
func (m *VMTemplateVMNicsItems0Vlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplateVMNicsItems0Vlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateVMNicsItems0Vlan) UnmarshalBinary(b []byte) error {
	var res VMTemplateVMNicsItems0Vlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
