// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type AddOnPricingCreate struct {

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// Allows up to 2 decimal places. Required unless `unit_amount_decimal` is provided.
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// Allows up to 9 decimal places. Only supported when `add_on_type` = `usage`.
	// If `unit_amount_decimal` is provided, `unit_amount` cannot be provided.
	UnitAmountDecimal *string `json:"unit_amount_decimal,omitempty"`
}
