// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type TaxDetail struct {
	recurlyResponse *ResponseMetadata

	// Provides the tax type for the region or type of Comminications tax when Avalara for Communications is enabled. For Canadian Sales Tax, this will be GST, HST, QST or PST.
	Type string `json:"type,omitempty"`

	// Provides the tax region applied on an invoice. For Canadian Sales Tax, this will be either the 2 letter province code or country code. Not present when Avalara for Communications is enabled.
	Region string `json:"region,omitempty"`

	// Provides the tax rate for the region.
	Rate float64 `json:"rate,omitempty"`

	// The total tax applied for this tax type.
	Tax float64 `json:"tax,omitempty"`

	// Provides the name of the Communications tax applied. Present only when Avalara for Communications is enabled.
	Name string `json:"name,omitempty"`

	// Provides the jurisdiction level for the Communications tax applied. Example values include city, state and federal. Present only when Avalara for Communications is enabled.
	Level string `json:"level,omitempty"`

	// Whether or not the line item is taxable. Only populated for a single LineItem fetch when Avalara for Communications is enabled.
	Billable bool `json:"billable,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TaxDetail) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TaxDetail) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type taxDetailList struct {
	ListMetadata
	Data            []TaxDetail `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *taxDetailList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *taxDetailList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TaxDetailList allows you to paginate TaxDetail objects
type TaxDetailList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []TaxDetail
}

func NewTaxDetailList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TaxDetailList {
	return &TaxDetailList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type TaxDetailLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []TaxDetail
	HasMore() bool
	Next() string
}

func (list *TaxDetailList) HasMore() bool {
	return list.hasMore
}

func (list *TaxDetailList) Next() string {
	return list.nextPagePath
}

func (list *TaxDetailList) Data() []TaxDetail {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *TaxDetailList) FetchWithContext(ctx context.Context) error {
	resources := &taxDetailList{}
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
func (list *TaxDetailList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TaxDetailList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &taxDetailList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TaxDetailList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
