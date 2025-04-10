// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type ExternalPaymentPhaseBase struct {

	// Started At
	StartedAt *time.Time `json:"started_at,omitempty"`

	// Ends At
	EndsAt *time.Time `json:"ends_at,omitempty"`

	// Starting Billing Period Index
	StartingBillingPeriodIndex *int `json:"starting_billing_period_index,omitempty"`

	// Ending Billing Period Index
	EndingBillingPeriodIndex *int `json:"ending_billing_period_index,omitempty"`

	// Type of discount offer given, e.g. "FREE_TRIAL"
	OfferType *string `json:"offer_type,omitempty"`

	// Name of the discount offer given, e.g. "introductory"
	OfferName *string `json:"offer_name,omitempty"`

	// Number of billing periods
	PeriodCount *int `json:"period_count,omitempty"`

	// Billing cycle length
	PeriodLength *string `json:"period_length,omitempty"`

	// Allows up to 9 decimal places
	Amount *string `json:"amount,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`
}
