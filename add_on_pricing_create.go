package recurly

import ()

type AddOnPricingCreate struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// Unit price
	UnitAmount *float64 `json:"unit_amount,omitempty"`
}

func (attr *AddOnPricingCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
