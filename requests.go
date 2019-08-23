package recurly

import "time"

type AccountCreate struct {
	Params `json:"-"`

	// The unique identifier of the account. This cannot be changed once the account is created.
	Code string `json:"code,omitempty"`

	Acquisition AccountAcquisitionUpdatable `json:"acquisition,omitempty"`

	ShippingAddresses []ShippingAddressCreate `json:"shipping_addresses,omitempty"`

	// A secondary value for the account.
	Username string `json:"username,omitempty"`

	// The email address used for communicating with this customer. The customer will also use this email address to log into your hosted account management pages. This value does not need to be unique.
	Email string `json:"email,omitempty"`

	// Used to determine the language and locale of emails sent on behalf of the merchant to the customer. The list of locales is restricted to those the merchant has enabled on the site.
	PreferredLocale string `json:"preferred_locale,omitempty"`

	// Additional email address that should receive account correspondence. These should be separated only by commas. These CC emails will receive all emails that the `email` field also receives.
	CcEmails string `json:"cc_emails,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	// The VAT number of the account (to avoid having the VAT applied). This is only used for manually collected invoices.
	VatNumber string `json:"vat_number,omitempty"`

	// The tax status of the account. `true` exempts tax on the account, `false` applies tax on the account.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// The tax exemption certificate number for the account. If the merchant has an integration for the Vertex tax provider, this optional value will be sent in any tax calculation requests for the account.
	ExemptionCertificate string `json:"exemption_certificate,omitempty"`

	// The account code of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountCode string `json:"parent_account_code,omitempty"`

	// The UUID of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountId string `json:"parent_account_id,omitempty"`

	// An enumerable describing the billing behavior of the account, specifically whether the account is self-paying or will rely on the parent account to pay.
	BillTo string `json:"bill_to,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`

	Address Address `json:"address,omitempty"`

	BillingInfo BillingInfoCreate `json:"billing_info,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

func (attr *AccountCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type AccountAcquisitionUpdatable struct {
	Params `json:"-"`

	// Account balance
	Cost AccountAcquisitionCost `json:"cost,omitempty"`

	// The channel through which the account was acquired.
	Channel string `json:"channel,omitempty"`

	// An arbitrary subchannel string representing a distinction/subcategory within a broader channel.
	Subchannel string `json:"subchannel,omitempty"`

	// An arbitrary identifier for the marketing campaign that led to the acquisition of this account.
	Campaign string `json:"campaign,omitempty"`
}

func (attr *AccountAcquisitionUpdatable) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type ShippingAddressCreate struct {
	Params `json:"-"`

	Nickname string `json:"nickname,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	Email string `json:"email,omitempty"`

	VatNumber string `json:"vat_number,omitempty"`

	Phone string `json:"phone,omitempty"`

	Street1 string `json:"street1,omitempty"`

	Street2 string `json:"street2,omitempty"`

	City string `json:"city,omitempty"`

	// State or province.
	Region string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO code.
	Country string `json:"country,omitempty"`
}

func (attr *ShippingAddressCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type BillingInfoCreate struct {
	Params `json:"-"`

	// A token [generated by Recurly.js](https://docs.recurly.com/js/#getting-a-token).
	TokenId string `json:"token_id,omitempty"`

	// First name
	FirstName string `json:"first_name,omitempty"`

	// Last name
	LastName string `json:"last_name,omitempty"`

	// Company name
	Company string `json:"company,omitempty"`

	Address Address `json:"address,omitempty"`

	// Credit card number, spaces and dashes are accepted.
	Number string `json:"number,omitempty"`

	// Expiration month
	Month string `json:"month,omitempty"`

	// Expiration year
	Year string `json:"year,omitempty"`

	// *STRONGLY RECOMMENDED*
	Cvv string `json:"cvv,omitempty"`

	// VAT number
	VatNumber string `json:"vat_number,omitempty"`

	// *STRONGLY RECOMMENDED* Customer's IP address when updating their billing information.
	IpAddress string `json:"ip_address,omitempty"`

	// A token used in place of a credit card in order to perform transactions. Must be used in conjunction with `gateway_code`.
	GatewayToken string `json:"gateway_token,omitempty"`

	// An identifier for a specific payment gateway. Must be used in conjunction with `gateway_token`.
	GatewayCode string `json:"gateway_code,omitempty"`

	// Amazon billing agreement ID
	AmazonBillingAgreementId string `json:"amazon_billing_agreement_id,omitempty"`

	// PayPal billing agreement ID
	PaypalBillingAgreementId string `json:"paypal_billing_agreement_id,omitempty"`

	// Fraud Session ID
	FraudSessionId string `json:"fraud_session_id,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`

	// A token generated by Recurly.js after completing a 3-D Secure device fingerprinting or authentication challenge.
	ThreeDSecureActionResultTokenId string `json:"three_d_secure_action_result_token_id,omitempty"`

	// The International Bank Account Number, up to 34 alphanumeric characters comprising a country code; two check digits; and a number that includes the domestic bank account number, branch identifier, and potential routing information. (SEPA only)
	Iban string `json:"iban,omitempty"`

	// The name associated with the bank account.
	NameOnAccount string `json:"name_on_account,omitempty"`

	// The bank account number. (ACH only)
	AccountNumber string `json:"account_number,omitempty"`

	// The bank's rounting number. (ACH only)
	RoutingNumber string `json:"routing_number,omitempty"`

	// The bank account type. (ACH only)
	AccountType string `json:"account_type,omitempty"`
}

func (attr *BillingInfoCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type AccountUpdate struct {
	Params `json:"-"`

	// A secondary value for the account.
	Username string `json:"username,omitempty"`

	// The email address used for communicating with this customer. The customer will also use this email address to log into your hosted account management pages. This value does not need to be unique.
	Email string `json:"email,omitempty"`

	// Used to determine the language and locale of emails sent on behalf of the merchant to the customer. The list of locales is restricted to those the merchant has enabled on the site.
	PreferredLocale string `json:"preferred_locale,omitempty"`

	// Additional email address that should receive account correspondence. These should be separated only by commas. These CC emails will receive all emails that the `email` field also receives.
	CcEmails string `json:"cc_emails,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	// The VAT number of the account (to avoid having the VAT applied). This is only used for manually collected invoices.
	VatNumber string `json:"vat_number,omitempty"`

	// The tax status of the account. `true` exempts tax on the account, `false` applies tax on the account.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// The tax exemption certificate number for the account. If the merchant has an integration for the Vertex tax provider, this optional value will be sent in any tax calculation requests for the account.
	ExemptionCertificate string `json:"exemption_certificate,omitempty"`

	// The account code of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountCode string `json:"parent_account_code,omitempty"`

	// The UUID of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountId string `json:"parent_account_id,omitempty"`

	// An enumerable describing the billing behavior of the account, specifically whether the account is self-paying or will rely on the parent account to pay.
	BillTo string `json:"bill_to,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`

	Address Address `json:"address,omitempty"`

	BillingInfo BillingInfoCreate `json:"billing_info,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

func (attr *AccountUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type CouponRedemptionCreate struct {
	Params `json:"-"`

	// Coupon ID
	CouponId string `json:"coupon_id,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`
}

func (attr *CouponRedemptionCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type InvoiceCreate struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// An automatic invoice means a corresponding transaction is run using the account's billing information at the same time the invoice is created. Manual invoices are created without a corresponding transaction. The merchant must enter a manual payment transaction or have the customer pay the invoice with an automatic method, like credit card, PayPal, Amazon, or ACH bank payment.
	CollectionMethod string `json:"collection_method,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings for charge invoices. Specify custom notes to add or override Customer Notes on charge invoices.
	ChargeCustomerNotes string `json:"charge_customer_notes,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings for credit invoices. Specify customer notes to add or override Customer Notes on credit invoices.
	CreditCustomerNotes string `json:"credit_customer_notes,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// This will default to the Terms and Conditions text specified on the Invoice Settings page in your Recurly admin. Specify custom notes to add or override Terms and Conditions.
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// VAT Reverse Charge Notes only appear if you have EU VAT enabled or are using your own Avalara AvaTax account and the customer is in the EU, has a VAT number, and is in a different country than your own. This will default to the VAT Reverse Charge Notes text specified on the Tax Settings page in your Recurly admin, unless custom notes were created with the original subscription.
	VatReverseChargeNotes string `json:"vat_reverse_charge_notes,omitempty"`
}

func (attr *InvoiceCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type LineItemCreate struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code. If `item_code`/`item_id` is part of the request then `currency` is optional, if the site has a single default currency. `currency` is required if `item_code`/`item_id` is present, and there are multiple currencies defined on the site. If `item_code`/`item_id` is not present `currency` is required.
	Currency string `json:"currency,omitempty"`

	// A positive or negative amount with `type=charge` will result in a positive `unit_amount`. A positive or negative amount with `type=credit` will result in a negative `unit_amount`. If `item_code`/`item_id` is present, `unit_amount` can be passed in, to override the `Item`'s `unit_amount`. If `item_code`/`item_id` is not present then `unit_amount` is required.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// This number will be multiplied by the unit amount to compute the subtotal before any discounts or taxes.
	Quantity int `json:"quantity,omitempty"`

	// Description that appears on the invoice. If `item_code`/`item_id` is part of the request then `description` must be absent.
	Description string `json:"description,omitempty"`

	// Unique code to identify an item. Avaliable when the Credit Invoices and Subscription Billing Terms features are enabled.
	ItemCode string `json:"item_code,omitempty"`

	// System-generated unique identifier for an item. Available when the Credit Invoices and Subscription Billing Terms features are enabled.
	ItemId string `json:"item_id,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Line item type. If `item_code`/`item_id` is present then `type` should not be present. If `item_code`/`item_id` is not present then `type` is required.
	Type string `json:"type,omitempty"`

	// The reason the credit was given when line item is `type=credit`. When the Credit Invoices feature is enabled, the value can be set and will default to `general`. When the Credit Invoices feature is not enabled, the value will always be `null`.
	CreditReasonCode string `json:"credit_reason_code,omitempty"`

	// Accounting Code for the `LineItem`. If `item_code`/`item_id` is part of the request then `accounting_code` must be absent.
	AccountingCode string `json:"accounting_code,omitempty"`

	// `true` exempts tax on charges, `false` applies tax on charges. If not defined, then defaults to the Plan and Site settings. This attribute does not work for credits (negative line items). Credits are always applied post-tax. Pre-tax discounts should use the Coupons feature.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's EU VAT tax feature to determine taxation rules. If you have your own AvaTax or Vertex account configured, use their tax codes to assign specific tax rules. If you are using Recurly's EU VAT feature, you can use values of `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// Optional field to track a product code or SKU for the line item. This can be used to later reporting on product purchases. For Vertex tax calculations, this field will be used as the Vertex `product` field. If `item_code`/`item_id` is part of the request then `product_code` must be absent.
	ProductCode string `json:"product_code,omitempty"`

	// Only allowed if the Gift Cards feature is enabled on your site and `type` is `credit`. Can only have a value of `external_gift_card`. Set this value in order to track gift card credits from external gift cards (like InComm). It also skips billing information requirements.
	Origin string `json:"origin,omitempty"`

	// If an end date is present, this is value indicates the beginning of a billing time range. If no end date is present it indicates billing for a specific date. Defaults to the current date-time.
	StartDate time.Time `json:"start_date,omitempty"`

	// If this date is provided, it indicates the end of a time range.
	EndDate time.Time `json:"end_date,omitempty"`
}

func (attr *LineItemCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type ShippingAddressUpdate struct {
	Params `json:"-"`

	// Shipping Address ID
	Id string `json:"id,omitempty"`

	Nickname string `json:"nickname,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	Email string `json:"email,omitempty"`

	VatNumber string `json:"vat_number,omitempty"`

	Phone string `json:"phone,omitempty"`

	Street1 string `json:"street1,omitempty"`

	Street2 string `json:"street2,omitempty"`

	City string `json:"city,omitempty"`

	// State or province.
	Region string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO code.
	Country string `json:"country,omitempty"`
}

func (attr *ShippingAddressUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type CouponCreate struct {
	Params `json:"-"`

	// The internal name for the coupon.
	Name string `json:"name,omitempty"`

	// A maximum number of redemptions for the coupon. The coupon will expire when it hits its maximum redemptions.
	MaxRedemptions int `json:"max_redemptions,omitempty"`

	// Redemptions per account is the number of times a specific account can redeem the coupon. Set redemptions per account to `1` if you want to keep customers from gaming the system and getting more than one discount from the coupon campaign.
	MaxRedemptionsPerAccount int `json:"max_redemptions_per_account,omitempty"`

	// This description will show up when a customer redeems a coupon on your Hosted Payment Pages, or if you choose to show the description on your own checkout page.
	HostedDescription string `json:"hosted_description,omitempty"`

	// Description of the coupon on the invoice.
	InvoiceDescription string `json:"invoice_description,omitempty"`

	// The date and time the coupon will expire and can no longer be redeemed. Time is always 11:59:59, the end-of-day Pacific time.
	RedeemByDate string `json:"redeem_by_date,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code string `json:"code,omitempty"`

	// The type of discount provided by the coupon (how the amount discounted is calculated)
	DiscountType string `json:"discount_type,omitempty"`

	// The percent of the price discounted by the coupon.  Required if `discount_type` is `percent`.
	DiscountPercent int `json:"discount_percent,omitempty"`

	// Description of the unit of time the coupon is for. Used with `free_trial_amount` to determine the duration of time the coupon is for.  Required if `discount_type` is `free_trial`.
	FreeTrialUnit string `json:"free_trial_unit,omitempty"`

	// Sets the duration of time the `free_trial_unit` is for. Required if `discount_type` is `free_trial`.
	FreeTrialAmount int `json:"free_trial_amount,omitempty"`

	// Fixed discount currencies by currency. Required if the coupon type is `fixed`. This parameter should contain the coupon discount values
	Currencies []CouponPricing `json:"currencies,omitempty"`

	// The coupon is valid for one-time, non-plan charges if true.
	AppliesToNonPlanCharges bool `json:"applies_to_non_plan_charges,omitempty"`

	// The coupon is valid for all plans if true. If false then `plans` and `plans_names` will list the applicable plans.
	AppliesToAllPlans bool `json:"applies_to_all_plans,omitempty"`

	// List of plan codes to which this coupon applies. See `applies_to_all_plans`
	PlanCodes []string `json:"plan_codes,omitempty"`

	// This field does not apply when the discount_type is `free_trial`. - "single_use" coupons applies to the first invoice only. - "temporal" coupons will apply to invoices for the duration determined by the `temporal_unit` and `temporal_amount` attributes. - "forever" coupons will apply to invoices forever.
	Duration string `json:"duration,omitempty"`

	// If `duration` is "temporal" than `temporal_amount` is an integer which is multiplied by `temporal_unit` to define the duration that the coupon will be applied to invoices for.
	TemporalAmount int `json:"temporal_amount,omitempty"`

	// If `duration` is "temporal" than `temporal_unit` is multiplied by `temporal_amount` to define the duration that the coupon will be applied to invoices for.
	TemporalUnit string `json:"temporal_unit,omitempty"`

	// Whether the coupon is "single_code" or "bulk". Bulk coupons will require a `unique_code_template` and will generate unique codes through the `/generate` endpoint.
	CouponType string `json:"coupon_type,omitempty"`

	// On a bulk coupon, the template from which unique coupon codes are generated. - You must start the template with your coupon_code wrapped in single quotes. - Outside of single quotes, use a 9 for a character that you want to be a random number. - Outside of single quotes, use an "x" for a character that you want to be a random letter. - Outside of single quotes, use an * for a character that you want to be a random number or letter. - Use single quotes ' ' for characters that you want to remain static. These strings can be alphanumeric and may contain a - _ or +. For example: "'abc-'****'-def'"
	UniqueCodeTemplate string `json:"unique_code_template,omitempty"`

	// Whether the discount is for all eligible charges on the account, or only a specific subscription.
	RedemptionResource string `json:"redemption_resource,omitempty"`
}

func (attr *CouponCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type CouponPricing struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// The fixed discount (in dollars) for the corresponding currency.
	Discount float64 `json:"discount,omitempty"`
}

func (attr *CouponPricing) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type CouponUpdate struct {
	Params `json:"-"`

	// The internal name for the coupon.
	Name string `json:"name,omitempty"`

	// A maximum number of redemptions for the coupon. The coupon will expire when it hits its maximum redemptions.
	MaxRedemptions int `json:"max_redemptions,omitempty"`

	// Redemptions per account is the number of times a specific account can redeem the coupon. Set redemptions per account to `1` if you want to keep customers from gaming the system and getting more than one discount from the coupon campaign.
	MaxRedemptionsPerAccount int `json:"max_redemptions_per_account,omitempty"`

	// This description will show up when a customer redeems a coupon on your Hosted Payment Pages, or if you choose to show the description on your own checkout page.
	HostedDescription string `json:"hosted_description,omitempty"`

	// Description of the coupon on the invoice.
	InvoiceDescription string `json:"invoice_description,omitempty"`

	// The date and time the coupon will expire and can no longer be redeemed. Time is always 11:59:59, the end-of-day Pacific time.
	RedeemByDate string `json:"redeem_by_date,omitempty"`
}

func (attr *CouponUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type CouponBulkCreate struct {
	Params `json:"-"`

	// The quantity of unique coupon codes to generate
	NumberOfUniqueCodes int `json:"number_of_unique_codes,omitempty"`
}

func (attr *CouponBulkCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type ItemCreate struct {
	Params `json:"-"`

	// Unique code to identify the item.
	Code string `json:"code,omitempty"`

	// This name describes your item and will appear on the invoice when it's purchased on a one time basis.
	Name string `json:"name,omitempty"`

	// Optional, description.
	Description string `json:"description,omitempty"`

	// Optional, stock keeping unit to link the item to other inventory systems.
	ExternalSku string `json:"external_sku,omitempty"`

	// Accounting code for invoice line items.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// `true` exempts tax on the item, `false` applies tax on the item.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Item Pricing
	Currencies []Pricing `json:"currencies,omitempty"`
}

func (attr *ItemCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type ItemUpdate struct {
	Params `json:"-"`

	// Unique code to identify the item.
	Code string `json:"code,omitempty"`

	// This name describes your item and will appear on the invoice when it's purchased on a one time basis.
	Name string `json:"name,omitempty"`

	// Optional, description.
	Description string `json:"description,omitempty"`

	// Optional, stock keeping unit to link the item to other inventory systems.
	ExternalSku string `json:"external_sku,omitempty"`

	// Accounting code for invoice line items.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// `true` exempts tax on the item, `false` applies tax on the item.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Item Pricing
	Currencies []Pricing `json:"currencies,omitempty"`
}

func (attr *ItemUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type InvoiceUpdatable struct {
	Params `json:"-"`

	// This identifies the PO number associated with the invoice. Not editable for credit invoices.
	PoNumber string `json:"po_number,omitempty"`

	// VAT Reverse Charge Notes are editable only if there was a VAT reverse charge applied to the invoice.
	VatReverseChargeNotes string `json:"vat_reverse_charge_notes,omitempty"`

	// Terms and conditions are an optional note field. Not editable for credit invoices.
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// Customer notes are an optional note field.
	CustomerNotes string `json:"customer_notes,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. Changing Net terms changes due_on, and the invoice could move between past due and pending.
	NetTerms int `json:"net_terms,omitempty"`

	Address InvoiceAddress `json:"address,omitempty"`
}

func (attr *InvoiceUpdatable) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type InvoiceCollect struct {
	Params `json:"-"`

	// A token generated by Recurly.js after completing a 3-D Secure device fingerprinting or authentication challenge.
	ThreeDSecureActionResultTokenId string `json:"three_d_secure_action_result_token_id,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`
}

func (attr *InvoiceCollect) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type InvoiceRefund struct {
	Params `json:"-"`

	// The type of refund. Amount and line items cannot both be specified in the request.
	Type string `json:"type,omitempty"`

	// The amount to be refunded. The amount will be split between the line items. If no amount is specified, it will default to refunding the total refundable amount on the invoice.
	Amount float64 `json:"amount,omitempty"`

	// The line items to be refunded. This is required when `type=line_items`.
	LineItems []LineItemRefund `json:"line_items,omitempty"`

	// Indicates how the invoice should be refunded when both a credit and transaction are present on the invoice: - `transaction_first` – Refunds the transaction first, then any amount is issued as credit back to the account. Default value when Credit Invoices feature is enabled. - `credit_first` – Issues credit back to the account first, then refunds any remaining amount back to the transaction. Default value when Credit Invoices feature is not enabled. - `all_credit` – Issues credit to the account for the entire amount of the refund. Only available when the Credit Invoices feature is enabled. - `all_transaction` – Refunds the entire amount back to transactions, using transactions from previous invoices if necessary. Only available when the Credit Invoices feature is enabled.
	RefundMethod string `json:"refund_method,omitempty"`

	// Used as the Customer Notes on the credit invoice.  This field can only be include when the Credit Invoices feature is enabled.
	CreditCustomerNotes string `json:"credit_customer_notes,omitempty"`

	// Indicates that the refund was settled outside of Recurly, and a manual transaction should be created to track it in Recurly.  Required when: - refunding a manually collected charge invoice, and `refund_method` is not `all_credit` - refunding a credit invoice that refunded manually collecting invoices - refunding a credit invoice for a partial amount  This field can only be included when the Credit Invoices feature is enabled.
	ExternalRefund ExternalRefund `json:"external_refund,omitempty"`
}

func (attr *InvoiceRefund) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type LineItemRefund struct {
	Params `json:"-"`

	// Line item ID
	Id string `json:"id,omitempty"`

	// Line item quantity to be refunded.
	Quantity int `json:"quantity,omitempty"`

	// Set to `true` if the line item should be prorated; set to `false` if not. This can only be used on line items that have a start and end date.
	Prorate bool `json:"prorate,omitempty"`
}

func (attr *LineItemRefund) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type ExternalRefund struct {
	Params `json:"-"`

	// Payment method used for external refund transaction.
	PaymentMethod string `json:"payment_method,omitempty"`

	// Used as the refund transactions' description.
	Description string `json:"description,omitempty"`

	// Date the external refund payment was made. Defaults to the current date-time.
	RefundedAt time.Time `json:"refunded_at,omitempty"`
}

func (attr *ExternalRefund) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type PlanCreate struct {
	Params `json:"-"`

	// Unique code to identify the plan. This is used in Hosted Payment Page URLs and in the invoice exports.
	Code string `json:"code,omitempty"`

	// This name describes your plan and will appear on the Hosted Payment Page and the subscriber's invoice.
	Name string `json:"name,omitempty"`

	// Optional description, not displayed.
	Description string `json:"description,omitempty"`

	// Accounting code for invoice line items for the plan. If no value is provided, it defaults to plan's code.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Unit for the plan's billing interval.
	IntervalUnit string `json:"interval_unit,omitempty"`

	// Length of the plan's billing interval in `interval_unit`.
	IntervalLength int `json:"interval_length,omitempty"`

	// Units for the plan's trial period.
	TrialUnit string `json:"trial_unit,omitempty"`

	// Length of plan's trial period in `trial_units`. `0` means `no trial`.
	TrialLength int `json:"trial_length,omitempty"`

	// Automatically terminate plans after a defined number of billing cycles.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// Subscriptions will automatically inherit this value once they are active. If `auto_renew` is `true`, then a subscription will automatically renew its term at renewal. If `auto_renew` is `false`, then a subscription will expire at the end of its term. `auto_renew` can be overridden on the subscription record itself.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Setup fee revenue schedule type
	SetupFeeRevenueScheduleType string `json:"setup_fee_revenue_schedule_type,omitempty"`

	// Accounting code for invoice line items for the plan's setup fee. If no value is provided, it defaults to plan's accounting code.
	SetupFeeAccountingCode string `json:"setup_fee_accounting_code,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's EU VAT tax feature to determine taxation rules. If you have your own AvaTax or Vertex account configured, use their tax codes to assign specific tax rules. If you are using Recurly's EU VAT feature, you can use values of `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// `true` exempts tax on the plan, `false` applies tax on the plan.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// Pricing
	Currencies []PlanPricing `json:"currencies,omitempty"`

	// Hosted pages settings
	HostedPages PlanHostedPages `json:"hosted_pages,omitempty"`

	// Add Ons
	AddOns []AddOnCreate `json:"add_ons,omitempty"`
}

func (attr *PlanCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type AddOnCreate struct {
	Params `json:"-"`

	// Unique code to identify an item. Avaliable when the `Catalog: Item Add-Ons` feature is enabled. If `item_id` and `item_code` are both present, `item_id` will be used.
	ItemCode string `json:"item_code,omitempty"`

	// System-generated unique identifier for an item. Available when the `Catalog: Item Add-Ons` feature is enabled. If `item_id` and `item_code` are both present, `item_id` will be used.
	ItemId string `json:"item_id,omitempty"`

	// The unique identifier for the add-on within its plan. If `item_code`/`item_id` is part of the request then `code` must be absent. If `item_code`/`item_id` is not present `code` is required.
	Code string `json:"code,omitempty"`

	// Describes your add-on and will appear in subscribers' invoices. If `item_code`/`item_id` is part of the request then `name` must be absent. If `item_code`/`item_id` is not present `name` is required.
	Name string `json:"name,omitempty"`

	// Plan ID
	PlanId string `json:"plan_id,omitempty"`

	// Accounting code for invoice line items for this add-on. If no value is provided, it defaults to add-on's code. If `item_code`/`item_id` is part of the request then `accounting_code` must be absent.
	AccountingCode string `json:"accounting_code,omitempty"`

	// When this add-on is invoiced, the line item will use this revenue schedule. If `item_code`/`item_id` is part of the request then `revenue_schedule_type` must be absent in the request as the value will be set from the item.
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Determines if the quantity field is displayed on the hosted pages for the add-on.
	DisplayQuantity bool `json:"display_quantity,omitempty"`

	// Default quantity for the hosted pages.
	DefaultQuantity int `json:"default_quantity,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's EU VAT tax feature to determine taxation rules. If you have your own AvaTax or Vertex account configured, use their tax codes to assign specific tax rules. If you are using Recurly's EU VAT feature, you can use values of `unknown`, `physical`, or `digital`. If `item_code`/`item_id` is part of the request then `tax_code` must be absent.
	TaxCode string `json:"tax_code,omitempty"`

	// If `item_code`/`item_id` is part of the request and the item has a default currency then `currencies` is optional. If the item does not have a default currency, then `currencies` is required. If `item_code`/`item_id` is not present `currencies` is required.
	Currencies []AddOnPricing `json:"currencies,omitempty"`
}

func (attr *AddOnCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type PlanUpdate struct {
	Params `json:"-"`

	// Plan ID
	Id string `json:"id,omitempty"`

	// Unique code to identify the plan. This is used in Hosted Payment Page URLs and in the invoice exports.
	Code string `json:"code,omitempty"`

	// This name describes your plan and will appear on the Hosted Payment Page and the subscriber's invoice.
	Name string `json:"name,omitempty"`

	// Optional description, not displayed.
	Description string `json:"description,omitempty"`

	// Accounting code for invoice line items for the plan. If no value is provided, it defaults to plan's code.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Units for the plan's trial period.
	TrialUnit string `json:"trial_unit,omitempty"`

	// Length of plan's trial period in `trial_units`. `0` means `no trial`.
	TrialLength int `json:"trial_length,omitempty"`

	// Automatically terminate plans after a defined number of billing cycles.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// Subscriptions will automatically inherit this value once they are active. If `auto_renew` is `true`, then a subscription will automatically renew its term at renewal. If `auto_renew` is `false`, then a subscription will expire at the end of its term. `auto_renew` can be overridden on the subscription record itself.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Setup fee revenue schedule type
	SetupFeeRevenueScheduleType string `json:"setup_fee_revenue_schedule_type,omitempty"`

	// Accounting code for invoice line items for the plan's setup fee. If no value is provided, it defaults to plan's accounting code.
	SetupFeeAccountingCode string `json:"setup_fee_accounting_code,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's EU VAT tax feature to determine taxation rules. If you have your own AvaTax or Vertex account configured, use their tax codes to assign specific tax rules. If you are using Recurly's EU VAT feature, you can use values of `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// `true` exempts tax on the plan, `false` applies tax on the plan.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// Pricing
	Currencies []PlanPricing `json:"currencies,omitempty"`

	// Hosted pages settings
	HostedPages PlanHostedPages `json:"hosted_pages,omitempty"`

	// Add Ons
	AddOns []AddOnCreate `json:"add_ons,omitempty"`
}

func (attr *PlanUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type AddOnUpdate struct {
	Params `json:"-"`

	// Add-on ID
	Id string `json:"id,omitempty"`

	// The unique identifier for the add-on within its plan. If an `Item` is associated to the `AddOn` then `code` must be absent.
	Code string `json:"code,omitempty"`

	// Describes your add-on and will appear in subscribers' invoices. If an `Item` is associated to the `AddOn` then `name` must be absent.
	Name string `json:"name,omitempty"`

	// Accounting code for invoice line items for this add-on. If no value is provided, it defaults to add-on's code. If an `Item` is associated to the `AddOn` then `accounting code` must be absent.
	AccountingCode string `json:"accounting_code,omitempty"`

	// When this add-on is invoiced, the line item will use this revenue schedule. If an `Item` is associated to the `AddOn` then `revenue_schedule_type` must be absent in the request as the value will be set from the item.
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's EU VAT tax feature to determine taxation rules. If you have your own AvaTax or Vertex account configured, use their tax codes to assign specific tax rules. If you are using Recurly's EU VAT feature, you can use values of `unknown`, `physical`, or `digital`. If an `Item` is associated to the `AddOn` then `tax code` must be absent.
	TaxCode string `json:"tax_code,omitempty"`

	// Determines if the quantity field is displayed on the hosted pages for the add-on.
	DisplayQuantity bool `json:"display_quantity,omitempty"`

	// Default quantity for the hosted pages.
	DefaultQuantity int `json:"default_quantity,omitempty"`

	// Add-on pricing
	Currencies []AddOnPricing `json:"currencies,omitempty"`
}

func (attr *AddOnUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionCreate struct {
	Params `json:"-"`

	// You must provide either a `plan_code` or `plan_id`. If both are provided the `plan_id` will be used.
	PlanCode string `json:"plan_code,omitempty"`

	// You must provide either a `plan_code` or `plan_id`. If both are provided the `plan_id` will be used.
	PlanId string `json:"plan_id,omitempty"`

	Account AccountCreate `json:"account,omitempty"`

	// Create a shipping address on the account and assign it to the subscription.
	Shipping SubscriptionShippingCreate `json:"shipping,omitempty"`

	// Collection method
	CollectionMethod string `json:"collection_method,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Override the unit amount of the subscription plan by setting this value. If not provided, the subscription will inherit the price from the subscription plan for the provided currency.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Optionally override the default quantity of 1.
	Quantity int `json:"quantity,omitempty"`

	// Add-ons
	AddOns []SubscriptionAddOnCreate `json:"add_ons,omitempty"`

	// Optional coupon code to redeem on the account and discount the subscription. Please note, the subscription request will fail if the coupon is invalid.
	CouponCode string `json:"coupon_code,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// If set, overrides the default trial behavior for the subscription. The date must be in the future.
	TrialEndsAt time.Time `json:"trial_ends_at,omitempty"`

	// If set, the subscription will begin in the future on this date. The subscription will apply the setup fee and trial period, unless the plan has no trial.
	StartsAt time.Time `json:"starts_at,omitempty"`

	// If present, this sets the date the subscription's next billing period will start (`current_period_ends_at`). This can be used to align the subscription’s billing to a specific day of the month. The initial invoice will be prorated for the period between the subscription's activation date and the billing period end date. Subsequent periods will be based off the plan interval. For a subscription with a trial period, this will change when the trial expires.
	NextBillDate time.Time `json:"next_bill_date,omitempty"`

	// The number of cycles/billing periods in a term. When `remaining_billing_cycles=0`, if `auto_renew=true` the subscription will renew and a new term will begin, otherwise the subscription will expire.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// If `auto_renew=true`, when a term completes, `total_billing_cycles` takes this value as the length of subsequent terms. Defaults to the plan's `total_billing_cycles`.
	RenewalBillingCycles int `json:"renewal_billing_cycles,omitempty"`

	// Whether the subscription renews at the end of its term.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// This will default to the Terms and Conditions text specified on the Invoice Settings page in your Recurly admin. Specify custom notes to add or override Terms and Conditions. Custom notes will stay with a subscription on all renewals.
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings. Specify custom notes to add or override Customer Notes. Custom notes will stay with a subscription on all renewals.
	CustomerNotes string `json:"customer_notes,omitempty"`

	// If there are pending credits on the account that will be invoiced during the subscription creation, these will be used as the Customer Notes on the credit invoice.
	CreditCustomerNotes string `json:"credit_customer_notes,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`
}

func (attr *SubscriptionCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionShippingCreate struct {
	Params `json:"-"`

	Address ShippingAddressCreate `json:"address,omitempty"`

	// Assign a shipping address from the account's existing shipping addresses. If `address_id` and `address` are both present, `address` will be used.
	AddressId string `json:"address_id,omitempty"`

	// The id of the shipping method used to deliver the subscription. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodId string `json:"method_id,omitempty"`

	// The code of the shipping method used to deliver the subscription. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodCode string `json:"method_code,omitempty"`

	// Assigns the subscription's shipping cost. If this is greater than zero then a `method_id` or `method_code` is required.
	Amount float64 `json:"amount,omitempty"`
}

func (attr *SubscriptionShippingCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionAddOnCreate struct {
	Params `json:"-"`

	// Add-on code
	Code string `json:"code,omitempty"`

	// Quantity
	Quantity int `json:"quantity,omitempty"`

	// Optionally, override the add-on's default unit amount.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`
}

func (attr *SubscriptionAddOnCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionUpdate struct {
	Params `json:"-"`

	// Change collection method
	CollectionMethod string `json:"collection_method,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// The remaining billing cycles in the current term.
	RemainingBillingCycles int `json:"remaining_billing_cycles,omitempty"`

	// If `auto_renew=true`, when a term completes, `total_billing_cycles` takes this value as the length of subsequent terms. Defaults to the plan's `total_billing_cycles`.
	RenewalBillingCycles int `json:"renewal_billing_cycles,omitempty"`

	// Whether the subscription renews at the end of its term.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// If present, this sets the date the subscription's next billing period will start (`current_period_ends_at`). This can be used to align the subscription’s billing to a specific day of the month. For a subscription in a trial period, this will change when the trial expires. This parameter is useful for postponement of a subscription to change its billing date without proration.
	NextBillDate time.Time `json:"next_bill_date,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Specify custom notes to add or override Terms and Conditions. Custom notes will stay with a subscription on all renewals.
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// Specify custom notes to add or override Customer Notes. Custom notes will stay with a subscription on all renewals.
	CustomerNotes string `json:"customer_notes,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	// Subscription shipping details
	Shipping SubscriptionShippingUpdate `json:"shipping,omitempty"`
}

func (attr *SubscriptionUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionShippingUpdate struct {
	Params `json:"-"`

	// Object type
	Object string `json:"object,omitempty"`

	Address ShippingAddressCreate `json:"address,omitempty"`

	// Assign a shipping address from the account's existing shipping addresses.
	AddressId string `json:"address_id,omitempty"`
}

func (attr *SubscriptionShippingUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionCancel struct {
	Params `json:"-"`

	// The timeframe parameter controls when the expiration takes place. The `bill_date` timeframe causes the subscription to expire when the subscription is scheduled to bill next. The `term_end` timeframe causes the subscription to continue to bill until the end of the subscription term, then expire.
	Timeframe string `json:"timeframe,omitempty"`
}

func (attr *SubscriptionCancel) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionPause struct {
	Params `json:"-"`

	// Number of billing cycles to pause the subscriptions.
	RemainingPauseCycles int `json:"remaining_pause_cycles,omitempty"`
}

func (attr *SubscriptionPause) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionChangeCreate struct {
	Params `json:"-"`

	// The timeframe parameter controls when the upgrade or downgrade takes place. The subscription change can occur now, when the subscription is next billed, or when the subscription term ends. Generally, if you're performing an upgrade, you will want the change to occur immediately (now). If you're performing a downgrade, you should set the timeframe to `term_end` or `bill_date` so the change takes effect at a scheduled billing date. The `renewal` timeframe option is accepted as an alias for `term_end`.
	Timeframe string `json:"timeframe,omitempty"`

	// If you want to change to a new plan, you can provide the plan's code or id. If both are provided the `plan_id` will be used.
	PlanId string `json:"plan_id,omitempty"`

	// If you want to change to a new plan, you can provide the plan's code or id. If both are provided the `plan_id` will be used.
	PlanCode string `json:"plan_code,omitempty"`

	// Optionally, sets custom pricing for the subscription, overriding the plan's default unit amount. The subscription's current currency will be used.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Optionally override the default quantity of 1.
	Quantity int `json:"quantity,omitempty"`

	// The shipping address can currently only be changed immediately, using SubscriptionUpdate.
	Shipping SubscriptionChangeShippingCreate `json:"shipping,omitempty"`

	// A list of coupon_codes to be redeemed on the subscription during the change. Only allowed if timeframe is now and you change something about the subscription that creates an invoice.
	CouponCodes []string `json:"coupon_codes,omitempty"`

	// If you provide a value for this field it will replace any existing add-ons. So, when adding or modifying an add-on, you need to include the existing subscription add-ons. Unchanged add-ons can be included just using the subscription add-on's ID: `{"id": "abc123"}`.
	AddOns []SubscriptionAddOnUpdate `json:"add_ons,omitempty"`

	// Collection method
	CollectionMethod string `json:"collection_method,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`
}

func (attr *SubscriptionChangeCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionChangeShippingCreate struct {
	Params `json:"-"`

	// The id of the shipping method used to deliver the subscription. To remove shipping set this to `null` and the `amount=0`. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodId string `json:"method_id,omitempty"`

	// The code of the shipping method used to deliver the subscription. To remove shipping set this to `null` and the `amount=0`. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodCode string `json:"method_code,omitempty"`

	// Assigns the subscription's shipping cost. If this is greater than zero then a `method_id` or `method_code` is required.
	Amount float64 `json:"amount,omitempty"`
}

func (attr *SubscriptionChangeShippingCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionAddOnUpdate struct {
	Params `json:"-"`

	// Set this to include or modify an existing subscription add-on.
	Id string `json:"id,omitempty"`

	// Add-on code
	Code string `json:"code,omitempty"`

	// Quantity
	Quantity int `json:"quantity,omitempty"`

	// Optionally, override the add-on's default unit amount.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`
}

func (attr *SubscriptionAddOnUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type PurchaseCreate struct {
	Params `json:"-"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	Account AccountPurchase `json:"account,omitempty"`

	// Must be set to manual in order to preview a purchase for an Account that does not have payment information associated with the Billing Info.
	CollectionMethod string `json:"collection_method,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	// Terms and conditions to be put on the purchase invoice.
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// Customer notes
	CustomerNotes string `json:"customer_notes,omitempty"`

	// VAT reverse charge notes for cross border European tax settlement.
	VatReverseChargeNotes string `json:"vat_reverse_charge_notes,omitempty"`

	// Notes to be put on the credit invoice resulting from credits in the purchase, if any.
	CreditCustomerNotes string `json:"credit_customer_notes,omitempty"`

	// The default payment gateway identifier to be used for the purchase transaction.  This will also be applied as the default for any subscriptions included in the purchase request.
	GatewayCode string `json:"gateway_code,omitempty"`

	Shipping ShippingPurchase `json:"shipping,omitempty"`

	// A list of one time charges or credits to be created with the purchase.
	LineItems []LineItemCreate `json:"line_items,omitempty"`

	// A list of subscriptions to be created with the purchase.
	Subscriptions []SubscriptionPurchase `json:"subscriptions,omitempty"`

	// A list of coupon_codes to be redeemed on the subscription or account during the purchase.
	CouponCodes []string `json:"coupon_codes,omitempty"`

	// A gift card redemption code to be redeemed on the purchase invoice.
	GiftCardRedemptionCode string `json:"gift_card_redemption_code,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`
}

func (attr *PurchaseCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type AccountPurchase struct {
	Params `json:"-"`

	// Optional, but if present allows an existing account to be used and updated as part of the purchase.
	Id string `json:"id,omitempty"`

	// The unique identifier of the account. This cannot be changed once the account is created.
	Code string `json:"code,omitempty"`

	Acquisition AccountAcquisitionUpdatable `json:"acquisition,omitempty"`

	// A secondary value for the account.
	Username string `json:"username,omitempty"`

	// The email address used for communicating with this customer. The customer will also use this email address to log into your hosted account management pages. This value does not need to be unique.
	Email string `json:"email,omitempty"`

	// Used to determine the language and locale of emails sent on behalf of the merchant to the customer. The list of locales is restricted to those the merchant has enabled on the site.
	PreferredLocale string `json:"preferred_locale,omitempty"`

	// Additional email address that should receive account correspondence. These should be separated only by commas. These CC emails will receive all emails that the `email` field also receives.
	CcEmails string `json:"cc_emails,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	// The VAT number of the account (to avoid having the VAT applied). This is only used for manually collected invoices.
	VatNumber string `json:"vat_number,omitempty"`

	// The tax status of the account. `true` exempts tax on the account, `false` applies tax on the account.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// The tax exemption certificate number for the account. If the merchant has an integration for the Vertex tax provider, this optional value will be sent in any tax calculation requests for the account.
	ExemptionCertificate string `json:"exemption_certificate,omitempty"`

	// The account code of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountCode string `json:"parent_account_code,omitempty"`

	// The UUID of the parent account to be associated with this account. Passing an empty value removes any existing parent association from this account. If both `parent_account_code` and `parent_account_id` are passed, the non-blank value in `parent_account_id` will be used. Only one level of parent child relationship is allowed. You cannot assign a parent account that itself has a parent account.
	ParentAccountId string `json:"parent_account_id,omitempty"`

	// An enumerable describing the billing behavior of the account, specifically whether the account is self-paying or will rely on the parent account to pay.
	BillTo string `json:"bill_to,omitempty"`

	// An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
	TransactionType string `json:"transaction_type,omitempty"`

	Address Address `json:"address,omitempty"`

	BillingInfo BillingInfoCreate `json:"billing_info,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

func (attr *AccountPurchase) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type ShippingPurchase struct {
	Params `json:"-"`

	// Assign a shipping address from the account's existing shipping addresses. If this and `shipping_address` are both present, `shipping_address` will take precedence.
	AddressId string `json:"address_id,omitempty"`

	Address ShippingAddressCreate `json:"address,omitempty"`

	// A list of shipping fees to be created as charges with the purchase.
	Fees []ShippingFeeCreate `json:"fees,omitempty"`
}

func (attr *ShippingPurchase) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type ShippingFeeCreate struct {
	Params `json:"-"`

	// The id of the shipping method used to deliver the purchase. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodId string `json:"method_id,omitempty"`

	// The code of the shipping method used to deliver the purchase. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodCode string `json:"method_code,omitempty"`

	// This is priced in the purchase's currency.
	Amount float64 `json:"amount,omitempty"`
}

func (attr *ShippingFeeCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionPurchase struct {
	Params `json:"-"`

	// Plan code
	PlanCode string `json:"plan_code,omitempty"`

	// Plan ID
	PlanId string `json:"plan_id,omitempty"`

	// Override the unit amount of the subscription plan by setting this value in cents. If not provided, the subscription will inherit the price from the subscription plan for the provided currency.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Optionally override the default quantity of 1.
	Quantity int `json:"quantity,omitempty"`

	// Add-ons
	AddOns []SubscriptionAddOnCreate `json:"add_ons,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Create a shipping address on the account and assign it to the subscription.
	Shipping SubscriptionShippingPurchase `json:"shipping,omitempty"`

	// If set, overrides the default trial behavior for the subscription. The date must be in the future.
	TrialEndsAt time.Time `json:"trial_ends_at,omitempty"`

	// If present, this sets the date the subscription's next billing period will start (`current_period_ends_at`). This can be used to align the subscription’s billing to a specific day of the month. The initial invoice will be prorated for the period between the subscription's activation date and the billing period end date. Subsequent periods will be based off the plan interval. For a subscription with a trial period, this will change when the trial expires.
	NextBillDate time.Time `json:"next_bill_date,omitempty"`

	// The number of cycles/billing periods in a term. When `remaining_billing_cycles=0`, if `auto_renew=true` the subscription will renew and a new term will begin, otherwise the subscription will expire.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// If `auto_renew=true`, when a term completes, `total_billing_cycles` takes this value as the length of subsequent terms. Defaults to the plan's `total_billing_cycles`.
	RenewalBillingCycles int `json:"renewal_billing_cycles,omitempty"`

	// Whether the subscription renews at the end of its term.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`
}

func (attr *SubscriptionPurchase) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

type SubscriptionShippingPurchase struct {
	Params `json:"-"`

	// The id of the shipping method used to deliver the subscription. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodId string `json:"method_id,omitempty"`

	// The code of the shipping method used to deliver the subscription. If `method_id` and `method_code` are both present, `method_id` will be used.
	MethodCode string `json:"method_code,omitempty"`

	// Assigns the subscription's shipping cost. If this is greater than zero then a `method_id` or `method_code` is required.
	Amount float64 `json:"amount,omitempty"`
}

func (attr *SubscriptionShippingPurchase) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
