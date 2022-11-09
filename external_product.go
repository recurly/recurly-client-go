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

type ExternalProduct struct {
	recurlyResponse *ResponseMetadata

	// System-generated unique identifier for an external product ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Name to identify the external product in Recurly.
	Name string `json:"name,omitempty"`

	// Just the important parts.
	Plan PlanMini `json:"plan,omitempty"`

	// When the external product was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the external product was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// List of external product references of the external product.
	ExternalProductReferences []ExternalProductReferenceMini `json:"external_product_references,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalProduct) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalProduct) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalProductList struct {
	ListMetadata
	Data            []ExternalProduct `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalProductList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalProductList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalProductList allows you to paginate ExternalProduct objects
type ExternalProductList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalProduct
}

func NewExternalProductList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalProductList {
	return &ExternalProductList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalProductLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalProduct
	HasMore() bool
	Next() string
}

func (list *ExternalProductList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalProductList) Next() string {
	return list.nextPagePath
}

func (list *ExternalProductList) Data() []ExternalProduct {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalProductList) FetchWithContext(ctx context.Context) error {
	resources := &externalProductList{}
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
func (list *ExternalProductList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalProductList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
