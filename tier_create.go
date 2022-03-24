// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type TierCreate struct {

	// Ending quantity for the tier.  This represents a unit amount for unit-priced add ons. Must be left empty if it is the final tier.
	EndingQuantity *int `json:"ending_quantity,omitempty"`

	// (deprecated) -- Use the percentage_tiers object instead.
	UsagePercentage *string `json:"usage_percentage,omitempty"`

	// Tier pricing
	Currencies []TierPricingCreate `json:"currencies,omitempty"`
}
