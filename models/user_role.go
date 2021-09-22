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

// UserRole user role
//
// swagger:model UserRole
type UserRole string

func NewUserRole(value UserRole) *UserRole {
	v := value
	return &v
}

const (

	// UserRoleADMIN captures enum value "ADMIN"
	UserRoleADMIN UserRole = "ADMIN"

	// UserRoleREADONLY captures enum value "READ_ONLY"
	UserRoleREADONLY UserRole = "READ_ONLY"

	// UserRoleROOT captures enum value "ROOT"
	UserRoleROOT UserRole = "ROOT"
)

// for schema
var userRoleEnum []interface{}

func init() {
	var res []UserRole
	if err := json.Unmarshal([]byte(`["ADMIN","READ_ONLY","ROOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRoleEnum = append(userRoleEnum, v)
	}
}

func (m UserRole) validateUserRoleEnum(path, location string, value UserRole) error {
	if err := validate.EnumCase(path, location, value, userRoleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user role
func (m UserRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user role based on context it is used
func (m UserRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}