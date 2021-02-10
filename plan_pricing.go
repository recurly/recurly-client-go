// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PlanPricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Amount of one-time setup fee automatically charged at the beginning of a subscription billing cycle. For subscription plans with a trial, the setup fee will be charged at the time of signup. Setup fees do not increase with the quantity of a subscription plan.
	SetupFee float64 `json:"setup_fee,omitempty"`

	// Unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanPricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanPricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planPricingList struct {
	ListMetadata
	Data            []PlanPricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanPricingList allows you to paginate PlanPricing objects
type PlanPricingList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PlanPricing
}

type PlanPricingLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PlanPricing
	HasMore() bool
	Next() string
}

func NewPlanPricingList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanPricingList {
	return &PlanPricingList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *PlanPricingList) HasMore() bool {
	return list.hasMore
}

func (list *PlanPricingList) Next() string {
	return list.nextPagePath
}

func (list *PlanPricingList) Data() []PlanPricing {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanPricingList) FetchWithContext(ctx context.Context) error {
	resources := &planPricingList{}
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
func (list *PlanPricingList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanPricingList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planPricingList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanPricingList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
