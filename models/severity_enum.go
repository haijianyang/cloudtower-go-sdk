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

// SeverityEnum severity enum
//
// swagger:model SeverityEnum
type SeverityEnum string

func NewSeverityEnum(value SeverityEnum) *SeverityEnum {
	v := value
	return &v
}

const (

	// SeverityEnumCRITICAL captures enum value "CRITICAL"
	SeverityEnumCRITICAL SeverityEnum = "CRITICAL"

	// SeverityEnumINFO captures enum value "INFO"
	SeverityEnumINFO SeverityEnum = "INFO"

	// SeverityEnumNOTICE captures enum value "NOTICE"
	SeverityEnumNOTICE SeverityEnum = "NOTICE"

	// SeverityEnumSEVERITYUNSPECIFIED captures enum value "SEVERITY_UNSPECIFIED"
	SeverityEnumSEVERITYUNSPECIFIED SeverityEnum = "SEVERITY_UNSPECIFIED"
)

// for schema
var severityEnumEnum []interface{}

func init() {
	var res []SeverityEnum
	if err := json.Unmarshal([]byte(`["CRITICAL","INFO","NOTICE","SEVERITY_UNSPECIFIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		severityEnumEnum = append(severityEnumEnum, v)
	}
}

func (m SeverityEnum) validateSeverityEnumEnum(path, location string, value SeverityEnum) error {
	if err := validate.EnumCase(path, location, value, severityEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this severity enum
func (m SeverityEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSeverityEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this severity enum based on context it is used
func (m SeverityEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
