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

type InvoiceTemplate struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Invoice template code.
	Code string `json:"code,omitempty"`

	// Invoice template name.
	Name string `json:"name,omitempty"`

	// Invoice template description.
	Description string `json:"description,omitempty"`

	// When the invoice template was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the invoice template was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *InvoiceTemplate) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *InvoiceTemplate) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceTemplateList struct {
	ListMetadata
	Data            []InvoiceTemplate `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceTemplateList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceTemplateList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceTemplateList allows you to paginate InvoiceTemplate objects
type InvoiceTemplateList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []InvoiceTemplate
}

func NewInvoiceTemplateList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *InvoiceTemplateList {
	return &InvoiceTemplateList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type InvoiceTemplateLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []InvoiceTemplate
	HasMore() bool
	Next() string
}

func (list *InvoiceTemplateList) HasMore() bool {
	return list.hasMore
}

func (list *InvoiceTemplateList) Next() string {
	return list.nextPagePath
}

func (list *InvoiceTemplateList) Data() []InvoiceTemplate {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceTemplateList) FetchWithContext(ctx context.Context) error {
	resources := &invoiceTemplateList{}
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
func (list *InvoiceTemplateList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceTemplateList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &invoiceTemplateList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceTemplateList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
