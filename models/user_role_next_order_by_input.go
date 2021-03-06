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

// UserRoleNextOrderByInput user role next order by input
//
// swagger:model UserRoleNextOrderByInput
type UserRoleNextOrderByInput string

func NewUserRoleNextOrderByInput(value UserRoleNextOrderByInput) *UserRoleNextOrderByInput {
	v := value
	return &v
}

const (

	// UserRoleNextOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	UserRoleNextOrderByInputCreatedAtASC UserRoleNextOrderByInput = "createdAt_ASC"

	// UserRoleNextOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	UserRoleNextOrderByInputCreatedAtDESC UserRoleNextOrderByInput = "createdAt_DESC"

	// UserRoleNextOrderByInputIDASC captures enum value "id_ASC"
	UserRoleNextOrderByInputIDASC UserRoleNextOrderByInput = "id_ASC"

	// UserRoleNextOrderByInputIDDESC captures enum value "id_DESC"
	UserRoleNextOrderByInputIDDESC UserRoleNextOrderByInput = "id_DESC"

	// UserRoleNextOrderByInputNameASC captures enum value "name_ASC"
	UserRoleNextOrderByInputNameASC UserRoleNextOrderByInput = "name_ASC"

	// UserRoleNextOrderByInputNameDESC captures enum value "name_DESC"
	UserRoleNextOrderByInputNameDESC UserRoleNextOrderByInput = "name_DESC"

	// UserRoleNextOrderByInputPlatformASC captures enum value "platform_ASC"
	UserRoleNextOrderByInputPlatformASC UserRoleNextOrderByInput = "platform_ASC"

	// UserRoleNextOrderByInputPlatformDESC captures enum value "platform_DESC"
	UserRoleNextOrderByInputPlatformDESC UserRoleNextOrderByInput = "platform_DESC"

	// UserRoleNextOrderByInputPresetASC captures enum value "preset_ASC"
	UserRoleNextOrderByInputPresetASC UserRoleNextOrderByInput = "preset_ASC"

	// UserRoleNextOrderByInputPresetDESC captures enum value "preset_DESC"
	UserRoleNextOrderByInputPresetDESC UserRoleNextOrderByInput = "preset_DESC"

	// UserRoleNextOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	UserRoleNextOrderByInputUpdatedAtASC UserRoleNextOrderByInput = "updatedAt_ASC"

	// UserRoleNextOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	UserRoleNextOrderByInputUpdatedAtDESC UserRoleNextOrderByInput = "updatedAt_DESC"
)

// for schema
var userRoleNextOrderByInputEnum []interface{}

func init() {
	var res []UserRoleNextOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","id_ASC","id_DESC","name_ASC","name_DESC","platform_ASC","platform_DESC","preset_ASC","preset_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRoleNextOrderByInputEnum = append(userRoleNextOrderByInputEnum, v)
	}
}

func (m UserRoleNextOrderByInput) validateUserRoleNextOrderByInputEnum(path, location string, value UserRoleNextOrderByInput) error {
	if err := validate.EnumCase(path, location, value, userRoleNextOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user role next order by input
func (m UserRoleNextOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserRoleNextOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user role next order by input based on context it is used
func (m UserRoleNextOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
