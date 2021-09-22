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

// FilterRuleAggregationEnum filter rule aggregation enum
//
// swagger:model FilterRuleAggregationEnum
type FilterRuleAggregationEnum string

func NewFilterRuleAggregationEnum(value FilterRuleAggregationEnum) *FilterRuleAggregationEnum {
	v := value
	return &v
}

const (

	// FilterRuleAggregationEnumAVG captures enum value "AVG"
	FilterRuleAggregationEnumAVG FilterRuleAggregationEnum = "AVG"

	// FilterRuleAggregationEnumMAX captures enum value "MAX"
	FilterRuleAggregationEnumMAX FilterRuleAggregationEnum = "MAX"

	// FilterRuleAggregationEnumMIN captures enum value "MIN"
	FilterRuleAggregationEnumMIN FilterRuleAggregationEnum = "MIN"

	// FilterRuleAggregationEnumQUANTILE captures enum value "QUANTILE"
	FilterRuleAggregationEnumQUANTILE FilterRuleAggregationEnum = "QUANTILE"

	// FilterRuleAggregationEnumSUM captures enum value "SUM"
	FilterRuleAggregationEnumSUM FilterRuleAggregationEnum = "SUM"
)

// for schema
var filterRuleAggregationEnumEnum []interface{}

func init() {
	var res []FilterRuleAggregationEnum
	if err := json.Unmarshal([]byte(`["AVG","MAX","MIN","QUANTILE","SUM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filterRuleAggregationEnumEnum = append(filterRuleAggregationEnumEnum, v)
	}
}

func (m FilterRuleAggregationEnum) validateFilterRuleAggregationEnumEnum(path, location string, value FilterRuleAggregationEnum) error {
	if err := validate.EnumCase(path, location, value, filterRuleAggregationEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this filter rule aggregation enum
func (m FilterRuleAggregationEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFilterRuleAggregationEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this filter rule aggregation enum based on context it is used
func (m FilterRuleAggregationEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}