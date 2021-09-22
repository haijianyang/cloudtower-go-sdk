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

// ConsistencyGroupCreationParams consistency group creation params
//
// swagger:model ConsistencyGroupCreationParams
type ConsistencyGroupCreationParams struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// description
	Description string `json:"description,omitempty"`

	// iscsi luns ids
	// Required: true
	IscsiLunsIds []string `json:"iscsi_luns_ids"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespaces ids
	// Required: true
	NamespacesIds []string `json:"namespaces_ids"`
}

// Validate validates this consistency group creation params
func (m *ConsistencyGroupCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiLunsIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespacesIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupCreationParams) validateIscsiLunsIds(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_luns_ids", "body", m.IscsiLunsIds); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ConsistencyGroupCreationParams) validateNamespacesIds(formats strfmt.Registry) error {

	if err := validate.Required("namespaces_ids", "body", m.NamespacesIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this consistency group creation params based on context it is used
func (m *ConsistencyGroupCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupCreationParams) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}