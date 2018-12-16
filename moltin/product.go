package moltin

import (
	"github.com/andrew-waters/gomo/entities"
)

func expandPrice(config []interface{}) []entities.ProductPrice {

	prices := make([]entities.ProductPrice, 0, len(config))

	for _, rawPrice := range config {
		price := rawPrice.(map[string]interface{})
		p := entities.ProductPrice{
			Amount:      price["amount"].(int),
			Currency:    price["currency"].(string),
			IncludesTax: price["includes_tax"].(bool),
		}

		prices = append(prices, p)
	}

	return prices
}
