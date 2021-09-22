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

// VMFolderOrderByInput Vm folder order by input
//
// swagger:model VmFolderOrderByInput
type VMFolderOrderByInput string

func NewVMFolderOrderByInput(value VMFolderOrderByInput) *VMFolderOrderByInput {
	v := value
	return &v
}

const (

	// VMFolderOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	VMFolderOrderByInputCreatedAtASC VMFolderOrderByInput = "createdAt_ASC"

	// VMFolderOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	VMFolderOrderByInputCreatedAtDESC VMFolderOrderByInput = "createdAt_DESC"

	// VMFolderOrderByInputIDASC captures enum value "id_ASC"
	VMFolderOrderByInputIDASC VMFolderOrderByInput = "id_ASC"

	// VMFolderOrderByInputIDDESC captures enum value "id_DESC"
	VMFolderOrderByInputIDDESC VMFolderOrderByInput = "id_DESC"

	// VMFolderOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VMFolderOrderByInputLocalIDASC VMFolderOrderByInput = "local_id_ASC"

	// VMFolderOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VMFolderOrderByInputLocalIDDESC VMFolderOrderByInput = "local_id_DESC"

	// VMFolderOrderByInputNameASC captures enum value "name_ASC"
	VMFolderOrderByInputNameASC VMFolderOrderByInput = "name_ASC"

	// VMFolderOrderByInputNameDESC captures enum value "name_DESC"
	VMFolderOrderByInputNameDESC VMFolderOrderByInput = "name_DESC"

	// VMFolderOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	VMFolderOrderByInputUpdatedAtASC VMFolderOrderByInput = "updatedAt_ASC"

	// VMFolderOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	VMFolderOrderByInputUpdatedAtDESC VMFolderOrderByInput = "updatedAt_DESC"

	// VMFolderOrderByInputVMNumASC captures enum value "vm_num_ASC"
	VMFolderOrderByInputVMNumASC VMFolderOrderByInput = "vm_num_ASC"

	// VMFolderOrderByInputVMNumDESC captures enum value "vm_num_DESC"
	VMFolderOrderByInputVMNumDESC VMFolderOrderByInput = "vm_num_DESC"
)

// for schema
var vmFolderOrderByInputEnum []interface{}

func init() {
	var res []VMFolderOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","updatedAt_ASC","updatedAt_DESC","vm_num_ASC","vm_num_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmFolderOrderByInputEnum = append(vmFolderOrderByInputEnum, v)
	}
}

func (m VMFolderOrderByInput) validateVMFolderOrderByInputEnum(path, location string, value VMFolderOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vmFolderOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm folder order by input
func (m VMFolderOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMFolderOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm folder order by input based on context it is used
func (m VMFolderOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
