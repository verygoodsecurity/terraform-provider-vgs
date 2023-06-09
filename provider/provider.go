package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	tfprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/verygoodsecurity/terraform-provider-vgs/model"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
)

type VgsProvider struct{}

func Provider() tfprovider.Provider {
	return &VgsProvider{}
}

func (*VgsProvider) Metadata(_ context.Context, _ tfprovider.MetadataRequest, resp *tfprovider.MetadataResponse) {
	resp.TypeName = "vgs"
}

func (*VgsProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

func (*VgsProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResourceVault,
		NewResourceRoute,
	}
}

func (*VgsProvider) Schema(ctx context.Context, req tfprovider.SchemaRequest, resp *tfprovider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"environment": schema.StringAttribute{
				Required:    true,
				Description: "VGS environment. One of 'sandbox', 'live', 'live-eu-1'",
				Validators: []validator.String{
					stringvalidator.OneOf("sandbox", "live", "live-eu-1"),
				},
			},
			"organization": schema.StringAttribute{
				Required:    true,
				Description: "VGS organization ID",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(24),
					stringvalidator.LengthAtMost(24),
				},
			},
		},
	}
}

func (*VgsProvider) Configure(ctx context.Context, req tfprovider.ConfigureRequest, resp *tfprovider.ConfigureResponse) {
	var config model.Provider
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		clientId, clientSecret string
	)

	if !config.ClientId.IsNull() {
		clientId = config.ClientId.ValueString()
	} else if existsEnv("VGS_CLIENT_ID") {
		clientId = os.Getenv("VGS_CLIENT_ID")
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing VGS client ID",
			"Make sure that either 'client_id' provider attribute or 'VGS_CLIENT_ID' environment variable is set.",
		)
	}

	if !config.ClientSecret.IsNull() {
		clientSecret = config.ClientSecret.ValueString()
	} else if existsEnv("VGS_CLIENT_SECRET") {
		clientSecret = os.Getenv("VGS_CLIENT_SECRET")
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing VGS client ID",
			"Make sure that either 'client_secret' provider attribute or 'VGS_CLIENT_SECRET' environment variable is set.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	// todo rewrite in a smarter way
	facade := contextualVgsFacade{
		vgs.NewVgsFacade(vgs.EnvironmentConfig().
			WithFallback(vgs.DynamicConfig().
				AddParameter("VGS_CLIENT_ID", clientId).
				AddParameter("VGS_CLIENT_SECRET", clientSecret).
				AddParameter("VGS_KEYCLOAK_URL", "https://auth.verygoodsecurity.com").
				AddParameter("VGS_KEYCLOAK_REALM", "vgs").
				AddParameter("VGS_ACCOUNT_MANAGEMENT_API_BASE_URL", "https://accounts.apps.verygoodsecurity.com").
				AddParameter("VGS_VAULT_MANAGEMENT_API_BASE_URL", fmt.Sprintf("https://api.%s.verygoodsecurity.com", config.Environment.ValueString())))),
		config.Organization.ValueString(),
		config.Environment.ValueString(),
	}
	resp.DataSourceData = facade
	resp.ResourceData = facade
}

func existsEnv(name string) bool {
	return os.Getenv(name) != ""
}

type contextualVgsFacade struct {
	vgs.VgsClientsFacade
	organizationId string
	evironment     string
}
