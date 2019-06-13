package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinStripeGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinStripeGatewayUpdate,
		Read:   resourceMoltinStripeGatewayRead,
		Update: resourceMoltinStripeGatewayUpdate,
		Delete: resourceMoltinStripeGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinStripeGatewayImport,
		},
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"login": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceMoltinStripeGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g, err := createStripeGatewayFromResourceData(d)
	if err != nil {
		return err
	}
	_, err = client.Put("gateways/stripe", &g)
	if err != nil {
		return fmt.Errorf("Error updating stripe gateway: %s", err)
	}

	d.SetId("stripe_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("login", g.Login)

	return resourceMoltinStripeGatewayRead(d, meta)
}

func resourceMoltinStripeGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	// gateways can't be deleted
	return nil
}

func resourceMoltinStripeGatewayRead(d *schema.ResourceData, meta interface{}) error {

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

func resourceMoltinStripeGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	err := resourceMoltinStripeGatewayRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func createStripeGatewayFromResourceData(d *schema.ResourceData) (*entities.StripeGateway, error) {
	g := entities.StripeGateway{}

	if v, ok := d.GetOk("enabled"); ok {
		g.Enabled = v.(bool)
	}
	if v, ok := d.GetOk("login"); ok {
		g.Login = v.(string)
	}

	return &g, nil
}
