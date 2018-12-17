package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinCategory() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinCategoryRead,
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

func dataSourceMoltinCategoryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	category := entities.Category{}
	_, err := client.Get(fmt.Sprintf("categories/%s", d.Get("id").(string)), &category)
	if err != nil {
		return err
	}

	d.SetId(category.ID)
	d.Set("id", category.ID)
	d.Set("name", category.Name)
	d.Set("slug", category.Name)
	d.Set("description", category.Description)
	d.Set("status", category.Status)

	return nil
}
