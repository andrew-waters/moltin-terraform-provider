package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinFlowCreate,
		Read:   resourceMoltinFlowRead,
		Update: resourceMoltinFlowUpdate,
		Delete: resourceMoltinFlowDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinFlowImport,
		},
		Schema: map[string]*schema.Schema{
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

func resourceMoltinFlowUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	flow, err := createFlowFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("flows/%s", d.Id()), flow)
	if err != nil {
		return fmt.Errorf("Error updating flow: %s", err)
	}

	return resourceMoltinFlowRead(d, meta)
}

func resourceMoltinFlowDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)
	_, err := client.Delete(fmt.Sprintf("flows/%s", d.Id()))

	return err
}

func resourceMoltinFlowRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	flow := entities.Flow{}
	_, err := client.Get(fmt.Sprintf("flows/%s", d.Id()), &flow)
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

func resourceMoltinFlowCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	flow, err := createFlowFromResourceData(d)
	if err != nil {
		return err
	}

	wrapper, err := client.Post("flows", flow)
	if err != nil {

		// if the flow already exists (it's impossible to delete your default one
		// after it has been created), get the ID instead
		if len(wrapper.Response.Errors) > 0 {
			if wrapper.Response.Errors[0].Title == "Flow already exists" {
				flows := []entities.Flow{}
				client.Get("flows", &flows)
				for _, f := range flows {
					if f.Slug == flow.Slug {
						flow.ID = f.ID
					}
				}
			}
		}

		if flow.ID == "" {
			return fmt.Errorf("Failed to create flow: %s", wrapper.Response.Errors[0])
		}
	}

	d.SetId(flow.ID)

	return resourceMoltinFlowRead(d, meta)
}

func resourceMoltinFlowImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinFlowRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find flow: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createFlowFromResourceData(d *schema.ResourceData) (*entities.Flow, error) {

	flow := entities.Flow{}
	if v, ok := d.GetOk("id"); ok {
		flow.ID = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		flow.Name = v.(string)
	}
	if v, ok := d.GetOk("slug"); ok {
		flow.Slug = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		flow.Description = v.(string)
	}
	if v, ok := d.GetOk("enabled"); ok {
		flow.Enabled = v.(bool)
	}

	return &flow, nil
}
