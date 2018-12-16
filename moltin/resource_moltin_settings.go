package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinSettingsCreate,
		Read:   resourceMoltinSettingsRead,
		Update: resourceMoltinSettingsUpdate,
		Delete: resourceMoltinSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinCurrencyImport,
		},
		Schema: map[string]*schema.Schema{
			"page_length": {
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
			},
			"list_child_products": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"additional_languages": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

func resourceMoltinSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	settings, err := createSettingsFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("currencies/%s", d.Id()), settings)
	if err != nil {
		return fmt.Errorf("Error updating Settings: %s", err)
	}

	return resourceMoltinCurrencyRead(d, meta)
}

func resourceMoltinSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	// settings can't be deleted
	return nil
}

func resourceMoltinSettingsRead(d *schema.ResourceData, meta interface{}) error {

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

func resourceMoltinSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	// Settings can't be created, you need to update them
	return nil
}

func resourceMoltinSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinSettingsRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find Settings: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createSettingsFromResourceData(d *schema.ResourceData) (*entities.Settings, error) {

	Settings := entities.Settings{}

	if v, ok := d.GetOk("page_length"); ok {
		Settings.PageLength = v.(int)
	}
	if v, ok := d.GetOk("list_child_products"); ok {
		Settings.ListChildProducts = v.(bool)
	}
	if v, ok := d.GetOk("additional_languages"); ok {
		Settings.AdditionalLanguages = v.([]string)
	}

	return &Settings, nil
}
