// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SnapshotPlanOrderByInput snapshot plan order by input
//
// swagger:model SnapshotPlanOrderByInput
type SnapshotPlanOrderByInput string

func NewSnapshotPlanOrderByInput(value SnapshotPlanOrderByInput) *SnapshotPlanOrderByInput {
	v := value
	return &v
}

const (

	// SnapshotPlanOrderByInputAutoDeleteNumASC captures enum value "auto_delete_num_ASC"
	SnapshotPlanOrderByInputAutoDeleteNumASC SnapshotPlanOrderByInput = "auto_delete_num_ASC"

	// SnapshotPlanOrderByInputAutoDeleteNumDESC captures enum value "auto_delete_num_DESC"
	SnapshotPlanOrderByInputAutoDeleteNumDESC SnapshotPlanOrderByInput = "auto_delete_num_DESC"

	// SnapshotPlanOrderByInputAutoExecuteNumASC captures enum value "auto_execute_num_ASC"
	SnapshotPlanOrderByInputAutoExecuteNumASC SnapshotPlanOrderByInput = "auto_execute_num_ASC"

	// SnapshotPlanOrderByInputAutoExecuteNumDESC captures enum value "auto_execute_num_DESC"
	SnapshotPlanOrderByInputAutoExecuteNumDESC SnapshotPlanOrderByInput = "auto_execute_num_DESC"

	// SnapshotPlanOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	SnapshotPlanOrderByInputCreatedAtASC SnapshotPlanOrderByInput = "createdAt_ASC"

	// SnapshotPlanOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	SnapshotPlanOrderByInputCreatedAtDESC SnapshotPlanOrderByInput = "createdAt_DESC"

	// SnapshotPlanOrderByInputDescriptionASC captures enum value "description_ASC"
	SnapshotPlanOrderByInputDescriptionASC SnapshotPlanOrderByInput = "description_ASC"

	// SnapshotPlanOrderByInputDescriptionDESC captures enum value "description_DESC"
	SnapshotPlanOrderByInputDescriptionDESC SnapshotPlanOrderByInput = "description_DESC"

	// SnapshotPlanOrderByInputEndTimeASC captures enum value "end_time_ASC"
	SnapshotPlanOrderByInputEndTimeASC SnapshotPlanOrderByInput = "end_time_ASC"

	// SnapshotPlanOrderByInputEndTimeDESC captures enum value "end_time_DESC"
	SnapshotPlanOrderByInputEndTimeDESC SnapshotPlanOrderByInput = "end_time_DESC"

	// SnapshotPlanOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	SnapshotPlanOrderByInputEntityAsyncStatusASC SnapshotPlanOrderByInput = "entityAsyncStatus_ASC"

	// SnapshotPlanOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	SnapshotPlanOrderByInputEntityAsyncStatusDESC SnapshotPlanOrderByInput = "entityAsyncStatus_DESC"

	// SnapshotPlanOrderByInputExechmASC captures enum value "exec_h_m_ASC"
	SnapshotPlanOrderByInputExechmASC SnapshotPlanOrderByInput = "exec_h_m_ASC"

	// SnapshotPlanOrderByInputExechmDESC captures enum value "exec_h_m_DESC"
	SnapshotPlanOrderByInputExechmDESC SnapshotPlanOrderByInput = "exec_h_m_DESC"

	// SnapshotPlanOrderByInputExecutePlanTypeASC captures enum value "execute_plan_type_ASC"
	SnapshotPlanOrderByInputExecutePlanTypeASC SnapshotPlanOrderByInput = "execute_plan_type_ASC"

	// SnapshotPlanOrderByInputExecutePlanTypeDESC captures enum value "execute_plan_type_DESC"
	SnapshotPlanOrderByInputExecutePlanTypeDESC SnapshotPlanOrderByInput = "execute_plan_type_DESC"

	// SnapshotPlanOrderByInputHealthyASC captures enum value "healthy_ASC"
	SnapshotPlanOrderByInputHealthyASC SnapshotPlanOrderByInput = "healthy_ASC"

	// SnapshotPlanOrderByInputHealthyDESC captures enum value "healthy_DESC"
	SnapshotPlanOrderByInputHealthyDESC SnapshotPlanOrderByInput = "healthy_DESC"

	// SnapshotPlanOrderByInputIDASC captures enum value "id_ASC"
	SnapshotPlanOrderByInputIDASC SnapshotPlanOrderByInput = "id_ASC"

	// SnapshotPlanOrderByInputIDDESC captures enum value "id_DESC"
	SnapshotPlanOrderByInputIDDESC SnapshotPlanOrderByInput = "id_DESC"

	// SnapshotPlanOrderByInputLastExecuteEndTimeASC captures enum value "last_execute_end_time_ASC"
	SnapshotPlanOrderByInputLastExecuteEndTimeASC SnapshotPlanOrderByInput = "last_execute_end_time_ASC"

	// SnapshotPlanOrderByInputLastExecuteEndTimeDESC captures enum value "last_execute_end_time_DESC"
	SnapshotPlanOrderByInputLastExecuteEndTimeDESC SnapshotPlanOrderByInput = "last_execute_end_time_DESC"

	// SnapshotPlanOrderByInputLastExecuteStatusASC captures enum value "last_execute_status_ASC"
	SnapshotPlanOrderByInputLastExecuteStatusASC SnapshotPlanOrderByInput = "last_execute_status_ASC"

	// SnapshotPlanOrderByInputLastExecuteStatusDESC captures enum value "last_execute_status_DESC"
	SnapshotPlanOrderByInputLastExecuteStatusDESC SnapshotPlanOrderByInput = "last_execute_status_DESC"

	// SnapshotPlanOrderByInputLastExecuteTimeASC captures enum value "last_execute_time_ASC"
	SnapshotPlanOrderByInputLastExecuteTimeASC SnapshotPlanOrderByInput = "last_execute_time_ASC"

	// SnapshotPlanOrderByInputLastExecuteTimeDESC captures enum value "last_execute_time_DESC"
	SnapshotPlanOrderByInputLastExecuteTimeDESC SnapshotPlanOrderByInput = "last_execute_time_DESC"

	// SnapshotPlanOrderByInputLocalIDASC captures enum value "local_id_ASC"
	SnapshotPlanOrderByInputLocalIDASC SnapshotPlanOrderByInput = "local_id_ASC"

	// SnapshotPlanOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	SnapshotPlanOrderByInputLocalIDDESC SnapshotPlanOrderByInput = "local_id_DESC"

	// SnapshotPlanOrderByInputLogicalSizeBytesASC captures enum value "logical_size_bytes_ASC"
	SnapshotPlanOrderByInputLogicalSizeBytesASC SnapshotPlanOrderByInput = "logical_size_bytes_ASC"

	// SnapshotPlanOrderByInputLogicalSizeBytesDESC captures enum value "logical_size_bytes_DESC"
	SnapshotPlanOrderByInputLogicalSizeBytesDESC SnapshotPlanOrderByInput = "logical_size_bytes_DESC"

	// SnapshotPlanOrderByInputManualDeleteNumASC captures enum value "manual_delete_num_ASC"
	SnapshotPlanOrderByInputManualDeleteNumASC SnapshotPlanOrderByInput = "manual_delete_num_ASC"

	// SnapshotPlanOrderByInputManualDeleteNumDESC captures enum value "manual_delete_num_DESC"
	SnapshotPlanOrderByInputManualDeleteNumDESC SnapshotPlanOrderByInput = "manual_delete_num_DESC"

	// SnapshotPlanOrderByInputManualExecuteNumASC captures enum value "manual_execute_num_ASC"
	SnapshotPlanOrderByInputManualExecuteNumASC SnapshotPlanOrderByInput = "manual_execute_num_ASC"

	// SnapshotPlanOrderByInputManualExecuteNumDESC captures enum value "manual_execute_num_DESC"
	SnapshotPlanOrderByInputManualExecuteNumDESC SnapshotPlanOrderByInput = "manual_execute_num_DESC"

	// SnapshotPlanOrderByInputMirrorASC captures enum value "mirror_ASC"
	SnapshotPlanOrderByInputMirrorASC SnapshotPlanOrderByInput = "mirror_ASC"

	// SnapshotPlanOrderByInputMirrorDESC captures enum value "mirror_DESC"
	SnapshotPlanOrderByInputMirrorDESC SnapshotPlanOrderByInput = "mirror_DESC"

	// SnapshotPlanOrderByInputNameASC captures enum value "name_ASC"
	SnapshotPlanOrderByInputNameASC SnapshotPlanOrderByInput = "name_ASC"

	// SnapshotPlanOrderByInputNameDESC captures enum value "name_DESC"
	SnapshotPlanOrderByInputNameDESC SnapshotPlanOrderByInput = "name_DESC"

	// SnapshotPlanOrderByInputNextExecuteTimeASC captures enum value "next_execute_time_ASC"
	SnapshotPlanOrderByInputNextExecuteTimeASC SnapshotPlanOrderByInput = "next_execute_time_ASC"

	// SnapshotPlanOrderByInputNextExecuteTimeDESC captures enum value "next_execute_time_DESC"
	SnapshotPlanOrderByInputNextExecuteTimeDESC SnapshotPlanOrderByInput = "next_execute_time_DESC"

	// SnapshotPlanOrderByInputObjectNumASC captures enum value "object_num_ASC"
	SnapshotPlanOrderByInputObjectNumASC SnapshotPlanOrderByInput = "object_num_ASC"

	// SnapshotPlanOrderByInputObjectNumDESC captures enum value "object_num_DESC"
	SnapshotPlanOrderByInputObjectNumDESC SnapshotPlanOrderByInput = "object_num_DESC"

	// SnapshotPlanOrderByInputPhysicalSizeBytesASC captures enum value "physical_size_bytes_ASC"
	SnapshotPlanOrderByInputPhysicalSizeBytesASC SnapshotPlanOrderByInput = "physical_size_bytes_ASC"

	// SnapshotPlanOrderByInputPhysicalSizeBytesDESC captures enum value "physical_size_bytes_DESC"
	SnapshotPlanOrderByInputPhysicalSizeBytesDESC SnapshotPlanOrderByInput = "physical_size_bytes_DESC"

	// SnapshotPlanOrderByInputRemainSnapshotNumASC captures enum value "remain_snapshot_num_ASC"
	SnapshotPlanOrderByInputRemainSnapshotNumASC SnapshotPlanOrderByInput = "remain_snapshot_num_ASC"

	// SnapshotPlanOrderByInputRemainSnapshotNumDESC captures enum value "remain_snapshot_num_DESC"
	SnapshotPlanOrderByInputRemainSnapshotNumDESC SnapshotPlanOrderByInput = "remain_snapshot_num_DESC"

	// SnapshotPlanOrderByInputSnapshotGroupNumASC captures enum value "snapshot_group_num_ASC"
	SnapshotPlanOrderByInputSnapshotGroupNumASC SnapshotPlanOrderByInput = "snapshot_group_num_ASC"

	// SnapshotPlanOrderByInputSnapshotGroupNumDESC captures enum value "snapshot_group_num_DESC"
	SnapshotPlanOrderByInputSnapshotGroupNumDESC SnapshotPlanOrderByInput = "snapshot_group_num_DESC"

	// SnapshotPlanOrderByInputStartTimeASC captures enum value "start_time_ASC"
	SnapshotPlanOrderByInputStartTimeASC SnapshotPlanOrderByInput = "start_time_ASC"

	// SnapshotPlanOrderByInputStartTimeDESC captures enum value "start_time_DESC"
	SnapshotPlanOrderByInputStartTimeDESC SnapshotPlanOrderByInput = "start_time_DESC"

	// SnapshotPlanOrderByInputStatusASC captures enum value "status_ASC"
	SnapshotPlanOrderByInputStatusASC SnapshotPlanOrderByInput = "status_ASC"

	// SnapshotPlanOrderByInputStatusDESC captures enum value "status_DESC"
	SnapshotPlanOrderByInputStatusDESC SnapshotPlanOrderByInput = "status_DESC"

	// SnapshotPlanOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	SnapshotPlanOrderByInputUpdatedAtASC SnapshotPlanOrderByInput = "updatedAt_ASC"

	// SnapshotPlanOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	SnapshotPlanOrderByInputUpdatedAtDESC SnapshotPlanOrderByInput = "updatedAt_DESC"
)

// for schema
var snapshotPlanOrderByInputEnum []interface{}

func init() {
	var res []SnapshotPlanOrderByInput
	if err := json.Unmarshal([]byte(`["auto_delete_num_ASC","auto_delete_num_DESC","auto_execute_num_ASC","auto_execute_num_DESC","createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","end_time_ASC","end_time_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","exec_h_m_ASC","exec_h_m_DESC","execute_plan_type_ASC","execute_plan_type_DESC","healthy_ASC","healthy_DESC","id_ASC","id_DESC","last_execute_end_time_ASC","last_execute_end_time_DESC","last_execute_status_ASC","last_execute_status_DESC","last_execute_time_ASC","last_execute_time_DESC","local_id_ASC","local_id_DESC","logical_size_bytes_ASC","logical_size_bytes_DESC","manual_delete_num_ASC","manual_delete_num_DESC","manual_execute_num_ASC","manual_execute_num_DESC","mirror_ASC","mirror_DESC","name_ASC","name_DESC","next_execute_time_ASC","next_execute_time_DESC","object_num_ASC","object_num_DESC","physical_size_bytes_ASC","physical_size_bytes_DESC","remain_snapshot_num_ASC","remain_snapshot_num_DESC","snapshot_group_num_ASC","snapshot_group_num_DESC","start_time_ASC","start_time_DESC","status_ASC","status_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotPlanOrderByInputEnum = append(snapshotPlanOrderByInputEnum, v)
	}
}

func (m SnapshotPlanOrderByInput) validateSnapshotPlanOrderByInputEnum(path, location string, value SnapshotPlanOrderByInput) error {
	if err := validate.EnumCase(path, location, value, snapshotPlanOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this snapshot plan order by input
func (m SnapshotPlanOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSnapshotPlanOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this snapshot plan order by input based on context it is used
func (m SnapshotPlanOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}