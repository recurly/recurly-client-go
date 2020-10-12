package recurly

import ()

type AccountAcquisitionCostCreate struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// The amount of the corresponding currency used to acquire the account.
	Amount *float64 `json:"amount,omitempty"`
}

func (attr *AccountAcquisitionCostCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
