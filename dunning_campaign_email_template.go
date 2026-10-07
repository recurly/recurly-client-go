// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type DunningCampaignEmailTemplate struct {
	recurlyResponse *ResponseMetadata

	// The id to assign under `intervals[].email_template_id`.
	Id string `json:"id,omitempty"`

	// Template name.
	Name string `json:"name,omitempty"`

	// The root template this custom template replaces, e.g. `payment_declined`, `invoice_past_due`, `post_trial_payment_declined`, `subscription_canceled_nonpayment`.
	Type string `json:"type,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *DunningCampaignEmailTemplate) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *DunningCampaignEmailTemplate) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type dunningCampaignEmailTemplateList struct {
	ListMetadata
	Data            []DunningCampaignEmailTemplate `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *dunningCampaignEmailTemplateList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *dunningCampaignEmailTemplateList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// DunningCampaignEmailTemplateList allows you to paginate DunningCampaignEmailTemplate objects
type DunningCampaignEmailTemplateList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []DunningCampaignEmailTemplate
}

func NewDunningCampaignEmailTemplateList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *DunningCampaignEmailTemplateList {
	return &DunningCampaignEmailTemplateList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type DunningCampaignEmailTemplateLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []DunningCampaignEmailTemplate
	HasMore() bool
	Next() string
}

func (list *DunningCampaignEmailTemplateList) HasMore() bool {
	return list.hasMore
}

func (list *DunningCampaignEmailTemplateList) Next() string {
	return list.nextPagePath
}

func (list *DunningCampaignEmailTemplateList) Data() []DunningCampaignEmailTemplate {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *DunningCampaignEmailTemplateList) FetchWithContext(ctx context.Context) error {
	resources := &dunningCampaignEmailTemplateList{}
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
func (list *DunningCampaignEmailTemplateList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *DunningCampaignEmailTemplateList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &dunningCampaignEmailTemplateList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *DunningCampaignEmailTemplateList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
