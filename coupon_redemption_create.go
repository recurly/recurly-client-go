// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type CouponRedemptionCreate struct {
	Params `json:"-"`

	// Coupon ID
	CouponId *string `json:"coupon_id,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// Subscription ID
	SubscriptionId *string `json:"subscription_id,omitempty"`
}

func (attr *CouponRedemptionCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
