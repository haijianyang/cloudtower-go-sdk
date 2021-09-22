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

// GraphWhereInput graph where input
//
// swagger:model GraphWhereInput
type GraphWhereInput struct {

	// a n d
	AND []*GraphWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*GraphWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*GraphWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster interface{} `json:"cluster,omitempty"`

	// disks every
	DisksEvery interface{} `json:"disks_every,omitempty"`

	// disks none
	DisksNone interface{} `json:"disks_none,omitempty"`

	// disks some
	DisksSome interface{} `json:"disks_some,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot interface{} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// hosts every
	HostsEvery interface{} `json:"hosts_every,omitempty"`

	// hosts none
	HostsNone interface{} `json:"hosts_none,omitempty"`

	// hosts some
	HostsSome interface{} `json:"hosts_some,omitempty"`

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

	// luns every
	LunsEvery interface{} `json:"luns_every,omitempty"`

	// luns none
	LunsNone interface{} `json:"luns_none,omitempty"`

	// luns some
	LunsSome interface{} `json:"luns_some,omitempty"`

	// metric count
	MetricCount *float64 `json:"metric_count,omitempty"`

	// metric count gt
	MetricCountGt *float64 `json:"metric_count_gt,omitempty"`

	// metric count gte
	MetricCountGte *float64 `json:"metric_count_gte,omitempty"`

	// metric count in
	MetricCountIn []float64 `json:"metric_count_in,omitempty"`

	// metric count lt
	MetricCountLt *float64 `json:"metric_count_lt,omitempty"`

	// metric count lte
	MetricCountLte *float64 `json:"metric_count_lte,omitempty"`

	// metric count not
	MetricCountNot *float64 `json:"metric_count_not,omitempty"`

	// metric count not in
	MetricCountNotIn []float64 `json:"metric_count_not_in,omitempty"`

	// metric name
	MetricName *string `json:"metric_name,omitempty"`

	// metric name contains
	MetricNameContains *string `json:"metric_name_contains,omitempty"`

	// metric name ends with
	MetricNameEndsWith *string `json:"metric_name_ends_with,omitempty"`

	// metric name gt
	MetricNameGt *string `json:"metric_name_gt,omitempty"`

	// metric name gte
	MetricNameGte *string `json:"metric_name_gte,omitempty"`

	// metric name in
	MetricNameIn []string `json:"metric_name_in,omitempty"`

	// metric name lt
	MetricNameLt *string `json:"metric_name_lt,omitempty"`

	// metric name lte
	MetricNameLte *string `json:"metric_name_lte,omitempty"`

	// metric name not
	MetricNameNot *string `json:"metric_name_not,omitempty"`

	// metric name not contains
	MetricNameNotContains *string `json:"metric_name_not_contains,omitempty"`

	// metric name not ends with
	MetricNameNotEndsWith *string `json:"metric_name_not_ends_with,omitempty"`

	// metric name not in
	MetricNameNotIn []string `json:"metric_name_not_in,omitempty"`

	// metric name not starts with
	MetricNameNotStartsWith *string `json:"metric_name_not_starts_with,omitempty"`

	// metric name starts with
	MetricNameStartsWith *string `json:"metric_name_starts_with,omitempty"`

	// metric type
	MetricType interface{} `json:"metric_type,omitempty"`

	// metric type in
	MetricTypeIn []MetricType `json:"metric_type_in,omitempty"`

	// metric type not
	MetricTypeNot interface{} `json:"metric_type_not,omitempty"`

	// metric type not in
	MetricTypeNotIn []MetricType `json:"metric_type_not_in,omitempty"`

	// namespaces every
	NamespacesEvery interface{} `json:"namespaces_every,omitempty"`

	// namespaces none
	NamespacesNone interface{} `json:"namespaces_none,omitempty"`

	// namespaces some
	NamespacesSome interface{} `json:"namespaces_some,omitempty"`

	// network
	Network interface{} `json:"network,omitempty"`

	// network in
	NetworkIn []NetworkType `json:"network_in,omitempty"`

	// network not
	NetworkNot interface{} `json:"network_not,omitempty"`

	// network not in
	NetworkNotIn []NetworkType `json:"network_not_in,omitempty"`

	// nics every
	NicsEvery interface{} `json:"nics_every,omitempty"`

	// nics none
	NicsNone interface{} `json:"nics_none,omitempty"`

	// nics some
	NicsSome interface{} `json:"nics_some,omitempty"`

	// resource type
	ResourceType *string `json:"resource_type,omitempty"`

	// resource type contains
	ResourceTypeContains *string `json:"resource_type_contains,omitempty"`

	// resource type ends with
	ResourceTypeEndsWith *string `json:"resource_type_ends_with,omitempty"`

	// resource type gt
	ResourceTypeGt *string `json:"resource_type_gt,omitempty"`

	// resource type gte
	ResourceTypeGte *string `json:"resource_type_gte,omitempty"`

	// resource type in
	ResourceTypeIn []string `json:"resource_type_in,omitempty"`

	// resource type lt
	ResourceTypeLt *string `json:"resource_type_lt,omitempty"`

	// resource type lte
	ResourceTypeLte *string `json:"resource_type_lte,omitempty"`

	// resource type not
	ResourceTypeNot *string `json:"resource_type_not,omitempty"`

	// resource type not contains
	ResourceTypeNotContains *string `json:"resource_type_not_contains,omitempty"`

	// resource type not ends with
	ResourceTypeNotEndsWith *string `json:"resource_type_not_ends_with,omitempty"`

	// resource type not in
	ResourceTypeNotIn []string `json:"resource_type_not_in,omitempty"`

	// resource type not starts with
	ResourceTypeNotStartsWith *string `json:"resource_type_not_starts_with,omitempty"`

	// resource type starts with
	ResourceTypeStartsWith *string `json:"resource_type_starts_with,omitempty"`

	// service
	Service *string `json:"service,omitempty"`

	// service contains
	ServiceContains *string `json:"service_contains,omitempty"`

	// service ends with
	ServiceEndsWith *string `json:"service_ends_with,omitempty"`

	// service gt
	ServiceGt *string `json:"service_gt,omitempty"`

	// service gte
	ServiceGte *string `json:"service_gte,omitempty"`

	// service in
	ServiceIn []string `json:"service_in,omitempty"`

	// service lt
	ServiceLt *string `json:"service_lt,omitempty"`

	// service lte
	ServiceLte *string `json:"service_lte,omitempty"`

	// service not
	ServiceNot *string `json:"service_not,omitempty"`

	// service not contains
	ServiceNotContains *string `json:"service_not_contains,omitempty"`

	// service not ends with
	ServiceNotEndsWith *string `json:"service_not_ends_with,omitempty"`

	// service not in
	ServiceNotIn []string `json:"service_not_in,omitempty"`

	// service not starts with
	ServiceNotStartsWith *string `json:"service_not_starts_with,omitempty"`

	// service starts with
	ServiceStartsWith *string `json:"service_starts_with,omitempty"`

	// title
	Title *string `json:"title,omitempty"`

	// title contains
	TitleContains *string `json:"title_contains,omitempty"`

	// title ends with
	TitleEndsWith *string `json:"title_ends_with,omitempty"`

	// title gt
	TitleGt *string `json:"title_gt,omitempty"`

	// title gte
	TitleGte *string `json:"title_gte,omitempty"`

	// title in
	TitleIn []string `json:"title_in,omitempty"`

	// title lt
	TitleLt *string `json:"title_lt,omitempty"`

	// title lte
	TitleLte *string `json:"title_lte,omitempty"`

	// title not
	TitleNot *string `json:"title_not,omitempty"`

	// title not contains
	TitleNotContains *string `json:"title_not_contains,omitempty"`

	// title not ends with
	TitleNotEndsWith *string `json:"title_not_ends_with,omitempty"`

	// title not in
	TitleNotIn []string `json:"title_not_in,omitempty"`

	// title not starts with
	TitleNotStartsWith *string `json:"title_not_starts_with,omitempty"`

	// title starts with
	TitleStartsWith *string `json:"title_starts_with,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// type in
	TypeIn []GraphType `json:"type_in,omitempty"`

	// type not
	TypeNot interface{} `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []GraphType `json:"type_not_in,omitempty"`

	// view
	View interface{} `json:"view,omitempty"`

	// vm nics every
	VMNicsEvery interface{} `json:"vmNics_every,omitempty"`

	// vm nics none
	VMNicsNone interface{} `json:"vmNics_none,omitempty"`

	// vm nics some
	VMNicsSome interface{} `json:"vmNics_some,omitempty"`

	// vm volumes every
	VMVolumesEvery interface{} `json:"vmVolumes_every,omitempty"`

	// vm volumes none
	VMVolumesNone interface{} `json:"vmVolumes_none,omitempty"`

	// vm volumes some
	VMVolumesSome interface{} `json:"vmVolumes_some,omitempty"`

	// vms every
	VmsEvery interface{} `json:"vms_every,omitempty"`

	// vms none
	VmsNone interface{} `json:"vms_none,omitempty"`

	// vms some
	VmsSome interface{} `json:"vms_some,omitempty"`

	// witnesses every
	WitnessesEvery interface{} `json:"witnesses_every,omitempty"`

	// witnesses none
	WitnessesNone interface{} `json:"witnesses_none,omitempty"`

	// witnesses some
	WitnessesSome interface{} `json:"witnesses_some,omitempty"`

	// zones every
	ZonesEvery interface{} `json:"zones_every,omitempty"`

	// zones none
	ZonesNone interface{} `json:"zones_none,omitempty"`

	// zones some
	ZonesSome interface{} `json:"zones_some,omitempty"`
}

// Validate validates this graph where input
func (m *GraphWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateMetricTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *GraphWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *GraphWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *GraphWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *GraphWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

func (m *GraphWhereInput) validateMetricTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricTypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.MetricTypeIn); i++ {

		if err := m.MetricTypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) validateMetricTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricTypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.MetricTypeNotIn); i++ {

		if err := m.MetricTypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) validateNetworkIn(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkIn) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkIn); i++ {

		if err := m.NetworkIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) validateNetworkNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkNotIn); i++ {

		if err := m.NetworkNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) validateTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this graph where input based on the context it is used
func (m *GraphWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateMetricTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetricTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GraphWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GraphWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GraphWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GraphWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GraphWhereInput) contextValidateMetricTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetricTypeIn); i++ {

		if err := m.MetricTypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) contextValidateMetricTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetricTypeNotIn); i++ {

		if err := m.MetricTypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) contextValidateNetworkIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetworkIn); i++ {

		if err := m.NetworkIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) contextValidateNetworkNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetworkNotIn); i++ {

		if err := m.NetworkNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GraphWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GraphWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphWhereInput) UnmarshalBinary(b []byte) error {
	var res GraphWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
