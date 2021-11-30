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

resource "vgs_route" "outbound_route" {
  environment = "sandbox"
  vault = "tntbcduzut5"
  inline_config = <<EOF
id: 37df6406-817f-4817-a72a-c675c50fb8ac
type: rule_chain
attributes:
  tags:
    name: my-awesome-outbound-route
    source: RouteContainer
  destination_override_endpoint: '*'
  host_endpoint: echo\.apps\.verygood\.systems
  id: 37df6406-817f-4817-a72a-c675c50fb8ac
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
    id: ea83a6fd-9d04-42b8-8ca0-d51b98380b2e
    id_selector: null
    operation: ENRICH
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
    type: null
EOF
}