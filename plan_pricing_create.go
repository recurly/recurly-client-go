// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type PlanPricingCreate struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// Amount of one-time setup fee automatically charged at the beginning of a subscription billing cycle. For subscription plans with a trial, the setup fee will be charged at the time of signup. Setup fees do not increase with the quantity of a subscription plan.
	SetupFee *float64 `json:"setup_fee,omitempty"`

	// Unit price
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// Determines whether or not tax is included in the unit amount. The Tax Inclusive Pricing feature (separate from the Mixed Tax Pricing feature) must be enabled to use this flag.
	TaxInclusive *bool `json:"tax_inclusive,omitempty"`
}

func (attr *PlanPricingCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
