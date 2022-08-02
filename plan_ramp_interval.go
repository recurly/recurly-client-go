// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PlanRampInterval struct {
	recurlyResponse *ResponseMetadata

	// Represents the first billing cycle of a ramp.
	StartingBillingCycle int `json:"starting_billing_cycle,omitempty"`

	// Represents the price for the ramp interval.
	Currencies []PlanRampPricing `json:"currencies,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanRampInterval) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanRampInterval) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planRampIntervalList struct {
	ListMetadata
	Data            []PlanRampInterval `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planRampIntervalList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planRampIntervalList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanRampIntervalList allows you to paginate PlanRampInterval objects
type PlanRampIntervalList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []PlanRampInterval
}

func NewPlanRampIntervalList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanRampIntervalList {
	return &PlanRampIntervalList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanRampIntervalList) FetchWithContext(ctx context.Context) error {
	resources := &planRampIntervalList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanRampIntervalList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanRampIntervalList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planRampIntervalList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanRampIntervalList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
