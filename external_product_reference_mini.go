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

type ExternalProductReferenceMini struct {
	recurlyResponse *ResponseMetadata

	// System-generated unique identifier for an external product ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// object
	Object string `json:"object,omitempty"`

	// A code which associates the external product to a corresponding object or resource in an external platform like the Apple App Store or Google Play Store.
	ReferenceCode string `json:"reference_code,omitempty"`

	// Source connection platform.
	ExternalConnectionType string `json:"external_connection_type,omitempty"`

	// When the external product was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the external product was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalProductReferenceMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalProductReferenceMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalProductReferenceMiniList struct {
	ListMetadata
	Data            []ExternalProductReferenceMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalProductReferenceMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalProductReferenceMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalProductReferenceMiniList allows you to paginate ExternalProductReferenceMini objects
type ExternalProductReferenceMiniList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalProductReferenceMini
}

func NewExternalProductReferenceMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalProductReferenceMiniList {
	return &ExternalProductReferenceMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalProductReferenceMiniLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalProductReferenceMini
	HasMore() bool
	Next() string
}

func (list *ExternalProductReferenceMiniList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalProductReferenceMiniList) Next() string {
	return list.nextPagePath
}

func (list *ExternalProductReferenceMiniList) Data() []ExternalProductReferenceMini {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalProductReferenceMiniList) FetchWithContext(ctx context.Context) error {
	resources := &externalProductReferenceMiniList{}
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
func (list *ExternalProductReferenceMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductReferenceMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalProductReferenceMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductReferenceMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
