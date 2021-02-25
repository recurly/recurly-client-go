// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type CouponDiscount struct {
	recurlyResponse *ResponseMetadata

	Type string `json:"type,omitempty"`

	// This is only present when `type=percent`.
	Percent int `json:"percent,omitempty"`

	// This is only present when `type=fixed`.
	Currencies []CouponDiscountPricing `json:"currencies,omitempty"`

	// This is only present when `type=free_trial`.
	Trial CouponDiscountTrial `json:"trial,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponDiscountList struct {
	ListMetadata
	Data            []CouponDiscount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponDiscountList allows you to paginate CouponDiscount objects
type CouponDiscountList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []CouponDiscount
}

func NewCouponDiscountList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CouponDiscountList {
	return &CouponDiscountList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type CouponDiscountLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []CouponDiscount
	HasMore() bool
	Next() string
}

func (list *CouponDiscountList) HasMore() bool {
	return list.hasMore
}

func (list *CouponDiscountList) Next() string {
	return list.nextPagePath
}

func (list *CouponDiscountList) Data() []CouponDiscount {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountList) FetchWithContext(ctx context.Context) error {
	resources := &couponDiscountList{}
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
func (list *CouponDiscountList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CouponDiscountList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &couponDiscountList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CouponDiscountList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
