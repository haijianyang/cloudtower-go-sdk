// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EntityFilter entity filter
//
// swagger:model EntityFilter
type EntityFilter struct {

	// apply to all clusters
	ApplyToAllClusters *bool `json:"apply_to_all_clusters,omitempty"`

	// clusters
	Clusters []*EntityFilterClustersItems0 `json:"clusters,omitempty"`

	// entity type
	// Required: true
	EntityType *EntityType `json:"entity_type"`

	// exclude ids
	// Required: true
	ExcludeIds []string `json:"exclude_ids"`

	// exec failed cluster
	ExecFailedCluster []*EntityFilterExecFailedClusterItems0 `json:"exec_failed_cluster,omitempty"`

	// filter error
	// Required: true
	FilterError []string `json:"filter_error"`

	// filter status
	// Required: true
	FilterStatus *FilterStatus `json:"filter_status"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ids
	// Required: true
	Ids []string `json:"ids"`

	// last executed at
	LastExecutedAt *string `json:"last_executed_at,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// preset
	Preset *string `json:"preset,omitempty"`

	// rules
	// Required: true
	Rules []*EntityFilterRulesItems0 `json:"rules"`
}

// Validate validates this entity filter
func (m *EntityFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludeIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecFailedCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityFilter) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntityFilter) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	if m.EntityType != nil {
		if err := m.EntityType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity_type")
			}
			return err
		}
	}

	return nil
}

func (m *EntityFilter) validateExcludeIds(formats strfmt.Registry) error {

	if err := validate.Required("exclude_ids", "body", m.ExcludeIds); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilter) validateExecFailedCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecFailedCluster) { // not required
		return nil
	}

	for i := 0; i < len(m.ExecFailedCluster); i++ {
		if swag.IsZero(m.ExecFailedCluster[i]) { // not required
			continue
		}

		if m.ExecFailedCluster[i] != nil {
			if err := m.ExecFailedCluster[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exec_failed_cluster" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntityFilter) validateFilterError(formats strfmt.Registry) error {

	if err := validate.Required("filter_error", "body", m.FilterError); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilter) validateFilterStatus(formats strfmt.Registry) error {

	if err := validate.Required("filter_status", "body", m.FilterStatus); err != nil {
		return err
	}

	if err := validate.Required("filter_status", "body", m.FilterStatus); err != nil {
		return err
	}

	if m.FilterStatus != nil {
		if err := m.FilterStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_status")
			}
			return err
		}
	}

	return nil
}

func (m *EntityFilter) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilter) validateIds(formats strfmt.Registry) error {

	if err := validate.Required("ids", "body", m.Ids); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilter) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("rules", "body", m.Rules); err != nil {
		return err
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this entity filter based on the context it is used
func (m *EntityFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExecFailedCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityFilter) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntityFilter) contextValidateEntityType(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityType != nil {
		if err := m.EntityType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity_type")
			}
			return err
		}
	}

	return nil
}

func (m *EntityFilter) contextValidateExecFailedCluster(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExecFailedCluster); i++ {

		if m.ExecFailedCluster[i] != nil {
			if err := m.ExecFailedCluster[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exec_failed_cluster" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntityFilter) contextValidateFilterStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterStatus != nil {
		if err := m.FilterStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_status")
			}
			return err
		}
	}

	return nil
}

func (m *EntityFilter) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {
			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityFilter) UnmarshalBinary(b []byte) error {
	var res EntityFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EntityFilterClustersItems0 entity filter clusters items0
//
// swagger:model EntityFilterClustersItems0
type EntityFilterClustersItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this entity filter clusters items0
func (m *EntityFilterClustersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityFilterClustersItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilterClustersItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this entity filter clusters items0 based on context it is used
func (m *EntityFilterClustersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntityFilterClustersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityFilterClustersItems0) UnmarshalBinary(b []byte) error {
	var res EntityFilterClustersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EntityFilterExecFailedClusterItems0 entity filter exec failed cluster items0
//
// swagger:model EntityFilterExecFailedClusterItems0
type EntityFilterExecFailedClusterItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this entity filter exec failed cluster items0
func (m *EntityFilterExecFailedClusterItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityFilterExecFailedClusterItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilterExecFailedClusterItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this entity filter exec failed cluster items0 based on context it is used
func (m *EntityFilterExecFailedClusterItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntityFilterExecFailedClusterItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityFilterExecFailedClusterItems0) UnmarshalBinary(b []byte) error {
	var res EntityFilterExecFailedClusterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EntityFilterRulesItems0 entity filter rules items0
//
// swagger:model EntityFilterRulesItems0
type EntityFilterRulesItems0 struct {

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

// Validate validates this entity filter rules items0
func (m *EntityFilterRulesItems0) Validate(formats strfmt.Registry) error {
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

func (m *EntityFilterRulesItems0) validateAggregation(formats strfmt.Registry) error {

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

func (m *EntityFilterRulesItems0) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilterRulesItems0) validateMetric(formats strfmt.Registry) error {

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

func (m *EntityFilterRulesItems0) validateOp(formats strfmt.Registry) error {

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

func (m *EntityFilterRulesItems0) validateQuantile(formats strfmt.Registry) error {

	if err := validate.Required("quantile", "body", m.Quantile); err != nil {
		return err
	}

	return nil
}

func (m *EntityFilterRulesItems0) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this entity filter rules items0 based on the context it is used
func (m *EntityFilterRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *EntityFilterRulesItems0) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EntityFilterRulesItems0) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EntityFilterRulesItems0) contextValidateOp(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EntityFilterRulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityFilterRulesItems0) UnmarshalBinary(b []byte) error {
	var res EntityFilterRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}