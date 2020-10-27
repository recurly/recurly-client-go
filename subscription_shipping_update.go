// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
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
