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

type ExternalCharge struct {
	recurlyResponse *ResponseMetadata

	// System-generated unique identifier for an external charge ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Unit Amount
	UnitAmount float64 `json:"unit_amount,omitempty"`

	Quantity int `json:"quantity,omitempty"`

	Description string `json:"description,omitempty"`

	// External Product Reference details
	ExternalProductReference ExternalProductReferenceMini `json:"external_product_reference,omitempty"`

	// When the external charge was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the external charge was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalCharge) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalCharge) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalChargeList struct {
	ListMetadata
	Data            []ExternalCharge `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalChargeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalChargeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalChargeList allows you to paginate ExternalCharge objects
type ExternalChargeList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalCharge
}

func NewExternalChargeList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalChargeList {
	return &ExternalChargeList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalChargeLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalCharge
	HasMore() bool
	Next() string
}

func (list *ExternalChargeList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalChargeList) Next() string {
	return list.nextPagePath
}

func (list *ExternalChargeList) Data() []ExternalCharge {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalChargeList) FetchWithContext(ctx context.Context) error {
	resources := &externalChargeList{}
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
func (list *ExternalChargeList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalChargeList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalChargeList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalChargeList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
