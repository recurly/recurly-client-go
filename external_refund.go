package recurly

import (
	"time"
)

type ExternalRefund struct {
	Params `json:"-"`

	// Payment method used for external refund transaction.
	PaymentMethod *string `json:"payment_method,omitempty"`

	// Used as the refund transactions' description.
	Description *string `json:"description,omitempty"`

	// Date the external refund payment was made. Defaults to the current date-time.
	RefundedAt *time.Time `json:"refunded_at,omitempty"`
}

func (attr *ExternalRefund) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
