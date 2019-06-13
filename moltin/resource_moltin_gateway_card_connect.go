package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinCardConnectGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinCardConnectGatewayUpdate,
		Read:   resourceMoltinCardConnectGatewayRead,
		Update: resourceMoltinCardConnectGatewayUpdate,
		Delete: resourceMoltinCardConnectGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinCardConnectGatewayImport,
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
			"merchant_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceMoltinCardConnectGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	g, err := createCardConnectGatewayFromResourceData(d)
	if err != nil {
		return err
	}
	_, err = client.Put("gateways/card_connect", &g)
	if err != nil {
		return fmt.Errorf("Error updating card_connect gateway: %s", err)
	}

	d.SetId("card_connect_gateway")
	d.Set("enabled", g.Enabled)
	d.Set("test", g.Test)
	d.Set("username", g.Username)
	d.Set("password", g.Password)
	d.Set("merchant_id", g.MerchantID)

	return resourceMoltinCardConnectGatewayRead(d, meta)
}

func resourceMoltinCardConnectGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	// gateways can't be deleted
	return nil
}

func resourceMoltinCardConnectGatewayRead(d *schema.ResourceData, meta interface{}) error {

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

func resourceMoltinCardConnectGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	err := resourceMoltinCardConnectGatewayRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func createCardConnectGatewayFromResourceData(d *schema.ResourceData) (*entities.CardConnectGateway, error) {
	g := entities.CardConnectGateway{}

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
	if v, ok := d.GetOk("merchant_id"); ok {
		g.MerchantID = v.(string)
	}

	return &g, nil
}
