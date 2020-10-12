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
}

func (attr *PlanPricingCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
