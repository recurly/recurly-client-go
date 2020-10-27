// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
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
