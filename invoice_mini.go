// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
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
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []InvoiceMini
}

func NewInvoiceMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *InvoiceMiniList {
	return &InvoiceMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type InvoiceMiniLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []InvoiceMini
	HasMore() bool
	Next() string
}

func (list *InvoiceMiniList) HasMore() bool {
	return list.hasMore
}

func (list *InvoiceMiniList) Next() string {
	return list.nextPagePath
}

func (list *InvoiceMiniList) Data() []InvoiceMini {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceMiniList) FetchWithContext(ctx context.Context) error {
	resources := &invoiceMiniList{}
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
func (list *InvoiceMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &invoiceMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
