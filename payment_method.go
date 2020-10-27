// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type PaymentMethod struct {
	recurlyResponse *ResponseMetadata

	Object string `json:"object,omitempty"`

	// Visa, MasterCard, American Express, Discover, JCB, etc.
	CardType string `json:"card_type,omitempty"`

	// Credit card number's first six digits.
	FirstSix string `json:"first_six,omitempty"`

	// Credit card number's last four digits. Will refer to bank account if payment method is ACH.
	LastFour string `json:"last_four,omitempty"`

	// The IBAN bank account's last two digits.
	LastTwo string `json:"last_two,omitempty"`

	// Expiration month.
	ExpMonth int `json:"exp_month,omitempty"`

	// Expiration year.
	ExpYear int `json:"exp_year,omitempty"`

	// A token used in place of a credit card in order to perform transactions.
	GatewayToken string `json:"gateway_token,omitempty"`

	// An identifier for a specific payment gateway.
	GatewayCode string `json:"gateway_code,omitempty"`

	// Billing Agreement identifier. Only present for Amazon or Paypal payment methods.
	BillingAgreementId string `json:"billing_agreement_id,omitempty"`

	// The name associated with the bank account.
	NameOnAccount string `json:"name_on_account,omitempty"`

	// The bank account type. Only present for ACH payment methods.
	AccountType string `json:"account_type,omitempty"`

	// The bank account's routing number. Only present for ACH payment methods.
	RoutingNumber string `json:"routing_number,omitempty"`

	// The bank name of this routing number.
	RoutingNumberBank string `json:"routing_number_bank,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PaymentMethod) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PaymentMethod) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type paymentMethodList struct {
	ListMetadata
	Data            []PaymentMethod `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *paymentMethodList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *paymentMethodList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PaymentMethodList allows you to paginate PaymentMethod objects
type PaymentMethodList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []PaymentMethod
}

func NewPaymentMethodList(client HttpCaller, nextPagePath string) *PaymentMethodList {
	return &PaymentMethodList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *PaymentMethodList) Fetch() error {
	resources := &paymentMethodList{}
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
func (list *PaymentMethodList) Count() (*int64, error) {
	resources := &paymentMethodList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
