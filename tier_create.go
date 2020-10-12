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
