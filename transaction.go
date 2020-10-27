// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type Transaction struct {
	recurlyResponse *ResponseMetadata

	// Transaction ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
	Uuid string `json:"uuid,omitempty"`

	// If this transaction is a refund (`type=refund`), this will be the ID of the original transaction on the invoice being refunded.
	OriginalTransactionId string `json:"original_transaction_id,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// Invoice mini details
	Invoice InvoiceMini `json:"invoice,omitempty"`

	// Invoice mini details
	VoidedByInvoice InvoiceMini `json:"voided_by_invoice,omitempty"`

	// If the transaction is charging or refunding for one or more subscriptions, these are their IDs.
	SubscriptionIds []string `json:"subscription_ids,omitempty"`

	// - `authorization` – verifies billing information and places a hold on money in the customer's account.
	// - `capture` – captures funds held by an authorization and completes a purchase.
	// - `purchase` – combines the authorization and capture in one transaction.
	// - `refund` – returns all or a portion of the money collected in a previous transaction to the customer.
	// - `verify` – a $0 or $1 transaction used to verify billing information which is immediately voided.
	Type string `json:"type,omitempty"`

	// Describes how the transaction was triggered.
	Origin string `json:"origin,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total transaction amount sent to the payment gateway.
	Amount float64 `json:"amount,omitempty"`

	// The current transaction status. Note that the status may change, e.g. a `pending` transaction may become `declined` or `success` may later become `void`.
	Status string `json:"status,omitempty"`

	// Did this transaction complete successfully?
	Success bool `json:"success,omitempty"`

	// Indicates if part or all of this transaction was refunded.
	Refunded bool `json:"refunded,omitempty"`

	BillingAddress Address `json:"billing_address,omitempty"`

	// The method by which the payment was collected.
	CollectionMethod string `json:"collection_method,omitempty"`

	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`

	// IP address provided when the billing information was collected:
	// - When the customer enters billing information into the Recurly.js or Hosted Payment Pages, Recurly records the IP address.
	// - When the merchant enters billing information using the API, the merchant may provide an IP address.
	// - When the merchant enters billing information using the UI, no IP address is recorded.
	IpAddressV4 string `json:"ip_address_v4,omitempty"`

	// IP address's country
	IpAddressCountry string `json:"ip_address_country,omitempty"`

	// Status code
	StatusCode string `json:"status_code,omitempty"`

	// For declined (`success=false`) transactions, the message displayed to the merchant.
	StatusMessage string `json:"status_message,omitempty"`

	// For declined (`success=false`) transactions, the message displayed to the customer.
	CustomerMessage string `json:"customer_message,omitempty"`

	// Language code for the message
	CustomerMessageLocale string `json:"customer_message_locale,omitempty"`

	PaymentGateway TransactionPaymentGateway `json:"payment_gateway,omitempty"`

	// Transaction message from the payment gateway.
	GatewayMessage string `json:"gateway_message,omitempty"`

	// Transaction reference number from the payment gateway.
	GatewayReference string `json:"gateway_reference,omitempty"`

	// Transaction approval code from the payment gateway.
	GatewayApprovalCode string `json:"gateway_approval_code,omitempty"`

	// For declined transactions (`success=false`), this field lists the gateway error code.
	GatewayResponseCode string `json:"gateway_response_code,omitempty"`

	// Time, in seconds, for gateway to process the transaction.
	GatewayResponseTime float64 `json:"gateway_response_time,omitempty"`

	// The values in this field will vary from gateway to gateway.
	GatewayResponseValues map[string]interface{} `json:"gateway_response_values,omitempty"`

	// When processed, result from checking the CVV/CVC value on the transaction.
	CvvCheck string `json:"cvv_check,omitempty"`

	// When processed, result from checking the overall AVS on the transaction.
	AvsCheck string `json:"avs_check,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Voided at
	VoidedAt time.Time `json:"voided_at,omitempty"`

	// Collected at, or if not collected yet, the time the transaction was created.
	CollectedAt time.Time `json:"collected_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Transaction) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Transaction) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type transactionList struct {
	ListMetadata
	Data            []Transaction `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *transactionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *transactionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TransactionList allows you to paginate Transaction objects
type TransactionList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []Transaction
}

func NewTransactionList(client HttpCaller, nextPagePath string) *TransactionList {
	return &TransactionList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *TransactionList) Fetch() error {
	resources := &transactionList{}
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
func (list *TransactionList) Count() (*int64, error) {
	resources := &transactionList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
