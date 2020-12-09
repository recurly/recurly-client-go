// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type SubscriptionShippingCreate struct {
	Address *ShippingAddressCreate `json:"address,omitempty"`

	// Assign a shipping address from the account's existing shipping addresses. If `address_id` and `address` are both present, `address` will be used.
	AddressId *string `json:"address_id,omitempty"`

	// The id of the shipping method used to deliver the subscription. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodId *string `json:"method_id,omitempty"`

	// The code of the shipping method used to deliver the subscription. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodCode *string `json:"method_code,omitempty"`

	// Assigns the subscription's shipping cost. If this is greater than zero then a `method_id` or `method_code` is required.
	Amount *float64 `json:"amount,omitempty"`
}
