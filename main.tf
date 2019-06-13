provider "moltin" {
  client_id     = "${var.moltin_client_id}"
  client_secret = "${var.moltin_client_secret}"
}


# Store Settings

resource "moltin_settings" "settings" {
  page_length          = 51
  list_child_products  = false
  additional_languages = ["fr"]
}


# Currencies

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


# Payment Gateways

## Adyen

resource "moltin_gateway_adyen" "adyen" {
  enabled          = true
  test             = true
  username         = "${var.adyen_username}"
  password         = "${var.adyen_password}"
  merchant_account = "${var.adyen_merchant_account}"
}

## Braintree

resource "moltin_gateway_braintree" "braintree" {
  enabled     = false
  environment = "sandbox"
  merchant_id = "${var.braintree_merchant_id}"
  private_key = "${var.braintree_private_key}"
  public_key  = "${var.braintree_public_key}"
}

## CardConnect

resource "moltin_gateway_card_connect" "cardconnect" {
  enabled     = false
  merchant_id = "${var.card_connect_merchant_id}"
  username    = "${var.card_connect_username}"
  password    = "${var.card_connect_password}"
}

## Stripe

resource "moltin_gateway_stripe" "stripe" {
  enabled = false
  login   = "${var.stripe_token}"
}


# Integrations

resource "moltin_integration" "order_placed_email" {
  enabled     = true
  name        = "Order Placed Email"
  description = "Notify a lambda function (via API Gateway) which sends an email to the customer when an order is placed"
  type        = "webhook"
  observes    = ["order.created"]

  configuration {
    url        = "https://path-to-lambda.aws.com/invoke"
    secret_key = "${var.integrations_order_placed_email_secret}"
  }
}


# Flows

## Flow

resource "moltin_flow" "wishlist" {
  name        = "Customer Wishlists"
  slug        = "customer-wishlists"
  description = "Allow logged in users to create and edit wishlists"
  enabled     = true
}

## Fields

// resource "moltin_field" "customer" {
//   required = true
//   unique   = false
//   type     = "relationship"

//   validation {
//     type = "one-to-one"
//     to   = "customer"
//   }

//   flow {
//     id = moltin_flow.wishlist.id
//   }
// }

// resource "moltin_field" "products" {
//   required = false
//   unique = false
//   type = "relationship"
//   validation {
//     type = "one-to-many"
//     to   = "product"
//   }
//   flow {
//     id = moltin_flow.wishlist.id
//   }
// }









// resource "moltin_brand" "DeWalt" {
//   name        = "DeWalt"
//   slug        = "dewalt"
//   description = "DeWalt is an American worldwide brand of power tools and hand tools for the construction, manufacturing and woodworking industries. DeWalt is a trade name of Black & Decker Inc., a subsidiary of Stanley Black & Decker."
//   status      = "live"
// }

// resource "moltin_brand" "Makita" {
//   name        = "Makita"
//   slug        = "makita"
//   description = "Makita Corporation is a Japanese manufacturer of power tools. Founded on March 21, 1915, it is based in Anjō, Japan, and operates factories in Brazil, Canada, China, Japan, Mexico, Romania, the United Kingdom, Germany and the United States."
//   status      = "live"
// }


// resource "moltin_category" "Drills" {
//   name        = "Drills"
//   slug        = "drills"
//   description = "Drills"
//   status      = "live"
// }

// resource "moltin_collection" "Pro" {
//   name        = "Pro"
//   slug        = "pro"
//   description = "Pro tools for the trade"
//   status      = "live"
// }



// resource "moltin_product" "DEWALT_DCD776S2T" {
//   name            = "DeWalt DCD776S2T-GB 18V 1.5Ah Li-Ion XR Cordless Combi Drill"
//   slug            = "dcd776S2t"
//   sku             = "DCD776S2T"
//   description     = "Powerful and versatile cordless combi drill with 15 torque settings as well as drill and hammer settings. Thermal overload protection increases durability and a fan-cooled, 2-speed variable motor provides long lasting performance. An integrated LED light improves visibility in dimly lit areas. A 13mm keyless chuck allows for secure, one-handed tightening. Lightweight compact design make it easy to handle and suitable for drilling in tight spaces. Supplied with 2 x batteries, charger and TSTAK kit box."
//   status          = "live"
//   commodity_type  = "physical"
//   manage_stock    = false

//   price {
//     amount        = 9999
//     currency      = "GBP"
//     includes_tax  = true
//   }

//   price {
//     amount        = 11999
//     currency      = "USD"
//     includes_tax  = true
//   }
// }

// resource "moltin_product" "DEWALT_DCD776D2T" {
//   name            = "DeWalt DCD776D2T-GB 18V 2.0Ah Li-Ion XR Cordless Combi Drill"
//   slug            = "dcd776d2t"
//   sku             = "DCD776D2T"
//   description     = "Compact drill driver with XR Li-ion battery technology, designed for efficiency and making applications faster. Improved ergonomic design and rubber grip increases user comfort and provides maximum control. Supplied with multi-voltage charger for all XR Li-ion slide pack batteries. Supplied in a TSTAK kit box. Manufactured in the UK."
//   status          = "live"
//   commodity_type  = "physical"
//   manage_stock    = true

//   price {
//     amount        = 11999
//     currency      = "GBP"
//     includes_tax  = true
//   }

//   price {
//     amount        = 13999
//     currency      = "USD"
//     includes_tax  = true
//   }
// }
