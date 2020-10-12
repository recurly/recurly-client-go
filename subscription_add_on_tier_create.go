package recurly

import ()

type SubscriptionAddOnTierCreate struct {
	Params `json:"-"`

	// Ending quantity
	EndingQuantity *int `json:"ending_quantity,omitempty"`

	// Unit amount
	UnitAmount *float64 `json:"unit_amount,omitempty"`
}

func (attr *SubscriptionAddOnTierCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
