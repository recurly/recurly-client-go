package recurly

import (
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []PlanHostedPages
}

func NewPlanHostedPagesList(client HttpCaller, nextPagePath string) *PlanHostedPagesList {
	return &PlanHostedPagesList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanHostedPagesList) Fetch() error {
	resources := &planHostedPagesList{}
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
func (list *PlanHostedPagesList) Count() (*int64, error) {
	resources := &planHostedPagesList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
