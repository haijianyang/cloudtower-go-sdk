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

// RackTopoOrderByInput rack topo order by input
//
// swagger:model RackTopoOrderByInput
type RackTopoOrderByInput string

func NewRackTopoOrderByInput(value RackTopoOrderByInput) *RackTopoOrderByInput {
	v := value
	return &v
}

const (

	// RackTopoOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	RackTopoOrderByInputCreatedAtASC RackTopoOrderByInput = "createdAt_ASC"

	// RackTopoOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	RackTopoOrderByInputCreatedAtDESC RackTopoOrderByInput = "createdAt_DESC"

	// RackTopoOrderByInputHeightASC captures enum value "height_ASC"
	RackTopoOrderByInputHeightASC RackTopoOrderByInput = "height_ASC"

	// RackTopoOrderByInputHeightDESC captures enum value "height_DESC"
	RackTopoOrderByInputHeightDESC RackTopoOrderByInput = "height_DESC"

	// RackTopoOrderByInputIDASC captures enum value "id_ASC"
	RackTopoOrderByInputIDASC RackTopoOrderByInput = "id_ASC"

	// RackTopoOrderByInputIDDESC captures enum value "id_DESC"
	RackTopoOrderByInputIDDESC RackTopoOrderByInput = "id_DESC"

	// RackTopoOrderByInputLocalIDASC captures enum value "local_id_ASC"
	RackTopoOrderByInputLocalIDASC RackTopoOrderByInput = "local_id_ASC"

	// RackTopoOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	RackTopoOrderByInputLocalIDDESC RackTopoOrderByInput = "local_id_DESC"

	// RackTopoOrderByInputNameASC captures enum value "name_ASC"
	RackTopoOrderByInputNameASC RackTopoOrderByInput = "name_ASC"

	// RackTopoOrderByInputNameDESC captures enum value "name_DESC"
	RackTopoOrderByInputNameDESC RackTopoOrderByInput = "name_DESC"

	// RackTopoOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	RackTopoOrderByInputUpdatedAtASC RackTopoOrderByInput = "updatedAt_ASC"

	// RackTopoOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	RackTopoOrderByInputUpdatedAtDESC RackTopoOrderByInput = "updatedAt_DESC"
)

// for schema
var rackTopoOrderByInputEnum []interface{}

func init() {
	var res []RackTopoOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","height_ASC","height_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackTopoOrderByInputEnum = append(rackTopoOrderByInputEnum, v)
	}
}

func (m RackTopoOrderByInput) validateRackTopoOrderByInputEnum(path, location string, value RackTopoOrderByInput) error {
	if err := validate.EnumCase(path, location, value, rackTopoOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this rack topo order by input
func (m RackTopoOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRackTopoOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this rack topo order by input based on context it is used
func (m RackTopoOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
