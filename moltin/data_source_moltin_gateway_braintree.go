package moltin

import (
	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinBraintreeGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinBraintreeGatewayRead,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key": {
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

func dataSourceMoltinBraintreeGatewayRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g := entities.BraintreeGateway{}
	_, err := client.Get("gateways/braintree", &g)
	if err != nil {
		return err
	}

	d.SetId("braintree_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("environment", g.Environment)
	d.Set("public_key", g.PublicKey)
	d.Set("private_key", g.PrivateKey)
	d.Set("merchant_id", g.MerchantID)

	return nil
}
