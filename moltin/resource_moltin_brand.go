package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceMoltinBrand() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinBrandCreate,
		Read:   resourceMoltinBrandRead,
		Update: resourceMoltinBrandUpdate,
		Delete: resourceMoltinBrandDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinBrandImport,
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
			"status": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"live",
					"draft",
				}, false),
			},
		},
	}
}

func resourceMoltinBrandUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	brand, err := createBrandFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("brands/%s", d.Id()), &brand)
	if err != nil {
		return fmt.Errorf("Error updating brand: %s", err)
	}

	return resourceMoltinBrandRead(d, meta)
}

func resourceMoltinBrandDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)
	_, err := client.Delete(fmt.Sprintf("brands/%s", d.Id()))

	return err
}

func resourceMoltinBrandRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	brand := entities.Brand{}
	_, err := client.Get(fmt.Sprintf("brands/%s", d.Id()), &brand)
	if err != nil {
		return err
	}

	d.SetId(brand.ID)
	d.Set("id", brand.ID)
	d.Set("name", brand.Name)
	d.Set("slug", brand.Slug)
	d.Set("description", brand.Description)
	d.Set("status", brand.Status)

	return nil
}

func resourceMoltinBrandCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	brand, err := createBrandFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Post("brands", &brand)
	if err != nil {
		return fmt.Errorf("Failed to create brand: %s", err)
	}

	d.SetId(brand.ID)

	return resourceMoltinBrandRead(d, meta)
}

func resourceMoltinBrandImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinBrandRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find brand: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createBrandFromResourceData(d *schema.ResourceData) (entities.Brand, error) {

	brand := entities.Brand{}
	brand.ID = d.Id()
	if v, ok := d.GetOk("name"); ok {
		brand.Name = v.(string)
	}
	if v, ok := d.GetOk("slug"); ok {
		brand.Slug = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		brand.Description = v.(string)
	}
	if v, ok := d.GetOk("status"); ok {
		brand.Status = v.(string)
	}

	return brand, nil
}
