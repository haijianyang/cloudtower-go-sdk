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

// ClusterImage cluster image
//
// swagger:model ClusterImage
type ClusterImage struct {

	// cluster
	// Required: true
	Cluster *ClusterImageCluster `json:"cluster"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// meta name
	// Required: true
	MetaName *string `json:"meta_name"`

	// meta size
	// Required: true
	MetaSize *float64 `json:"meta_size"`

	// name
	// Required: true
	Name *string `json:"name"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this cluster image
func (m *ClusterImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterImage) validateCluster(formats strfmt.Registry) error {

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

func (m *ClusterImage) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterImage) validateMetaName(formats strfmt.Registry) error {

	if err := validate.Required("meta_name", "body", m.MetaName); err != nil {
		return err
	}

	return nil
}

func (m *ClusterImage) validateMetaSize(formats strfmt.Registry) error {

	if err := validate.Required("meta_size", "body", m.MetaSize); err != nil {
		return err
	}

	return nil
}

func (m *ClusterImage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ClusterImage) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *ClusterImage) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster image based on the context it is used
func (m *ClusterImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterImage) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ClusterImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterImage) UnmarshalBinary(b []byte) error {
	var res ClusterImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterImageCluster cluster image cluster
//
// swagger:model ClusterImageCluster
type ClusterImageCluster struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this cluster image cluster
func (m *ClusterImageCluster) Validate(formats strfmt.Registry) error {
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

func (m *ClusterImageCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterImageCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster image cluster based on context it is used
func (m *ClusterImageCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterImageCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterImageCluster) UnmarshalBinary(b []byte) error {
	var res ClusterImageCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
