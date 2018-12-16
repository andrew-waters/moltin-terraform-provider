package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMoltinCurrency() *schema.Resource {
	return &schema.Resource{
		Create: resourceMoltinCurrencyCreate,
		Read:   resourceMoltinCurrencyRead,
		Update: resourceMoltinCurrencyUpdate,
		Delete: resourceMoltinCurrencyDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMoltinCurrencyImport,
		},
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
			"decimal_point": {
				Type:     schema.TypeString,
				Required: true,
			},
			"decimal_places": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"thousand_separator": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceMoltinCurrencyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	currency, err := createCurrencyFromResourceData(d)
	if err != nil {
		return err
	}

	_, err = client.Put(fmt.Sprintf("currencies/%s", d.Id()), currency)
	if err != nil {
		return fmt.Errorf("Error updating currency: %s", err)
	}

	return resourceMoltinCurrencyRead(d, meta)
}

func resourceMoltinCurrencyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)
	_, err := client.Delete(fmt.Sprintf("currencies/%s", d.Id()))

	return err
}

func resourceMoltinCurrencyRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	currency := entities.Currency{}
	_, err := client.Get(fmt.Sprintf("currencies/%s", d.Id()), &currency)
	if err != nil {
		return err
	}

	d.SetId(currency.ID)
	d.Set("id", currency.ID)
	d.Set("code", currency.Code)
	d.Set("exchange_rate", currency.ExchangeRate)
	d.Set("format", currency.Format)
	d.Set("decimal_point", currency.DecimalPoint)
	d.Set("decimal_places", currency.DecimalPlaces)
	d.Set("thousand_separator", currency.ThousandSeparator)
	d.Set("default", currency.Default)
	d.Set("enabled", currency.Enabled)

	return nil
}

func resourceMoltinCurrencyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*gomo.Client)

	currency, err := createCurrencyFromResourceData(d)
	if err != nil {
		return err
	}

	wrapper, err := client.Post("currencies", currency)
	if err != nil {

		// if the currency already exists (it's impossible to delete your default one
		// after it has been created), get the ID instead
		if len(wrapper.Response.Errors) > 0 {
			if wrapper.Response.Errors[0].Title == "Currency already exists" {
				currencies := []entities.Currency{}
				client.Get("currencies", &currencies)
				currency.ID = currencies[0].ID
			}
		}

		if currency.ID == "" {
			return fmt.Errorf("Failed to create currency: %s", err)
		}
	}

	d.SetId(currency.ID)

	return resourceMoltinCurrencyRead(d, meta)
}

func resourceMoltinCurrencyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	err := resourceMoltinCurrencyRead(d, meta)
	if err != nil {
		return nil, err
	}

	if d.Id() == "" {
		return nil, fmt.Errorf("Couldn't find currency: %s", id)
	}

	results := []*schema.ResourceData{d}

	return results, nil
}

func createCurrencyFromResourceData(d *schema.ResourceData) (*entities.Currency, error) {

	currency := entities.Currency{}
	if v, ok := d.GetOk("id"); ok {
		currency.ID = v.(string)
	}
	if v, ok := d.GetOk("code"); ok {
		currency.Code = v.(string)
	}
	if v, ok := d.GetOk("exchange_rate"); ok {
		currency.ExchangeRate = v.(float64)
	}
	if v, ok := d.GetOk("format"); ok {
		currency.Format = v.(string)
	}
	if v, ok := d.GetOk("decimal_point"); ok {
		currency.DecimalPoint = v.(string)
	}
	if v, ok := d.GetOk("decimal_places"); ok {
		currency.DecimalPlaces = int64(v.(int))
	}
	if v, ok := d.GetOk("thousand_separator"); ok {
		currency.ThousandSeparator = v.(string)
	}
	if v, ok := d.GetOk("default"); ok {
		currency.Default = v.(bool)
	}
	if v, ok := d.GetOk("enabled"); ok {
		currency.Enabled = v.(bool)
	}

	return &currency, nil
}
