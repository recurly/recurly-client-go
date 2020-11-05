// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type CouponCreate struct {
	Params `json:"-"`

	// The internal name for the coupon.
	Name *string `json:"name,omitempty"`

	// A maximum number of redemptions for the coupon. The coupon will expire when it hits its maximum redemptions.
	MaxRedemptions *int `json:"max_redemptions,omitempty"`

	// Redemptions per account is the number of times a specific account can redeem the coupon. Set redemptions per account to `1` if you want to keep customers from gaming the system and getting more than one discount from the coupon campaign.
	MaxRedemptionsPerAccount *int `json:"max_redemptions_per_account,omitempty"`

	// This description will show up when a customer redeems a coupon on your Hosted Payment Pages, or if you choose to show the description on your own checkout page.
	HostedDescription *string `json:"hosted_description,omitempty"`

	// Description of the coupon on the invoice.
	InvoiceDescription *string `json:"invoice_description,omitempty"`

	// The date and time the coupon will expire and can no longer be redeemed. Time is always 11:59:59, the end-of-day Pacific time.
	RedeemByDate *string `json:"redeem_by_date,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code *string `json:"code,omitempty"`

	// The type of discount provided by the coupon (how the amount discounted is calculated)
	DiscountType *string `json:"discount_type,omitempty"`

	// The percent of the price discounted by the coupon.  Required if `discount_type` is `percent`.
	DiscountPercent *int `json:"discount_percent,omitempty"`

	// Description of the unit of time the coupon is for. Used with `free_trial_amount` to determine the duration of time the coupon is for.  Required if `discount_type` is `free_trial`.
	FreeTrialUnit *string `json:"free_trial_unit,omitempty"`

	// Sets the duration of time the `free_trial_unit` is for. Required if `discount_type` is `free_trial`.
	FreeTrialAmount *int `json:"free_trial_amount,omitempty"`

	// Fixed discount currencies by currency. Required if the coupon type is `fixed`. This parameter should contain the coupon discount values
	Currencies []CouponPricing `json:"currencies,omitempty"`

	// The coupon is valid for one-time, non-plan charges if true.
	AppliesToNonPlanCharges *bool `json:"applies_to_non_plan_charges,omitempty"`

	// The coupon is valid for all plans if true. If false then
	// `plans` and `plans_names` will list the applicable plans.
	AppliesToAllPlans *bool `json:"applies_to_all_plans,omitempty"`

	// To apply coupon to Items in your Catalog, include a list
	// of `item_codes` in the request that the coupon will apply to. Or set value
	// to true to apply to all Items in your Catalog. The following values
	// are not permitted when `applies_to_all_items` is included: `free_trial_amount`
	// and `free_trial_unit`.
	AppliesToAllItems *bool `json:"applies_to_all_items,omitempty"`

	// List of plan codes to which this coupon applies. Required
	// if `applies_to_all_plans` is false. Overrides `applies_to_all_plans`
	// when `applies_to_all_plans` is true.
	PlanCodes []string `json:"plan_codes,omitempty"`

	// List of item codes to which this coupon applies. Sending
	// `item_codes` is only permitted when `applies_to_all_items` is set to false.
	// The following values are not permitted when `item_codes` is included:
	// `free_trial_amount` and `free_trial_unit`.
	ItemCodes []string `json:"item_codes,omitempty"`

	// This field does not apply when the discount_type is `free_trial`.
	// - "single_use" coupons applies to the first invoice only.
	// - "temporal" coupons will apply to invoices for the duration determined by the `temporal_unit` and `temporal_amount` attributes.
	// - "forever" coupons will apply to invoices forever.
	Duration *string `json:"duration,omitempty"`

	// If `duration` is "temporal" than `temporal_amount` is an integer which is multiplied by `temporal_unit` to define the duration that the coupon will be applied to invoices for.
	TemporalAmount *int `json:"temporal_amount,omitempty"`

	// If `duration` is "temporal" than `temporal_unit` is multiplied by `temporal_amount` to define the duration that the coupon will be applied to invoices for.
	TemporalUnit *string `json:"temporal_unit,omitempty"`

	// Whether the coupon is "single_code" or "bulk". Bulk coupons will require a `unique_code_template` and will generate unique codes through the `/generate` endpoint.
	CouponType *string `json:"coupon_type,omitempty"`

	// On a bulk coupon, the template from which unique coupon codes are generated.
	// - You must start the template with your coupon_code wrapped in single quotes.
	// - Outside of single quotes, use a 9 for a character that you want to be a random number.
	// - Outside of single quotes, use an "x" for a character that you want to be a random letter.
	// - Outside of single quotes, use an * for a character that you want to be a random number or letter.
	// - Use single quotes ' ' for characters that you want to remain static. These strings can be alphanumeric and may contain a - _ or +.
	// For example: "'abc-'****'-def'"
	UniqueCodeTemplate *string `json:"unique_code_template,omitempty"`

	// Whether the discount is for all eligible charges on the account, or only a specific subscription.
	RedemptionResource *string `json:"redemption_resource,omitempty"`
}

func (attr *CouponCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
