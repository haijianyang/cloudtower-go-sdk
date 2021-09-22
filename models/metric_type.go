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

// MetricType metric type
//
// swagger:model MetricType
type MetricType string

func NewMetricType(value MetricType) *MetricType {
	v := value
	return &v
}

const (

	// MetricTypeBOTTOMK captures enum value "BOTTOMK"
	MetricTypeBOTTOMK MetricType = "BOTTOMK"

	// MetricTypeNORMAL captures enum value "NORMAL"
	MetricTypeNORMAL MetricType = "NORMAL"

	// MetricTypeTOPK captures enum value "TOPK"
	MetricTypeTOPK MetricType = "TOPK"
)

// for schema
var metricTypeEnum []interface{}

func init() {
	var res []MetricType
	if err := json.Unmarshal([]byte(`["BOTTOMK","NORMAL","TOPK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricTypeEnum = append(metricTypeEnum, v)
	}
}

func (m MetricType) validateMetricTypeEnum(path, location string, value MetricType) error {
	if err := validate.EnumCase(path, location, value, metricTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this metric type
func (m MetricType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMetricTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this metric type based on context it is used
func (m MetricType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
