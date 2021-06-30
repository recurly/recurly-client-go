// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type BillingInfoVerify struct {
	Params `json:"-"`

	// An identifier for a specific payment gateway.
	GatewayCode *string `json:"gateway_code,omitempty"`
}

func (attr *BillingInfoVerify) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
