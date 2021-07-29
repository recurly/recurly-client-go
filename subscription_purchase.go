// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type SubscriptionPurchase struct {

	// Plan code
	PlanCode *string `json:"plan_code,omitempty"`

	// Plan ID
	PlanId *string `json:"plan_id,omitempty"`

	// Override the unit amount of the subscription plan by setting this value. If not provided, the subscription will inherit the price from the subscription plan for the provided currency.
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// Optionally override the default quantity of 1.
	Quantity *int `json:"quantity,omitempty"`

	// Add-ons
	AddOns []SubscriptionAddOnCreate `json:"add_ons,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomFieldCreate `json:"custom_fields,omitempty"`

	// Create a shipping address on the account and assign it to the subscription.
	Shipping *SubscriptionShippingPurchase `json:"shipping,omitempty"`

	// If set, overrides the default trial behavior for the subscription. When the current date time or a past date time is provided the subscription will begin with no trial phase (overriding any plan default trial). When a future date time is provided the subscription will begin with a trial phase ending at the specified date time.
	TrialEndsAt *time.Time `json:"trial_ends_at,omitempty"`

	// If set, the subscription will begin in the future on this date. The subscription will apply the setup fee and trial period, unless the plan has no trial.
	StartsAt *time.Time `json:"starts_at,omitempty"`

	// If present, this sets the date the subscription's next billing period will start (`current_period_ends_at`). This can be used to align the subscription’s billing to a specific day of the month. The initial invoice will be prorated for the period between the subscription's activation date and the billing period end date. Subsequent periods will be based off the plan interval. For a subscription with a trial period, this will change when the trial expires.
	NextBillDate *time.Time `json:"next_bill_date,omitempty"`

	// The number of cycles/billing periods in a term. When `remaining_billing_cycles=0`, if `auto_renew=true` the subscription will renew and a new term will begin, otherwise the subscription will expire.
	TotalBillingCycles *int `json:"total_billing_cycles,omitempty"`

	// If `auto_renew=true`, when a term completes, `total_billing_cycles` takes this value as the length of subsequent terms. Defaults to the plan's `total_billing_cycles`.
	RenewalBillingCycles *int `json:"renewal_billing_cycles,omitempty"`

	// Whether the subscription renews at the end of its term.
	AutoRenew *bool `json:"auto_renew,omitempty"`

	// Revenue schedule type
	RevenueScheduleType *string `json:"revenue_schedule_type,omitempty"`
}
