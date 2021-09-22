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

// BrickTopoUpdationParams brick topo updation params
//
// swagger:model BrickTopoUpdationParams
type BrickTopoUpdationParams struct {

	// data
	// Required: true
	Data *BrickTopoUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *BrickTopoWhereInput `json:"where"`
}

// Validate validates this brick topo updation params
func (m *BrickTopoUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoUpdationParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoUpdationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this brick topo updation params based on the context it is used
func (m *BrickTopoUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoUpdationParams) UnmarshalBinary(b []byte) error {
	var res BrickTopoUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BrickTopoUpdationParamsData brick topo updation params data
//
// swagger:model BrickTopoUpdationParamsData
type BrickTopoUpdationParamsData struct {

	// capacity
	Capacity *BrickTopoUpdationParamsDataCapacity `json:"capacity,omitempty"`

	// height
	Height float64 `json:"height,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// node topoes
	NodeTopoes *NodeTopoWhereInput `json:"node_topoes,omitempty"`

	// position
	Position float64 `json:"position,omitempty"`

	// tag position in brick
	TagPositionInBrick []*BrickTopoUpdationParamsDataTagPositionInBrickItems0 `json:"tag_position_in_brick"`
}

// Validate validates this brick topo updation params data
func (m *BrickTopoUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeTopoes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagPositionInBrick(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoUpdationParamsData) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "capacity")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoUpdationParamsData) validateNodeTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeTopoes) { // not required
		return nil
	}

	if m.NodeTopoes != nil {
		if err := m.NodeTopoes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "node_topoes")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoUpdationParamsData) validateTagPositionInBrick(formats strfmt.Registry) error {
	if swag.IsZero(m.TagPositionInBrick) { // not required
		return nil
	}

	for i := 0; i < len(m.TagPositionInBrick); i++ {
		if swag.IsZero(m.TagPositionInBrick[i]) { // not required
			continue
		}

		if m.TagPositionInBrick[i] != nil {
			if err := m.TagPositionInBrick[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "tag_position_in_brick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this brick topo updation params data based on the context it is used
func (m *BrickTopoUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTagPositionInBrick(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoUpdationParamsData) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if m.Capacity != nil {
		if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "capacity")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoUpdationParamsData) contextValidateNodeTopoes(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeTopoes != nil {
		if err := m.NodeTopoes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "node_topoes")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoUpdationParamsData) contextValidateTagPositionInBrick(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagPositionInBrick); i++ {

		if m.TagPositionInBrick[i] != nil {
			if err := m.TagPositionInBrick[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "tag_position_in_brick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res BrickTopoUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BrickTopoUpdationParamsDataCapacity brick topo updation params data capacity
//
// swagger:model BrickTopoUpdationParamsDataCapacity
type BrickTopoUpdationParamsDataCapacity struct {

	// column
	Column *float64 `json:"column,omitempty"`

	// row
	Row *float64 `json:"row,omitempty"`
}

// Validate validates this brick topo updation params data capacity
func (m *BrickTopoUpdationParamsDataCapacity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this brick topo updation params data capacity based on context it is used
func (m *BrickTopoUpdationParamsDataCapacity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoUpdationParamsDataCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoUpdationParamsDataCapacity) UnmarshalBinary(b []byte) error {
	var res BrickTopoUpdationParamsDataCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BrickTopoUpdationParamsDataTagPositionInBrickItems0 brick topo updation params data tag position in brick items0
//
// swagger:model BrickTopoUpdationParamsDataTagPositionInBrickItems0
type BrickTopoUpdationParamsDataTagPositionInBrickItems0 struct {

	// column
	// Required: true
	Column *float64 `json:"column"`

	// row
	// Required: true
	Row *float64 `json:"row"`

	// tag
	// Required: true
	Tag *string `json:"tag"`
}

// Validate validates this brick topo updation params data tag position in brick items0
func (m *BrickTopoUpdationParamsDataTagPositionInBrickItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoUpdationParamsDataTagPositionInBrickItems0) validateColumn(formats strfmt.Registry) error {

	if err := validate.Required("column", "body", m.Column); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoUpdationParamsDataTagPositionInBrickItems0) validateRow(formats strfmt.Registry) error {

	if err := validate.Required("row", "body", m.Row); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoUpdationParamsDataTagPositionInBrickItems0) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this brick topo updation params data tag position in brick items0 based on context it is used
func (m *BrickTopoUpdationParamsDataTagPositionInBrickItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoUpdationParamsDataTagPositionInBrickItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoUpdationParamsDataTagPositionInBrickItems0) UnmarshalBinary(b []byte) error {
	var res BrickTopoUpdationParamsDataTagPositionInBrickItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}