// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type InvoiceCollection struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	ChargeInvoice Invoice `json:"charge_invoice,omitempty"`

	// Credit invoices
	CreditInvoices []Invoice `json:"credit_invoices,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *InvoiceCollection) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *InvoiceCollection) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceCollectionList struct {
	ListMetadata
	Data            []InvoiceCollection `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceCollectionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceCollectionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceCollectionList allows you to paginate InvoiceCollection objects
type InvoiceCollectionList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []InvoiceCollection
}

func NewInvoiceCollectionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *InvoiceCollectionList {
	return &InvoiceCollectionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceCollectionList) FetchWithContext(ctx context.Context) error {
	resources := &invoiceCollectionList{}
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
func (list *InvoiceCollectionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceCollectionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &invoiceCollectionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceCollectionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
