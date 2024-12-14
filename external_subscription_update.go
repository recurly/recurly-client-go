// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type ExternalSubscriptionUpdate struct {
	ExternalProductReference *ExternalProductReferenceUpdate `json:"external_product_reference,omitempty"`

	// Id of the subscription in the external system, i.e. Apple App Store or Google Play Store.
	ExternalId *string `json:"external_id,omitempty"`

	// When a new billing event occurred on the external subscription in conjunction with a recent billing period, reactivation or upgrade/downgrade.
	LastPurchased *time.Time `json:"last_purchased,omitempty"`

	// An indication of whether or not the external subscription will auto-renew at the expiration date.
	AutoRenew *bool `json:"auto_renew,omitempty"`

	// External subscriptions can be active, canceled, expired, past_due, voided, revoked, or paused.
	State *string `json:"state,omitempty"`

	// Identifier of the app that generated the external subscription.
	AppIdentifier *string `json:"app_identifier,omitempty"`

	// An indication of the quantity of a subscribed item's quantity.
	Quantity *int `json:"quantity,omitempty"`

	// When the external subscription was activated in the external platform.
	ActivatedAt *time.Time `json:"activated_at,omitempty"`

	// When the external subscription expires in the external platform.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// When the external subscription trial period started in the external platform.
	TrialStartedAt *time.Time `json:"trial_started_at,omitempty"`

	// When the external subscription trial period ends in the external platform.
	TrialEndsAt *time.Time `json:"trial_ends_at,omitempty"`

	// An indication of whether or not the external subscription was being created by a historical data import.
	Imported *bool `json:"imported,omitempty"`
}
