// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ReferenceOnlyCurrencyConversion struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// The subtotal converted to the currency.
	SubtotalInCents float64 `json:"subtotal_in_cents,omitempty"`

	// The tax converted to the currency.
	TaxInCents float64 `json:"tax_in_cents,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ReferenceOnlyCurrencyConversion) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ReferenceOnlyCurrencyConversion) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type referenceOnlyCurrencyConversionList struct {
	ListMetadata
	Data            []ReferenceOnlyCurrencyConversion `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *referenceOnlyCurrencyConversionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *referenceOnlyCurrencyConversionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ReferenceOnlyCurrencyConversionList allows you to paginate ReferenceOnlyCurrencyConversion objects
type ReferenceOnlyCurrencyConversionList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ReferenceOnlyCurrencyConversion
}

func NewReferenceOnlyCurrencyConversionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ReferenceOnlyCurrencyConversionList {
	return &ReferenceOnlyCurrencyConversionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ReferenceOnlyCurrencyConversionLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ReferenceOnlyCurrencyConversion
	HasMore() bool
	Next() string
}

func (list *ReferenceOnlyCurrencyConversionList) HasMore() bool {
	return list.hasMore
}

func (list *ReferenceOnlyCurrencyConversionList) Next() string {
	return list.nextPagePath
}

func (list *ReferenceOnlyCurrencyConversionList) Data() []ReferenceOnlyCurrencyConversion {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ReferenceOnlyCurrencyConversionList) FetchWithContext(ctx context.Context) error {
	resources := &referenceOnlyCurrencyConversionList{}
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
func (list *ReferenceOnlyCurrencyConversionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ReferenceOnlyCurrencyConversionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &referenceOnlyCurrencyConversionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ReferenceOnlyCurrencyConversionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
