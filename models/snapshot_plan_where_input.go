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
)

// SnapshotPlanWhereInput snapshot plan where input
//
// swagger:model SnapshotPlanWhereInput
type SnapshotPlanWhereInput struct {

	// a n d
	AND []*SnapshotPlanWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*SnapshotPlanWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*SnapshotPlanWhereInput `json:"OR,omitempty"`

	// auto delete num
	AutoDeleteNum *float64 `json:"auto_delete_num,omitempty"`

	// auto delete num gt
	AutoDeleteNumGt *float64 `json:"auto_delete_num_gt,omitempty"`

	// auto delete num gte
	AutoDeleteNumGte *float64 `json:"auto_delete_num_gte,omitempty"`

	// auto delete num in
	AutoDeleteNumIn []float64 `json:"auto_delete_num_in,omitempty"`

	// auto delete num lt
	AutoDeleteNumLt *float64 `json:"auto_delete_num_lt,omitempty"`

	// auto delete num lte
	AutoDeleteNumLte *float64 `json:"auto_delete_num_lte,omitempty"`

	// auto delete num not
	AutoDeleteNumNot *float64 `json:"auto_delete_num_not,omitempty"`

	// auto delete num not in
	AutoDeleteNumNotIn []float64 `json:"auto_delete_num_not_in,omitempty"`

	// auto execute num
	AutoExecuteNum *float64 `json:"auto_execute_num,omitempty"`

	// auto execute num gt
	AutoExecuteNumGt *float64 `json:"auto_execute_num_gt,omitempty"`

	// auto execute num gte
	AutoExecuteNumGte *float64 `json:"auto_execute_num_gte,omitempty"`

	// auto execute num in
	AutoExecuteNumIn []float64 `json:"auto_execute_num_in,omitempty"`

	// auto execute num lt
	AutoExecuteNumLt *float64 `json:"auto_execute_num_lt,omitempty"`

	// auto execute num lte
	AutoExecuteNumLte *float64 `json:"auto_execute_num_lte,omitempty"`

	// auto execute num not
	AutoExecuteNumNot *float64 `json:"auto_execute_num_not,omitempty"`

	// auto execute num not in
	AutoExecuteNumNotIn []float64 `json:"auto_execute_num_not_in,omitempty"`

	// cluster
	Cluster interface{} `json:"cluster,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

	// end time
	EndTime *string `json:"end_time,omitempty"`

	// end time gt
	EndTimeGt *string `json:"end_time_gt,omitempty"`

	// end time gte
	EndTimeGte *string `json:"end_time_gte,omitempty"`

	// end time in
	EndTimeIn []string `json:"end_time_in,omitempty"`

	// end time lt
	EndTimeLt *string `json:"end_time_lt,omitempty"`

	// end time lte
	EndTimeLte *string `json:"end_time_lte,omitempty"`

	// end time not
	EndTimeNot *string `json:"end_time_not,omitempty"`

	// end time not in
	EndTimeNotIn []string `json:"end_time_not_in,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot interface{} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// execute plan type
	ExecutePlanType interface{} `json:"execute_plan_type,omitempty"`

	// execute plan type in
	ExecutePlanTypeIn []SnapshotPlanExecuteType `json:"execute_plan_type_in,omitempty"`

	// execute plan type not
	ExecutePlanTypeNot interface{} `json:"execute_plan_type_not,omitempty"`

	// execute plan type not in
	ExecutePlanTypeNotIn []SnapshotPlanExecuteType `json:"execute_plan_type_not_in,omitempty"`

	// execution tasks every
	ExecutionTasksEvery interface{} `json:"execution_tasks_every,omitempty"`

	// execution tasks none
	ExecutionTasksNone interface{} `json:"execution_tasks_none,omitempty"`

	// execution tasks some
	ExecutionTasksSome interface{} `json:"execution_tasks_some,omitempty"`

	// healthy
	Healthy *bool `json:"healthy,omitempty"`

	// healthy not
	HealthyNot *bool `json:"healthy_not,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// last execute end time
	LastExecuteEndTime *string `json:"last_execute_end_time,omitempty"`

	// last execute end time gt
	LastExecuteEndTimeGt *string `json:"last_execute_end_time_gt,omitempty"`

	// last execute end time gte
	LastExecuteEndTimeGte *string `json:"last_execute_end_time_gte,omitempty"`

	// last execute end time in
	LastExecuteEndTimeIn []string `json:"last_execute_end_time_in,omitempty"`

	// last execute end time lt
	LastExecuteEndTimeLt *string `json:"last_execute_end_time_lt,omitempty"`

	// last execute end time lte
	LastExecuteEndTimeLte *string `json:"last_execute_end_time_lte,omitempty"`

	// last execute end time not
	LastExecuteEndTimeNot *string `json:"last_execute_end_time_not,omitempty"`

	// last execute end time not in
	LastExecuteEndTimeNotIn []string `json:"last_execute_end_time_not_in,omitempty"`

	// last execute status
	LastExecuteStatus interface{} `json:"last_execute_status,omitempty"`

	// last execute status in
	LastExecuteStatusIn []SnapshotPlanExecuteStatus `json:"last_execute_status_in,omitempty"`

	// last execute status not
	LastExecuteStatusNot interface{} `json:"last_execute_status_not,omitempty"`

	// last execute status not in
	LastExecuteStatusNotIn []SnapshotPlanExecuteStatus `json:"last_execute_status_not_in,omitempty"`

	// last execute time
	LastExecuteTime *string `json:"last_execute_time,omitempty"`

	// last execute time gt
	LastExecuteTimeGt *string `json:"last_execute_time_gt,omitempty"`

	// last execute time gte
	LastExecuteTimeGte *string `json:"last_execute_time_gte,omitempty"`

	// last execute time in
	LastExecuteTimeIn []string `json:"last_execute_time_in,omitempty"`

	// last execute time lt
	LastExecuteTimeLt *string `json:"last_execute_time_lt,omitempty"`

	// last execute time lte
	LastExecuteTimeLte *string `json:"last_execute_time_lte,omitempty"`

	// last execute time not
	LastExecuteTimeNot *string `json:"last_execute_time_not,omitempty"`

	// last execute time not in
	LastExecuteTimeNotIn []string `json:"last_execute_time_not_in,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// logical size bytes
	LogicalSizeBytes *float64 `json:"logical_size_bytes,omitempty"`

	// logical size bytes gt
	LogicalSizeBytesGt *float64 `json:"logical_size_bytes_gt,omitempty"`

	// logical size bytes gte
	LogicalSizeBytesGte *float64 `json:"logical_size_bytes_gte,omitempty"`

	// logical size bytes in
	LogicalSizeBytesIn []float64 `json:"logical_size_bytes_in,omitempty"`

	// logical size bytes lt
	LogicalSizeBytesLt *float64 `json:"logical_size_bytes_lt,omitempty"`

	// logical size bytes lte
	LogicalSizeBytesLte *float64 `json:"logical_size_bytes_lte,omitempty"`

	// logical size bytes not
	LogicalSizeBytesNot *float64 `json:"logical_size_bytes_not,omitempty"`

	// logical size bytes not in
	LogicalSizeBytesNotIn []float64 `json:"logical_size_bytes_not_in,omitempty"`

	// manual delete num
	ManualDeleteNum *float64 `json:"manual_delete_num,omitempty"`

	// manual delete num gt
	ManualDeleteNumGt *float64 `json:"manual_delete_num_gt,omitempty"`

	// manual delete num gte
	ManualDeleteNumGte *float64 `json:"manual_delete_num_gte,omitempty"`

	// manual delete num in
	ManualDeleteNumIn []float64 `json:"manual_delete_num_in,omitempty"`

	// manual delete num lt
	ManualDeleteNumLt *float64 `json:"manual_delete_num_lt,omitempty"`

	// manual delete num lte
	ManualDeleteNumLte *float64 `json:"manual_delete_num_lte,omitempty"`

	// manual delete num not
	ManualDeleteNumNot *float64 `json:"manual_delete_num_not,omitempty"`

	// manual delete num not in
	ManualDeleteNumNotIn []float64 `json:"manual_delete_num_not_in,omitempty"`

	// manual execute num
	ManualExecuteNum *float64 `json:"manual_execute_num,omitempty"`

	// manual execute num gt
	ManualExecuteNumGt *float64 `json:"manual_execute_num_gt,omitempty"`

	// manual execute num gte
	ManualExecuteNumGte *float64 `json:"manual_execute_num_gte,omitempty"`

	// manual execute num in
	ManualExecuteNumIn []float64 `json:"manual_execute_num_in,omitempty"`

	// manual execute num lt
	ManualExecuteNumLt *float64 `json:"manual_execute_num_lt,omitempty"`

	// manual execute num lte
	ManualExecuteNumLte *float64 `json:"manual_execute_num_lte,omitempty"`

	// manual execute num not
	ManualExecuteNumNot *float64 `json:"manual_execute_num_not,omitempty"`

	// manual execute num not in
	ManualExecuteNumNotIn []float64 `json:"manual_execute_num_not_in,omitempty"`

	// mirror
	Mirror *bool `json:"mirror,omitempty"`

	// mirror not
	MirrorNot *bool `json:"mirror_not,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// next execute time
	NextExecuteTime *string `json:"next_execute_time,omitempty"`

	// next execute time gt
	NextExecuteTimeGt *string `json:"next_execute_time_gt,omitempty"`

	// next execute time gte
	NextExecuteTimeGte *string `json:"next_execute_time_gte,omitempty"`

	// next execute time in
	NextExecuteTimeIn []string `json:"next_execute_time_in,omitempty"`

	// next execute time lt
	NextExecuteTimeLt *string `json:"next_execute_time_lt,omitempty"`

	// next execute time lte
	NextExecuteTimeLte *string `json:"next_execute_time_lte,omitempty"`

	// next execute time not
	NextExecuteTimeNot *string `json:"next_execute_time_not,omitempty"`

	// next execute time not in
	NextExecuteTimeNotIn []string `json:"next_execute_time_not_in,omitempty"`

	// object num
	ObjectNum *float64 `json:"object_num,omitempty"`

	// object num gt
	ObjectNumGt *float64 `json:"object_num_gt,omitempty"`

	// object num gte
	ObjectNumGte *float64 `json:"object_num_gte,omitempty"`

	// object num in
	ObjectNumIn []float64 `json:"object_num_in,omitempty"`

	// object num lt
	ObjectNumLt *float64 `json:"object_num_lt,omitempty"`

	// object num lte
	ObjectNumLte *float64 `json:"object_num_lte,omitempty"`

	// object num not
	ObjectNumNot *float64 `json:"object_num_not,omitempty"`

	// object num not in
	ObjectNumNotIn []float64 `json:"object_num_not_in,omitempty"`

	// physical size bytes
	PhysicalSizeBytes *float64 `json:"physical_size_bytes,omitempty"`

	// physical size bytes gt
	PhysicalSizeBytesGt *float64 `json:"physical_size_bytes_gt,omitempty"`

	// physical size bytes gte
	PhysicalSizeBytesGte *float64 `json:"physical_size_bytes_gte,omitempty"`

	// physical size bytes in
	PhysicalSizeBytesIn []float64 `json:"physical_size_bytes_in,omitempty"`

	// physical size bytes lt
	PhysicalSizeBytesLt *float64 `json:"physical_size_bytes_lt,omitempty"`

	// physical size bytes lte
	PhysicalSizeBytesLte *float64 `json:"physical_size_bytes_lte,omitempty"`

	// physical size bytes not
	PhysicalSizeBytesNot *float64 `json:"physical_size_bytes_not,omitempty"`

	// physical size bytes not in
	PhysicalSizeBytesNotIn []float64 `json:"physical_size_bytes_not_in,omitempty"`

	// remain snapshot num
	RemainSnapshotNum *float64 `json:"remain_snapshot_num,omitempty"`

	// remain snapshot num gt
	RemainSnapshotNumGt *float64 `json:"remain_snapshot_num_gt,omitempty"`

	// remain snapshot num gte
	RemainSnapshotNumGte *float64 `json:"remain_snapshot_num_gte,omitempty"`

	// remain snapshot num in
	RemainSnapshotNumIn []float64 `json:"remain_snapshot_num_in,omitempty"`

	// remain snapshot num lt
	RemainSnapshotNumLt *float64 `json:"remain_snapshot_num_lt,omitempty"`

	// remain snapshot num lte
	RemainSnapshotNumLte *float64 `json:"remain_snapshot_num_lte,omitempty"`

	// remain snapshot num not
	RemainSnapshotNumNot *float64 `json:"remain_snapshot_num_not,omitempty"`

	// remain snapshot num not in
	RemainSnapshotNumNotIn []float64 `json:"remain_snapshot_num_not_in,omitempty"`

	// snapshot group num
	SnapshotGroupNum *float64 `json:"snapshot_group_num,omitempty"`

	// snapshot group num gt
	SnapshotGroupNumGt *float64 `json:"snapshot_group_num_gt,omitempty"`

	// snapshot group num gte
	SnapshotGroupNumGte *float64 `json:"snapshot_group_num_gte,omitempty"`

	// snapshot group num in
	SnapshotGroupNumIn []float64 `json:"snapshot_group_num_in,omitempty"`

	// snapshot group num lt
	SnapshotGroupNumLt *float64 `json:"snapshot_group_num_lt,omitempty"`

	// snapshot group num lte
	SnapshotGroupNumLte *float64 `json:"snapshot_group_num_lte,omitempty"`

	// snapshot group num not
	SnapshotGroupNumNot *float64 `json:"snapshot_group_num_not,omitempty"`

	// snapshot group num not in
	SnapshotGroupNumNotIn []float64 `json:"snapshot_group_num_not_in,omitempty"`

	// start time
	StartTime *string `json:"start_time,omitempty"`

	// start time gt
	StartTimeGt *string `json:"start_time_gt,omitempty"`

	// start time gte
	StartTimeGte *string `json:"start_time_gte,omitempty"`

	// start time in
	StartTimeIn []string `json:"start_time_in,omitempty"`

	// start time lt
	StartTimeLt *string `json:"start_time_lt,omitempty"`

	// start time lte
	StartTimeLte *string `json:"start_time_lte,omitempty"`

	// start time not
	StartTimeNot *string `json:"start_time_not,omitempty"`

	// start time not in
	StartTimeNotIn []string `json:"start_time_not_in,omitempty"`

	// status
	Status interface{} `json:"status,omitempty"`

	// status in
	StatusIn []SnapshotPlanStatus `json:"status_in,omitempty"`

	// status not
	StatusNot interface{} `json:"status_not,omitempty"`

	// status not in
	StatusNotIn []SnapshotPlanStatus `json:"status_not_in,omitempty"`

	// vms every
	VmsEvery interface{} `json:"vms_every,omitempty"`

	// vms none
	VmsNone interface{} `json:"vms_none,omitempty"`

	// vms some
	VmsSome interface{} `json:"vms_some,omitempty"`
}

// Validate validates this snapshot plan where input
func (m *SnapshotPlanWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutePlanTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutePlanTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastExecuteStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastExecuteStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPlanWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateExecutePlanTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecutePlanTypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ExecutePlanTypeIn); i++ {

		if err := m.ExecutePlanTypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execute_plan_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateExecutePlanTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecutePlanTypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ExecutePlanTypeNotIn); i++ {

		if err := m.ExecutePlanTypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execute_plan_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateLastExecuteStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LastExecuteStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LastExecuteStatusIn); i++ {

		if err := m.LastExecuteStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execute_status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateLastExecuteStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LastExecuteStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LastExecuteStatusNotIn); i++ {

		if err := m.LastExecuteStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execute_status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) validateStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this snapshot plan where input based on the context it is used
func (m *SnapshotPlanWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExecutePlanTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExecutePlanTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastExecuteStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastExecuteStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateExecutePlanTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExecutePlanTypeIn); i++ {

		if err := m.ExecutePlanTypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execute_plan_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateExecutePlanTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExecutePlanTypeNotIn); i++ {

		if err := m.ExecutePlanTypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execute_plan_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateLastExecuteStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LastExecuteStatusIn); i++ {

		if err := m.LastExecuteStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execute_status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateLastExecuteStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LastExecuteStatusNotIn); i++ {

		if err := m.LastExecuteStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_execute_status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotPlanWhereInput) contextValidateStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPlanWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPlanWhereInput) UnmarshalBinary(b []byte) error {
	var res SnapshotPlanWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}