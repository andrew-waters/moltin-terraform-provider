package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinIntegration() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinIntegrationCreate,
		Read:   resourceMoltinIntegrationRead,
		Update: resourceMoltinIntegrationUpdate,
		Delete: resourceMoltinIntegrationDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinIntegrationImport,
		},
		Schema: map[string]*schema.Schema{
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

func resourceMoltinIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	integration, err := createIntegrationFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("integrations/%s", d.Id()), integration)
	if err != nil {
		return fmt.Errorf("Error updating integration: %s", err)
	}

	return resourceMoltinIntegrationRead(d, meta)
}

func resourceMoltinIntegrationDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)
	_, err := client.Delete(fmt.Sprintf("integrations/%s", d.Id()))

	return err
}

func resourceMoltinIntegrationRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	integration := entities.Integration{}
	_, err := client.Get(fmt.Sprintf("integrations/%s", d.Id()), &integration)
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

func resourceMoltinIntegrationCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	integration, err := createIntegrationFromResourceData(d)
	if err != nil {
		return err
	}

	wrapper, err := client.Post("integrations", integration)
	if err != nil {

		// if the integration already exists (it's impossible to delete your default one
		// after it has been created), get the ID instead
		if len(wrapper.Response.Errors) > 0 {
			if wrapper.Response.Errors[0].Title == "Integration already exists" {
				integrations := []entities.Integration{}
				client.Get("integrations", &integrations)
				integration.ID = integrations[0].ID
			}
		}

		if integration.ID == "" {
			return fmt.Errorf("Failed to create integration: %s", err)
		}
	}

	d.SetId(integration.ID)

	return resourceMoltinIntegrationRead(d, meta)
}

func resourceMoltinIntegrationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinIntegrationRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find integration: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createIntegrationFromResourceData(d *schema.ResourceData) (*entities.Integration, error) {

	integration := entities.Integration{}
	if v, ok := d.GetOk("id"); ok {
		integration.ID = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		integration.Name = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		integration.Description = v.(string)
	}
	if v, ok := d.GetOk("enabled"); ok {
		integration.Enabled = v.(bool)
	}
	if v, ok := d.GetOk("type"); ok {
		integration.IntegrationType = v.(string)
	}
	if v, ok := d.GetOk("observes"); ok {
		watch := make([]string, 0)
		switch v := v.(type) {
		case []interface{}:
			for _, event := range v {
				watch = append(watch, event.(string))
			}
		}
		integration.Observes = watch
	}
	if v, ok := d.GetOk("configuration"); ok {
		config := make(map[string]interface{}, 0)
		switch v := v.(type) {
		case map[string]interface{}:
			for k, val := range v {
				config[k] = val
			}
		}
		integration.Configuration = config
	}

	return &integration, nil
}
