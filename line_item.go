// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
        "time"
)

type LineItem struct {
  recurlyResponse *ResponseMetadata

  
        // Line item ID
        Id string `json:"id,omitempty"`

  
        // Object type
        Object string `json:"object,omitempty"`

  
        // The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
        Uuid string `json:"uuid,omitempty"`

  
        // Charges are positive line items that debit the account. Credits are negative line items that credit the account.
        Type string `json:"type,omitempty"`

  
        // Unique code to identify an item. Available when the Credit Invoices feature is enabled.
        ItemCode string `json:"item_code,omitempty"`

  
        // System-generated unique identifier for an item. Available when the Credit Invoices feature is enabled.
        ItemId string `json:"item_id,omitempty"`

  
        // Optional Stock Keeping Unit assigned to an item. Available when the Credit Invoices feature is enabled.
        ExternalSku string `json:"external_sku,omitempty"`

  
        // Revenue schedule type
        RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

  
        // Pending line items are charges or credits on an account that have not been applied to an invoice yet. Invoiced line items will always have an `invoice_id` value.
        State string `json:"state,omitempty"`

  
        // Category to describe the role of a line item on a legacy invoice:
 // - "charges" refers to charges being billed for on this invoice.
 // - "credits" refers to refund or proration credits. This portion of the invoice can be considered a credit memo.
 // - "applied_credits" refers to previous credits applied to this invoice. See their original_line_item_id to determine where the credit first originated.
 // - "carryforwards" can be ignored. They exist to consume any remaining credit balance. A new credit with the same amount will be created and placed back on the account.
        LegacyCategory string `json:"legacy_category,omitempty"`

  
        // Account mini details
        Account AccountMini `json:"account,omitempty"`

  
        // The UUID of the account responsible for originating the line item.
        BillForAccountId string `json:"bill_for_account_id,omitempty"`

  
        // If the line item is a charge or credit for a subscription, this is its ID.
        SubscriptionId string `json:"subscription_id,omitempty"`

  
        // If the line item is a charge or credit for a plan or add-on, this is the plan's ID.
        PlanId string `json:"plan_id,omitempty"`

  
        // If the line item is a charge or credit for a plan or add-on, this is the plan's code.
        PlanCode string `json:"plan_code,omitempty"`

  
        // If the line item is a charge or credit for an add-on this is its ID.
        AddOnId string `json:"add_on_id,omitempty"`

  
        // If the line item is a charge or credit for an add-on, this is its code.
        AddOnCode string `json:"add_on_code,omitempty"`

  
        // Once the line item has been invoiced this will be the invoice's ID.
        InvoiceId string `json:"invoice_id,omitempty"`

  
        // Once the line item has been invoiced this will be the invoice's number. If VAT taxation and the Country Invoice Sequencing feature are enabled, invoices will have country-specific invoice numbers for invoices billed to EU countries (ex: FR1001). Non-EU invoices will continue to use the site-level invoice number sequence.
        InvoiceNumber string `json:"invoice_number,omitempty"`

  
        // Will only have a value if the line item is a credit created from a previous credit, or if the credit was created from a charge refund.
        PreviousLineItemId string `json:"previous_line_item_id,omitempty"`

  
        // The invoice where the credit originated. Will only have a value if the line item is a credit created from a previous credit, or if the credit was created from a charge refund.
        OriginalLineItemInvoiceId string `json:"original_line_item_invoice_id,omitempty"`

  
        // A credit created from an original charge will have the value of the charge's origin.
        Origin string `json:"origin,omitempty"`

  
        // Internal accounting code to help you reconcile your revenue to the correct ledger. Line items created as part of a subscription invoice will use the plan or add-on's accounting code, otherwise the value will only be present if you define an accounting code when creating the line item.
        AccountingCode string `json:"accounting_code,omitempty"`

  
        // For plan-related line items this will be the plan's code, for add-on related line items it will be the add-on's code. For item-related line items it will be the item's `external_sku`.
        ProductCode string `json:"product_code,omitempty"`

  
        // The reason the credit was given when line item is `type=credit`.
        CreditReasonCode string `json:"credit_reason_code,omitempty"`

  
        // 3-letter ISO 4217 currency code.
        Currency string `json:"currency,omitempty"`

  
        // `(quantity * unit_amount) - (discount + tax)`
        Amount float64 `json:"amount,omitempty"`

  
        // Description that appears on the invoice. For subscription related items this will be filled in automatically.
        Description string `json:"description,omitempty"`

  
        // This number will be multiplied by the unit amount to compute the subtotal before any discounts or taxes.
        Quantity int `json:"quantity,omitempty"`

  
        // Positive amount for a charge, negative amount for a credit.
        UnitAmount float64 `json:"unit_amount,omitempty"`

  
        // Positive amount for a charge, negative amount for a credit.
        UnitAmountDecimal string `json:"unit_amount_decimal,omitempty"`

  
        // Determines whether or not tax is included in the unit amount. The Tax Inclusive Pricing feature (separate from the Mixed Tax Pricing feature) must be enabled to utilize this flag.
        TaxInclusive bool `json:"tax_inclusive,omitempty"`

  
        // `quantity * unit_amount`
        Subtotal float64 `json:"subtotal,omitempty"`

  
        // The discount applied to the line item.
        Discount float64 `json:"discount,omitempty"`

  
        // The tax amount for the line item.
        Tax float64 `json:"tax,omitempty"`

  
        // `true` if the line item is taxable, `false` if it is not.
        Taxable bool `json:"taxable,omitempty"`

  
        // `true` exempts tax on charges, `false` applies tax on charges. If not defined, then defaults to the Plan and Site settings. This attribute does not work for credits (negative line items). Credits are always applied post-tax. Pre-tax discounts should use the Coupons feature.
        TaxExempt bool `json:"tax_exempt,omitempty"`

  
        // Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the line item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
        AvalaraTransactionType int `json:"avalara_transaction_type,omitempty"`

  
        // Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the line item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
        AvalaraServiceType int `json:"avalara_service_type,omitempty"`

  
        // Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
        TaxCode string `json:"tax_code,omitempty"`

  
        // Tax info
        TaxInfo TaxInfo `json:"tax_info,omitempty"`

  
        // When a line item has been prorated, this is the rate of the proration. Proration rates were made available for line items created after March 30, 2017. For line items created prior to that date, the proration rate will be `null`, even if the line item was prorated.
        ProrationRate float64 `json:"proration_rate,omitempty"`

  
        // Refund?
        Refund bool `json:"refund,omitempty"`

  
        // For refund charges, the quantity being refunded. For non-refund charges, the total quantity refunded (possibly over multiple refunds).
        RefundedQuantity int `json:"refunded_quantity,omitempty"`

  
        // The amount of credit from this line item that was applied to the invoice.
        CreditApplied float64 `json:"credit_applied,omitempty"`

  
        ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`

  
        // If an end date is present, this is value indicates the beginning of a billing time range. If no end date is present it indicates billing for a specific date.
        StartDate time.Time `json:"start_date,omitempty"`

  
        // If this date is provided, it indicates the end of a time range.
        EndDate time.Time `json:"end_date,omitempty"`

  
        // When the line item was created.
        CreatedAt time.Time `json:"created_at,omitempty"`

  
        // When the line item was last changed.
        UpdatedAt time.Time `json:"updated_at,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *LineItem) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *LineItem) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type lineItemList struct {
	ListMetadata
  Data []LineItem `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *lineItemList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *lineItemList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// LineItemList allows you to paginate LineItem objects
type LineItemList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []LineItem
}

func NewLineItemList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *LineItemList {
  return &LineItemList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type LineItemLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []LineItem
  HasMore() bool
  Next() string
}

func (list  *LineItemList) HasMore() bool {
    return list.hasMore
}

func (list  *LineItemList) Next() string {
    return list.nextPagePath
}

func (list *LineItemList) Data() []LineItem {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *LineItemList) FetchWithContext(ctx context.Context) error {
	resources := &lineItemList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
  // copy over properties from the response
  list.nextPagePath = resources.Next
	list.hasMore = resources.HasMore
	list.data = resources.Data
  return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *LineItemList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *LineItemList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &lineItemList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *LineItemList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
