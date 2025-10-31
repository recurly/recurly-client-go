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

	// Must be sent for one-time transactions in order to provide context on which entity is submitting the transaction to ensure proper fraud checks are observed, such as 3DS. If the customer is in session, send `customer`. If this is a merchant initiated one-time transaction, send `merchant`.
	Initiator string `json:"initiator,omitempty"`

	// Invoice mini details
	Invoice InvoiceMini `json:"invoice,omitempty"`

	// This conditional parameter is useful for merchants in specific industries who need to submit one-time Merchant Initiated transactions in specific cases.
	// Not all gateways support these methods, but will support a generic one-time Merchant Initiated transaction.
	// Only use this if the initiator value is "merchant". Otherwise, it will be ignored.
	//   - Incremental: Send `incremental` with an additional purchase if the original authorization amount is not sufficient to cover the costs of your service or product. For example, if the customer adds goods or services or there are additional expenses.
	//   - No Show: Send `no_show` if you charge customers a fee due to an agreed-upon cancellation policy in your industry.
	//   - Resubmission: Send `resubmission` if you need to attempt collection on a declined transaction. You may also use the force collection behavior which has the same effect.
	//   - Service Extension: Send `service_extension` if you are in a service industry and the customer has increased/extended their service in some way. For example: adding a day onto a car rental agreement.
	//   - Split Shipment: Send `split_shipment` if you sell physical product and need to split up a shipment into multiple transactions when the customer is no longer in session.
	//   - Top Up: Send `top_up` if you process one-time transactions based on a pre-arranged agreement with your customer where there is a pre-arranged account balance that needs maintaining. For example, if the customer has agreed to maintain an account balance of 30.00 and their current balance is 20.00, the MIT amount would be at least 10.00 to meet that 30.00 threshold.
	MerchantReasonCode string `json:"merchant_reason_code,omitempty"`

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

	// Indicates if the transaction was completed using a backup payment
	BackupPaymentMethodUsed bool `json:"backup_payment_method_used,omitempty"`

	// Indicates if part or all of this transaction was refunded.
	Refunded bool `json:"refunded,omitempty"`

	BillingAddress AddressWithName `json:"billing_address,omitempty"`

	// The method by which the payment was collected.
	CollectionMethod string `json:"collection_method,omitempty"`

	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`

	// IP address provided when the billing information was collected:
	// - When the customer enters billing information into the Recurly.js or Hosted Payment Pages, Recurly records the IP address.
	// - When the merchant enters billing information using the API, the merchant may provide an IP address.
	// - When the merchant enters billing information using the UI, no IP address is recorded.
	IpAddressV4 string `json:"ip_address_v4,omitempty"`

	// Origin IP address country, 2-letter ISO 3166-1 alpha-2 code, if known by Recurly.
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
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// Voided at
	VoidedAt *time.Time `json:"voided_at,omitempty"`

	// Collected at, or if not collected yet, the time the transaction was created.
	CollectedAt *time.Time `json:"collected_at,omitempty"`

	// Action result params to be used in Recurly-JS to complete a payment when using asynchronous payment methods, e.g., Boleto, iDEAL and Sofort.
	ActionResult map[string]interface{} `json:"action_result,omitempty"`

	// VAT number for the customer on this transaction. If the customer's Billing Info country is BR or AR, then this will be their Tax Identifier. For all other countries this will come from the VAT Number field in the Billing Info.
	VatNumber string `json:"vat_number,omitempty"`

	// Fraud information
	FraudInfo TransactionFraudInfo `json:"fraud_info,omitempty"`

	// Next action values are used for any required customer follow-up action. Currently, this is supported for Ebanx when using Pix Automatico.
	NextAction TransactionNextAction `json:"next_action,omitempty"`
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
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Transaction
}

func NewTransactionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TransactionList {
	return &TransactionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type TransactionLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Transaction
	HasMore() bool
	Next() string
}

func (list *TransactionList) HasMore() bool {
	return list.hasMore
}

func (list *TransactionList) Next() string {
	return list.nextPagePath
}

func (list *TransactionList) Data() []Transaction {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *TransactionList) FetchWithContext(ctx context.Context) error {
	resources := &transactionList{}
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
func (list *TransactionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TransactionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &transactionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TransactionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
