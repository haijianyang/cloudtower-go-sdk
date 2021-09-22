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

// RackTopoCreationParams rack topo creation params
//
// swagger:model RackTopoCreationParams
type RackTopoCreationParams struct {

	// brick topoes
	BrickTopoes *BrickTopoWhereInput `json:"brick_topoes,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// height
	// Required: true
	Height *float64 `json:"height"`

	// name
	// Required: true
	Name *string `json:"name"`

	// zone topo id
	// Required: true
	ZoneTopoID *string `json:"zone_topo_id"`
}

// Validate validates this rack topo creation params
func (m *RackTopoCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrickTopoes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneTopoID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopoCreationParams) validateBrickTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.BrickTopoes) { // not required
		return nil
	}

	if m.BrickTopoes != nil {
		if err := m.BrickTopoes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("brick_topoes")
			}
			return err
		}
	}

	return nil
}

func (m *RackTopoCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *RackTopoCreationParams) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *RackTopoCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RackTopoCreationParams) validateZoneTopoID(formats strfmt.Registry) error {

	if err := validate.Required("zone_topo_id", "body", m.ZoneTopoID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rack topo creation params based on the context it is used
func (m *RackTopoCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBrickTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopoCreationParams) contextValidateBrickTopoes(ctx context.Context, formats strfmt.Registry) error {

	if m.BrickTopoes != nil {
		if err := m.BrickTopoes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("brick_topoes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackTopoCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTopoCreationParams) UnmarshalBinary(b []byte) error {
	var res RackTopoCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
