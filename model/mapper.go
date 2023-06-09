package model

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pkg/errors"
	"github.com/verygoodsecurity/vgs-api-client-go/dto"
)

func (m *Vault) Merge(v dto.Vault) {
	m.ID = types.StringValue(v.Id)
	m.Name = types.StringValue(v.Name)
}

func (m *Route) Merge(r dto.Route) error {
	m.ID = types.StringValue(r.ID)
	m.Protocol = types.StringValue(r.Protocol)
	m.Filters = make([]Filter, len(r.Filters))
	m.CreatedAt = types.StringValue(r.CreatedAt.Format("2006-01-02 15:04:05"))
	m.UpdatedAt = types.StringValue(r.UpdatedAt.Format("2006-01-02 15:04:05"))
	m.Port = types.Int64Value(int64(r.Port))
	m.Ordinal = types.Int64Value(int64(r.Ordinal))
	
	tags := make(map[string]attr.Value)
	for k, v := range r.Tags {
		tags[k] = types.StringValue(v)
	}

	m.Tags = types.MapValueMust(types.StringType, tags)
	for _, f := range r.Filters {
		var filter Filter
		if err := filter.Merge(f); err != nil {
			return err
		}
		m.Filters = append(m.Filters, filter)
	}

	return nil
}

func (m *Route) ToApiObject() (dto.Route, error) {
	var d dto.Route
	var err error
	d.ID = m.ID.String()
	if d.CreatedAt, err = time.Parse("2006-01-02 15:04:05", m.CreatedAt.String()); err != nil {
		return d, errors.Wrap(err, "Failed to parse created_at")
	}
	if d.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", m.UpdatedAt.String()); err != nil {
		return d, errors.Wrap(err, "Failed to parse updated_at")
	}
	d.SourceEndpoint = m.SourceEndpoint.String()
	d.DestinationOverrideEndpoint = m.DestinationOverrideEndpoint.String()
	d.HostEndpoint = m.HostEnpoint.String()
	d.Port = int32(m.Port.ValueInt64())
	d.Ordinal = int32(m.Ordinal.ValueInt64())
	d.Protocol = m.Protocol.String()

	d.Filters = make([]dto.Filter, len(m.Filters))
	for _, f := range m.Filters {
		filter, err := f.ToApiObject()
		if err != nil {
			return d, err
		}
		d.Filters = append(d.Filters, filter)
	}

	return d, nil
}

func (m *Filter) Merge(f dto.Filter) error {
	m.ID = types.StringValue(f.ID)
	m.Phase = types.StringValue(f.Phase)
	m.Operation = types.StringValue(f.Operation)
	m.AliasFormat = types.StringValue(f.PublicTokenGenerator)
	m.Transformer = types.ObjectValueMust(map[string]attr.Type{
		"content_type": types.StringType,
		"config": types.ObjectType{},
	}, map[string]attr.Value{
		"content_type": types.StringValue(f.RuleTransformer),
		//"config": 
	})
	m.Targets = make([]types.String, len(f.Targets))
	for _, t := range f.Targets {
		m.Targets = append(m.Targets, types.StringValue(t))
	}
	m.IdSelector = types.StringValue(f.IdSelector)

	
	conditions, err := json.Marshal(f.Config)
	if err != nil {
		return errors.Wrap(err, "Failed to parse conditions")
	}
	m.ConditionsInline = types.StringValue(string(conditions))
	operations, err := json.Marshal(f.Operations)
	if err != nil {
		return errors.Wrap(err, "Failed to parse operations")
	}
	m.Operations = types.StringValue(string(operations))
	return nil
}

func (m *Filter) ToApiObject() (dto.Filter, error) {
	var d dto.Filter
	d.ID = m.ID.String()
	d.Phase = m.Phase.String()
	d.Operation = m.Phase.String()
	d.PublicTokenGenerator = m.AliasFormat.String()
	attrs := m.Transformer.Attributes()
	d.RuleTransformer = attrs["content_type"].String()
	return d, nil
}