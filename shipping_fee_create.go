package recurly

import ()

type ShippingFeeCreate struct {
	Params `json:"-"`

	// The id of the shipping method used to deliver the purchase. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodId *string `json:"method_id,omitempty"`

	// The code of the shipping method used to deliver the purchase. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodCode *string `json:"method_code,omitempty"`

	// This is priced in the purchase's currency.
	Amount *float64 `json:"amount,omitempty"`
}

func (attr *ShippingFeeCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
