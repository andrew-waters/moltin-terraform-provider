package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinProduct() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinProductRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sku": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"price": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"amount": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"currency": {
							Type:     schema.TypeString,
							Required: true,
						},
						"includes_tax": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceMoltinProductRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	product := entities.Product{}
	_, err := client.Get(fmt.Sprintf("products/%s", d.Get("id").(string)), &product)
	if err != nil {
		return err
	}

	d.SetId(product.ID)
	d.Set("id", product.ID)
	d.Set("name", product.Name)
	d.Set("slug", product.Name)
	d.Set("sku", product.Name)
	d.Set("description", product.Description)
	d.Set("price", product.Price)

	return nil
}
