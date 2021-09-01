// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
	"time"
)

type DunningCampaign struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Campaign code.
	Code string `json:"code,omitempty"`

	// Campaign name.
	Name string `json:"name,omitempty"`

	// Campaign description.
	Description string `json:"description,omitempty"`

	// Whether or not this is the default campaign for accounts or plans without an assigned dunning campaign.
	DefaultCampaign bool `json:"default_campaign,omitempty"`

	// Dunning Cycle settings.
	DunningCycles []DunningCycle `json:"dunning_cycles,omitempty"`

	// When the current campaign was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the current campaign was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// When the current campaign was deleted in Recurly.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *DunningCampaign) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *DunningCampaign) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type dunningCampaignList struct {
	ListMetadata
	Data            []DunningCampaign `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *dunningCampaignList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *dunningCampaignList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// DunningCampaignList allows you to paginate DunningCampaign objects
type DunningCampaignList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []DunningCampaign
}

func NewDunningCampaignList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *DunningCampaignList {
	return &DunningCampaignList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *DunningCampaignList) FetchWithContext(ctx context.Context) error {
	resources := &dunningCampaignList{}
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
func (list *DunningCampaignList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *DunningCampaignList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &dunningCampaignList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *DunningCampaignList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
