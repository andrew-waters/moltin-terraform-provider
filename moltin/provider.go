package moltin

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MOLTIN_CLIENT_ID", nil),
				Description: "Your Moltin Client ID.",
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MOLTIN_CLIENT_SECRET", nil),
				Description: "Your Moltin Client Secret.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"moltin_product":  dataSourceMoltinProduct(),
			"moltin_currency": dataSourceMoltinCurrency(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"moltin_product":  resourceMoltinProduct(),
			"moltin_currency": resourceMoltinCurrency(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := config{
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
	}
	return config.client()
}
