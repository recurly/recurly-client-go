// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type LineItemCreate struct {

	// 3-letter ISO 4217 currency code. If `item_code`/`item_id` is part of the request then `currency` is optional, if the site has a single default currency. `currency` is required if `item_code`/`item_id` is present, and there are multiple currencies defined on the site. If `item_code`/`item_id` is not present `currency` is required.
	Currency *string `json:"currency,omitempty"`

	// A positive or negative amount with `type=charge` will result in a positive `unit_amount`.
	// A positive or negative amount with `type=credit` will result in a negative `unit_amount`.
	// If `item_code`/`item_id` is present, `unit_amount` can be passed in, to override the
	// `Item`'s `unit_amount`. If `item_code`/`item_id` is not present then `unit_amount` is required.
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// This number will be multiplied by the unit amount to compute the subtotal before any discounts or taxes.
	Quantity *int `json:"quantity,omitempty"`

	// Description that appears on the invoice. If `item_code`/`item_id` is part of the request then `description` must be absent.
	Description *string `json:"description,omitempty"`

	// Unique code to identify an item. Available when the Credit Invoices and Subscription Billing Terms features are enabled.
	ItemCode *string `json:"item_code,omitempty"`

	// System-generated unique identifier for an item. Available when the Credit Invoices and Subscription Billing Terms features are enabled.
	ItemId *string `json:"item_id,omitempty"`

	// Revenue schedule type
	RevenueScheduleType *string `json:"revenue_schedule_type,omitempty"`

	// Line item type. If `item_code`/`item_id` is present then `type` should not be present. If `item_code`/`item_id` is not present then `type` is required.
	Type *string `json:"type,omitempty"`

	// The reason the credit was given when line item is `type=credit`. When the Credit Invoices feature is enabled, the value can be set and will default to `general`. When the Credit Invoices feature is not enabled, the value will always be `null`.
	CreditReasonCode *string `json:"credit_reason_code,omitempty"`

	// Accounting Code for the `LineItem`. If `item_code`/`item_id` is part of the request then `accounting_code` must be absent.
	AccountingCode *string `json:"accounting_code,omitempty"`

	// `true` exempts tax on charges, `false` applies tax on charges. If not defined, then defaults to the Plan and Site settings. This attribute does not work for credits (negative line items). Credits are always applied post-tax. Pre-tax discounts should use the Coupons feature.
	TaxExempt *bool `json:"tax_exempt,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the line item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types. If an `Item` is associated to the `LineItem`, then the `avalara_transaction_type` must be absent.
	AvalaraTransactionType *int `json:"avalara_transaction_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the line item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types. If an `Item` is associated to the `LineItem`, then the `avalara_service_type` must be absent.
	AvalaraServiceType *int `json:"avalara_service_type,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's EU VAT tax feature to determine taxation rules. If you have your own AvaTax or Vertex account configured, use their tax codes to assign specific tax rules. If you are using Recurly's EU VAT feature, you can use values of `unknown`, `physical`, or `digital`.
	TaxCode *string `json:"tax_code,omitempty"`

	// Optional field to track a product code or SKU for the line item. This can be used to later reporting on product purchases. For Vertex tax calculations, this field will be used as the Vertex `product` field. If `item_code`/`item_id` is part of the request then `product_code` must be absent.
	ProductCode *string `json:"product_code,omitempty"`

	// Origin `external_gift_card` is allowed if the Gift Cards feature is enabled on your site and `type` is `credit`. Set this value in order to track gift card credits from external gift cards (like InComm). It also skips billing information requirements.  Origin `prepayment` is only allowed if `type` is `charge` and `tax_exempt` is left blank or set to true.  This origin creates a charge and opposite credit on the account to be used for future invoices.
	Origin *string `json:"origin,omitempty"`

	// If an end date is present, this is value indicates the beginning of a billing time range. If no end date is present it indicates billing for a specific date. Defaults to the current date-time.
	StartDate *time.Time `json:"start_date,omitempty"`

	// If this date is provided, it indicates the end of a time range.
	EndDate *time.Time `json:"end_date,omitempty"`
}
