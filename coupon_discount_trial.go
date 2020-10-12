package recurly

import (
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []CouponDiscountTrial
}

func NewCouponDiscountTrialList(client HttpCaller, nextPagePath string) *CouponDiscountTrialList {
	return &CouponDiscountTrialList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountTrialList) Fetch() error {
	resources := &couponDiscountTrialList{}
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
func (list *CouponDiscountTrialList) Count() (*int64, error) {
	resources := &couponDiscountTrialList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
