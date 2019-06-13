# Terraform Moltin Provider

This provider enables Moltin stores to manage the resources on their store - settings, currencies, products, categories, collections, brands, integrations, promotions payments gateways, customers, addresses and files.

## Basic Usage

```hcl
provider "moltin" {
  client_id     = "${var.moltin_client_id}"
  client_secret = "${var.moltin_client_secret}"
}

resource "moltin_settings" "settings" {
  page_length          = 25
  list_child_products  = false
  additional_languages = ["de"]
}

resource "moltin_currency" "GBP" {
  code                = "GBP"
  exchange_rate       = 1
  format              = "£{price}"
  decimal_point       = "."
  decimal_places      = 2
  thousand_separator  = ","
  default             = true
  enabled             = true
}

resource "moltin_product" "your_product" {
  name            = "Your Product"
  slug            = "your-product"
  sku             = "y.p.001"
  description     = "An amazing product"
  status          = "live"
  commodity_type  = "physical"
  manage_stock    = false

  price {
    amount        = 9999
    currency      = "GBP"
    includes_tax  = true
  }
}

```

## Supported Resources

- [ ] [Products](https://docs.moltin.com/catalog/products)
  - [x] name
  - [x] slug
  - [x] sku
  - [x] manage_stock
  - [x] description
  - [x] price
  - [x] status
  - [x] commodity_type
  - [ ] relationships
  - [ ] variations
- [x] [Settings](https://docs.moltin.com/advanced/settings)
  - [x] page_length
  - [x] list_child_products
  - [x] additional_languages
- [ ] [Files](https://docs.moltin.com/advanced/files)
  - [ ] file_name
  - [ ] mime_type
  - [ ] public
- [ ] [Currencies](https://docs.moltin.com/advanced/currencies)
  - [x] code
  - [x] exchange_rate
  - [x] format
  - [x] decimal_point
  - [x] thousand_separator
  - [x] decimal_places
  - [x] default
  - [x] enabled
- [ ] [Brands](https://docs.moltin.com/catalog/brands)
  - [x] name
  - [x] slug
  - [x] description
  - [x] status
  - [ ] relationships
- [ ] [Categories](https://docs.moltin.com/catalog/categories)
  - [x] name
  - [x] slug
  - [x] description
  - [x] status
  - [ ] relationships
- [ ] [Collections](https://docs.moltin.com/catalog/collections)
  - [x] name
  - [x] slug
  - [x] description
  - [x] status
  - [ ] relationships
- [x] [Payment Gateways](https://docs.moltin.com/payments/gateways)
  - [x] [Adyen](https://docs.moltin.com/payments/gateways/configure-adyen)
    - [x] enabled
    - [x] test
    - [x] username
    - [x] password
    - [x] merchant_account
  - [x] [Braintree](https://docs.moltin.com/payments/gateways/configure-braintree)
    - [x] enabled
    - [x] environment
    - [x] merchant_id
    - [x] private_key
    - [x] public_key
  - [x] [CardConnect](https://docs.moltin.com/payments/gateways/configure-cardconnect)
    - [x] enabled
    - [x] merchant_id
    - [x] username
    - [x] password
  - [x] [Stripe](https://docs.moltin.com/payments/gateways/configure-stripe)
    - [x] enabled
    - [x] login
- [x] [Integrations](https://docs.moltin.com/advanced/events)
  - [x] enabled
  - [x] name
  - [x] description
  - [x] integration_type
  - [x] observes
  - [x] configuration
- [ ] [Flows](https://docs.moltin.com/advanced/custom-data/flows)
  - [ ] @TODO
- [ ] [Fields](https://docs.moltin.com/advanced/custom-data/fields)
  - [ ] @TODO
- [ ] [Entries](https://docs.moltin.com/advanced/custom-data/entries)
  - [ ] @TODO




