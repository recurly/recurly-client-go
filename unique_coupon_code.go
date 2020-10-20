package recurly

import (
	"net/http"
	"time"
)

type UniqueCouponCode struct {
	recurlyResponse *ResponseMetadata

	// Unique Coupon Code ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code string `json:"code,omitempty"`

	// Indicates if the unique coupon code is redeemable or why not.
	State string `json:"state,omitempty"`

	// The Coupon ID of the parent Bulk Coupon
	BulkCouponId string `json:"bulk_coupon_id,omitempty"`

	// The Coupon code of the parent Bulk Coupon
	BulkCouponCode string `json:"bulk_coupon_code,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date and time the unique coupon code was redeemed.
	RedeemedAt time.Time `json:"redeemed_at,omitempty"`

	// The date and time the coupon was expired early or reached its `max_redemptions`.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *UniqueCouponCode) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *UniqueCouponCode) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type uniqueCouponCodeList struct {
	ListMetadata
	Data            []UniqueCouponCode `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// UniqueCouponCodeList allows you to paginate UniqueCouponCode objects
type UniqueCouponCodeList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []UniqueCouponCode
}

func NewUniqueCouponCodeList(client HttpCaller, nextPagePath string) *UniqueCouponCodeList {
	return &UniqueCouponCodeList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *UniqueCouponCodeList) Fetch() error {
	resources := &uniqueCouponCodeList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *UniqueCouponCodeList) Count() (*int64, error) {
	resources := &uniqueCouponCodeList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
