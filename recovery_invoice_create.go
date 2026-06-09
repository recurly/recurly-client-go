// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type RecoveryInvoiceCreate struct {

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// Date invoice was originally due. Must be in the past.
	DueAt *time.Time `json:"due_at,omitempty"`

	// This identifies the PO number associated with the subscription.
	PoNumber *string `json:"po_number,omitempty"`

	// Must be set to `true` to acknowledge that the invoice is eligible for external recovery. Requests with `false`, omitted, or non-boolean values will be rejected.
	ExternalRecoveryEligible *bool `json:"external_recovery_eligible,omitempty"`

	Account *RecoveryAccountCreate `json:"account,omitempty"`

	// Line items to include on the invoice. Currency is specified at the root level and must not be included in individual line items.
	LineItems *[]RecoveryLineItemCreate `json:"line_items,omitempty"`
}
