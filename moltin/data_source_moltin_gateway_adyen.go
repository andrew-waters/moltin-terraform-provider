package moltin

import (
	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinAdyenGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinAdyenGatewayRead,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"test": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"merchant_account": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceMoltinAdyenGatewayRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g := entities.AdyenGateway{}
	_, err := client.Get("gateways/adyen", &g)
	if err != nil {
		return err
	}

	d.SetId("adyen_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("test", g.Test)
	d.Set("username", g.Username)
	d.Set("password", g.Password)
	d.Set("merchant_account", g.MerchantAccount)

	return nil
}
