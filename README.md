# VGS Terraform Provider [WIP]

## Build
## For your system
```shell
~ make default
~ ls ./bin
terraform-provider-vgs_v0.1.0
```

## All targets
```shell
~ make all
```

## Use
1. Put the binary under Terraform plugin directory, usually like
```shell
~ mv ./bin/terraform-provider-vgs_v0.1.0 ~/.terraform.d/plugins/terraform-provider-vgs_v0.1.0
```

2. Create a Vault through VGS dashboard and get your Vault ID.

3. Create a directory for your TF files (or use `/examples`) and create `main.tf`
```terraform
provider "vgs" {
  version = "~> 0.1"
}
resource "vgs_route" "inbound_route" {
  environment = "sandbox"
  vault = "tntbcduzut5"
  inline_config = <<EOF
id: 04b2e1b7-fb60-472f-a79f-af7e2353f122
type: rule_chain
attributes:
  tags:
    name: my-awesome-inbound-route
    source: RouteContainer
  destination_override_endpoint: 'https://echo.apps.verygood.systems'
  host_endpoint: (.*)\.verygoodproxy\.io
  id: 04b2e1b7-fb60-472f-a79f-af7e2353f122
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
```

4. Install and use [vgs-cli](https://github.com/verygoodsecurity/vgs-cli) to create a [ServiceAccount](https://www.verygoodsecurity.com/docs/vgs-cli/service-account#create).
5. Set the `VGS_CLIENT_ID` and `VGS_CLIENT_SECRET` environment variables from ServiceAccount and run
```shell
~ terraform init
~ VGS_CLIENT_ID=xxx VGS_CLIENT_SECRET=yyy terraform apply
```

## Develop

Useful overrides for development
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

```shell
~ TF_ACC=true VGS_CLIENT_ID=xxx VGS_CLIENT_SECRET=yyy go test ./...
?   	github.com/verygoodsecurity/terraform-provider-vgs	[no test files]
ok  	github.com/verygoodsecurity/terraform-provider-vgs/provider	66.337s
```
