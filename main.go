package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/verygoodsecurity/terraform-provider-vgs/provider"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
//go:generate terraform fmt -recursive ./examples/
func main() {

	if err := providerserver.Serve(
		context.Background(),
		provider.Provider,
		providerserver.ServeOpts{
			Address: "registry.terraform.io/<namespace>/<provider_name>",
		},
	); err != nil {
		log.Fatal(err)
	}

}
