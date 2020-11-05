// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type SubscriptionUpdate struct {
	Params `json:"-"`

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

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms *int `json:"net_terms,omitempty"`

	// Subscription shipping details
	Shipping *SubscriptionShippingUpdate `json:"shipping,omitempty"`

	// The `billing_info_id` is the value that represents a specific billing info for an end customer. When `billing_info_id` is used to assign billing info to the subscription, all future billing events for the subscription will bill to the specified billing info.
	BillingInfoId *string `json:"billing_info_id,omitempty"`
}

func (attr *SubscriptionUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
