package recurly

import ()

type CouponBulkCreate struct {
	Params `json:"-"`

	// The quantity of unique coupon codes to generate
	NumberOfUniqueCodes *int `json:"number_of_unique_codes,omitempty"`
}

func (attr *CouponBulkCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
