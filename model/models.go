package model

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Provider struct {
	ClientId     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	Environment  types.String `tfsdk:"environment"`
	Organization types.String `tfsdk:"organization"`
}

type Vault struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Preferences types.Map    `tfsdk:"preferences"`
}

type Route struct {
	ID                          types.String `tfsdk:"id"`
	CreatedAt                   types.String `tfsdk:"created_at"`
	UpdatedAt                   types.String `tfsdk:"updated_at"`
	Vault                       types.String `tfsdk:"vault"`
	Protocol                    types.String `tfsdk:"protocol"`
	SourceEndpoint              types.String `tfsdk:"source_endpoint"`
	DestinationOverrideEndpoint types.String `tfsdk:"destination_override_endpoint"`
	HostEnpoint                 types.String `tfsdk:"host_endpoint"`
	Port                        types.Int64  `tfsdk:"port"`
	Ordinal                     types.Int64  `tfsdk:"ordinal"`
	Filters                     []Filter     `tfsdk:"filters"`
	Tags                        types.Map    `tfsdk:"tags"`
}

type Filter struct {
	ID               types.String   `tfsdk:"id"`
	Phase            types.String   `tfsdk:"phase"`
	Operation        types.String   `tfsdk:"operation"`
	AliasFormat      types.String   `tfsdk:"alias_format"`
	Transformer      types.Object   `tfsdk:"transformer"`
	Targets          []types.String `tfsdk:"targets"`
	IdSelector       types.String   `tfsdk:"id_selector"`
	ConditionsInline types.String   `tfsdk:"conditions_inline"`
	Operations       types.String   `tfsdk:"operations"`
	Classifiers      types.Map      `tfsdk:"classifiers"`
}

type Upstream struct {
	ID          types.String `tfsdk:"id"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdateAt    types.String `tfsdk:"updated_at"`
	Route       types.String `tfsdk:"route"`
	Protocol    types.String `tfsdk:"protocol"`
	Host        types.String `tfsdk:"host"`
	Port        types.Int64  `tfsdk:"port"`
	Username    types.String `tfsdk:"username"`
	PrivateKey  types.String `tfsdk:"private_key"`
	PublicKey   types.String `tfsdk:"public_key"`
	Password    types.String `tfsdk:"password"`
	Credentials types.List   `tfsdk:"credentials"`
}
