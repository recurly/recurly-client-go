// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
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
  Data []CouponMini `json:"data"`
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
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []CouponMini
}

func NewCouponMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CouponMiniList {
  return &CouponMiniList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type CouponMiniLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []CouponMini
  HasMore() bool
  Next() string
}

func (list  *CouponMiniList) HasMore() bool {
    return list.hasMore
}

func (list  *CouponMiniList) Next() string {
    return list.nextPagePath
}

func (list *CouponMiniList) Data() []CouponMini {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *CouponMiniList) FetchWithContext(ctx context.Context) error {
	resources := &couponMiniList{}
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
func (list *CouponMiniList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CouponMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &couponMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CouponMiniList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
