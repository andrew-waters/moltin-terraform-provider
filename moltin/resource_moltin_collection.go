package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceMoltinCollection() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinCollectionCreate,
		Read:   resourceMoltinCollectionRead,
		Update: resourceMoltinCollectionUpdate,
		Delete: resourceMoltinCollectionDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinCollectionImport,
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

func resourceMoltinCollectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	collection, err := createCollectionFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("collections/%s", d.Id()), &collection)
	if err != nil {
		return fmt.Errorf("Error updating collection: %s", err)
	}

	return resourceMoltinCollectionRead(d, meta)
}

func resourceMoltinCollectionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)
	_, err := client.Delete(fmt.Sprintf("collections/%s", d.Id()))

	return err
}

func resourceMoltinCollectionRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	collection := entities.Collection{}
	_, err := client.Get(fmt.Sprintf("collections/%s", d.Id()), &collection)
	if err != nil {
		return err
	}

	d.SetId(collection.ID)
	d.Set("id", collection.ID)
	d.Set("name", collection.Name)
	d.Set("slug", collection.Slug)
	d.Set("description", collection.Description)
	d.Set("status", collection.Status)

	return nil
}

func resourceMoltinCollectionCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	collection, err := createCollectionFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Post("collections", &collection)
	if err != nil {
		return fmt.Errorf("Failed to create collection: %s", err)
	}

	d.SetId(collection.ID)

	return resourceMoltinCollectionRead(d, meta)
}

func resourceMoltinCollectionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinCollectionRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find collection: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createCollectionFromResourceData(d *schema.ResourceData) (entities.Collection, error) {

	collection := entities.Collection{}
	collection.ID = d.Id()
	if v, ok := d.GetOk("name"); ok {
		collection.Name = v.(string)
	}
	if v, ok := d.GetOk("slug"); ok {
		collection.Slug = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		collection.Description = v.(string)
	}
	if v, ok := d.GetOk("status"); ok {
		collection.Status = v.(string)
	}

	return collection, nil
}
