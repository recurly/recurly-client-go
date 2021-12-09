// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type PricingCreate struct {

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// Unit price
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// Determines whether or not tax is included in the unit amount. The Tax Inclusive Pricing feature (separate from the Mixed Tax Pricing feature) must be enabled to use this flag.
	TaxInclusive *bool `json:"tax_inclusive,omitempty"`
}
