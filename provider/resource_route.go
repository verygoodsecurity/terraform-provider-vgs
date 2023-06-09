package provider

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &resourceRoute{}
	_ resource.ResourceWithConfigure = &resourceRoute{}
)

type resourceRoute struct{}

func NewResourceRoute() resource.Resource {
	return &resourceRoute{}
}

func (*resourceRoute) Configure(context.Context, resource.ConfigureRequest, *resource.ConfigureResponse) {
	//todo
}

func (*resourceRoute) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route"
}

func (*resourceRoute) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "", // TODO
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vault": schema.StringAttribute{
				Required:    true,
				Description: "VGS Vault ID (identifier). Usually looks like 'tntxxxxxxxx'.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^tnt[a-z0-9]{8}$"), ""),
				},
			},
			"protocol": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP", "HTTPS", "SFTP", "TCP"),
				},
			},
			"source_endpoint": schema.StringAttribute{
				Optional:    true,
			},
			"destination_override_endpoint": schema.StringAttribute{
				Optional:    true,
			},
			"host_endpoint": schema.StringAttribute{
				Optional:    true,
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65536),
				},
			},
			"ordinal": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"tags": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"filters": schema.ListNestedAttribute{
				Required:    true,
				Description: "todo",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"phase": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("REQUEST", "RESPONSE"),
							},
						},
						"operation": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("REDACT", "REVEAL"),
							},
						},
						"alias_format": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"UUID",
									"RAW_UUID",
									"FPE_SIX_T_FOUR",
									"FPE_T_FOUR",
									"NUM_LENGTH_PRESERVING",
									"PFPT",
									"FPE_SSN_T_FOUR",
									"FPE_ACC_NUM_T_FOUR",
									"FPE_ALPHANUMERIC_ACC_NUM_T_FOUR",
									"GENERIC_T_FOUR",
									"ALPHANUMERIC_SIX_T_FOUR",
									"NON_LUHN_FPE_ALPHANUMERIC",
									"VGS_FIXED_LEN_GENERIC",
									"CUSTOM"),
							},
						},
						"transformer": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"content_type": schema.StringAttribute{
									Required: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"FORM_FIELD",
											"HTML",
											"JSON_PATH",
											"REGEX",
											"MAPPING_REGEX",
											"XPATH",
											"SAMPLE",
											"CSV",
											"TRANSFORMER",
											"PDF",
											"FILE_STORAGE",
											"PDF_METADATA_TOKEN",
											"ZIP"),
									},
								},
								"config": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"has_header": schema.BoolAttribute{
											Optional: true,
										},
										"column_indices": schema.StringAttribute{
											Optional: true,
										},
										// TODO all fields from transfromerConfig and transformerConfigMap
									},
								},
							},
						},
						"targets": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"id_selector": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.LengthAtMost(100),
							},
						},
						"conditions_inline": schema.StringAttribute{
							Required:    false,
							Description: "Expressions for filters (JSON)", // TODO seems there's no way represent recursive schema
						},
						"classifiers": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"include": schema.ListAttribute{
									Optional:    true,
									ElementType: types.StringType,
								},
								"exclude": schema.ListAttribute{
									Optional:    true,
									ElementType: types.StringType,
								},
								"tags": schema.ListAttribute{
									Optional:    true,
									ElementType: types.StringType,
								},
							},
						},
						"operations": schema.StringAttribute{
							Optional:    true,
							Description: "VGS Operation Pipeline configuration", // TODO?
						},
					},
				},
			},
		},
	}
}

// Create implements resource.Resource
func (*resourceRoute) Create(context.Context, resource.CreateRequest, *resource.CreateResponse) {
	panic("unimplemented")
}

// Read implements resource.Resource
func (*resourceRoute) Read(context.Context, resource.ReadRequest, *resource.ReadResponse) {
	panic("unimplemented")
}

// Update implements resource.Resource
func (*resourceRoute) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {
	panic("unimplemented")
}

// Delete implements resource.Resource
func (*resourceRoute) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
	panic("unimplemented")
}
