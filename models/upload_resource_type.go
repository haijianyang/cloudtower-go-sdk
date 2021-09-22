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

// UploadResourceType upload resource type
//
// swagger:model UploadResourceType
type UploadResourceType string

func NewUploadResourceType(value UploadResourceType) *UploadResourceType {
	v := value
	return &v
}

const (

	// UploadResourceTypeCLUSTERIMAGE captures enum value "CLUSTER_IMAGE"
	UploadResourceTypeCLUSTERIMAGE UploadResourceType = "CLUSTER_IMAGE"

	// UploadResourceTypeCLUSTERIMAGEMETA captures enum value "CLUSTER_IMAGE_META"
	UploadResourceTypeCLUSTERIMAGEMETA UploadResourceType = "CLUSTER_IMAGE_META"

	// UploadResourceTypeELFIMAGE captures enum value "ELF_IMAGE"
	UploadResourceTypeELFIMAGE UploadResourceType = "ELF_IMAGE"

	// UploadResourceTypeMONITORIMAGE captures enum value "MONITOR_IMAGE"
	UploadResourceTypeMONITORIMAGE UploadResourceType = "MONITOR_IMAGE"

	// UploadResourceTypeSVTIMAGE captures enum value "SVT_IMAGE"
	UploadResourceTypeSVTIMAGE UploadResourceType = "SVT_IMAGE"
)

// for schema
var uploadResourceTypeEnum []interface{}

func init() {
	var res []UploadResourceType
	if err := json.Unmarshal([]byte(`["CLUSTER_IMAGE","CLUSTER_IMAGE_META","ELF_IMAGE","MONITOR_IMAGE","SVT_IMAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		uploadResourceTypeEnum = append(uploadResourceTypeEnum, v)
	}
}

func (m UploadResourceType) validateUploadResourceTypeEnum(path, location string, value UploadResourceType) error {
	if err := validate.EnumCase(path, location, value, uploadResourceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this upload resource type
func (m UploadResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUploadResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this upload resource type based on context it is used
func (m UploadResourceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
