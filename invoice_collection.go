// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []InvoiceCollection
}

func NewInvoiceCollectionList(client HttpCaller, nextPagePath string) *InvoiceCollectionList {
	return &InvoiceCollectionList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceCollectionList) Fetch() error {
	resources := &invoiceCollectionList{}
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
func (list *InvoiceCollectionList) Count() (*int64, error) {
	resources := &invoiceCollectionList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
