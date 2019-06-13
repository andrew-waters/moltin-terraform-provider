package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceMoltinProduct() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinProductCreate,
		Read:   resourceMoltinProductRead,
		Update: resourceMoltinProductUpdate,
		Delete: resourceMoltinProductDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinProductImport,
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
			"sku": {
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
			"commodity_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"physical",
					"digital",
				}, false),
			},
			"manage_stock": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"price": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"amount": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"currency": {
							Type:     schema.TypeString,
							Required: true,
						},
						"includes_tax": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceMoltinProductUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	product, err := createProductFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("products/%s", product.ID), &product)
	if err != nil {
		return fmt.Errorf("Error updating product (%s): %s", product.ID, err)
	}

	return resourceMoltinProductRead(d, meta)
}

func resourceMoltinProductDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)
	client.Delete(fmt.Sprintf("products/%s", d.Id()))

	return nil
}

func resourceMoltinProductRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)
	product := entities.Product{}
	_, err := client.Get(fmt.Sprintf("products/%s", d.Id()), &product)
	if err != nil {
		return err
	}

	d.SetId(product.ID)
	d.Set("name", product.Name)
	d.Set("slug", product.Slug)
	d.Set("sku", product.SKU)
	d.Set("description", product.Description)
	d.Set("status", product.Status)
	d.Set("commodity_type", product.CommodityType)
	d.Set("manage_stock", product.ManageStock)
	d.Set("price", product.Price)

	return nil
}

func resourceMoltinProductCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	product, err := createProductFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Post("products", &product)
	if err != nil {
		return fmt.Errorf("Failed to create product: %s", err)
	}

	d.SetId(product.ID)

	return resourceMoltinProductRead(d, meta)
}

func resourceMoltinProductImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinProductRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find product: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createProductFromResourceData(d *schema.ResourceData) (entities.Product, error) {

	product := entities.Product{}

	product.ID = d.Id()
	if v, ok := d.GetOk("name"); ok {
		product.Name = v.(string)
	}
	if v, ok := d.GetOk("slug"); ok {
		product.Slug = v.(string)
	}
	if v, ok := d.GetOk("sku"); ok {
		product.SKU = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		product.Description = v.(string)
	}
	if v, ok := d.GetOk("status"); ok {
		product.Status = v.(string)
	}
	if v, ok := d.GetOk("commodity_type"); ok {
		product.CommodityType = v.(string)
	}
	if v, ok := d.GetOk("manage_stock"); ok {
		product.ManageStock = v.(bool)
	}

	product.Price = expandPrice(d.Get("price").([]interface{}))

	return product, nil
}
