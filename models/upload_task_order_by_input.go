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

// UploadTaskOrderByInput upload task order by input
//
// swagger:model UploadTaskOrderByInput
type UploadTaskOrderByInput string

func NewUploadTaskOrderByInput(value UploadTaskOrderByInput) *UploadTaskOrderByInput {
	v := value
	return &v
}

const (

	// UploadTaskOrderByInputArgsASC captures enum value "args_ASC"
	UploadTaskOrderByInputArgsASC UploadTaskOrderByInput = "args_ASC"

	// UploadTaskOrderByInputArgsDESC captures enum value "args_DESC"
	UploadTaskOrderByInputArgsDESC UploadTaskOrderByInput = "args_DESC"

	// UploadTaskOrderByInputChunkSizeASC captures enum value "chunk_size_ASC"
	UploadTaskOrderByInputChunkSizeASC UploadTaskOrderByInput = "chunk_size_ASC"

	// UploadTaskOrderByInputChunkSizeDESC captures enum value "chunk_size_DESC"
	UploadTaskOrderByInputChunkSizeDESC UploadTaskOrderByInput = "chunk_size_DESC"

	// UploadTaskOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	UploadTaskOrderByInputCreatedAtASC UploadTaskOrderByInput = "createdAt_ASC"

	// UploadTaskOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	UploadTaskOrderByInputCreatedAtDESC UploadTaskOrderByInput = "createdAt_DESC"

	// UploadTaskOrderByInputCurrentChunkASC captures enum value "current_chunk_ASC"
	UploadTaskOrderByInputCurrentChunkASC UploadTaskOrderByInput = "current_chunk_ASC"

	// UploadTaskOrderByInputCurrentChunkDESC captures enum value "current_chunk_DESC"
	UploadTaskOrderByInputCurrentChunkDESC UploadTaskOrderByInput = "current_chunk_DESC"

	// UploadTaskOrderByInputFinishedAtASC captures enum value "finished_at_ASC"
	UploadTaskOrderByInputFinishedAtASC UploadTaskOrderByInput = "finished_at_ASC"

	// UploadTaskOrderByInputFinishedAtDESC captures enum value "finished_at_DESC"
	UploadTaskOrderByInputFinishedAtDESC UploadTaskOrderByInput = "finished_at_DESC"

	// UploadTaskOrderByInputIDASC captures enum value "id_ASC"
	UploadTaskOrderByInputIDASC UploadTaskOrderByInput = "id_ASC"

	// UploadTaskOrderByInputIDDESC captures enum value "id_DESC"
	UploadTaskOrderByInputIDDESC UploadTaskOrderByInput = "id_DESC"

	// UploadTaskOrderByInputResourceTypeASC captures enum value "resource_type_ASC"
	UploadTaskOrderByInputResourceTypeASC UploadTaskOrderByInput = "resource_type_ASC"

	// UploadTaskOrderByInputResourceTypeDESC captures enum value "resource_type_DESC"
	UploadTaskOrderByInputResourceTypeDESC UploadTaskOrderByInput = "resource_type_DESC"

	// UploadTaskOrderByInputSizeASC captures enum value "size_ASC"
	UploadTaskOrderByInputSizeASC UploadTaskOrderByInput = "size_ASC"

	// UploadTaskOrderByInputSizeDESC captures enum value "size_DESC"
	UploadTaskOrderByInputSizeDESC UploadTaskOrderByInput = "size_DESC"

	// UploadTaskOrderByInputStartedAtASC captures enum value "started_at_ASC"
	UploadTaskOrderByInputStartedAtASC UploadTaskOrderByInput = "started_at_ASC"

	// UploadTaskOrderByInputStartedAtDESC captures enum value "started_at_DESC"
	UploadTaskOrderByInputStartedAtDESC UploadTaskOrderByInput = "started_at_DESC"

	// UploadTaskOrderByInputStatusASC captures enum value "status_ASC"
	UploadTaskOrderByInputStatusASC UploadTaskOrderByInput = "status_ASC"

	// UploadTaskOrderByInputStatusDESC captures enum value "status_DESC"
	UploadTaskOrderByInputStatusDESC UploadTaskOrderByInput = "status_DESC"

	// UploadTaskOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	UploadTaskOrderByInputUpdatedAtASC UploadTaskOrderByInput = "updatedAt_ASC"

	// UploadTaskOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	UploadTaskOrderByInputUpdatedAtDESC UploadTaskOrderByInput = "updatedAt_DESC"
)

// for schema
var uploadTaskOrderByInputEnum []interface{}

func init() {
	var res []UploadTaskOrderByInput
	if err := json.Unmarshal([]byte(`["args_ASC","args_DESC","chunk_size_ASC","chunk_size_DESC","createdAt_ASC","createdAt_DESC","current_chunk_ASC","current_chunk_DESC","finished_at_ASC","finished_at_DESC","id_ASC","id_DESC","resource_type_ASC","resource_type_DESC","size_ASC","size_DESC","started_at_ASC","started_at_DESC","status_ASC","status_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		uploadTaskOrderByInputEnum = append(uploadTaskOrderByInputEnum, v)
	}
}

func (m UploadTaskOrderByInput) validateUploadTaskOrderByInputEnum(path, location string, value UploadTaskOrderByInput) error {
	if err := validate.EnumCase(path, location, value, uploadTaskOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this upload task order by input
func (m UploadTaskOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUploadTaskOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this upload task order by input based on context it is used
func (m UploadTaskOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
