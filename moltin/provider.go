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
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MOLTIN_CLIENT_ID", nil),
				Description: "Your Moltin Client ID.",
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MOLTIN_CLIENT_SECRET", nil),
				Description: "Your Moltin Client Secret.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"moltin_brand":      dataSourceMoltinBrand(),
			"moltin_category":   dataSourceMoltinCategory(),
			"moltin_collection": dataSourceMoltinCollection(),
			"moltin_currency":   dataSourceMoltinCurrency(),
			"moltin_product":    dataSourceMoltinProduct(),
			"moltin_settings":   dataSourceMoltinSettings(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"moltin_brand":      resourceMoltinBrand(),
			"moltin_category":   resourceMoltinCategory(),
			"moltin_collection": resourceMoltinCollection(),
			"moltin_currency":   resourceMoltinCurrency(),
			"moltin_product":    resourceMoltinProduct(),
			"moltin_settings":   resourceMoltinSettings(),
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
