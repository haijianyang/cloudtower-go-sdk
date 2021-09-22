// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVdsesConnectionRequestBody get vdses connection request body
//
// swagger:model GetVdsesConnectionRequestBody
type GetVdsesConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *float64 `json:"first,omitempty"`

	// last
	Last *float64 `json:"last,omitempty"`

	// order by
	OrderBy interface{} `json:"orderBy,omitempty"`

	// skip
	Skip *float64 `json:"skip,omitempty"`

	// where
	Where interface{} `json:"where,omitempty"`
}

// Validate validates this get vdses connection request body
func (m *GetVdsesConnectionRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get vdses connection request body based on context it is used
func (m *GetVdsesConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetVdsesConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVdsesConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVdsesConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
