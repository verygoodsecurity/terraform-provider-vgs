package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
	vgstools "github.com/verygoodsecurity/vgs-api-client-go/tools"
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

	cfg := config(d)
	c := m.(vgs.VgsClientsFacade)
	if _, err := c.Vaults(cfg).RetrieveVault(vaultId); err != nil {
		return diag.FromErr(err)
	}
	routeId, err := vgstools.RouteIdFromYaml(d.Get("inline_config").(string))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to extract route ID"))
	}
	_, err = c.Routes(cfg).GetRoute(vaultId, routeId)

	// TODO set yaml?

	return diag.FromErr(err)
}
func createRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	vaultId := d.Get("vault").(string)

	cfg := config(d)
	c := m.(vgs.VgsClientsFacade)
	if _, err := c.Vaults(cfg).RetrieveVault(vaultId); err != nil {
		return diag.FromErr(err)
	}
	routeYaml := d.Get("inline_config").(string)
	routeId, err := c.Routes(cfg).ImportRoute(vaultId, strings.NewReader(routeYaml))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(routeId)
	return nil
}

func updateRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChange("vault") {
		return diag.Errorf("Vault ID update is not allowed")
	}
	if d.HasChange("environment") {
		return diag.Errorf("Environment update is not allowed")
	}
	return createRoute(ctx, d, m)
}

func deleteRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	vaultId := d.Get("vault").(string)

	cfg := config(d)
	c := m.(vgs.VgsClientsFacade)
	if _, err := c.Vaults(cfg).RetrieveVault(vaultId); err != nil {
		return diag.FromErr(err)
	}
	id, err := vgstools.RouteIdFromYaml(d.Get("inline_config").(string))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to extract route ID"))
	}

	return diag.FromErr(c.Routes(cfg).DeleteRoute(vaultId, id))
}

func config(d *schema.ResourceData) vgs.ClientConfig {
	dataEnv := strings.ToLower(d.Get("environment").(string))
	return vgs.EnvironmentConfig().
		WithFallback(vgs.DynamicConfig().
			AddParameter("VGS_KEYCLOAK_URL", "https://auth.verygoodsecurity.com").
			AddParameter("VGS_KEYCLOAK_REALM", "vgs").
			AddParameter("VGS_ACCOUNT_MANAGEMENT_API_BASE_URL", "https://accounts.apps.verygoodsecurity.com").
			AddParameter("VGS_VAULT_MANAGEMENT_API_BASE_URL", fmt.Sprintf("https://api.%s.verygoodsecurity.com", dataEnv)))
}
