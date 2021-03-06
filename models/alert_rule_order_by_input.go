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

// AlertRuleOrderByInput alert rule order by input
//
// swagger:model AlertRuleOrderByInput
type AlertRuleOrderByInput string

func NewAlertRuleOrderByInput(value AlertRuleOrderByInput) *AlertRuleOrderByInput {
	v := value
	return &v
}

const (

	// AlertRuleOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	AlertRuleOrderByInputCreatedAtASC AlertRuleOrderByInput = "createdAt_ASC"

	// AlertRuleOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	AlertRuleOrderByInputCreatedAtDESC AlertRuleOrderByInput = "createdAt_DESC"

	// AlertRuleOrderByInputCustomizedASC captures enum value "customized_ASC"
	AlertRuleOrderByInputCustomizedASC AlertRuleOrderByInput = "customized_ASC"

	// AlertRuleOrderByInputCustomizedDESC captures enum value "customized_DESC"
	AlertRuleOrderByInputCustomizedDESC AlertRuleOrderByInput = "customized_DESC"

	// AlertRuleOrderByInputDisabledASC captures enum value "disabled_ASC"
	AlertRuleOrderByInputDisabledASC AlertRuleOrderByInput = "disabled_ASC"

	// AlertRuleOrderByInputDisabledDESC captures enum value "disabled_DESC"
	AlertRuleOrderByInputDisabledDESC AlertRuleOrderByInput = "disabled_DESC"

	// AlertRuleOrderByInputIDASC captures enum value "id_ASC"
	AlertRuleOrderByInputIDASC AlertRuleOrderByInput = "id_ASC"

	// AlertRuleOrderByInputIDDESC captures enum value "id_DESC"
	AlertRuleOrderByInputIDDESC AlertRuleOrderByInput = "id_DESC"

	// AlertRuleOrderByInputLocalIDASC captures enum value "local_id_ASC"
	AlertRuleOrderByInputLocalIDASC AlertRuleOrderByInput = "local_id_ASC"

	// AlertRuleOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	AlertRuleOrderByInputLocalIDDESC AlertRuleOrderByInput = "local_id_DESC"

	// AlertRuleOrderByInputThresholdsASC captures enum value "thresholds_ASC"
	AlertRuleOrderByInputThresholdsASC AlertRuleOrderByInput = "thresholds_ASC"

	// AlertRuleOrderByInputThresholdsDESC captures enum value "thresholds_DESC"
	AlertRuleOrderByInputThresholdsDESC AlertRuleOrderByInput = "thresholds_DESC"

	// AlertRuleOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	AlertRuleOrderByInputUpdatedAtASC AlertRuleOrderByInput = "updatedAt_ASC"

	// AlertRuleOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	AlertRuleOrderByInputUpdatedAtDESC AlertRuleOrderByInput = "updatedAt_DESC"
)

// for schema
var alertRuleOrderByInputEnum []interface{}

func init() {
	var res []AlertRuleOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","customized_ASC","customized_DESC","disabled_ASC","disabled_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","thresholds_ASC","thresholds_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertRuleOrderByInputEnum = append(alertRuleOrderByInputEnum, v)
	}
}

func (m AlertRuleOrderByInput) validateAlertRuleOrderByInputEnum(path, location string, value AlertRuleOrderByInput) error {
	if err := validate.EnumCase(path, location, value, alertRuleOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this alert rule order by input
func (m AlertRuleOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAlertRuleOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this alert rule order by input based on context it is used
func (m AlertRuleOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
