package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/verygoodsecurity/terraform-provider-vgs/provider/internal"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
	vgstools "github.com/verygoodsecurity/vgs-api-client-go/tools"
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
		Importer: &schema.ResourceImporter{
			StateContext: resourceRouteImport,
		},
		CustomizeDiff: customdiff.ForceNewIfChange("inline_config", isIdUpdated),

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
				Type:                  schema.TypeString,
				Required:              true,
				Description:           "YAML route configuration https://www.verygoodsecurity.com/docs/features/yaml",
				DiffSuppressFunc:      suppressEquivalentRouteDiffs,
				DiffSuppressOnRefresh: true,
			},
		},
	}
}

func suppressEquivalentRouteDiffs(k string, oldValue string, newValue string, d *schema.ResourceData) bool {

	var oldMap interface{}
	if err := yaml.Unmarshal([]byte(oldValue), &oldMap); err != nil {
		return false
	}

	var newMap interface{}
	if err := yaml.Unmarshal([]byte(newValue), &newMap); err != nil {
		return false
	}

	oldMap = internal.Convert(oldMap)
	_ = internal.SafeDeleteKey(oldMap, "attributes.created_at")
	_ = internal.SafeDeleteKey(oldMap, "attributes.updated_at")
	_ = internal.SafeDeleteKey(oldMap, "attributes.id")

	newMap = internal.Convert(newMap)
	_ = internal.SafeDeleteKey(newMap, "attributes.created_at")
	_ = internal.SafeDeleteKey(newMap, "attributes.updated_at")
	_ = internal.SafeDeleteKey(newMap, "attributes.id")

	return reflect.DeepEqual(oldMap, newMap)
}

func isIdUpdated(ctx context.Context, oldValue, newValue, meta interface{}) bool {
	oldRouteId, _ := vgstools.RouteIdFromYaml(oldValue.(string))
	newRouteId, _ := vgstools.RouteIdFromYaml(newValue.(string))

	return oldRouteId != newRouteId
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

	routeJson, err := c.Routes(cfg).GetRoute(vaultId, routeId)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to get route"))
	}

	routeYaml, err := jsonResponse2YamlRoute(routeJson)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to parse response"))
	}

	d.SetId(routeId)
	d.Set("inline_config", routeYaml)

	return diag.FromErr(err)
}
func createRoute(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	vaultId := d.Get("vault").(string)

	cfg := config(d)
	c := m.(vgs.VgsClientsFacade)
	if _, err := c.Vaults(cfg).RetrieveVault(vaultId); err != nil {
		return diag.FromErr(errors.Wrap(err, "Failed to retrieve vault"))
	}
	routeYaml := d.Get("inline_config").(string)
	routeId, err := c.Routes(cfg).ImportRoute(vaultId, strings.NewReader(routeYaml))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "Failed to import route"))
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

func resourceRouteImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	d.Set("environment", parts[0])
	d.Set("vault", parts[1])
	d.Set("inline_config", fmt.Sprintf("id: %s", parts[2]))
	d.SetId(parts[2])

	return []*schema.ResourceData{d}, nil
}

func jsonResponse2YamlRoute(jsonStr string) (yamlStr string, err error) {
	var body map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &body); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal JSON")
	}
	unwrappedBody := body["data"].(map[string]interface{})
	b, err := yaml.Marshal(unwrappedBody)
	return string(b), errors.Wrap(err, "failed to marshal YAML")
}
