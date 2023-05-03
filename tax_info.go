// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type TaxInfo struct {
	recurlyResponse *ResponseMetadata

	// Provides the tax type as "vat" for EU VAT, "usst" for U.S. Sales Tax, or the 2 letter country code for country level tax types like Canada, Australia, New Zealand, Israel, and all non-EU European countries. Not present when Avalara for Communications is enabled.
	Type string `json:"type,omitempty"`

	// Provides the tax region applied on an invoice. For U.S. Sales Tax, this will be the 2 letter state code. For EU VAT this will be the 2 letter country code. For all country level tax types, this will display the regional tax, like VAT, GST, or PST. Not present when Avalara for Communications is enabled.
	Region string `json:"region,omitempty"`

	// The combined tax rate. Not present when Avalara for Communications is enabled.
	Rate float64 `json:"rate,omitempty"`

	// Provides additional tax details for Communications taxes when Avalara for Communications is enabled or Canadian Sales Tax when there is tax applied at both the country and province levels. This will only be populated for the Invoice response when fetching a single invoice and not for the InvoiceList or LineItemList. Only populated for a single LineItem fetch when Avalara for Communications is enabled.
	TaxDetails []TaxDetail `json:"tax_details,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TaxInfo) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TaxInfo) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type taxInfoList struct {
	ListMetadata
	Data            []TaxInfo `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *taxInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *taxInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TaxInfoList allows you to paginate TaxInfo objects
type TaxInfoList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []TaxInfo
}

func NewTaxInfoList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TaxInfoList {
	return &TaxInfoList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *TaxInfoList) FetchWithContext(ctx context.Context) error {
	resources := &taxInfoList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *TaxInfoList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TaxInfoList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &taxInfoList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TaxInfoList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
