// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
)

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

  
        // Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
        AvalaraTransactionType *int `json:"avalara_transaction_type,omitempty"`

  
        // Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
        AvalaraServiceType *int `json:"avalara_service_type,omitempty"`

  
        // Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
        TaxCode *string `json:"tax_code,omitempty"`

  
        // `true` exempts tax on the item, `false` applies tax on the item.
        TaxExempt *bool `json:"tax_exempt,omitempty"`

  
        // The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
        CustomFields []CustomFieldCreate `json:"custom_fields,omitempty"`

  
        // Item Pricing
        Currencies []PricingCreate `json:"currencies,omitempty"`

  
}


