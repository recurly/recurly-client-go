// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type AccountInvoiceTemplate struct {
	recurlyResponse *ResponseMetadata

	// Unique ID to identify the invoice template.
	Id string `json:"id,omitempty"`

	// Template name
	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountInvoiceTemplate) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountInvoiceTemplate) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountInvoiceTemplateList struct {
	ListMetadata
	Data            []AccountInvoiceTemplate `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountInvoiceTemplateList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountInvoiceTemplateList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountInvoiceTemplateList allows you to paginate AccountInvoiceTemplate objects
type AccountInvoiceTemplateList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []AccountInvoiceTemplate
}

func NewAccountInvoiceTemplateList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AccountInvoiceTemplateList {
	return &AccountInvoiceTemplateList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type AccountInvoiceTemplateLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []AccountInvoiceTemplate
	HasMore() bool
	Next() string
}

func (list *AccountInvoiceTemplateList) HasMore() bool {
	return list.hasMore
}

func (list *AccountInvoiceTemplateList) Next() string {
	return list.nextPagePath
}

func (list *AccountInvoiceTemplateList) Data() []AccountInvoiceTemplate {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountInvoiceTemplateList) FetchWithContext(ctx context.Context) error {
	resources := &accountInvoiceTemplateList{}
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
func (list *AccountInvoiceTemplateList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AccountInvoiceTemplateList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &accountInvoiceTemplateList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountInvoiceTemplateList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
