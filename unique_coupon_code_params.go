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

type UniqueCouponCodeParams struct {
	recurlyResponse *ResponseMetadata

	// The number of UniqueCouponCodes that will be generated
	Limit int `json:"limit,omitempty"`

	// Sort order to list newly generated UniqueCouponCodes (should always be `asc`)
	Order string `json:"order,omitempty"`

	// Sort field to list newly generated UniqueCouponCodes (should always be `created_at`)
	Sort string `json:"sort,omitempty"`

	// The date-time to be included when listing UniqueCouponCodes
	BeginTime time.Time `json:"begin_time,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *UniqueCouponCodeParams) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *UniqueCouponCodeParams) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type uniqueCouponCodeParamsList struct {
	ListMetadata
	Data            []UniqueCouponCodeParams `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeParamsList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeParamsList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// UniqueCouponCodeParamsList allows you to paginate UniqueCouponCodeParams objects
type UniqueCouponCodeParamsList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []UniqueCouponCodeParams
}

func NewUniqueCouponCodeParamsList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *UniqueCouponCodeParamsList {
	return &UniqueCouponCodeParamsList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type UniqueCouponCodeParamsLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []UniqueCouponCodeParams
	HasMore() bool
	Next() string
}

func (list *UniqueCouponCodeParamsList) HasMore() bool {
	return list.hasMore
}

func (list *UniqueCouponCodeParamsList) Next() string {
	return list.nextPagePath
}

func (list *UniqueCouponCodeParamsList) Data() []UniqueCouponCodeParams {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *UniqueCouponCodeParamsList) FetchWithContext(ctx context.Context) error {
	resources := &uniqueCouponCodeParamsList{}
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
func (list *UniqueCouponCodeParamsList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *UniqueCouponCodeParamsList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &uniqueCouponCodeParamsList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *UniqueCouponCodeParamsList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
