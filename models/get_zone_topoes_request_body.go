// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetZoneTopoesRequestBody get zone topoes request body
//
// swagger:model GetZoneTopoesRequestBody
type GetZoneTopoesRequestBody struct {

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

// Validate validates this get zone topoes request body
func (m *GetZoneTopoesRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get zone topoes request body based on context it is used
func (m *GetZoneTopoesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetZoneTopoesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetZoneTopoesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetZoneTopoesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}