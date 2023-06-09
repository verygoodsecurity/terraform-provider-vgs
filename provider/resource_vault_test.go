package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	asserting "github.com/stretchr/testify/assert"
	vgs "github.com/verygoodsecurity/vgs-api-client-go/clients"
)

const (
	testTerraform = `
	provider "vgs" {
		client_id = "ACs55NN6p-vgs-cli-zUjmf"
		client_secret = "4971ef94-7d12-460b-bfc7-5a2ed75eb037"
		environment = "sandbox"
		organization = "ACrF5NY4vxzj8BvhU3q8btsN"
	}
	
	resource "vgs_vault" "sandbox_test_vault" {
		name = "Sandbox Test Vault"
	}
`
)

func TestAccCreatingVault(t *testing.T) {
	assert := asserting.New(t)

	assert.Nil(os.Setenv("TF_ACC", "true"))
	assert.Nil(os.Setenv("VGS_KEYCLOAK_REALM", "vgs"))
	assert.Nil(os.Setenv("VGS_KEYCLOAK_URL", "https://auth.verygoodsecurity.com"))
	assert.Nil(os.Setenv("VGS_ACCOUNT_MANAGEMENT_API_BASE_URL", "https://accounts.apps.verygoodsecurity.com"))
	assert.Nil(os.Setenv("VGS_VAULT_MANAGEMENT_API_BASE_URL", "https://api.sandbox.verygoodsecurity.com"))

	vaults := vgs.NewVgsFacade(vgs.EnvironmentConfig()).Vaults()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"vgs": providerserver.NewProtocol6WithError(Provider()),
		},
		CheckDestroy: func(state *terraform.State) error {
			for _, rs := range state.RootModule().Resources {
				if rs.Type == "vgs_vault" {
					r, err := vaults.RetrieveVault(rs.Primary.ID)
					assert.Empty(r)
					if assert.NotNil(err) {
						assert.Contains(err.Error(), "Vault not found")
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
						for _, rs := range state.RootModule().Resources {
							if rs.Type == "vgs_vault" {
								assert.NotEmpty(rs.Primary.ID)
								t.Logf("Primary.ID = %s", rs.Primary.ID)
								r, err := vaults.RetrieveVault(rs.Primary.ID)
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
