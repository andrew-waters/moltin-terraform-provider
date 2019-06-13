package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinSettingsUpdate,
		Read:   resourceMoltinSettingsRead,
		Update: resourceMoltinSettingsUpdate,
		Delete: resourceMoltinSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinSettingsImport,
		},
		Schema: map[string]*schema.Schema{
			"page_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"list_child_products": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"additional_languages": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				// This is commented out because TF doesn't support validation on lists or maps yet
				// ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
				// 	v := val.(string)
				// 	s := entities.Settings{}
				// 	for _, b := range s.SupportedLanguages() {
				// 		if b == v {
				// 			return
				// 		}
				// 	}
				// 	errs = append(errs, fmt.Errorf("%q is not a valid language", key, v))
				// 	return
				// },
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
	_, err = client.Put("settings", settings)
	if err != nil {
		return fmt.Errorf("Error updating Settings: %s", err)
	}

	d.SetId("settings")
	d.Set("page_length", settings.PageLength)
	d.Set("list_child_products", settings.ListChildProducts)
	d.Set("additional_languages", settings.AdditionalLanguages)

	return resourceMoltinSettingsRead(d, meta)
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

func resourceMoltinSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	err := resourceMoltinSettingsRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func createSettingsFromResourceData(d *schema.ResourceData) (*entities.Settings, error) {
	s := entities.Settings{}

	if v, ok := d.GetOk("page_length"); ok {
		s.PageLength = v.(int)
	}
	if v, ok := d.GetOk("list_child_products"); ok {
		s.ListChildProducts = v.(bool)
	}
	if v, ok := d.GetOk("additional_languages"); ok {
		languages := make([]string, 0)
		switch v := v.(type) {
		case []interface{}:
			for _, lang := range v {
				languages = append(languages, lang.(string))
			}
		}
		s.AdditionalLanguages = languages
	}

	return &s, nil
}
