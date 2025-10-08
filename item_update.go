// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type ItemUpdate struct {

	// Unique code to identify the item.
	Code *string `json:"code,omitempty"`

	// This name describes your item and will appear on the invoice when it's purchased on a one time basis.
	Name *string `json:"name,omitempty"`

	// Optional, description.
	Description *string `json:"description,omitempty"`

	// Optional, stock keeping unit to link the item to other inventory systems.
	ExternalSku *string `json:"external_sku,omitempty"`

	// Accounting code for invoice line items.
	AccountingCode *string `json:"accounting_code,omitempty"`

	// Revenue schedule type
	RevenueScheduleType *string `json:"revenue_schedule_type,omitempty"`

	// The ID of a performance obligation. Performance obligations are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	PerformanceObligationId *string `json:"performance_obligation_id,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	LiabilityGlAccountId *string `json:"liability_gl_account_id,omitempty"`

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	RevenueGlAccountId *string `json:"revenue_gl_account_id,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraTransactionType *int `json:"avalara_transaction_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraServiceType *int `json:"avalara_service_type,omitempty"`

	// Optional field used by Avalara, Vertex, and Recurly's In-the-Box tax solution to determine taxation rules. You can pass in specific tax codes using any of these tax integrations. For Recurly's In-the-Box tax offering you can also choose to instead use simple values of `unknown`, `physical`, or `digital` tax codes.
	TaxCode *string `json:"tax_code,omitempty"`

	// The Harmonized System (HS) code is an internationally standardized system of names and numbers to classify traded products. The HS code, sometimes called Commodity Code, is used by customs authorities around the world to identify products when assessing duties and taxes. The HS code may also be referred to as the tariff code or customs code. Values should contain only digits and decimals.
	HarmonizedSystemCode *string `json:"harmonized_system_code,omitempty"`

	// `true` exempts tax on the item, `false` applies tax on the item.
	TaxExempt *bool `json:"tax_exempt,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields *[]CustomFieldCreate `json:"custom_fields,omitempty"`

	// Item Pricing
	Currencies *[]PricingCreate `json:"currencies,omitempty"`
}
