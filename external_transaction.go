package recurly

import (
	"time"
)

type ExternalTransaction struct {
	Params `json:"-"`

	// Payment method used for the external transaction.
	PaymentMethod *string `json:"payment_method,omitempty"`

	// Used as the transaction's description.
	Description *string `json:"description,omitempty"`

	// The total amount of the transcaction. Cannot excceed the invoice total.
	Amount *float64 `json:"amount,omitempty"`

	// Datetime that the external payment was collected. Defaults to current datetime.
	CollectedAt *time.Time `json:"collected_at,omitempty"`
}

func (attr *ExternalTransaction) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
