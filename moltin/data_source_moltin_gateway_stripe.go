package moltin

import (
	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinStripeGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinStripeGatewayRead,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"login": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceMoltinStripeGatewayRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g := entities.StripeGateway{}
	_, err := client.Get("gateways/stripe", &g)
	if err != nil {
		return err
	}

	d.SetId("stripe_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("login", g.Login)

	return nil
}
