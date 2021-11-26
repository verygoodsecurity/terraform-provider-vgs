package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
	"os"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// VGS_CLIENT_ID and VGS_CLIENT_SECRET set only via environment
		},
		ResourcesMap: map[string]*schema.Resource{
			"vgs_route": resourceRoute(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	if !existsEnv("VGS_CLIENT_ID") || !existsEnv("VGS_CLIENT_SECRET") {
		return nil, diag.Diagnostics{
			{
				Severity: diag.Error,
				Summary:  "Unable to create VGS client",
				Detail:   "VGS Client requires VGS_CLIENT_ID and VGS_CLIENT_SECRET to be set in OS environment",
			},
		}
	}
	return vgs.NewVgsFacade(), diag.Diagnostics{}
}

func existsEnv(name string) bool {
	return os.Getenv(name) != ""
}
