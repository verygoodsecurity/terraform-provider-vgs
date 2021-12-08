# VGS Terraform Provider Examples
Create `main.tf` file and declare use of `vgs` provider and its version in this file.
```terraform
provider "vgs" {

}
```

Then declare `vgs_route` resources. There are separate resources for inbound and outbound routes.

**environment** (String) VGS Environment. One of [sandbox, live, live-eu-1]  
**vault** (String) Vault Identifier. Must match ^tnt[a-z0-9]{8}$  
**inline_config** (String) YAML route configuration https://www.verygoodsecurity.com/docs/features/yaml  

```terraform
resource "vgs_route" "inbound_route" {
  environment   = "sandbox"
  vault         = "tntbcduzut5"
  inline_config = "TBD"
}
```

Finally, specify route ``inline_config``. NOTE: YAML Configuration of the route can be [exported](https://www.verygoodsecurity.com/docs/features/yaml#export-a-single-route) from VGS Dashboard.

See the whole declaration:
```terraform
terraform {
  required_providers {
    vgs = {
      source  = "local.terraform.com/user/vgs"
      version = "0.1.2"
    }
  }
}

provider "vgs" {

}

resource "vgs_route" "inbound_route" {
  environment   = "sandbox"
  vault         = "tntbcduzut5"
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

See more examples:
* [main.tf](/examples/main.tf) for terraform ver 0.12.
* [main_0.13.tf](/examples/main_0.13.tf) for terraform ver 0.13 and later.