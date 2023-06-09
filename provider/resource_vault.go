package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	model "github.com/verygoodsecurity/terraform-provider-vgs/model"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
)

var (
	_ resource.Resource              = &resourceVault{}
	_ resource.ResourceWithConfigure = &resourceVault{}
)

type resourceVault struct {
	client vgs.VgsClientsFacade
	orgId string
	env string
}

func NewResourceVault() resource.Resource {
	return &resourceVault{}
}

func (r *resourceVault) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(contextualVgsFacade)
	if !ok {
		panic("Assertion failed")
	}
	r.client = c
	r.orgId = c.organizationId
	r.env = c.evironment
}

func (*resourceVault) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vault"
}

func (*resourceVault) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "User-friendly name for your VGS vault.",
				Validators:  []validator.String{
					stringvalidator.LengthBetween(5, 100),
				},
			},
			"preferences": schema.MapAttribute{
				Optional: true,
				Description: "Vault Preferences as a key-value map",
				ElementType: types.StringType,
			},
		},
	}
}


func (r *resourceVault) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.Vault
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiVault, err := r.client.Vaults().ProvisionVault(r.orgId, vgs.CreateVaultForm{
		Name: plan.Name.ValueString(),
		Environment: r.env,
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Could not create vault",
			"Failed to create vault: "+err.Error(),
		)
		return
	}

	if !plan.Preferences.IsNull() {
		// update preferences
	}

	plan.Merge(apiVault)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *resourceVault) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var vault model.Vault
	diags := req.State.Get(ctx, &vault)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := vault.ID.ValueString()
	apiVault, err := r.client.Vaults().RetrieveVault(id)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Could not find vault by ID '%v'", id),
			"Failed to fetch vault: "+err.Error(),
		)
		return
	}

	vault.Merge(apiVault)
	diags = resp.State.Set(ctx, &vault)
	resp.Diagnostics.Append(diags...)
}

func (*resourceVault) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {
	panic("unimplemented")
}

func (*resourceVault) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
	panic("unimplemented")
}
