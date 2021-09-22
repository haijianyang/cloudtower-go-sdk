// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FilterRuleInput filter rule input
//
// swagger:model FilterRuleInput
type FilterRuleInput struct {

	// aggregation
	// Required: true
	Aggregation *FilterRuleAggregationEnum `json:"aggregation"`

	// duration
	// Required: true
	Duration *float64 `json:"duration"`

	// metric
	// Required: true
	Metric *FilterRuleMetricEnum `json:"metric"`

	// op
	// Required: true
	Op *FilterRuleOpEnum `json:"op"`

	// quantile
	// Required: true
	Quantile *float64 `json:"quantile"`

	// threshold
	// Required: true
	Threshold *float64 `json:"threshold"`
}

// Validate validates this filter rule input
func (m *FilterRuleInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilterRuleInput) validateAggregation(formats strfmt.Registry) error {

	if err := validate.Required("aggregation", "body", m.Aggregation); err != nil {
		return err
	}

	if err := validate.Required("aggregation", "body", m.Aggregation); err != nil {
		return err
	}

	if m.Aggregation != nil {
		if err := m.Aggregation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *FilterRuleInput) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *FilterRuleInput) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *FilterRuleInput) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	if m.Op != nil {
		if err := m.Op.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

func (m *FilterRuleInput) validateQuantile(formats strfmt.Registry) error {

	if err := validate.Required("quantile", "body", m.Quantile); err != nil {
		return err
	}

	return nil
}

func (m *FilterRuleInput) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this filter rule input based on the context it is used
func (m *FilterRuleInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilterRuleInput) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregation != nil {
		if err := m.Aggregation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *FilterRuleInput) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

	if m.Metric != nil {
		if err := m.Metric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *FilterRuleInput) contextValidateOp(ctx context.Context, formats strfmt.Registry) error {

	if m.Op != nil {
		if err := m.Op.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilterRuleInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterRuleInput) UnmarshalBinary(b []byte) error {
	var res FilterRuleInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
