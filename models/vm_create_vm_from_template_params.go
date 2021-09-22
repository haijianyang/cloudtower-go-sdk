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

// VMCreateVMFromTemplateParams Vm create Vm from template params
//
// swagger:model VmCreateVmFromTemplateParams
type VMCreateVMFromTemplateParams struct {

	// cloud init
	CloudInit *VMCreateVMFromTemplateParamsCloudInit `json:"cloud_init,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// cpu cores
	CPUCores float64 `json:"cpu_cores,omitempty"`

	// cpu sockets
	CPUSockets float64 `json:"cpu_sockets,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// disk operate
	DiskOperate *VMCreateVMFromTemplateParamsDiskOperate `json:"disk_operate,omitempty"`

	// firmware
	Firmware VMFirmware `json:"firmware,omitempty"`

	// folder id
	FolderID string `json:"folder_id,omitempty"`

	// guest os type
	GuestOsType VMGuestsOperationSystem `json:"guest_os_type,omitempty"`

	// ha
	Ha bool `json:"ha,omitempty"`

	// host id
	HostID string `json:"host_id,omitempty"`

	// io policy
	IoPolicy VMDiskIoPolicy `json:"io_policy,omitempty"`

	// is full copy
	// Required: true
	IsFullCopy *bool `json:"is_full_copy"`

	// max bandwidth
	MaxBandwidth float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops float64 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// memory
	Memory float64 `json:"memory,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// status
	Status VMStatus `json:"status,omitempty"`

	// template id
	// Required: true
	TemplateID *string `json:"template_id"`

	// vcpu
	Vcpu float64 `json:"vcpu,omitempty"`

	// vm nics
	VMNics VMNicParams `json:"vm_nics,omitempty"`
}

// Validate validates this Vm create Vm from template params
func (m *VMCreateVMFromTemplateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudInit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskOperate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsFullCopy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParams) validateCloudInit(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudInit) { // not required
		return nil
	}

	if m.CloudInit != nil {
		if err := m.CloudInit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_init")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateDiskOperate(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskOperate) { // not required
		return nil
	}

	if m.DiskOperate != nil {
		if err := m.DiskOperate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateFirmware(formats strfmt.Registry) error {
	if swag.IsZero(m.Firmware) { // not required
		return nil
	}

	if err := m.Firmware.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("firmware")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateGuestOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.GuestOsType) { // not required
		return nil
	}

	if err := m.GuestOsType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("guest_os_type")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateIoPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicy) { // not required
		return nil
	}

	if err := m.IoPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("io_policy")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateIsFullCopy(formats strfmt.Registry) error {

	if err := validate.Required("is_full_copy", "body", m.IsFullCopy); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicy) { // not required
		return nil
	}

	if err := m.MaxBandwidthPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_bandwidth_policy")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateMaxIopsPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicy) { // not required
		return nil
	}

	if err := m.MaxIopsPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_iops_policy")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("template_id", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	if err := m.VMNics.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vm_nics")
		}
		return err
	}

	return nil
}

// ContextValidate validate this Vm create Vm from template params based on the context it is used
func (m *VMCreateVMFromTemplateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudInit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiskOperate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestOsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
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

func (m *VMCreateVMFromTemplateParams) contextValidateCloudInit(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudInit != nil {
		if err := m.CloudInit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_init")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateDiskOperate(ctx context.Context, formats strfmt.Registry) error {

	if m.DiskOperate != nil {
		if err := m.DiskOperate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("firmware")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateGuestOsType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.GuestOsType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("guest_os_type")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IoPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("io_policy")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaxBandwidthPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_bandwidth_policy")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaxIopsPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_iops_policy")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParams) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VMNics.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vm_nics")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParams) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromTemplateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMCreateVMFromTemplateParamsCloudInit VM create VM from template params cloud init
//
// swagger:model VMCreateVMFromTemplateParamsCloudInit
type VMCreateVMFromTemplateParamsCloudInit struct {

	// default user password
	DefaultUserPassword string `json:"default_user_password,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// nameservers
	Nameservers []string `json:"nameservers,omitempty"`

	// networks
	Networks []*VMCreateVMFromTemplateParamsCloudInitNetworksItems0 `json:"networks,omitempty"`

	// public keys
	PublicKeys []string `json:"public_keys,omitempty"`

	// user data
	UserData string `json:"user_data,omitempty"`
}

// Validate validates this VM create VM from template params cloud init
func (m *VMCreateVMFromTemplateParamsCloudInit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInit) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloud_init" + "." + "networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this VM create VM from template params cloud init based on the context it is used
func (m *VMCreateVMFromTemplateParamsCloudInit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInit) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {
			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloud_init" + "." + "networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsCloudInit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsCloudInit) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromTemplateParamsCloudInit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMCreateVMFromTemplateParamsCloudInitNetworksItems0 VM create VM from template params cloud init networks items0
//
// swagger:model VMCreateVMFromTemplateParamsCloudInitNetworksItems0
type VMCreateVMFromTemplateParamsCloudInitNetworksItems0 struct {

	// ip address
	IPAddress string `json:"ip_address,omitempty"`

	// netmask
	Netmask string `json:"netmask,omitempty"`

	// nic index
	// Required: true
	NicIndex *float64 `json:"nic_index"`

	// routes
	Routes []*VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0 `json:"routes,omitempty"`

	// type
	// Required: true
	Type *CloudInitNetworkTypeEnum `json:"type"`
}

// Validate validates this VM create VM from template params cloud init networks items0
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNicIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
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

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) validateNicIndex(formats strfmt.Registry) error {

	if err := validate.Required("nic_index", "body", m.NicIndex); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) validateRoutes(formats strfmt.Registry) error {
	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) validateType(formats strfmt.Registry) error {

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

// ContextValidate validate this VM create VM from template params cloud init networks items0 based on the context it is used
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoutes(ctx, formats); err != nil {
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

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) contextValidateRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Routes); i++ {

		if m.Routes[i] != nil {
			if err := m.Routes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromTemplateParamsCloudInitNetworksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0 VM create VM from template params cloud init networks items0 routes items0
//
// swagger:model VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0
type VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0 struct {

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// netmask
	// Required: true
	Netmask *string `json:"netmask"`

	// network
	// Required: true
	Network *string `json:"network"`
}

// Validate validates this VM create VM from template params cloud init networks items0 routes items0
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0) validateNetmask(formats strfmt.Registry) error {

	if err := validate.Required("netmask", "body", m.Netmask); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("network", "body", m.Network); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM create VM from template params cloud init networks items0 routes items0 based on context it is used
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromTemplateParamsCloudInitNetworksItems0RoutesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMCreateVMFromTemplateParamsDiskOperate VM create VM from template params disk operate
//
// swagger:model VMCreateVMFromTemplateParamsDiskOperate
type VMCreateVMFromTemplateParamsDiskOperate struct {

	// modify disks
	ModifyDisks []*VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0 `json:"modify_disks,omitempty"`

	// new disks
	NewDisks *VMDiskParams `json:"new_disks,omitempty"`

	// remove disks
	RemoveDisks *VMCreateVMFromTemplateParamsDiskOperateRemoveDisks `json:"remove_disks,omitempty"`
}

// Validate validates this VM create VM from template params disk operate
func (m *VMCreateVMFromTemplateParamsDiskOperate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifyDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoveDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperate) validateModifyDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifyDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.ModifyDisks); i++ {
		if swag.IsZero(m.ModifyDisks[i]) { // not required
			continue
		}

		if m.ModifyDisks[i] != nil {
			if err := m.ModifyDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disk_operate" + "." + "modify_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperate) validateNewDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.NewDisks) { // not required
		return nil
	}

	if m.NewDisks != nil {
		if err := m.NewDisks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate" + "." + "new_disks")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperate) validateRemoveDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoveDisks) { // not required
		return nil
	}

	if m.RemoveDisks != nil {
		if err := m.RemoveDisks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate" + "." + "remove_disks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this VM create VM from template params disk operate based on the context it is used
func (m *VMCreateVMFromTemplateParamsDiskOperate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModifyDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoveDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperate) contextValidateModifyDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModifyDisks); i++ {

		if m.ModifyDisks[i] != nil {
			if err := m.ModifyDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disk_operate" + "." + "modify_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperate) contextValidateNewDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.NewDisks != nil {
		if err := m.NewDisks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate" + "." + "new_disks")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperate) contextValidateRemoveDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.RemoveDisks != nil {
		if err := m.RemoveDisks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate" + "." + "remove_disks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsDiskOperate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsDiskOperate) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromTemplateParamsDiskOperate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0 VM create VM from template params disk operate modify disks items0
//
// swagger:model VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0
type VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0 struct {

	// bus
	Bus Bus `json:"bus,omitempty"`

	// disk index
	// Required: true
	DiskIndex *float64 `json:"disk_index,omitempty"`

	// vm volume id
	VMVolumeID string `json:"vm_volume_id,omitempty"`
}

// Validate validates this VM create VM from template params disk operate modify disks items0
func (m *VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0) validateBus(formats strfmt.Registry) error {
	if swag.IsZero(m.Bus) { // not required
		return nil
	}

	if err := m.Bus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bus")
		}
		return err
	}

	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0) validateDiskIndex(formats strfmt.Registry) error {

	if err := validate.Required("disk_index", "body", m.DiskIndex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this VM create VM from template params disk operate modify disks items0 based on the context it is used
func (m *VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Bus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bus")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromTemplateParamsDiskOperateModifyDisksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMCreateVMFromTemplateParamsDiskOperateRemoveDisks VM create VM from template params disk operate remove disks
//
// swagger:model VMCreateVMFromTemplateParamsDiskOperateRemoveDisks
type VMCreateVMFromTemplateParamsDiskOperateRemoveDisks struct {

	// disk index
	// Required: true
	DiskIndex []float64 `json:"disk_index"`
}

// Validate validates this VM create VM from template params disk operate remove disks
func (m *VMCreateVMFromTemplateParamsDiskOperateRemoveDisks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromTemplateParamsDiskOperateRemoveDisks) validateDiskIndex(formats strfmt.Registry) error {

	if err := validate.Required("disk_operate"+"."+"remove_disks"+"."+"disk_index", "body", m.DiskIndex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM create VM from template params disk operate remove disks based on context it is used
func (m *VMCreateVMFromTemplateParamsDiskOperateRemoveDisks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsDiskOperateRemoveDisks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromTemplateParamsDiskOperateRemoveDisks) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromTemplateParamsDiskOperateRemoveDisks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
