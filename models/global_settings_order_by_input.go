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

// GlobalSettingsOrderByInput global settings order by input
//
// swagger:model GlobalSettingsOrderByInput
type GlobalSettingsOrderByInput string

func NewGlobalSettingsOrderByInput(value GlobalSettingsOrderByInput) *GlobalSettingsOrderByInput {
	v := value
	return &v
}

const (

	// GlobalSettingsOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	GlobalSettingsOrderByInputCreatedAtASC GlobalSettingsOrderByInput = "createdAt_ASC"

	// GlobalSettingsOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	GlobalSettingsOrderByInputCreatedAtDESC GlobalSettingsOrderByInput = "createdAt_DESC"

	// GlobalSettingsOrderByInputIDASC captures enum value "id_ASC"
	GlobalSettingsOrderByInputIDASC GlobalSettingsOrderByInput = "id_ASC"

	// GlobalSettingsOrderByInputIDDESC captures enum value "id_DESC"
	GlobalSettingsOrderByInputIDDESC GlobalSettingsOrderByInput = "id_DESC"

	// GlobalSettingsOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	GlobalSettingsOrderByInputUpdatedAtASC GlobalSettingsOrderByInput = "updatedAt_ASC"

	// GlobalSettingsOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	GlobalSettingsOrderByInputUpdatedAtDESC GlobalSettingsOrderByInput = "updatedAt_DESC"

	// GlobalSettingsOrderByInputVMRecycleBinASC captures enum value "vm_recycle_bin_ASC"
	GlobalSettingsOrderByInputVMRecycleBinASC GlobalSettingsOrderByInput = "vm_recycle_bin_ASC"

	// GlobalSettingsOrderByInputVMRecycleBinDESC captures enum value "vm_recycle_bin_DESC"
	GlobalSettingsOrderByInputVMRecycleBinDESC GlobalSettingsOrderByInput = "vm_recycle_bin_DESC"
)

// for schema
var globalSettingsOrderByInputEnum []interface{}

func init() {
	var res []GlobalSettingsOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","id_ASC","id_DESC","updatedAt_ASC","updatedAt_DESC","vm_recycle_bin_ASC","vm_recycle_bin_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalSettingsOrderByInputEnum = append(globalSettingsOrderByInputEnum, v)
	}
}

func (m GlobalSettingsOrderByInput) validateGlobalSettingsOrderByInputEnum(path, location string, value GlobalSettingsOrderByInput) error {
	if err := validate.EnumCase(path, location, value, globalSettingsOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this global settings order by input
func (m GlobalSettingsOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGlobalSettingsOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this global settings order by input based on context it is used
func (m GlobalSettingsOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
