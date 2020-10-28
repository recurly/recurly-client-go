// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type TierCreate struct {
	Params `json:"-"`

	// Ending quantity
	EndingQuantity *int `json:"ending_quantity,omitempty"`

	// Tier pricing
	Currencies []PricingCreate `json:"currencies,omitempty"`
}

func (attr *TierCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
