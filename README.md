# VGS Terraform Provider

[![CircleCI](https://circleci.com/gh/verygoodsecurity/terraform-provider-vgs.svg?style=svg&circle-token=8ae379e820e61ec6f8e8451ebaf5ed6958fa7c13)](https://circleci.com/gh/verygoodsecurity/terraform-provider-vgs)

Custom Terraform provider that allows provisioning [VGS Proxy Routes](https://www.verygoodsecurity.com/docs/guides/managing-your-routes).

# THIS LIBRARY IS PRE-RELEASE AND NOT READY FOR CUSTOMER CONSUMPTION

It will be ready when we provide

- [x] Tests
- [x] CICD
- [x] Documentation

# How to Install

## Manual (in-house provider)
1. Navigate to the latest release of the provider.
2. Download archive for appropriate OS and Architecture. You can run `terraform --version` on your environment to see which variant from the list to use.
3. Unzip the archive and copy the provider's binary into `~/terraform.d/plugin/...` according to [official documentation](https://www.terraform.io/docs/cloud/run/install-software.html#in-house-providers).

## Terraform Registry
NOTE: The current version of the provider has not been published to [Terraform Registry](https://registry.terraform.io/) yet.

# How to Use
1. Create a Vault through VGS dashboard and get your Vault ID.
2. Prepare terraform configuration for `vgs` provider. See [/examples](/examples/README.md) for more information on how to write the configuration.
3. Install and use [vgs-cli](https://github.com/verygoodsecurity/vgs-cli) to create a [ServiceAccount](https://www.verygoodsecurity.com/docs/vgs-cli/service-account#create).
4. Set the `VGS_CLIENT_ID` and `VGS_CLIENT_SECRET` environment variables from ServiceAccount and run.
```shell
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