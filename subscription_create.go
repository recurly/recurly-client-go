// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type SubscriptionCreate struct {
	Params `json:"-"`

	// You must provide either a `plan_code` or `plan_id`. If both are provided the `plan_id` will be used.
	PlanCode *string `json:"plan_code,omitempty"`

	// You must provide either a `plan_code` or `plan_id`. If both are provided the `plan_id` will be used.
	PlanId *string `json:"plan_id,omitempty"`

	Account *AccountCreate `json:"account,omitempty"`

	// Create a shipping address on the account and assign it to the subscription.
	Shipping *SubscriptionShippingCreate `json:"shipping,omitempty"`

	// Collection method
	CollectionMethod *string `json:"collection_method,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// Override the unit amount of the subscription plan by setting this value. If not provided, the subscription will inherit the price from the subscription plan for the provided currency.
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// Optionally override the default quantity of 1.
	Quantity *int `json:"quantity,omitempty"`

	// Add-ons
	AddOns []SubscriptionAddOnCreate `json:"add_ons,omitempty"`

	// Optional coupon code to redeem on the account and discount the subscription. Please note, the subscription request will fail if the coupon is invalid.
	CouponCode *string `json:"coupon_code,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomFieldCreate `json:"custom_fields,omitempty"`

	// If set, overrides the default trial behavior for the subscription. The date must be in the future.
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

	// This will default to the Terms and Conditions text specified on the Invoice Settings page in your Recurly admin. Specify custom notes to add or override Terms and Conditions. Custom notes will stay with a subscription on all renewals.
	TermsAndConditions *string `json:"terms_and_conditions,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings. Specify custom notes to add or override Customer Notes. Custom notes will stay with a subscription on all renewals.
	CustomerNotes *string `json:"customer_notes,omitempty"`

	// If there are pending credits on the account that will be invoiced during the subscription creation, these will be used as the Customer Notes on the credit invoice.
	CreditCustomerNotes *string `json:"credit_customer_notes,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber *string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms *int `json:"net_terms,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType *string `json:"transaction_type,omitempty"`
}

func (attr *SubscriptionCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
