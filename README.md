# VGS Terraform Provider

[![CircleCI](https://circleci.com/gh/verygoodsecurity/terraform-provider-vgs.svg?style=svg&circle-token=8ae379e820e61ec6f8e8451ebaf5ed6958fa7c13)](https://circleci.com/gh/verygoodsecurity/terraform-provider-vgs)

Custom Terraform provider that allows provisioning [VGS Proxy Routes](https://www.verygoodsecurity.com/docs/guides/managing-your-routes).

## Provider Status

 <b style="background-color:#d2382c">Deprecation Notice:</b> We do not currently provide support for this provider. VGS is committed to providing developer tooling and recommends using the [VGS CLI and its related tools and patterns](https://www.verygoodsecurity.com/docs/vgs-cli/getting-started) for managing your vaults. Look forward to future updates at our [blog](https://www.verygoodsecurity.com/blog/). 


## How to Install
Requirements: `terraform` ver **0.12 or later**
## Manual (in-house provider)
1. Navigate to the [latest release](https://github.com/verygoodsecurity/terraform-provider-vgs/releases) of the provider.
2. Download archive for appropriate OS and Architecture. You can run `terraform --version` on your environment to see which variant from the list to use.
3. Unzip the archive and copy the provider's binary into `~/.terraform.d/plugin/...` according to [official documentation](https://www.terraform.io/docs/cloud/run/install-software.html#in-house-providers).  

Example for `terraform` 0.13 and later:
```shell
~ mkdir -p ~/.terraform.d/plugins/local.terraform.com/user/vgs/{ver}/darwin_amd64
~ cp ./bin/terraform-provider-vgs_{ver} ~/.terraform.d/plugins/local.terraform.com/user/vgs/{ver}/darwin_amd64/terraform-provider-vgs_{ver}
```
For `terraform` 0.12:
```shell
~ mkdir -p ~/.terraform.d/plugins
~ cp ./bin/terraform-provider-vgs_{ver} ~/.terraform.d/plugins/terraform-provider-vgs_{ver}
```

# How to Use
1. Create a Vault through VGS dashboard and get your Vault ID.
2. Prepare terraform configuration for `vgs` provider in separate folder (e.g. `/vgs`). See [/examples](/examples/README.md) for more information on how to write the configuration.
3. Install and use [vgs-cli](https://github.com/verygoodsecurity/vgs-cli) to create a [ServiceAccount](https://www.verygoodsecurity.com/docs/vgs-cli/service-account#create).
4. Set the `VGS_CLIENT_ID` and `VGS_CLIENT_SECRET` environment variables from ServiceAccount and run.
```shell
~ cd /vgs
~ terraform init
~ VGS_CLIENT_ID=xxx VGS_CLIENT_SECRET=yyy terraform apply
```

# How to build from source
Requirements: Go

## For your system
To compile binaries:
```shell
~ make build
~ ls ./bin
terraform-provider-vgs_v<ver>
```

## Develop
Useful overrides for development:
```shell
~ VGS_VAULT_MANAGEMENT_API_BASE_URL=https://api.verygoodsecurity.io \
VGS_ACCOUNT_MANAGEMENT_API_BASE_URL=https://accounts.verygoodsecurity.io \
VGS_KEYCLOAK_URL=https://auth.verygoodsecurity.io \
VGS_CLIENT_ID=XXX \
VGS_CLIENT_SECRET=YYY \
terraform apply
```

API client located under https://github.com/verygoodsecurity/vgs-api-client-go

## Test
To run tests locally:
```shell
~ TF_ACC=true VGS_CLIENT_ID=xxx VGS_CLIENT_SECRET=yyy go test ./...
?   	github.com/verygoodsecurity/terraform-provider-vgs	[no test files]
ok  	github.com/verygoodsecurity/terraform-provider-vgs/provider	66.337s
```
