// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type CouponDiscountPricing struct {
  recurlyResponse *ResponseMetadata

  
        // 3-letter ISO 4217 currency code.
        Currency string `json:"currency,omitempty"`

  
        // Value of the fixed discount that this coupon applies.
        Amount float64 `json:"amount,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscountPricing) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscountPricing) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type couponDiscountPricingList struct {
	ListMetadata
  Data []CouponDiscountPricing `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// CouponDiscountPricingList allows you to paginate CouponDiscountPricing objects
type CouponDiscountPricingList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []CouponDiscountPricing
}

func NewCouponDiscountPricingList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CouponDiscountPricingList {
  return &CouponDiscountPricingList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type CouponDiscountPricingLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []CouponDiscountPricing
  HasMore() bool
  Next() string
}

func (list  *CouponDiscountPricingList) HasMore() bool {
    return list.hasMore
}

func (list  *CouponDiscountPricingList) Next() string {
    return list.nextPagePath
}

func (list *CouponDiscountPricingList) Data() []CouponDiscountPricing {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountPricingList) FetchWithContext(ctx context.Context) error {
	resources := &couponDiscountPricingList{}
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
func (list *CouponDiscountPricingList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CouponDiscountPricingList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &couponDiscountPricingList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CouponDiscountPricingList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
