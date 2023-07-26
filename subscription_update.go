// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type SubscriptionUpdate struct {

	// Change collection method
	CollectionMethod *string `json:"collection_method,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomFieldCreate `json:"custom_fields,omitempty"`

	// The remaining billing cycles in the current term.
	RemainingBillingCycles *int `json:"remaining_billing_cycles,omitempty"`

	// If `auto_renew=true`, when a term completes, `total_billing_cycles` takes this value as the length of subsequent terms. Defaults to the plan's `total_billing_cycles`.
	RenewalBillingCycles *int `json:"renewal_billing_cycles,omitempty"`

	// Whether the subscription renews at the end of its term.
	AutoRenew *bool `json:"auto_renew,omitempty"`

	// If present, this sets the date the subscription's next billing period will start (`current_period_ends_at`). This can be used to align the subscription’s billing to a specific day of the month. For a subscription in a trial period, this will change when the trial expires. This parameter is useful for postponement of a subscription to change its billing date without proration.
	NextBillDate *time.Time `json:"next_bill_date,omitempty"`

	// Revenue schedule type
	RevenueScheduleType *string `json:"revenue_schedule_type,omitempty"`

	// Specify custom notes to add or override Terms and Conditions. Custom notes will stay with a subscription on all renewals.
	TermsAndConditions *string `json:"terms_and_conditions,omitempty"`

	// Specify custom notes to add or override Customer Notes. Custom notes will stay with a subscription on all renewals.
	CustomerNotes *string `json:"customer_notes,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber *string `json:"po_number,omitempty"`

	// Integer paired with `Net Terms Type` and representing the number
	// of days past the current date (for `net` Net Terms Type) or days after
	// the last day of the current month (for `eom` Net Terms Type) that the
	// invoice will become past due. For any value, an additional 24 hours is
	// added to ensure the customer has the entire last day to make payment before
	// becoming past due.  For example:
	// If an invoice is due `net 0`, it is due 'On Receipt' and will become past due 24 hours after it's created.
	// If an invoice is due `net 30`, it will become past due at 31 days exactly.
	// If an invoice is due `eom 30`, it will become past due 31 days from the last day of the current month.
	// When `eom` Net Terms Type is passed, the value for `Net Terms` is restricted to `0, 15, 30, 45, 60, or 90`.
	// For more information please visit our docs page (https://docs.recurly.com/docs/manual-payments#section-collection-terms)
	NetTerms *int `json:"net_terms,omitempty"`

	// Optionally supplied string that may be either `net` or `eom` (end-of-month).
	// When `net`, an invoice becomes past due the specified number of `Net Terms` days from the current date.
	// When `eom` an invoice becomes past due the specified number of `Net Terms` days from the last day of the current month.
	// This field is only available when the EOM Net Terms feature is enabled.
	NetTermsType *string `json:"net_terms_type,omitempty"`

	// If present, this subscription's transactions will use the payment gateway with this code.
	GatewayCode *string `json:"gateway_code,omitempty"`

	// This field is deprecated. Please do not use it.
	TaxInclusive *bool `json:"tax_inclusive,omitempty"`

	// Subscription shipping details
	Shipping *SubscriptionShippingUpdate `json:"shipping,omitempty"`

	// The `billing_info_id` is the value that represents a specific billing info for an end customer. When `billing_info_id` is used to assign billing info to the subscription, all future billing events for the subscription will bill to the specified billing info. `billing_info_id` can ONLY be used for sites utilizing the Wallet feature.
	BillingInfoId *string `json:"billing_info_id,omitempty"`
}
