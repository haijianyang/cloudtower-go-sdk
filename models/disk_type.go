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

// DiskType disk type
//
// swagger:model DiskType
type DiskType string

func NewDiskType(value DiskType) *DiskType {
	v := value
	return &v
}

const (

	// DiskTypeHDD captures enum value "HDD"
	DiskTypeHDD DiskType = "HDD"

	// DiskTypePMem captures enum value "PMem"
	DiskTypePMem DiskType = "PMem"

	// DiskTypeSSD captures enum value "SSD"
	DiskTypeSSD DiskType = "SSD"
)

// for schema
var diskTypeEnum []interface{}

func init() {
	var res []DiskType
	if err := json.Unmarshal([]byte(`["HDD","PMem","SSD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diskTypeEnum = append(diskTypeEnum, v)
	}
}

func (m DiskType) validateDiskTypeEnum(path, location string, value DiskType) error {
	if err := validate.EnumCase(path, location, value, diskTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this disk type
func (m DiskType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDiskTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this disk type based on context it is used
func (m DiskType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
