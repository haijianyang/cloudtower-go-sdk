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

// RackTopo rack topo
//
// swagger:model RackTopo
type RackTopo struct {

	// brick topoes
	BrickTopoes []*RackTopoBrickTopoesItems0 `json:"brick_topoes,omitempty"`

	// cluster
	// Required: true
	Cluster *RackTopoCluster `json:"cluster"`

	// height
	// Required: true
	Height *float64 `json:"height"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// zone topo
	// Required: true
	ZoneTopo *RackTopoZoneTopo `json:"zone_topo"`
}

// Validate validates this rack topo
func (m *RackTopo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrickTopoes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneTopo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopo) validateBrickTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.BrickTopoes) { // not required
		return nil
	}

	for i := 0; i < len(m.BrickTopoes); i++ {
		if swag.IsZero(m.BrickTopoes[i]) { // not required
			continue
		}

		if m.BrickTopoes[i] != nil {
			if err := m.BrickTopoes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("brick_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RackTopo) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *RackTopo) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *RackTopo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RackTopo) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *RackTopo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RackTopo) validateZoneTopo(formats strfmt.Registry) error {

	if err := validate.Required("zone_topo", "body", m.ZoneTopo); err != nil {
		return err
	}

	if m.ZoneTopo != nil {
		if err := m.ZoneTopo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone_topo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rack topo based on the context it is used
func (m *RackTopo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBrickTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneTopo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopo) contextValidateBrickTopoes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BrickTopoes); i++ {

		if m.BrickTopoes[i] != nil {
			if err := m.BrickTopoes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("brick_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RackTopo) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *RackTopo) contextValidateZoneTopo(ctx context.Context, formats strfmt.Registry) error {

	if m.ZoneTopo != nil {
		if err := m.ZoneTopo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone_topo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackTopo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTopo) UnmarshalBinary(b []byte) error {
	var res RackTopo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTopoBrickTopoesItems0 rack topo brick topoes items0
//
// swagger:model RackTopoBrickTopoesItems0
type RackTopoBrickTopoesItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this rack topo brick topoes items0
func (m *RackTopoBrickTopoesItems0) Validate(formats strfmt.Registry) error {
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

func (m *RackTopoBrickTopoesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RackTopoBrickTopoesItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack topo brick topoes items0 based on context it is used
func (m *RackTopoBrickTopoesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackTopoBrickTopoesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTopoBrickTopoesItems0) UnmarshalBinary(b []byte) error {
	var res RackTopoBrickTopoesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTopoCluster rack topo cluster
//
// swagger:model RackTopoCluster
type RackTopoCluster struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this rack topo cluster
func (m *RackTopoCluster) Validate(formats strfmt.Registry) error {
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

func (m *RackTopoCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RackTopoCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack topo cluster based on context it is used
func (m *RackTopoCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackTopoCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTopoCluster) UnmarshalBinary(b []byte) error {
	var res RackTopoCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTopoZoneTopo rack topo zone topo
//
// swagger:model RackTopoZoneTopo
type RackTopoZoneTopo struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this rack topo zone topo
func (m *RackTopoZoneTopo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopoZoneTopo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("zone_topo"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rack topo zone topo based on context it is used
func (m *RackTopoZoneTopo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RackTopoZoneTopo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTopoZoneTopo) UnmarshalBinary(b []byte) error {
	var res RackTopoZoneTopo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}