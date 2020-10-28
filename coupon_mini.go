// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"time"
)

type CouponMini struct {
	recurlyResponse *ResponseMetadata

	// Coupon ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code string `json:"code,omitempty"`

	// The internal name for the coupon.
	Name string `json:"name,omitempty"`

	// Indicates if the coupon is redeemable, and if it is not, why.
	State string `json:"state,omitempty"`

	// Details of the discount a coupon applies. Will contain a `type`
	// property and one of the following properties: `percent`, `fixed`, `trial`.
	Discount CouponDiscount `json:"discount,omitempty"`

	// Whether the coupon is "single_code" or "bulk". Bulk coupons will require a `unique_code_template` and will generate unique codes through the `/generate` endpoint.
	CouponType string `json:"coupon_type,omitempty"`

	// The date and time the coupon was expired early or reached its `max_redemptions`.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponMiniList struct {
	ListMetadata
	Data            []CouponMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponMiniList allows you to paginate CouponMini objects
type CouponMiniList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []CouponMini
}

func NewCouponMiniList(client HttpCaller, nextPagePath string) *CouponMiniList {
	return &CouponMiniList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponMiniList) Fetch() error {
	resources := &couponMiniList{}
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
func (list *CouponMiniList) Count() (*int64, error) {
	resources := &couponMiniList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
