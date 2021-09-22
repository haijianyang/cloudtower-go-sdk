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

// MigrateTransmitterOrderByInput migrate transmitter order by input
//
// swagger:model MigrateTransmitterOrderByInput
type MigrateTransmitterOrderByInput string

func NewMigrateTransmitterOrderByInput(value MigrateTransmitterOrderByInput) *MigrateTransmitterOrderByInput {
	v := value
	return &v
}

const (

	// MigrateTransmitterOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	MigrateTransmitterOrderByInputCreatedAtASC MigrateTransmitterOrderByInput = "createdAt_ASC"

	// MigrateTransmitterOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	MigrateTransmitterOrderByInputCreatedAtDESC MigrateTransmitterOrderByInput = "createdAt_DESC"

	// MigrateTransmitterOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	MigrateTransmitterOrderByInputEntityAsyncStatusASC MigrateTransmitterOrderByInput = "entityAsyncStatus_ASC"

	// MigrateTransmitterOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	MigrateTransmitterOrderByInputEntityAsyncStatusDESC MigrateTransmitterOrderByInput = "entityAsyncStatus_DESC"

	// MigrateTransmitterOrderByInputIDASC captures enum value "id_ASC"
	MigrateTransmitterOrderByInputIDASC MigrateTransmitterOrderByInput = "id_ASC"

	// MigrateTransmitterOrderByInputIDDESC captures enum value "id_DESC"
	MigrateTransmitterOrderByInputIDDESC MigrateTransmitterOrderByInput = "id_DESC"

	// MigrateTransmitterOrderByInputIPASC captures enum value "ip_ASC"
	MigrateTransmitterOrderByInputIPASC MigrateTransmitterOrderByInput = "ip_ASC"

	// MigrateTransmitterOrderByInputIPDESC captures enum value "ip_DESC"
	MigrateTransmitterOrderByInputIPDESC MigrateTransmitterOrderByInput = "ip_DESC"

	// MigrateTransmitterOrderByInputNameASC captures enum value "name_ASC"
	MigrateTransmitterOrderByInputNameASC MigrateTransmitterOrderByInput = "name_ASC"

	// MigrateTransmitterOrderByInputNameDESC captures enum value "name_DESC"
	MigrateTransmitterOrderByInputNameDESC MigrateTransmitterOrderByInput = "name_DESC"

	// MigrateTransmitterOrderByInputPasswordASC captures enum value "password_ASC"
	MigrateTransmitterOrderByInputPasswordASC MigrateTransmitterOrderByInput = "password_ASC"

	// MigrateTransmitterOrderByInputPasswordDESC captures enum value "password_DESC"
	MigrateTransmitterOrderByInputPasswordDESC MigrateTransmitterOrderByInput = "password_DESC"

	// MigrateTransmitterOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	MigrateTransmitterOrderByInputUpdatedAtASC MigrateTransmitterOrderByInput = "updatedAt_ASC"

	// MigrateTransmitterOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	MigrateTransmitterOrderByInputUpdatedAtDESC MigrateTransmitterOrderByInput = "updatedAt_DESC"

	// MigrateTransmitterOrderByInputUsernameASC captures enum value "username_ASC"
	MigrateTransmitterOrderByInputUsernameASC MigrateTransmitterOrderByInput = "username_ASC"

	// MigrateTransmitterOrderByInputUsernameDESC captures enum value "username_DESC"
	MigrateTransmitterOrderByInputUsernameDESC MigrateTransmitterOrderByInput = "username_DESC"
)

// for schema
var migrateTransmitterOrderByInputEnum []interface{}

func init() {
	var res []MigrateTransmitterOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","ip_ASC","ip_DESC","name_ASC","name_DESC","password_ASC","password_DESC","updatedAt_ASC","updatedAt_DESC","username_ASC","username_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		migrateTransmitterOrderByInputEnum = append(migrateTransmitterOrderByInputEnum, v)
	}
}

func (m MigrateTransmitterOrderByInput) validateMigrateTransmitterOrderByInputEnum(path, location string, value MigrateTransmitterOrderByInput) error {
	if err := validate.EnumCase(path, location, value, migrateTransmitterOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this migrate transmitter order by input
func (m MigrateTransmitterOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMigrateTransmitterOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this migrate transmitter order by input based on context it is used
func (m MigrateTransmitterOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
