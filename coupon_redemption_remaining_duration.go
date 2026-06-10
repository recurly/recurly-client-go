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

type CouponRedemptionRemainingDuration struct {
	recurlyResponse *ResponseMetadata

	// The coupon's duration type. `temporal` includes an `expires_at` timestamp. `forever` and `single_use` have no additional fields.
	Type string `json:"type,omitempty"`

	// Present when `type` is `temporal`. The datetime after which this redemption will no longer apply.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponRedemptionRemainingDuration) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponRedemptionRemainingDuration) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponRedemptionRemainingDurationList struct {
	ListMetadata
	Data            []CouponRedemptionRemainingDuration `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponRedemptionRemainingDurationList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponRedemptionRemainingDurationList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponRedemptionRemainingDurationList allows you to paginate CouponRedemptionRemainingDuration objects
type CouponRedemptionRemainingDurationList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []CouponRedemptionRemainingDuration
}

func NewCouponRedemptionRemainingDurationList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CouponRedemptionRemainingDurationList {
	return &CouponRedemptionRemainingDurationList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type CouponRedemptionRemainingDurationLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []CouponRedemptionRemainingDuration
	HasMore() bool
	Next() string
}

func (list *CouponRedemptionRemainingDurationList) HasMore() bool {
	return list.hasMore
}

func (list *CouponRedemptionRemainingDurationList) Next() string {
	return list.nextPagePath
}

func (list *CouponRedemptionRemainingDurationList) Data() []CouponRedemptionRemainingDuration {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponRedemptionRemainingDurationList) FetchWithContext(ctx context.Context) error {
	resources := &couponRedemptionRemainingDurationList{}
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
func (list *CouponRedemptionRemainingDurationList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CouponRedemptionRemainingDurationList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &couponRedemptionRemainingDurationList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CouponRedemptionRemainingDurationList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
