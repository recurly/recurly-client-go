// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PaymentGatewayReferences struct {
	recurlyResponse *ResponseMetadata

	// Reference value used when the external token was created. If a Stripe gateway or Ebanx gateway is used, this value will need to be accompanied by its reference_type.
	Token string `json:"token,omitempty"`

	// The type of reference token. Required if token is passed in for Stripe Gateway or Ebanx UPI.
	ReferenceType string `json:"reference_type,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PaymentGatewayReferences) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PaymentGatewayReferences) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type paymentGatewayReferencesList struct {
	ListMetadata
	Data            []PaymentGatewayReferences `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *paymentGatewayReferencesList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *paymentGatewayReferencesList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PaymentGatewayReferencesList allows you to paginate PaymentGatewayReferences objects
type PaymentGatewayReferencesList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PaymentGatewayReferences
}

func NewPaymentGatewayReferencesList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PaymentGatewayReferencesList {
	return &PaymentGatewayReferencesList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PaymentGatewayReferencesLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PaymentGatewayReferences
	HasMore() bool
	Next() string
}

func (list *PaymentGatewayReferencesList) HasMore() bool {
	return list.hasMore
}

func (list *PaymentGatewayReferencesList) Next() string {
	return list.nextPagePath
}

func (list *PaymentGatewayReferencesList) Data() []PaymentGatewayReferences {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PaymentGatewayReferencesList) FetchWithContext(ctx context.Context) error {
	resources := &paymentGatewayReferencesList{}
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
func (list *PaymentGatewayReferencesList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PaymentGatewayReferencesList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &paymentGatewayReferencesList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PaymentGatewayReferencesList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
