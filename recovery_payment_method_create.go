// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type RecoveryPaymentMethodCreate struct {

	// The payment method type.
	Object *string `json:"object,omitempty"`

	// The card brand (e.g. `Visa`, `MasterCard`). Present for `credit_card`, `apple_pay`, and `google_pay`/`google_pay_device_pan`; omitted for `paypal_billing_agreement`.
	CardType *string `json:"card_type,omitempty"`

	// For a plain card, the card's own first six digits (BIN).
	// For a tokenized wallet payment (`apple_pay`, `google_pay`, or
	// `google_pay_device_pan`), this is the DPAN's (the wallet/device token's own
	// number) first six digits — **not** the underlying card's (FPAN). The FPAN is
	// never accepted or represented; no separate wallet-specific field is provided.
	FirstSix *string `json:"first_six,omitempty"`

	// The card's (or, for a tokenized wallet payment, the DPAN's) last four digits. See `first_six` for the DPAN-vs-FPAN distinction on wallets.
	LastFour *string `json:"last_four,omitempty"`

	// Expiration month.
	ExpMonth *int `json:"exp_month,omitempty"`

	// Expiration year.
	ExpYear *int `json:"exp_year,omitempty"`
}
