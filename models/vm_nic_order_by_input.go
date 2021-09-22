// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VMNicOrderByInput Vm nic order by input
//
// swagger:model VmNicOrderByInput
type VMNicOrderByInput string

func NewVMNicOrderByInput(value VMNicOrderByInput) *VMNicOrderByInput {
	v := value
	return &v
}

const (

	// VMNicOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	VMNicOrderByInputCreatedAtASC VMNicOrderByInput = "createdAt_ASC"

	// VMNicOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	VMNicOrderByInputCreatedAtDESC VMNicOrderByInput = "createdAt_DESC"

	// VMNicOrderByInputEnabledASC captures enum value "enabled_ASC"
	VMNicOrderByInputEnabledASC VMNicOrderByInput = "enabled_ASC"

	// VMNicOrderByInputEnabledDESC captures enum value "enabled_DESC"
	VMNicOrderByInputEnabledDESC VMNicOrderByInput = "enabled_DESC"

	// VMNicOrderByInputGatewayASC captures enum value "gateway_ASC"
	VMNicOrderByInputGatewayASC VMNicOrderByInput = "gateway_ASC"

	// VMNicOrderByInputGatewayDESC captures enum value "gateway_DESC"
	VMNicOrderByInputGatewayDESC VMNicOrderByInput = "gateway_DESC"

	// VMNicOrderByInputIDASC captures enum value "id_ASC"
	VMNicOrderByInputIDASC VMNicOrderByInput = "id_ASC"

	// VMNicOrderByInputIDDESC captures enum value "id_DESC"
	VMNicOrderByInputIDDESC VMNicOrderByInput = "id_DESC"

	// VMNicOrderByInputInterfaceIDASC captures enum value "interface_id_ASC"
	VMNicOrderByInputInterfaceIDASC VMNicOrderByInput = "interface_id_ASC"

	// VMNicOrderByInputInterfaceIDDESC captures enum value "interface_id_DESC"
	VMNicOrderByInputInterfaceIDDESC VMNicOrderByInput = "interface_id_DESC"

	// VMNicOrderByInputIPAddressASC captures enum value "ip_address_ASC"
	VMNicOrderByInputIPAddressASC VMNicOrderByInput = "ip_address_ASC"

	// VMNicOrderByInputIPAddressDESC captures enum value "ip_address_DESC"
	VMNicOrderByInputIPAddressDESC VMNicOrderByInput = "ip_address_DESC"

	// VMNicOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VMNicOrderByInputLocalIDASC VMNicOrderByInput = "local_id_ASC"

	// VMNicOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VMNicOrderByInputLocalIDDESC VMNicOrderByInput = "local_id_DESC"

	// VMNicOrderByInputMacAddressASC captures enum value "mac_address_ASC"
	VMNicOrderByInputMacAddressASC VMNicOrderByInput = "mac_address_ASC"

	// VMNicOrderByInputMacAddressDESC captures enum value "mac_address_DESC"
	VMNicOrderByInputMacAddressDESC VMNicOrderByInput = "mac_address_DESC"

	// VMNicOrderByInputMirrorASC captures enum value "mirror_ASC"
	VMNicOrderByInputMirrorASC VMNicOrderByInput = "mirror_ASC"

	// VMNicOrderByInputMirrorDESC captures enum value "mirror_DESC"
	VMNicOrderByInputMirrorDESC VMNicOrderByInput = "mirror_DESC"

	// VMNicOrderByInputModelASC captures enum value "model_ASC"
	VMNicOrderByInputModelASC VMNicOrderByInput = "model_ASC"

	// VMNicOrderByInputModelDESC captures enum value "model_DESC"
	VMNicOrderByInputModelDESC VMNicOrderByInput = "model_DESC"

	// VMNicOrderByInputOrderASC captures enum value "order_ASC"
	VMNicOrderByInputOrderASC VMNicOrderByInput = "order_ASC"

	// VMNicOrderByInputOrderDESC captures enum value "order_DESC"
	VMNicOrderByInputOrderDESC VMNicOrderByInput = "order_DESC"

	// VMNicOrderByInputSubnetMaskASC captures enum value "subnet_mask_ASC"
	VMNicOrderByInputSubnetMaskASC VMNicOrderByInput = "subnet_mask_ASC"

	// VMNicOrderByInputSubnetMaskDESC captures enum value "subnet_mask_DESC"
	VMNicOrderByInputSubnetMaskDESC VMNicOrderByInput = "subnet_mask_DESC"

	// VMNicOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	VMNicOrderByInputUpdatedAtASC VMNicOrderByInput = "updatedAt_ASC"

	// VMNicOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	VMNicOrderByInputUpdatedAtDESC VMNicOrderByInput = "updatedAt_DESC"
)

// for schema
var vmNicOrderByInputEnum []interface{}

func init() {
	var res []VMNicOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","enabled_ASC","enabled_DESC","gateway_ASC","gateway_DESC","id_ASC","id_DESC","interface_id_ASC","interface_id_DESC","ip_address_ASC","ip_address_DESC","local_id_ASC","local_id_DESC","mac_address_ASC","mac_address_DESC","mirror_ASC","mirror_DESC","model_ASC","model_DESC","order_ASC","order_DESC","subnet_mask_ASC","subnet_mask_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmNicOrderByInputEnum = append(vmNicOrderByInputEnum, v)
	}
}

func (m VMNicOrderByInput) validateVMNicOrderByInputEnum(path, location string, value VMNicOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vmNicOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm nic order by input
func (m VMNicOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMNicOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm nic order by input based on context it is used
func (m VMNicOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}