// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type BusinessEntityMini struct {
	recurlyResponse *ResponseMetadata

	// Business entity ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The entity code of the business entity.
	Code string `json:"code,omitempty"`

	// This name describes your business entity and will appear on the invoice.
	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *BusinessEntityMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *BusinessEntityMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type businessEntityMiniList struct {
	ListMetadata
	Data            []BusinessEntityMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *businessEntityMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *businessEntityMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// BusinessEntityMiniList allows you to paginate BusinessEntityMini objects
type BusinessEntityMiniList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []BusinessEntityMini
}

func NewBusinessEntityMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *BusinessEntityMiniList {
	return &BusinessEntityMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type BusinessEntityMiniLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []BusinessEntityMini
	HasMore() bool
	Next() string
}

func (list *BusinessEntityMiniList) HasMore() bool {
	return list.hasMore
}

func (list *BusinessEntityMiniList) Next() string {
	return list.nextPagePath
}

func (list *BusinessEntityMiniList) Data() []BusinessEntityMini {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *BusinessEntityMiniList) FetchWithContext(ctx context.Context) error {
	resources := &businessEntityMiniList{}
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
func (list *BusinessEntityMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *BusinessEntityMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &businessEntityMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *BusinessEntityMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
