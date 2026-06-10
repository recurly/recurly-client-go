// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type RecoveryLineItemCreate struct {

	// The tax amount for the line item.
	Tax *float64 `json:"tax,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields *[]CustomFieldCreate `json:"custom_fields,omitempty"`

	// The Harmonized System (HS) code is an internationally standardized system of names and numbers to classify traded products. The HS code, sometimes called Commodity Code, is used by customs authorities around the world to identify products when assessing duties and taxes. The HS code may also be referred to as the tariff code or customs code. Values should contain only digits and decimals.
	HarmonizedSystemCode *string `json:"harmonized_system_code,omitempty"`

	// Optional field to track a product code or SKU for the line item. This can be used to later reporting on product purchases.
	ProductCode *string `json:"product_code,omitempty"`

	// This number will be multiplied by the unit amount to compute the subtotal before any discounts or taxes.
	Quantity *int `json:"quantity,omitempty"`

	// Description that appears on the invoice.
	Description *string `json:"description,omitempty"`

	// A positive or negative amount will result in a positive `unit_amount`.
	UnitAmount *float64 `json:"unit_amount,omitempty"`
}
