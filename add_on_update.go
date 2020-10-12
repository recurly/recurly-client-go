package recurly

import ()

type AddOnUpdate struct {
	Params `json:"-"`

	// Add-on ID
	Id *string `json:"id,omitempty"`

	// The unique identifier for the add-on within its plan. If an `Item` is associated to the `AddOn` then `code` must be absent.
	Code *string `json:"code,omitempty"`

	// Describes your add-on and will appear in subscribers' invoices. If an `Item` is associated to the `AddOn` then `name` must be absent.
	Name *string `json:"name,omitempty"`

	// The percentage taken of the monetary amount of usage tracked. This can be up to 4 decimal places. A value between 0.0 and 100.0. Required if `add_on_type` is usage and `usage_type` is percentage. Must be omitted otherwise. `usage_percentage` does not support tiers.
	UsagePercentage *float64 `json:"usage_percentage,omitempty"`

	// System-generated unique identifier for a measured unit to be associated with the add-on. Either `measured_unit_id` or `measured_unit_name` are required when `add_on_type` is `usage`. If `measured_unit_id` and `measured_unit_name` are both present, `measured_unit_id` will be used.
	MeasuredUnitId *string `json:"measured_unit_id,omitempty"`

	// Name of a measured unit to be associated with the add-on. Either `measured_unit_id` or `measured_unit_name` are required when `add_on_type` is `usage`. If `measured_unit_id` and `measured_unit_name` are both present, `measured_unit_id` will be used.
	MeasuredUnitName *string `json:"measured_unit_name,omitempty"`

	// Accounting code for invoice line items for this add-on. If no value is provided, it defaults to add-on's code. If an `Item` is associated to the `AddOn` then `accounting code` must be absent.
	AccountingCode *string `json:"accounting_code,omitempty"`

	// When this add-on is invoiced, the line item will use this revenue schedule. If an `Item` is associated to the `AddOn` then `revenue_schedule_type` must be absent in the request as the value will be set from the item.
	RevenueScheduleType *string `json:"revenue_schedule_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the add-on is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types. If an `Item` is associated to the `AddOn`, then the `avalara_transaction_type` must be absent.
	AvalaraTransactionType *int `json:"avalara_transaction_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the add-on is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types. If an `Item` is associated to the `AddOn`, then the `avalara_service_type` must be absent.
	AvalaraServiceType *int `json:"avalara_service_type,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's EU VAT tax feature to determine taxation rules. If you have your own AvaTax or Vertex account configured, use their tax codes to assign specific tax rules. If you are using Recurly's EU VAT feature, you can use values of `unknown`, `physical`, or `digital`. If an `Item` is associated to the `AddOn` then `tax code` must be absent.
	TaxCode *string `json:"tax_code,omitempty"`

	// Determines if the quantity field is displayed on the hosted pages for the add-on.
	DisplayQuantity *bool `json:"display_quantity,omitempty"`

	// Default quantity for the hosted pages.
	DefaultQuantity *int `json:"default_quantity,omitempty"`

	// Whether the add-on is optional for the customer to include in their purchase on the hosted payment page. If false, the add-on will be included when a subscription is created through the Recurly UI. However, the add-on will not be included when a subscription is created through the API.
	Optional *bool `json:"optional,omitempty"`

	// If the add-on's `tier_type` is `tiered`, `volume` or `stairstep`,
	// then `currencies` must be absent.
	Currencies []AddOnPricingCreate `json:"currencies,omitempty"`

	// If the tier_type is `flat`, then `tiers` must be absent. The `tiers` object
	// must include one to many tiers with `ending_quantity` and `unit_amount` for
	// the desired `currencies`. There must be one tier with an `ending_quantity` of
	// 999999999 which is the default if not provided.
	Tiers []TierCreate `json:"tiers,omitempty"`
}

func (attr *AddOnUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
