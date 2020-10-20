package recurly

import ()

type CustomFieldCreate struct {
	Params `json:"-"`

	// Fields must be created in the UI before values can be assigned to them.
	Name *string `json:"name,omitempty"`

	// Any values that resemble a credit card number or security code (CVV/CVC) will be rejected.
	Value *string `json:"value,omitempty"`
}

func (attr *CustomFieldCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
