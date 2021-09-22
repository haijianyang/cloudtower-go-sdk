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

// ViewOrderByInput view order by input
//
// swagger:model ViewOrderByInput
type ViewOrderByInput string

func NewViewOrderByInput(value ViewOrderByInput) *ViewOrderByInput {
	v := value
	return &v
}

const (

	// ViewOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ViewOrderByInputCreatedAtASC ViewOrderByInput = "createdAt_ASC"

	// ViewOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ViewOrderByInputCreatedAtDESC ViewOrderByInput = "createdAt_DESC"

	// ViewOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	ViewOrderByInputEntityAsyncStatusASC ViewOrderByInput = "entityAsyncStatus_ASC"

	// ViewOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	ViewOrderByInputEntityAsyncStatusDESC ViewOrderByInput = "entityAsyncStatus_DESC"

	// ViewOrderByInputIDASC captures enum value "id_ASC"
	ViewOrderByInputIDASC ViewOrderByInput = "id_ASC"

	// ViewOrderByInputIDDESC captures enum value "id_DESC"
	ViewOrderByInputIDDESC ViewOrderByInput = "id_DESC"

	// ViewOrderByInputLocalIDASC captures enum value "local_id_ASC"
	ViewOrderByInputLocalIDASC ViewOrderByInput = "local_id_ASC"

	// ViewOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	ViewOrderByInputLocalIDDESC ViewOrderByInput = "local_id_DESC"

	// ViewOrderByInputNameASC captures enum value "name_ASC"
	ViewOrderByInputNameASC ViewOrderByInput = "name_ASC"

	// ViewOrderByInputNameDESC captures enum value "name_DESC"
	ViewOrderByInputNameDESC ViewOrderByInput = "name_DESC"

	// ViewOrderByInputTimeSpanASC captures enum value "time_span_ASC"
	ViewOrderByInputTimeSpanASC ViewOrderByInput = "time_span_ASC"

	// ViewOrderByInputTimeSpanDESC captures enum value "time_span_DESC"
	ViewOrderByInputTimeSpanDESC ViewOrderByInput = "time_span_DESC"

	// ViewOrderByInputTimeUnitASC captures enum value "time_unit_ASC"
	ViewOrderByInputTimeUnitASC ViewOrderByInput = "time_unit_ASC"

	// ViewOrderByInputTimeUnitDESC captures enum value "time_unit_DESC"
	ViewOrderByInputTimeUnitDESC ViewOrderByInput = "time_unit_DESC"

	// ViewOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	ViewOrderByInputUpdatedAtASC ViewOrderByInput = "updatedAt_ASC"

	// ViewOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	ViewOrderByInputUpdatedAtDESC ViewOrderByInput = "updatedAt_DESC"
)

// for schema
var viewOrderByInputEnum []interface{}

func init() {
	var res []ViewOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","time_span_ASC","time_span_DESC","time_unit_ASC","time_unit_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewOrderByInputEnum = append(viewOrderByInputEnum, v)
	}
}

func (m ViewOrderByInput) validateViewOrderByInputEnum(path, location string, value ViewOrderByInput) error {
	if err := validate.EnumCase(path, location, value, viewOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this view order by input
func (m ViewOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateViewOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this view order by input based on context it is used
func (m ViewOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}