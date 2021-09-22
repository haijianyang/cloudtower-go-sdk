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

// DeleteSnmpTrapReceiver delete snmp trap receiver
//
// swagger:model DeleteSnmpTrapReceiver
type DeleteSnmpTrapReceiver struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this delete snmp trap receiver
func (m *DeleteSnmpTrapReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteSnmpTrapReceiver) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete snmp trap receiver based on context it is used
func (m *DeleteSnmpTrapReceiver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteSnmpTrapReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteSnmpTrapReceiver) UnmarshalBinary(b []byte) error {
	var res DeleteSnmpTrapReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}