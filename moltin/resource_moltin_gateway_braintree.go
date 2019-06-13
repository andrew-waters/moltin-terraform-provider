package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceMoltinBraintreeGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinBraintreeGatewayUpdate,
		Read:   resourceMoltinBraintreeGatewayRead,
		Update: resourceMoltinBraintreeGatewayUpdate,
		Delete: resourceMoltinBraintreeGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinBraintreeGatewayImport,
		},
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"production",
					"sandbox",
				}, false),
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key": {
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

func resourceMoltinBraintreeGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g, err := createBraintreeGatewayFromResourceData(d)
	if err != nil {
		return err
	}
	_, err = client.Put("gateways/braintree", &g)
	if err != nil {
		return fmt.Errorf("Error updating braintree gateway: %s", err)
	}

	d.SetId("braintree_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("environment", g.Environment)
	d.Set("public_key", g.PublicKey)
	d.Set("private_key", g.PrivateKey)
	d.Set("merchant_id", g.MerchantID)

	return resourceMoltinBraintreeGatewayRead(d, meta)
}

func resourceMoltinBraintreeGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	// gateways can't be deleted
	return nil
}

func resourceMoltinBraintreeGatewayRead(d *schema.ResourceData, meta interface{}) error {

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

func resourceMoltinBraintreeGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	err := resourceMoltinBraintreeGatewayRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func createBraintreeGatewayFromResourceData(d *schema.ResourceData) (*entities.BraintreeGateway, error) {
	g := entities.BraintreeGateway{}

	if v, ok := d.GetOk("enabled"); ok {
		g.Enabled = v.(bool)
	}
	if v, ok := d.GetOk("environment"); ok {
		g.Environment = v.(string)
	}
	if v, ok := d.GetOk("public_key"); ok {
		g.PublicKey = v.(string)
	}
	if v, ok := d.GetOk("private_key"); ok {
		g.PrivateKey = v.(string)
	}
	if v, ok := d.GetOk("merchant_id"); ok {
		g.MerchantID = v.(string)
	}

	return &g, nil
}
