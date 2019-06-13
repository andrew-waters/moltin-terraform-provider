package moltin

import (
	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinCardConnectGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinCardConnectGatewayRead,
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
			"merchant_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceMoltinCardConnectGatewayRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g := entities.CardConnectGateway{}
	_, err := client.Get("gateways/card_connect", &g)
	if err != nil {
		return err
	}

	d.SetId("card_connect_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("test", g.Test)
	d.Set("username", g.Username)
	d.Set("password", g.Password)
	d.Set("merchant_id", g.MerchantID)

	return nil
}
