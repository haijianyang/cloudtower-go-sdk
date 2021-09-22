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
)

// GlobalAlertRuleWhereInput global alert rule where input
//
// swagger:model GlobalAlertRuleWhereInput
type GlobalAlertRuleWhereInput struct {

	// a n d
	AND []*GlobalAlertRuleWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*GlobalAlertRuleWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*GlobalAlertRuleWhereInput `json:"OR,omitempty"`

	// alert rules every
	AlertRulesEvery interface{} `json:"alert_rules_every,omitempty"`

	// alert rules none
	AlertRulesNone interface{} `json:"alert_rules_none,omitempty"`

	// alert rules some
	AlertRulesSome interface{} `json:"alert_rules_some,omitempty"`

	// boolean
	Boolean *bool `json:"boolean,omitempty"`

	// boolean not
	BooleanNot *bool `json:"boolean_not,omitempty"`

	// cause
	Cause *string `json:"cause,omitempty"`

	// cause contains
	CauseContains *string `json:"cause_contains,omitempty"`

	// cause ends with
	CauseEndsWith *string `json:"cause_ends_with,omitempty"`

	// cause gt
	CauseGt *string `json:"cause_gt,omitempty"`

	// cause gte
	CauseGte *string `json:"cause_gte,omitempty"`

	// cause in
	CauseIn []string `json:"cause_in,omitempty"`

	// cause lt
	CauseLt *string `json:"cause_lt,omitempty"`

	// cause lte
	CauseLte *string `json:"cause_lte,omitempty"`

	// cause not
	CauseNot *string `json:"cause_not,omitempty"`

	// cause not contains
	CauseNotContains *string `json:"cause_not_contains,omitempty"`

	// cause not ends with
	CauseNotEndsWith *string `json:"cause_not_ends_with,omitempty"`

	// cause not in
	CauseNotIn []string `json:"cause_not_in,omitempty"`

	// cause not starts with
	CauseNotStartsWith *string `json:"cause_not_starts_with,omitempty"`

	// cause starts with
	CauseStartsWith *string `json:"cause_starts_with,omitempty"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// disabled not
	DisabledNot *bool `json:"disabled_not,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// impact
	Impact *string `json:"impact,omitempty"`

	// impact contains
	ImpactContains *string `json:"impact_contains,omitempty"`

	// impact ends with
	ImpactEndsWith *string `json:"impact_ends_with,omitempty"`

	// impact gt
	ImpactGt *string `json:"impact_gt,omitempty"`

	// impact gte
	ImpactGte *string `json:"impact_gte,omitempty"`

	// impact in
	ImpactIn []string `json:"impact_in,omitempty"`

	// impact lt
	ImpactLt *string `json:"impact_lt,omitempty"`

	// impact lte
	ImpactLte *string `json:"impact_lte,omitempty"`

	// impact not
	ImpactNot *string `json:"impact_not,omitempty"`

	// impact not contains
	ImpactNotContains *string `json:"impact_not_contains,omitempty"`

	// impact not ends with
	ImpactNotEndsWith *string `json:"impact_not_ends_with,omitempty"`

	// impact not in
	ImpactNotIn []string `json:"impact_not_in,omitempty"`

	// impact not starts with
	ImpactNotStartsWith *string `json:"impact_not_starts_with,omitempty"`

	// impact starts with
	ImpactStartsWith *string `json:"impact_starts_with,omitempty"`

	// message
	Message *string `json:"message,omitempty"`

	// message contains
	MessageContains *string `json:"message_contains,omitempty"`

	// message ends with
	MessageEndsWith *string `json:"message_ends_with,omitempty"`

	// message gt
	MessageGt *string `json:"message_gt,omitempty"`

	// message gte
	MessageGte *string `json:"message_gte,omitempty"`

	// message in
	MessageIn []string `json:"message_in,omitempty"`

	// message lt
	MessageLt *string `json:"message_lt,omitempty"`

	// message lte
	MessageLte *string `json:"message_lte,omitempty"`

	// message not
	MessageNot *string `json:"message_not,omitempty"`

	// message not contains
	MessageNotContains *string `json:"message_not_contains,omitempty"`

	// message not ends with
	MessageNotEndsWith *string `json:"message_not_ends_with,omitempty"`

	// message not in
	MessageNotIn []string `json:"message_not_in,omitempty"`

	// message not starts with
	MessageNotStartsWith *string `json:"message_not_starts_with,omitempty"`

	// message starts with
	MessageStartsWith *string `json:"message_starts_with,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// object
	Object interface{} `json:"object,omitempty"`

	// object in
	ObjectIn []AlertRuleObject `json:"object_in,omitempty"`

	// object not
	ObjectNot interface{} `json:"object_not,omitempty"`

	// object not in
	ObjectNotIn []AlertRuleObject `json:"object_not_in,omitempty"`

	// operator
	Operator *string `json:"operator,omitempty"`

	// operator contains
	OperatorContains *string `json:"operator_contains,omitempty"`

	// operator ends with
	OperatorEndsWith *string `json:"operator_ends_with,omitempty"`

	// operator gt
	OperatorGt *string `json:"operator_gt,omitempty"`

	// operator gte
	OperatorGte *string `json:"operator_gte,omitempty"`

	// operator in
	OperatorIn []string `json:"operator_in,omitempty"`

	// operator lt
	OperatorLt *string `json:"operator_lt,omitempty"`

	// operator lte
	OperatorLte *string `json:"operator_lte,omitempty"`

	// operator not
	OperatorNot *string `json:"operator_not,omitempty"`

	// operator not contains
	OperatorNotContains *string `json:"operator_not_contains,omitempty"`

	// operator not ends with
	OperatorNotEndsWith *string `json:"operator_not_ends_with,omitempty"`

	// operator not in
	OperatorNotIn []string `json:"operator_not_in,omitempty"`

	// operator not starts with
	OperatorNotStartsWith *string `json:"operator_not_starts_with,omitempty"`

	// operator starts with
	OperatorStartsWith *string `json:"operator_starts_with,omitempty"`

	// solution
	Solution *string `json:"solution,omitempty"`

	// solution contains
	SolutionContains *string `json:"solution_contains,omitempty"`

	// solution ends with
	SolutionEndsWith *string `json:"solution_ends_with,omitempty"`

	// solution gt
	SolutionGt *string `json:"solution_gt,omitempty"`

	// solution gte
	SolutionGte *string `json:"solution_gte,omitempty"`

	// solution in
	SolutionIn []string `json:"solution_in,omitempty"`

	// solution lt
	SolutionLt *string `json:"solution_lt,omitempty"`

	// solution lte
	SolutionLte *string `json:"solution_lte,omitempty"`

	// solution not
	SolutionNot *string `json:"solution_not,omitempty"`

	// solution not contains
	SolutionNotContains *string `json:"solution_not_contains,omitempty"`

	// solution not ends with
	SolutionNotEndsWith *string `json:"solution_not_ends_with,omitempty"`

	// solution not in
	SolutionNotIn []string `json:"solution_not_in,omitempty"`

	// solution not starts with
	SolutionNotStartsWith *string `json:"solution_not_starts_with,omitempty"`

	// solution starts with
	SolutionStartsWith *string `json:"solution_starts_with,omitempty"`

	// unit
	Unit interface{} `json:"unit,omitempty"`

	// unit in
	UnitIn []AlertRuleUnit `json:"unit_in,omitempty"`

	// unit not
	UnitNot interface{} `json:"unit_not,omitempty"`

	// unit not in
	UnitNotIn []AlertRuleUnit `json:"unit_not_in,omitempty"`
}

// Validate validates this global alert rule where input
func (m *GlobalAlertRuleWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalAlertRuleWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) validateObjectIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ObjectIn); i++ {

		if err := m.ObjectIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) validateObjectNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ObjectNotIn); i++ {

		if err := m.ObjectNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) validateUnitIn(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitIn) { // not required
		return nil
	}

	for i := 0; i < len(m.UnitIn); i++ {

		if err := m.UnitIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) validateUnitNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.UnitNotIn); i++ {

		if err := m.UnitNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this global alert rule where input based on the context it is used
func (m *GlobalAlertRuleWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnitIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnitNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalAlertRuleWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) contextValidateObjectIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ObjectIn); i++ {

		if err := m.ObjectIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) contextValidateObjectNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ObjectNotIn); i++ {

		if err := m.ObjectNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) contextValidateUnitIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UnitIn); i++ {

		if err := m.UnitIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GlobalAlertRuleWhereInput) contextValidateUnitNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UnitNotIn); i++ {

		if err := m.UnitNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalAlertRuleWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalAlertRuleWhereInput) UnmarshalBinary(b []byte) error {
	var res GlobalAlertRuleWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}