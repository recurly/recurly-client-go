package recurly

import ()

type SubscriptionShippingUpdate struct {
	Params `json:"-"`

	// Object type
	Object *string `json:"object,omitempty"`

	Address *ShippingAddressCreate `json:"address,omitempty"`

	// Assign a shipping address from the account's existing shipping addresses.
	AddressId *string `json:"address_id,omitempty"`
}

func (attr *SubscriptionShippingUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
