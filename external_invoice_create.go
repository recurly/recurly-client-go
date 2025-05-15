// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type ExternalInvoiceCreate struct {

	// An identifier which associates the external invoice to a corresponding object in an external platform.
	ExternalId *string `json:"external_id,omitempty"`

	State *string `json:"state,omitempty"`

	Total *string `json:"total,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// When the invoice was created in the external platform.
	PurchasedAt *time.Time `json:"purchased_at,omitempty"`

	LineItems *[]ExternalChargeCreate `json:"line_items,omitempty"`

	ExternalPaymentPhase *ExternalPaymentPhaseBase `json:"external_payment_phase,omitempty"`

	// External payment phase ID, e.g. `a34ypb2ef9w1`.
	ExternalPaymentPhaseId *string `json:"external_payment_phase_id,omitempty"`
}
