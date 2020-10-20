package recurly

import ()

type ShippingPurchase struct {
	Params `json:"-"`

	// Assign a shipping address from the account's existing shipping addresses. If this and `address` are both present, `address` will take precedence.
	AddressId *string `json:"address_id,omitempty"`

	Address *ShippingAddressCreate `json:"address,omitempty"`

	// A list of shipping fees to be created as charges with the purchase.
	Fees []ShippingFeeCreate `json:"fees,omitempty"`
}

func (attr *ShippingPurchase) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
