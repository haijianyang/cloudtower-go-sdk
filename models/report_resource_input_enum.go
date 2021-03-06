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

// ReportResourceInputEnum report resource input enum
//
// swagger:model ReportResourceInputEnum
type ReportResourceInputEnum string

func NewReportResourceInputEnum(value ReportResourceInputEnum) *ReportResourceInputEnum {
	v := value
	return &v
}

const (

	// ReportResourceInputEnumALERT captures enum value "ALERT"
	ReportResourceInputEnumALERT ReportResourceInputEnum = "ALERT"

	// ReportResourceInputEnumALL captures enum value "ALL"
	ReportResourceInputEnumALL ReportResourceInputEnum = "ALL"

	// ReportResourceInputEnumCLUSTER captures enum value "CLUSTER"
	ReportResourceInputEnumCLUSTER ReportResourceInputEnum = "CLUSTER"

	// ReportResourceInputEnumDATACENTER captures enum value "DATA_CENTER"
	ReportResourceInputEnumDATACENTER ReportResourceInputEnum = "DATA_CENTER"

	// ReportResourceInputEnumDISK captures enum value "DISK"
	ReportResourceInputEnumDISK ReportResourceInputEnum = "DISK"

	// ReportResourceInputEnumELFIMAGE captures enum value "ELF_IMAGE"
	ReportResourceInputEnumELFIMAGE ReportResourceInputEnum = "ELF_IMAGE"

	// ReportResourceInputEnumENTITYFILTERS captures enum value "ENTITY_FILTERS"
	ReportResourceInputEnumENTITYFILTERS ReportResourceInputEnum = "ENTITY_FILTERS"

	// ReportResourceInputEnumGLOBALALERTRULE captures enum value "GLOBAL_ALERT_RULE"
	ReportResourceInputEnumGLOBALALERTRULE ReportResourceInputEnum = "GLOBAL_ALERT_RULE"

	// ReportResourceInputEnumHOST captures enum value "HOST"
	ReportResourceInputEnumHOST ReportResourceInputEnum = "HOST"

	// ReportResourceInputEnumTASK captures enum value "TASK"
	ReportResourceInputEnumTASK ReportResourceInputEnum = "TASK"

	// ReportResourceInputEnumVDS captures enum value "VDS"
	ReportResourceInputEnumVDS ReportResourceInputEnum = "VDS"

	// ReportResourceInputEnumVLAN captures enum value "VLAN"
	ReportResourceInputEnumVLAN ReportResourceInputEnum = "VLAN"

	// ReportResourceInputEnumVM captures enum value "VM"
	ReportResourceInputEnumVM ReportResourceInputEnum = "VM"

	// ReportResourceInputEnumVMTEMPLATE captures enum value "VM_TEMPLATE"
	ReportResourceInputEnumVMTEMPLATE ReportResourceInputEnum = "VM_TEMPLATE"
)

// for schema
var reportResourceInputEnumEnum []interface{}

func init() {
	var res []ReportResourceInputEnum
	if err := json.Unmarshal([]byte(`["ALERT","ALL","CLUSTER","DATA_CENTER","DISK","ELF_IMAGE","ENTITY_FILTERS","GLOBAL_ALERT_RULE","HOST","TASK","VDS","VLAN","VM","VM_TEMPLATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportResourceInputEnumEnum = append(reportResourceInputEnumEnum, v)
	}
}

func (m ReportResourceInputEnum) validateReportResourceInputEnumEnum(path, location string, value ReportResourceInputEnum) error {
	if err := validate.EnumCase(path, location, value, reportResourceInputEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this report resource input enum
func (m ReportResourceInputEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReportResourceInputEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this report resource input enum based on context it is used
func (m ReportResourceInputEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
