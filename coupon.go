// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type Coupon struct {
	recurlyResponse *ResponseMetadata

	// Coupon ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code string `json:"code,omitempty"`

	// The internal name for the coupon.
	Name string `json:"name,omitempty"`

	// Indicates if the coupon is redeemable, and if it is not, why.
	State string `json:"state,omitempty"`

	// A maximum number of redemptions for the coupon. The coupon will expire when it hits its maximum redemptions.
	MaxRedemptions int `json:"max_redemptions,omitempty"`

	// Redemptions per account is the number of times a specific account can redeem the coupon. Set redemptions per account to `1` if you want to keep customers from gaming the system and getting more than one discount from the coupon campaign.
	MaxRedemptionsPerAccount int `json:"max_redemptions_per_account,omitempty"`

	// When this number reaches `max_redemptions` the coupon will no longer be redeemable.
	UniqueCouponCodesCount int `json:"unique_coupon_codes_count,omitempty"`

	// On a bulk coupon, the template from which unique coupon codes are generated.
	UniqueCodeTemplate string `json:"unique_code_template,omitempty"`

	// - "single_use" coupons applies to the first invoice only.
	// - "temporal" coupons will apply to invoices for the duration determined by the `temporal_unit` and `temporal_amount` attributes.
	Duration string `json:"duration,omitempty"`

	// If `duration` is "temporal" than `temporal_amount` is an integer which is multiplied by `temporal_unit` to define the duration that the coupon will be applied to invoices for.
	TemporalAmount int `json:"temporal_amount,omitempty"`

	// If `duration` is "temporal" than `temporal_unit` is multiplied by `temporal_amount` to define the duration that the coupon will be applied to invoices for.
	TemporalUnit string `json:"temporal_unit,omitempty"`

	// Description of the unit of time the coupon is for. Used with `free_trial_amount` to determine the duration of time the coupon is for.
	FreeTrialUnit string `json:"free_trial_unit,omitempty"`

	// Sets the duration of time the `free_trial_unit` is for.
	FreeTrialAmount int `json:"free_trial_amount,omitempty"`

	// The coupon is valid for all plans if true. If false then `plans` and `plans_names` will list the applicable plans.
	AppliesToAllPlans bool `json:"applies_to_all_plans,omitempty"`

	// The coupon is valid for all items if true. If false then `items`
	// will list the applicable items.
	AppliesToAllItems bool `json:"applies_to_all_items,omitempty"`

	// The coupon is valid for one-time, non-plan charges if true.
	AppliesToNonPlanCharges bool `json:"applies_to_non_plan_charges,omitempty"`

	// A list of plan names for which this coupon applies.
	PlansNames []string `json:"plans_names,omitempty"`

	// A list of plans for which this coupon applies. This will be `null` if `applies_to_all_plans=true`.
	Plans []PlanMini `json:"plans,omitempty"`

	// A list of items for which this coupon applies. This will be
	// `null` if `applies_to_all_items=true`.
	Items []ItemMini `json:"items,omitempty"`

	// Whether the discount is for all eligible charges on the account, or only a specific subscription.
	RedemptionResource string `json:"redemption_resource,omitempty"`

	// Details of the discount a coupon applies. Will contain a `type`
	// property and one of the following properties: `percent`, `fixed`, `trial`.
	Discount CouponDiscount `json:"discount,omitempty"`

	// Whether the coupon is "single_code" or "bulk". Bulk coupons will require a `unique_code_template` and will generate unique codes through the `/generate` endpoint.
	CouponType string `json:"coupon_type,omitempty"`

	// This description will show up when a customer redeems a coupon on your Hosted Payment Pages, or if you choose to show the description on your own checkout page.
	HostedPageDescription string `json:"hosted_page_description,omitempty"`

	// Description of the coupon on the invoice.
	InvoiceDescription string `json:"invoice_description,omitempty"`

	// The date and time the coupon will expire and can no longer be redeemed. Time is always 11:59:59, the end-of-day Pacific time.
	RedeemBy time.Time `json:"redeem_by,omitempty"`

	// The Coupon ID of the parent Bulk Coupon
	BulkCouponId string `json:"bulk_coupon_id,omitempty"`

	// The Coupon code of the parent Bulk Coupon
	BulkCouponCode string `json:"bulk_coupon_code,omitempty"`

	// The date and time the unique coupon code was redeemed. This is only present for bulk coupons.
	RedeemedAt time.Time `json:"redeemed_at,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date and time the coupon was expired early or reached its `max_redemptions`.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Coupon) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Coupon) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponList struct {
	ListMetadata
	Data            []Coupon `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponList allows you to paginate Coupon objects
type CouponList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []Coupon
}

func NewCouponList(client HttpCaller, nextPagePath string) *CouponList {
	return &CouponList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponList) Fetch() error {
	resources := &couponList{}
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
func (list *CouponList) Count() (*int64, error) {
	resources := &couponList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
