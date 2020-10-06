package main

import (
	"github.com/Shopify/terraform-provider-gsuite/gsuite"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return gsuite.Provider()
		},
	})
}
