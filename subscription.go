// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type Subscription struct {
	recurlyResponse *ResponseMetadata

	// Subscription ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
	Uuid string `json:"uuid,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// Just the important parts.
	Plan PlanMini `json:"plan,omitempty"`

	// State
	State string `json:"state,omitempty"`

	// Subscription shipping details
	Shipping SubscriptionShipping `json:"shipping,omitempty"`

	// Coupon redemptions
	CouponRedemptions []CouponRedemptionMini `json:"coupon_redemptions,omitempty"`

	// Subscription Change
	PendingChange SubscriptionChange `json:"pending_change,omitempty"`

	// Current billing period started at
	CurrentPeriodStartedAt time.Time `json:"current_period_started_at,omitempty"`

	// Current billing period ends at
	CurrentPeriodEndsAt time.Time `json:"current_period_ends_at,omitempty"`

	// The start date of the term when the first billing period starts. The subscription term is the length of time that a customer will be committed to a subscription. A term can span multiple billing periods.
	CurrentTermStartedAt time.Time `json:"current_term_started_at,omitempty"`

	// When the term ends. This is calculated by a plan's interval and `total_billing_cycles` in a term. Subscription changes with a `timeframe=renewal` will be applied on this date.
	CurrentTermEndsAt time.Time `json:"current_term_ends_at,omitempty"`

	// Trial period started at
	TrialStartedAt time.Time `json:"trial_started_at,omitempty"`

	// Trial period ends at
	TrialEndsAt time.Time `json:"trial_ends_at,omitempty"`

	// The remaining billing cycles in the current term.
	RemainingBillingCycles int `json:"remaining_billing_cycles,omitempty"`

	// The number of cycles/billing periods in a term. When `remaining_billing_cycles=0`, if `auto_renew=true` the subscription will renew and a new term will begin, otherwise the subscription will expire.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// If `auto_renew=true`, when a term completes, `total_billing_cycles` takes this value as the length of subsequent terms. Defaults to the plan's `total_billing_cycles`.
	RenewalBillingCycles int `json:"renewal_billing_cycles,omitempty"`

	// Whether the subscription renews at the end of its term.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Null unless subscription is paused or will pause at the end of the current billing period.
	PausedAt time.Time `json:"paused_at,omitempty"`

	// Null unless subscription is paused or will pause at the end of the current billing period.
	RemainingPauseCycles int `json:"remaining_pause_cycles,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Subscription unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Subscription quantity
	Quantity int `json:"quantity,omitempty"`

	// Add-ons
	AddOns []SubscriptionAddOn `json:"add_ons,omitempty"`

	// Total price of add-ons
	AddOnsTotal float64 `json:"add_ons_total,omitempty"`

	// Estimated total, before tax.
	Subtotal float64 `json:"subtotal,omitempty"`

	// Collection method
	CollectionMethod string `json:"collection_method,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	// Terms and conditions
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// Customer notes
	CustomerNotes string `json:"customer_notes,omitempty"`

	// Expiration reason
	ExpirationReason string `json:"expiration_reason,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Activated at
	ActivatedAt time.Time `json:"activated_at,omitempty"`

	// Canceled at
	CanceledAt time.Time `json:"canceled_at,omitempty"`

	// Expires at
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	// Recurring subscriptions paid with ACH will have this attribute set. This timestamp is used for alerting customers to reauthorize in 3 years in accordance with NACHA rules. If a subscription becomes inactive or the billing info is no longer a bank account, this timestamp is cleared.
	BankAccountAuthorizedAt time.Time `json:"bank_account_authorized_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Subscription) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Subscription) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionList struct {
	ListMetadata
	Data            []Subscription `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionList allows you to paginate Subscription objects
type SubscriptionList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []Subscription
}

func NewSubscriptionList(client HttpCaller, nextPagePath string) *SubscriptionList {
	return &SubscriptionList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionList) Fetch() error {
	resources := &subscriptionList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionList) Count() (*int64, error) {
	resources := &subscriptionList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
