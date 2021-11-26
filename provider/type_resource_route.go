package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
	"regexp"
	"strings"
)

const (
	validVaultIdRegexRaw = `^tnt[a-z0-9]{8}$`
)

var (
	validVaultIdRegex     = regexp.MustCompile(validVaultIdRegexRaw)
	validDataEnvironments = []string{
		"sandbox",
		"live",
		"live-eu-1",
	}
)

func resourceRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRoute,
		ReadContext:   readRoute,
		UpdateContext: updateRoute,
		DeleteContext: deleteRoute,

		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "VGS Environment. One of [sandbox, live, live-eu-1]",
				ValidateDiagFunc: validation.ToDiagFunc(validation.
					StringInSlice(validDataEnvironments, true)),
			},
			"vault": {
				Type:        schema.TypeString,
				Required:    true,
				Description: fmt.Sprintf("Vault Identifier. Must match %s", validVaultIdRegexRaw),
				ValidateDiagFunc: validation.ToDiagFunc(validation.
					StringMatch(validVaultIdRegex, fmt.
						Sprintf("Vault identifier must match %s", validVaultIdRegexRaw))),
			},
			"inline_config": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "YAML route configuration https://www.verygoodsecurity.com/docs/features/yaml",
			},
		},
	}
}

func readRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	vaultId := d.Get("vault").(string)
	dataEnv := strings.ToLower(d.Get("environment").(string))

	config := vgs.DynamicConfig().
		AddParameter("VGS_ACCOUNT_MANAGEMENT_API_BASE_URL", "https://accounts.verygoodsecurity.com").
		AddParameter("VGS_VAULT_MANAGEMENT_API_BASE_URL", fmt.Sprintf("https://api.%s.verygoodsecurity.com", dataEnv))

	c := m.(vgs.VgsClientsFacade)
	if _, err := c.Vaults(config).RetrieveVault(vaultId); err != nil {
		return diag.FromErr(err)
	}

	// c.Route(config).RetrieveRoute(...)

	return nil
}
func createRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
func updateRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
func deleteRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
