// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type LineItemRefund struct {
	Params `json:"-"`

	// Line item ID
	Id *string `json:"id,omitempty"`

	// Line item quantity to be refunded.
	Quantity *int `json:"quantity,omitempty"`

	// Set to `true` if the line item should be prorated; set to `false` if not.
	// This can only be used on line items that have a start and end date.
	Prorate *bool `json:"prorate,omitempty"`
}

func (attr *LineItemRefund) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
