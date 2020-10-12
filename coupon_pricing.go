package recurly

import ()

type CouponPricing struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// The fixed discount (in dollars) for the corresponding currency.
	Discount *float64 `json:"discount,omitempty"`
}

func (attr *CouponPricing) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
