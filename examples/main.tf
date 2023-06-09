terraform {
  required_providers {
    vgs = {
      source  = "local.terraform.com/user/vgs"
      version = "0.2.3"
    }
  }
}

provider "vgs" {
  environment  = "sandbox"
  organization = "ACrF5NY4vxzj8BvhU3q8btsN"
}

resource "vgs_vault" "my_sandbox_vault" {
  name = "My Testing Vault"
  preferences = {
    "interceptCORS" : "true"
  }
}

resource "vgs_route" "test_route" {
  vault                         = "tntbcduzut5"
  protocol                      = "HTTP"
  source_endpoint               = "blab"
  destination_override_endpoint = "vvvvv"
  host_endpoint                 = "xxxx"
  port                          = 80
  filters = [
    {
      phase = "request"
      alias_format = "UUID"
      transformer = {
        content_type = "csv"
        config = {
          column_indices : "[1, 2]"
        }
      }

      targets  = ["body"]
      function = <<EOL
      load('@stdlib/json', 'json')
      def process(input, ctx):
              body = json.decode(input.body())
              as_str = json.encode(body['fields'])
              token = vault.put(as_str)
              body['fields'] = token
              input.set_body(json.encode(body))
              return input
      EOL
      conditions_inline = jsonencode({
        condition = "AND"
        rules = [
          {
            expression = {
              field    = "PathInfo"
              operator = "matches"
              type     = "string"
              values = [
                "/post"
              ]
            }
          },
          {
            condition = "OR"
            rules = [
              {
                expression = {
                  field    = "ContentType"
                  operator = "equals"
                  type     = "string"
                  values = [
                    "application/json"
                  ]
                }
              },
              {
                expression = {
                  field    = "ContentType"
                  operator = "equals"
                  type     = "string"
                  values = [
                    "application/xml"
                  ]
                }
              }
            ]
          }
        ]
      })
      classifiers = {
        "include" : ["pci-data"]
        "tags" : ["my-filter"]
      }
    }
  ]
}
