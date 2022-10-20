// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type PlanRampPricing struct {
  recurlyResponse *ResponseMetadata

  
        // 3-letter ISO 4217 currency code.
        Currency string `json:"currency,omitempty"`

  
        // Represents the price for the Ramp Interval.
        UnitAmount float64 `json:"unit_amount,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanRampPricing) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanRampPricing) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type planRampPricingList struct {
	ListMetadata
  Data []PlanRampPricing `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planRampPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planRampPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// PlanRampPricingList allows you to paginate PlanRampPricing objects
type PlanRampPricingList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []PlanRampPricing
}

func NewPlanRampPricingList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanRampPricingList {
  return &PlanRampPricingList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type PlanRampPricingLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []PlanRampPricing
  HasMore() bool
  Next() string
}

func (list  *PlanRampPricingList) HasMore() bool {
    return list.hasMore
}

func (list  *PlanRampPricingList) Next() string {
    return list.nextPagePath
}

func (list *PlanRampPricingList) Data() []PlanRampPricing {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *PlanRampPricingList) FetchWithContext(ctx context.Context) error {
	resources := &planRampPricingList{}
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
func (list *PlanRampPricingList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanRampPricingList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planRampPricingList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanRampPricingList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
