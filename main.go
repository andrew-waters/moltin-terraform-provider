package main

import (
	"github.com/andrew-waters/moltin-terraform-provider/moltin"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: moltin.Provider,
	})
}
