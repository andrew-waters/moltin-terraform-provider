package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinAdyenGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinAdyenGatewayUpdate,
		Read:   resourceMoltinAdyenGatewayRead,
		Update: resourceMoltinAdyenGatewayUpdate,
		Delete: resourceMoltinAdyenGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinAdyenGatewayImport,
		},
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
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

func resourceMoltinAdyenGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g, err := createAdyenGatewayFromResourceData(d)
	if err != nil {
		return err
	}
	_, err = client.Put("gateways/adyen", &g)
	if err != nil {
		return fmt.Errorf("Error updating adyen gateway: %s", err)
	}

	d.SetId("adyen_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("test", g.Test)
	d.Set("username", g.Username)
	d.Set("password", g.Password)
	d.Set("merchant_account", g.MerchantAccount)

	return resourceMoltinAdyenGatewayRead(d, meta)
}

func resourceMoltinAdyenGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	// gateways can't be deleted
	return nil
}

func resourceMoltinAdyenGatewayRead(d *schema.ResourceData, meta interface{}) error {

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

func resourceMoltinAdyenGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	err := resourceMoltinAdyenGatewayRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func createAdyenGatewayFromResourceData(d *schema.ResourceData) (*entities.AdyenGateway, error) {
	g := entities.AdyenGateway{}

	if v, ok := d.GetOk("enabled"); ok {
		g.Enabled = v.(bool)
	}
	if v, ok := d.GetOk("test"); ok {
		g.Test = v.(bool)
	}
	if v, ok := d.GetOk("username"); ok {
		g.Username = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		g.Password = v.(string)
	}
	if v, ok := d.GetOk("merchant_account"); ok {
		g.MerchantAccount = v.(string)
	}

	return &g, nil
}
