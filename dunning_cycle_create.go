// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type DunningCycleCreate struct {

	// The type of invoice this cycle applies to.
	Type *string `json:"type,omitempty"`

	// Only meaningful on the `trial` cycle, where sending `false` removes it. Other cycle types cannot be deactivated.
	Active *bool `json:"active,omitempty"`

	// Whether the dunning settings will be applied to manual trials. Only applies to trial cycles.
	AppliesToManualTrial *bool `json:"applies_to_manual_trial,omitempty"`

	// Whether the subscription(s) should be cancelled at the end of the dunning cycle.
	ExpireSubscription *bool `json:"expire_subscription,omitempty"`

	// Whether the invoice should be failed at the end of the dunning cycle.
	FailInvoice *bool `json:"fail_invoice,omitempty"`

	// Whether or not to send an extra email immediately to customers whose initial payment attempt fails with either a hard decline or invalid billing info.
	SendImmediatelyOnHardDecline *bool `json:"send_immediately_on_hard_decline,omitempty"`

	// Number of days to extend external payment recovery. Only available when the site has external payment retries enabled.
	ExternalPaymentRecoveryExtensionDays *int `json:"external_payment_recovery_extension_days,omitempty"`

	// Dunning intervals. Required unless `active` is `false`.
	Intervals *[]DunningIntervalCreate `json:"intervals,omitempty"`
}
