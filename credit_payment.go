// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type CreditPayment struct {
	recurlyResponse *ResponseMetadata

	// Credit Payment ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
	Uuid string `json:"uuid,omitempty"`

	// The action for which the credit was created.
	Action string `json:"action,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// Invoice mini details
	AppliedToInvoice InvoiceMini `json:"applied_to_invoice,omitempty"`

	// Invoice mini details
	OriginalInvoice InvoiceMini `json:"original_invoice,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total credit payment amount applied to the charge invoice.
	Amount float64 `json:"amount,omitempty"`

	// For credit payments with action `refund`, this is the credit payment that was refunded.
	OriginalCreditPaymentId string `json:"original_credit_payment_id,omitempty"`

	RefundTransaction Transaction `json:"refund_transaction,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Voided at
	VoidedAt time.Time `json:"voided_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CreditPayment) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CreditPayment) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type creditPaymentList struct {
	ListMetadata
	Data            []CreditPayment `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *creditPaymentList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *creditPaymentList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CreditPaymentList allows you to paginate CreditPayment objects
type CreditPaymentList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []CreditPayment
}

func NewCreditPaymentList(client HttpCaller, nextPagePath string) *CreditPaymentList {
	return &CreditPaymentList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CreditPaymentList) Fetch() error {
	resources := &creditPaymentList{}
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
func (list *CreditPaymentList) Count() (*int64, error) {
	resources := &creditPaymentList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
