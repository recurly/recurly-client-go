// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type InvoiceMini struct {
	recurlyResponse *ResponseMetadata

	// Invoice ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Invoice number
	Number string `json:"number,omitempty"`

	// Invoice type
	Type string `json:"type,omitempty"`

	// Invoice state
	State string `json:"state,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *InvoiceMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *InvoiceMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceMiniList struct {
	ListMetadata
	Data            []InvoiceMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceMiniList allows you to paginate InvoiceMini objects
type InvoiceMiniList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []InvoiceMini
}

func NewInvoiceMiniList(client HttpCaller, nextPagePath string) *InvoiceMiniList {
	return &InvoiceMiniList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceMiniList) Fetch() error {
	resources := &invoiceMiniList{}
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
func (list *InvoiceMiniList) Count() (*int64, error) {
	resources := &invoiceMiniList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
