package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinBrand() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinBrandRead,
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
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceMoltinBrandRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	brand := entities.Brand{}
	_, err := client.Get(fmt.Sprintf("brands/%s", d.Get("id").(string)), &brand)
	if err != nil {
		return err
	}

	d.SetId(brand.ID)
	d.Set("id", brand.ID)
	d.Set("name", brand.Name)
	d.Set("slug", brand.Name)
	d.Set("description", brand.Description)
	d.Set("status", brand.Status)

	return nil
}
