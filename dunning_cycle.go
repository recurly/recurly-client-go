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

type DunningCycle struct {
	recurlyResponse *ResponseMetadata

	// The type of invoice this cycle applies to.
	Type string `json:"type,omitempty"`

	// Whether the dunning settings will be applied to manual trials. Only applies to trial cycles.
	AppliesToManualTrial bool `json:"applies_to_manual_trial,omitempty"`

	// The number of days after a transaction failure before the first dunning email is sent.
	FirstCommunicationInterval int `json:"first_communication_interval,omitempty"`

	// Whether or not to send an extra email immediately to customers whose initial payment attempt fails with either a hard decline or invalid billing info.
	SendImmediatelyOnHardDecline bool `json:"send_immediately_on_hard_decline,omitempty"`

	// Dunning intervals.
	Intervals []DunningInterval `json:"intervals,omitempty"`

	// Whether the subscription(s) should be cancelled at the end of the dunning cycle.
	ExpireSubscription bool `json:"expire_subscription,omitempty"`

	// Whether the invoice should be failed at the end of the dunning cycle.
	FailInvoice bool `json:"fail_invoice,omitempty"`

	// The number of days between the first dunning email being sent and the end of the dunning cycle.
	TotalDunningDays int `json:"total_dunning_days,omitempty"`

	// The number of days between a transaction failure and the end of the dunning cycle.
	TotalRecyclingDays int `json:"total_recycling_days,omitempty"`

	// Current campaign version.
	Version int `json:"version,omitempty"`

	// When the current settings were created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the current settings were updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *DunningCycle) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *DunningCycle) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type dunningCycleList struct {
	ListMetadata
	Data            []DunningCycle `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *dunningCycleList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *dunningCycleList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// DunningCycleList allows you to paginate DunningCycle objects
type DunningCycleList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []DunningCycle
}

func NewDunningCycleList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *DunningCycleList {
	return &DunningCycleList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *DunningCycleList) FetchWithContext(ctx context.Context) error {
	resources := &dunningCycleList{}
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
func (list *DunningCycleList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *DunningCycleList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &dunningCycleList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *DunningCycleList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
