provider "moltin" {
  client_id     = "${var.moltin_client_id}"
  client_secret = "${var.moltin_client_secret}"
}

// resource "moltin_settings" "settings" {
//   page_length          = "GBP"
//   list_child_products  = 1
//   additional_languages = []
// }

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

resource "moltin_currency" "USD" {
  code                = "USD"
  exchange_rate       = 1
  format              = "$${price}" // $$ escapes the required $ sign for this currency
  decimal_point       = "."
  decimal_places      = 2
  thousand_separator  = ","
  default             = false
  enabled             = true
}

resource "moltin_product" "DEWALT_DCD776S2T" {
  name            = "DeWalt DCD776S2T-GB 18V 1.5Ah Li-Ion XR Cordless Combi Drill"
  slug            = "dcd776S2t"
  sku             = "DCD776S2T"
  description     = "Powerful and versatile cordless combi drill with 15 torque settings as well as drill and hammer settings. Thermal overload protection increases durability and a fan-cooled, 2-speed variable motor provides long lasting performance. An integrated LED light improves visibility in dimly lit areas. A 13mm keyless chuck allows for secure, one-handed tightening. Lightweight compact design make it easy to handle and suitable for drilling in tight spaces. Supplied with 2 x batteries, charger and TSTAK kit box."
  status          = "live"
  commodity_type  = "physical"

  price {
    amount        = 9999
    currency      = "GBP"
    includes_tax  = true
  }

  price {
    amount        = 11999
    currency      = "USD"
    includes_tax  = true
  }
}

resource "moltin_product" "DEWALT_DCD776D2T" {
  name            = "DeWalt DCD776D2T-GB 18V 2.0Ah Li-Ion XR Cordless Combi Drill"
  slug            = "dcd776d2t"
  sku             = "DCD776D2T"
  description     = "Compact drill driver with XR Li-ion battery technology, designed for efficiency and making applications faster. Improved ergonomic design and rubber grip increases user comfort and provides maximum control. Supplied with multi-voltage charger for all XR Li-ion slide pack batteries. Supplied in a TSTAK kit box. Manufactured in the UK."
  status          = "live"
  commodity_type  = "physical"

  price {
    amount        = 11999
    currency      = "GBP"
    includes_tax  = true
  }

  price {
    amount        = 13999
    currency      = "USD"
    includes_tax  = true
  }
}
