package moltin

import (
	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinSettingsRead,
		Schema: map[string]*schema.Schema{
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"exchange_rate": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"format": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceMoltinSettingsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	settings := entities.Settings{}
	_, err := client.Get("settings", &settings)
	if err != nil {
		return err
	}

	d.Set("page_length", settings.PageLength)
	d.Set("list_child_products", settings.ListChildProducts)
	d.Set("additional_languages", settings.AdditionalLanguages)

	return nil
}
