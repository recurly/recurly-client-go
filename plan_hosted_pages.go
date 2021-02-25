// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PlanHostedPages struct {
	recurlyResponse *ResponseMetadata

	// URL to redirect to after signup on the hosted payment pages.
	SuccessUrl string `json:"success_url,omitempty"`

	// URL to redirect to on canceled signup on the hosted payment pages.
	CancelUrl string `json:"cancel_url,omitempty"`

	// If `true`, the customer will be sent directly to your `success_url` after a successful signup, bypassing Recurly's hosted confirmation page.
	BypassConfirmation bool `json:"bypass_confirmation,omitempty"`

	// Determines if the quantity field is displayed on the hosted pages for the plan.
	DisplayQuantity bool `json:"display_quantity,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanHostedPages) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanHostedPages) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planHostedPagesList struct {
	ListMetadata
	Data            []PlanHostedPages `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planHostedPagesList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planHostedPagesList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanHostedPagesList allows you to paginate PlanHostedPages objects
type PlanHostedPagesList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PlanHostedPages
}

func NewPlanHostedPagesList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanHostedPagesList {
	return &PlanHostedPagesList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PlanHostedPagesLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PlanHostedPages
	HasMore() bool
	Next() string
}

func (list *PlanHostedPagesList) HasMore() bool {
	return list.hasMore
}

func (list *PlanHostedPagesList) Next() string {
	return list.nextPagePath
}

func (list *PlanHostedPagesList) Data() []PlanHostedPages {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanHostedPagesList) FetchWithContext(ctx context.Context) error {
	resources := &planHostedPagesList{}
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
func (list *PlanHostedPagesList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanHostedPagesList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planHostedPagesList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanHostedPagesList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
