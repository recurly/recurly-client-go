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

type CouponRedemptionMini struct {
	recurlyResponse *ResponseMetadata

	// Coupon Redemption ID
	Id string `json:"id,omitempty"`

	// Will always be `coupon`.
	Object string `json:"object,omitempty"`

	Coupon CouponMini `json:"coupon,omitempty"`

	// Invoice state
	State string `json:"state,omitempty"`

	// The amount that was discounted upon the application of the coupon, formatted with the currency.
	Discounted float64 `json:"discounted,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponRedemptionMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponRedemptionMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponRedemptionMiniList struct {
	ListMetadata
	Data            []CouponRedemptionMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponRedemptionMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponRedemptionMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponRedemptionMiniList allows you to paginate CouponRedemptionMini objects
type CouponRedemptionMiniList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []CouponRedemptionMini
}

type CouponRedemptionMiniLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []CouponRedemptionMini
	HasMore() bool
	Next() string
}

func NewCouponRedemptionMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CouponRedemptionMiniList {
	return &CouponRedemptionMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *CouponRedemptionMiniList) HasMore() bool {
	return list.hasMore
}

func (list *CouponRedemptionMiniList) Next() string {
	return list.nextPagePath
}

func (list *CouponRedemptionMiniList) Data() []CouponRedemptionMini {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponRedemptionMiniList) FetchWithContext(ctx context.Context) error {
	resources := &couponRedemptionMiniList{}
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
func (list *CouponRedemptionMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CouponRedemptionMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &couponRedemptionMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CouponRedemptionMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
