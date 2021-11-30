package provider

import (
	asserting "github.com/stretchr/testify/assert"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
	"testing"
)

var (
	provider   = Provider()
	testConfig = vgs.EnvironmentConfig()
)

const (
	tenant        = "tntbcduzut5"
	testTerraform = `
resource "vgs_route" "inbound_route" {
  environment = "sandbox"
  vault = "tntbcduzut5"
  inline_config = <<EOF
id: eccf1542-820b-4e7e-8b3f-03c4d8639f9d
type: rule_chain
attributes:
  tags:
    name: test-inbound-route
    source: RouteContainer
  destination_override_endpoint: 'https://echo.apps.verygood.systems'
  host_endpoint: (.*)\.verygoodproxy\.io
  id: eccf1542-820b-4e7e-8b3f-03c4d8639f9d
  ordinal: null
  port: 80
  protocol: http
  source_endpoint: '*'
  entries:
  - classifiers: {}
    config:
      condition: AND
      rules:
        - expression:
            field: PathInfo
            operator: matches
            type: string
            values:
              - /post
        - expression:
            field: ContentType
            operator: equals
            type: string
            values:
              - application/json
          rules: null
    id: 39f2f5db-06a0-461d-9387-dd9a7ab19035
    id_selector: null
    operation: REDACT
    operations: null
    phase: REQUEST
    public_token_generator: UUID
    targets:
      - body
    token_manager: PERSISTENT
    transformer: JSON_PATH
    transformer_config:
      - $.account_number
    transformer_config_map: null
EOF
}
`
)

func Test(t *testing.T) {
	assert := asserting.New(t)

	assert.Nil(os.Setenv("VGS_KEYCLOAK_REALM", "vgs"))
	assert.Nil(os.Setenv("VGS_KEYCLOAK_URL", "https://auth.verygoodsecurity.io"))
	assert.Nil(os.Setenv("VGS_ACCOUNT_MANAGEMENT_API_BASE_URL", "https://accounts.verygoodsecurity.io"))
	assert.Nil(os.Setenv("VGS_VAULT_MANAGEMENT_API_BASE_URL", "https://api.verygoodsecurity.io"))

	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"vgs": func() (*schema.Provider, error) {
				return provider, nil
			},
		},
		CheckDestroy: func(state *terraform.State) error {
			routes := provider.Meta().(vgs.VgsClientsFacade).Routes(testConfig)
			for _, rs := range state.RootModule().Resources {
				if rs.Type == "vgs_route" {
					r, err := routes.GetRoute(tenant, rs.Primary.ID)
					assert.Empty(r)
					if assert.NotNil(err) {
						assert.Contains(err.Error(), "Route not found")
					}
				}
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testTerraform,
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						routes := provider.Meta().(vgs.VgsClientsFacade).Routes(testConfig)
						for _, rs := range state.RootModule().Resources {
							if rs.Type == "vgs_route" {
								assert.NotEmpty(rs.Primary.ID)
								r, err := routes.GetRoute(tenant, rs.Primary.ID)
								assert.NotEmpty(r)
								assert.Nil(err)
							}
						}
						return nil
					},
				),
			},
		},
	})
}
