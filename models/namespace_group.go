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

// NamespaceGroup namespace group
//
// swagger:model NamespaceGroup
type NamespaceGroup struct {

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// labels
	Labels []*NamespaceGroupLabelsItems0 `json:"labels,omitempty"`

	// local create time
	// Required: true
	LocalCreateTime *string `json:"local_create_time"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespaces
	Namespaces []*NamespaceGroupNamespacesItems0 `json:"namespaces,omitempty"`

	// nvmf subsystem
	// Required: true
	NvmfSubsystem *NamespaceGroupNvmfSubsystem `json:"nvmf_subsystem"`
}

// Validate validates this namespace group
func (m *NamespaceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamespaceGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceGroup) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NamespaceGroup) validateLocalCreateTime(formats strfmt.Registry) error {

	if err := validate.Required("local_create_time", "body", m.LocalCreateTime); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceGroup) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceGroup) validateNamespaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Namespaces); i++ {
		if swag.IsZero(m.Namespaces[i]) { // not required
			continue
		}

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NamespaceGroup) validateNvmfSubsystem(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem", "body", m.NvmfSubsystem); err != nil {
		return err
	}

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this namespace group based on the context it is used
func (m *NamespaceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamespaceGroup) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NamespaceGroup) contextValidateNamespaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Namespaces); i++ {

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NamespaceGroup) contextValidateNvmfSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceGroup) UnmarshalBinary(b []byte) error {
	var res NamespaceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NamespaceGroupLabelsItems0 namespace group labels items0
//
// swagger:model NamespaceGroupLabelsItems0
type NamespaceGroupLabelsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this namespace group labels items0
func (m *NamespaceGroupLabelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamespaceGroupLabelsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this namespace group labels items0 based on context it is used
func (m *NamespaceGroupLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceGroupLabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceGroupLabelsItems0) UnmarshalBinary(b []byte) error {
	var res NamespaceGroupLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NamespaceGroupNamespacesItems0 namespace group namespaces items0
//
// swagger:model NamespaceGroupNamespacesItems0
type NamespaceGroupNamespacesItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this namespace group namespaces items0
func (m *NamespaceGroupNamespacesItems0) Validate(formats strfmt.Registry) error {
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

func (m *NamespaceGroupNamespacesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceGroupNamespacesItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this namespace group namespaces items0 based on context it is used
func (m *NamespaceGroupNamespacesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceGroupNamespacesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceGroupNamespacesItems0) UnmarshalBinary(b []byte) error {
	var res NamespaceGroupNamespacesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NamespaceGroupNvmfSubsystem namespace group nvmf subsystem
//
// swagger:model NamespaceGroupNvmfSubsystem
type NamespaceGroupNvmfSubsystem struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this namespace group nvmf subsystem
func (m *NamespaceGroupNvmfSubsystem) Validate(formats strfmt.Registry) error {
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

func (m *NamespaceGroupNvmfSubsystem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NamespaceGroupNvmfSubsystem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this namespace group nvmf subsystem based on context it is used
func (m *NamespaceGroupNvmfSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceGroupNvmfSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceGroupNvmfSubsystem) UnmarshalBinary(b []byte) error {
	var res NamespaceGroupNvmfSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
