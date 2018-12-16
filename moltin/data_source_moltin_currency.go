package moltin

import (
	"fmt"

	"github.com/andrew-waters/gomo"
	"github.com/andrew-waters/gomo/entities"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceMoltinCurrency() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMoltinCurrencyRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
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

func dataSourceMoltinCurrencyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*gomo.Client)

	currency := entities.Currency{}
	_, err := client.Get(fmt.Sprintf("currencies/%s", d.Get("id").(string)), &currency)
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
