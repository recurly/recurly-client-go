// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type CouponDiscountTrial struct {
	recurlyResponse *ResponseMetadata

	// Temporal unit of the free trial
	Unit string `json:"unit,omitempty"`

	// Trial length measured in the units specified by the sibling `unit` property
	Length int `json:"length,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscountTrial) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscountTrial) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponDiscountTrialList struct {
	ListMetadata
	Data            []CouponDiscountTrial `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountTrialList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountTrialList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponDiscountTrialList allows you to paginate CouponDiscountTrial objects
type CouponDiscountTrialList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []CouponDiscountTrial
}

type CouponDiscountTrialLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []CouponDiscountTrial
	HasMore() bool
	Next() string
}

func NewCouponDiscountTrialList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CouponDiscountTrialList {
	return &CouponDiscountTrialList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *CouponDiscountTrialList) HasMore() bool {
	return list.hasMore
}

func (list *CouponDiscountTrialList) Next() string {
	return list.nextPagePath
}

func (list *CouponDiscountTrialList) Data() []CouponDiscountTrial {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountTrialList) FetchWithContext(ctx context.Context) error {
	resources := &couponDiscountTrialList{}
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
func (list *CouponDiscountTrialList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CouponDiscountTrialList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &couponDiscountTrialList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CouponDiscountTrialList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
