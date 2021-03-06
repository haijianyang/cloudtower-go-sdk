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

// Task task
//
// swagger:model Task
type Task struct {

	// args
	// Required: true
	Args interface{} `json:"args"`

	// cluster
	Cluster *TaskCluster `json:"cluster,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// error code
	ErrorCode *string `json:"error_code,omitempty"`

	// error message
	ErrorMessage *string `json:"error_message,omitempty"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// progress
	// Required: true
	Progress *float64 `json:"progress"`

	// resource id
	ResourceID *string `json:"resource_id,omitempty"`

	// resource mutation
	ResourceMutation *string `json:"resource_mutation,omitempty"`

	// resource rollback error
	ResourceRollbackError *string `json:"resource_rollback_error,omitempty"`

	// resource rollbacked
	ResourceRollbacked *bool `json:"resource_rollbacked,omitempty"`

	// resource type
	ResourceType *string `json:"resource_type,omitempty"`

	// snapshot
	// Required: true
	Snapshot *string `json:"snapshot"`

	// started at
	StartedAt *string `json:"started_at,omitempty"`

	// status
	// Required: true
	Status *TaskStatus `json:"status"`

	// steps
	// Required: true
	Steps []*TaskStepsItems0 `json:"steps"`

	// user
	User *TaskUser `json:"user,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateArgs(formats strfmt.Registry) error {

	if m.Args == nil {
		return errors.Required("args", "body", nil)
	}

	return nil
}

func (m *Task) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
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

func (m *Task) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateSnapshot(formats strfmt.Registry) error {

	if err := validate.Required("snapshot", "body", m.Snapshot); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Task) validateSteps(formats strfmt.Registry) error {

	if err := validate.Required("steps", "body", m.Steps); err != nil {
		return err
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task based on the context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Task) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Task) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {
			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskCluster task cluster
//
// swagger:model TaskCluster
type TaskCluster struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this task cluster
func (m *TaskCluster) Validate(formats strfmt.Registry) error {
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

func (m *TaskCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *TaskCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task cluster based on context it is used
func (m *TaskCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskCluster) UnmarshalBinary(b []byte) error {
	var res TaskCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskStepsItems0 task steps items0
//
// swagger:model TaskStepsItems0
type TaskStepsItems0 struct {

	// current
	Current *float64 `json:"current,omitempty"`

	// finished
	Finished *bool `json:"finished,omitempty"`

	// key
	Key *string `json:"key,omitempty"`

	// per second
	PerSecond *float64 `json:"per_second,omitempty"`

	// total
	Total *float64 `json:"total,omitempty"`

	// unit
	Unit interface{} `json:"unit,omitempty"`
}

// Validate validates this task steps items0
func (m *TaskStepsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this task steps items0 based on context it is used
func (m *TaskStepsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskStepsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskStepsItems0) UnmarshalBinary(b []byte) error {
	var res TaskStepsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskUser task user
//
// swagger:model TaskUser
type TaskUser struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this task user
func (m *TaskUser) Validate(formats strfmt.Registry) error {
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

func (m *TaskUser) validateID(formats strfmt.Registry) error {

	if err := validate.Required("user"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *TaskUser) validateName(formats strfmt.Registry) error {

	if err := validate.Required("user"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task user based on context it is used
func (m *TaskUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskUser) UnmarshalBinary(b []byte) error {
	var res TaskUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
