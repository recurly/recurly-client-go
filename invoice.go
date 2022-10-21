// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
	"time"
)

type Invoice struct {
	recurlyResponse *ResponseMetadata

	// Invoice ID
	Id string `json:"id,omitempty"`

	// Invoice UUID
	Uuid string `json:"uuid,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Invoices are either charge, credit, or legacy invoices.
	Type string `json:"type,omitempty"`

	// The event that created the invoice.
	Origin string `json:"origin,omitempty"`

	// Invoice state
	State string `json:"state,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// The `billing_info_id` is the value that represents a specific billing info for an end customer. When `billing_info_id` is used to assign billing info to the subscription, all future billing events for the subscription will bill to the specified billing info. `billing_info_id` can ONLY be used for sites utilizing the Wallet feature.
	BillingInfoId string `json:"billing_info_id,omitempty"`

	// If the invoice is charging or refunding for one or more subscriptions, these are their IDs.
	SubscriptionIds []string `json:"subscription_ids,omitempty"`

	// On refund invoices, this value will exist and show the invoice ID of the purchase invoice the refund was created from.
	PreviousInvoiceId string `json:"previous_invoice_id,omitempty"`

	// If VAT taxation and the Country Invoice Sequencing feature are enabled, invoices will have country-specific invoice numbers for invoices billed to EU countries (ex: FR1001). Non-EU invoices will continue to use the site-level invoice number sequence.
	Number string `json:"number,omitempty"`

	// An automatic invoice means a corresponding transaction is run using the account's billing information at the same time the invoice is created. Manual invoices are created without a corresponding transaction. The merchant must enter a manual payment transaction or have the customer pay the invoice with an automatic method, like credit card, PayPal, Amazon, or ACH bank payment.
	CollectionMethod string `json:"collection_method,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	Address InvoiceAddress `json:"address,omitempty"`

	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total discounts applied to this invoice.
	Discount float64 `json:"discount,omitempty"`

	// The summation of charges and credits, before discounts and taxes.
	Subtotal float64 `json:"subtotal,omitempty"`

	// The total tax on this invoice.
	Tax float64 `json:"tax,omitempty"`

	// The final total on this invoice. The summation of invoice charges, discounts, credits, and tax.
	Total float64 `json:"total,omitempty"`

	// The refundable amount on a charge invoice. It will be null for all other invoices.
	RefundableAmount float64 `json:"refundable_amount,omitempty"`

	// The total amount of successful payments transaction on this invoice.
	Paid float64 `json:"paid,omitempty"`

	// The outstanding balance remaining on this invoice.
	Balance float64 `json:"balance,omitempty"`

	// Tax info
	TaxInfo TaxInfo `json:"tax_info,omitempty"`

	// VAT registration number for the customer on this invoice. This will come from the VAT Number field in the Billing Info or the Account Info depending on your tax settings and the invoice collection method.
	VatNumber string `json:"vat_number,omitempty"`

	// VAT Reverse Charge Notes only appear if you have EU VAT enabled or are using your own Avalara AvaTax account and the customer is in the EU, has a VAT number, and is in a different country than your own. This will default to the VAT Reverse Charge Notes text specified on the Tax Settings page in your Recurly admin, unless custom notes were created with the original subscription.
	VatReverseChargeNotes string `json:"vat_reverse_charge_notes,omitempty"`

	// This will default to the Terms and Conditions text specified on the Invoice Settings page in your Recurly admin. Specify custom notes to add or override Terms and Conditions.
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings. Specify custom notes to add or override Customer Notes.
	CustomerNotes string `json:"customer_notes,omitempty"`

	// Line Items
	LineItems []LineItem `json:"line_items,omitempty"`

	// Identifies if the invoice has more line items than are returned in `line_items`. If `has_more_line_items` is `true`, then a request needs to be made to the `list_invoice_line_items` endpoint.
	HasMoreLineItems bool `json:"has_more_line_items,omitempty"`

	// Transactions
	Transactions []Transaction `json:"transactions,omitempty"`

	// Credit payments
	CreditPayments []CreditPayment `json:"credit_payments,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Date invoice is due. This is the date the net terms are reached.
	DueAt time.Time `json:"due_at,omitempty"`

	// Date invoice was marked paid or failed.
	ClosedAt time.Time `json:"closed_at,omitempty"`

	// Unique ID to identify the dunning campaign used when dunning the invoice. For sites without multiple dunning campaigns enabled, this will always be the default dunning campaign.
	DunningCampaignId string `json:"dunning_campaign_id,omitempty"`

	// Number of times the event was sent.
	DunningEventsSent int `json:"dunning_events_sent,omitempty"`

	// Last communication attempt.
	FinalDunningEvent bool `json:"final_dunning_event,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Invoice) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Invoice) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceList struct {
	ListMetadata
	Data            []Invoice `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceList allows you to paginate Invoice objects
type InvoiceList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Invoice
}

func NewInvoiceList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *InvoiceList {
	return &InvoiceList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type InvoiceLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Invoice
	HasMore() bool
	Next() string
}

func (list *InvoiceList) HasMore() bool {
	return list.hasMore
}

func (list *InvoiceList) Next() string {
	return list.nextPagePath
}

func (list *InvoiceList) Data() []Invoice {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceList) FetchWithContext(ctx context.Context) error {
	resources := &invoiceList{}
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
func (list *InvoiceList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &invoiceList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
