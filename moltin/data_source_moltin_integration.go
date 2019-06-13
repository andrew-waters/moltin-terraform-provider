package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinIntegration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinIntegrationRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"observes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"configuration": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func dataSourceMoltinIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	integration := entities.Integration{}
	_, err := client.Get(fmt.Sprintf("integrations/%s", d.Get("id").(string)), &integration)
	if err != nil {
		return err
	}

	d.SetId(integration.ID)
	d.Set("id", integration.ID)
	d.Set("name", integration.Name)
	d.Set("description", integration.Description)
	d.Set("enabled", integration.Enabled)
	d.Set("type", integration.IntegrationType)
	d.Set("observes", integration.Observes)
	d.Set("configuration", integration.Configuration)

	return nil
}
