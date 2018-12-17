package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceMoltinCategory() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinCategoryCreate,
		Read:   resourceMoltinCategoryRead,
		Update: resourceMoltinCategoryUpdate,
		Delete: resourceMoltinCategoryDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinCategoryImport,
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

func resourceMoltinCategoryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	category, err := createCategoryFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("categories/%s", d.Id()), &category)
	if err != nil {
		return fmt.Errorf("Error updating category: %s", err)
	}

	return resourceMoltinCategoryRead(d, meta)
}

func resourceMoltinCategoryDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)
	_, err := client.Delete(fmt.Sprintf("categories/%s", d.Id()))

	return err
}

func resourceMoltinCategoryRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	category := entities.Category{}
	_, err := client.Get(fmt.Sprintf("categories/%s", d.Id()), &category)
	if err != nil {
		return err
	}

	d.SetId(category.ID)
	d.Set("id", category.ID)
	d.Set("name", category.Name)
	d.Set("slug", category.Slug)
	d.Set("description", category.Description)
	d.Set("status", category.Status)

	return nil
}

func resourceMoltinCategoryCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	category, err := createCategoryFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Post("categories", &category)
	if err != nil {
		return fmt.Errorf("Failed to create category: %s", err)
	}

	d.SetId(category.ID)

	return resourceMoltinCategoryRead(d, meta)
}

func resourceMoltinCategoryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinCategoryRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find category: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createCategoryFromResourceData(d *schema.ResourceData) (entities.Category, error) {

	category := entities.Category{}
	category.ID = d.Id()
	if v, ok := d.GetOk("name"); ok {
		category.Name = v.(string)
	}
	if v, ok := d.GetOk("slug"); ok {
		category.Slug = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		category.Description = v.(string)
	}
	if v, ok := d.GetOk("status"); ok {
		category.Status = v.(string)
	}

	return category, nil
}
