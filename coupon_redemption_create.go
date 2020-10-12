package recurly

import ()

type CouponRedemptionCreate struct {
	Params `json:"-"`

	// Coupon ID
	CouponId *string `json:"coupon_id,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`
}

func (attr *CouponRedemptionCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
