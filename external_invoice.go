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

type ExternalInvoice struct {
	recurlyResponse *ResponseMetadata

	// System-generated unique identifier for an external invoice ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// Subscription from an external resource such as Apple App Store or Google Play Store.
	ExternalSubscription ExternalSubscription `json:"external_subscription,omitempty"`

	// An identifier which associates the external invoice to a corresponding object in an external platform.
	ExternalId string `json:"external_id,omitempty"`

	State string `json:"state,omitempty"`

	// Total
	Total string `json:"total,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	LineItems []ExternalCharge `json:"line_items,omitempty"`

	// When the invoice was created in the external platform.
	PurchasedAt time.Time `json:"purchased_at,omitempty"`

	// When the external invoice was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the external invoice was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalInvoice) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalInvoice) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalInvoiceList struct {
	ListMetadata
	Data            []ExternalInvoice `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalInvoiceList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalInvoiceList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalInvoiceList allows you to paginate ExternalInvoice objects
type ExternalInvoiceList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalInvoice
}

func NewExternalInvoiceList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalInvoiceList {
	return &ExternalInvoiceList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalInvoiceLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalInvoice
	HasMore() bool
	Next() string
}

func (list *ExternalInvoiceList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalInvoiceList) Next() string {
	return list.nextPagePath
}

func (list *ExternalInvoiceList) Data() []ExternalInvoice {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalInvoiceList) FetchWithContext(ctx context.Context) error {
	resources := &externalInvoiceList{}
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
func (list *ExternalInvoiceList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalInvoiceList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalInvoiceList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalInvoiceList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
