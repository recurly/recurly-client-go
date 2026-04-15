// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type UniqueCouponCodeGenerationResponse struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	// An array containing the newly generated unique coupon codes.
	UniqueCouponCodes []UniqueCouponCode `json:"unique_coupon_codes,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *UniqueCouponCodeGenerationResponse) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *UniqueCouponCodeGenerationResponse) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type uniqueCouponCodeGenerationResponseList struct {
	ListMetadata
	Data            []UniqueCouponCodeGenerationResponse `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeGenerationResponseList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeGenerationResponseList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// UniqueCouponCodeGenerationResponseList allows you to paginate UniqueCouponCodeGenerationResponse objects
type UniqueCouponCodeGenerationResponseList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []UniqueCouponCodeGenerationResponse
}

func NewUniqueCouponCodeGenerationResponseList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *UniqueCouponCodeGenerationResponseList {
	return &UniqueCouponCodeGenerationResponseList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type UniqueCouponCodeGenerationResponseLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []UniqueCouponCodeGenerationResponse
	HasMore() bool
	Next() string
}

func (list *UniqueCouponCodeGenerationResponseList) HasMore() bool {
	return list.hasMore
}

func (list *UniqueCouponCodeGenerationResponseList) Next() string {
	return list.nextPagePath
}

func (list *UniqueCouponCodeGenerationResponseList) Data() []UniqueCouponCodeGenerationResponse {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *UniqueCouponCodeGenerationResponseList) FetchWithContext(ctx context.Context) error {
	resources := &uniqueCouponCodeGenerationResponseList{}
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
func (list *UniqueCouponCodeGenerationResponseList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *UniqueCouponCodeGenerationResponseList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &uniqueCouponCodeGenerationResponseList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *UniqueCouponCodeGenerationResponseList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
