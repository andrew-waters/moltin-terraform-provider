package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinCollection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinCollectionRead,
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

func dataSourceMoltinCollectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	collection := entities.Collection{}
	_, err := client.Get(fmt.Sprintf("collections/%s", d.Get("id").(string)), &collection)
	if err != nil {
		return err
	}

	d.SetId(collection.ID)
	d.Set("id", collection.ID)
	d.Set("name", collection.Name)
	d.Set("slug", collection.Name)
	d.Set("description", collection.Description)
	d.Set("status", collection.Status)

	return nil
}
