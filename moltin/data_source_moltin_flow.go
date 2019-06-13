package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinFlow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinFlowRead,
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
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func dataSourceMoltinFlowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	flow := entities.Flow{}
	_, err := client.Get(fmt.Sprintf("flows/%s", d.Get("id").(string)), &flow)
	if err != nil {
		return err
	}

	d.SetId(flow.ID)
	d.Set("id", flow.ID)
	d.Set("name", flow.Name)
	d.Set("slug", flow.Slug)
	d.Set("description", flow.Description)
	d.Set("enabled", flow.Enabled)

	return nil
}
