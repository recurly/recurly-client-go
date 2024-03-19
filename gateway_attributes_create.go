// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type GatewayAttributesCreate struct {
	Params `json:"-"`

	// Used by Adyen and Braintree gateways. For Adyen The Shopper Reference value used when the external token was created. Must be used in conjunction with gateway_token and gateway_code. For Braintree the PayPal PayerID is populated in the response.
	AccountReference *string `json:"account_reference,omitempty"`
}

func (attr *GatewayAttributesCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
