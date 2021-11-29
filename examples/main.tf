provider "vgs" {}

resource "vgs_route" "my_route" {
  environment = "sandbox"
  vault = "tntbcduzut5"
  inline_config = <<EOF
id: 04b2e1b7-fb60-472f-a79f-af7e2353f122
type: rule_chain
attributes:
  created_at: '2021-11-26T18:10:08'
  destination_override_endpoint: 'https://echon.apps.verygood.systems'
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
  host_endpoint: (.*)\.verygoodproxy\.io
  id: 04b2e1b7-fb60-472f-a79f-af7e2353f122
  ordinal: null
  port: 80
  protocol: http
  source_endpoint: '*'
  tags:
    name: echo.apps.verygood.systems-beige-crescent
    source: RouteContainer
  updated_at: '2021-11-26T18:10:08'
EOF
}