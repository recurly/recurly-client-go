// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type LineItemDiscount struct {
	recurlyResponse *ResponseMetadata

	// Will always be `line_item_discount`.
	Object string `json:"object,omitempty"`

	// The ID of the coupon that generated this discount.
	CouponId string `json:"coupon_id,omitempty"`

	// The ID of the coupon redemption that generated this discount.
	CouponRedemptionId string `json:"coupon_redemption_id,omitempty"`

	// The order in which this discount was applied when multiple coupons were redeemed.
	OrderApplied int `json:"order_applied,omitempty"`

	// The amount discounted on this line item by this coupon redemption.
	DiscountAmount float64 `json:"discount_amount,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *LineItemDiscount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *LineItemDiscount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type lineItemDiscountList struct {
	ListMetadata
	Data            []LineItemDiscount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *lineItemDiscountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *lineItemDiscountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// LineItemDiscountList allows you to paginate LineItemDiscount objects
type LineItemDiscountList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []LineItemDiscount
}

func NewLineItemDiscountList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *LineItemDiscountList {
	return &LineItemDiscountList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type LineItemDiscountLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []LineItemDiscount
	HasMore() bool
	Next() string
}

func (list *LineItemDiscountList) HasMore() bool {
	return list.hasMore
}

func (list *LineItemDiscountList) Next() string {
	return list.nextPagePath
}

func (list *LineItemDiscountList) Data() []LineItemDiscount {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *LineItemDiscountList) FetchWithContext(ctx context.Context) error {
	resources := &lineItemDiscountList{}
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
func (list *LineItemDiscountList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *LineItemDiscountList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &lineItemDiscountList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *LineItemDiscountList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
