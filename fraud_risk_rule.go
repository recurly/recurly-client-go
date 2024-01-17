// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type FraudRiskRule struct {
	recurlyResponse *ResponseMetadata

	// The Kount rule number.
	Code string `json:"code,omitempty"`

	// Description of why the rule was triggered
	Message string `json:"message,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *FraudRiskRule) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *FraudRiskRule) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type fraudRiskRuleList struct {
	ListMetadata
	Data            []FraudRiskRule `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *fraudRiskRuleList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *fraudRiskRuleList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// FraudRiskRuleList allows you to paginate FraudRiskRule objects
type FraudRiskRuleList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []FraudRiskRule
}

func NewFraudRiskRuleList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *FraudRiskRuleList {
	return &FraudRiskRuleList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type FraudRiskRuleLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []FraudRiskRule
	HasMore() bool
	Next() string
}

func (list *FraudRiskRuleList) HasMore() bool {
	return list.hasMore
}

func (list *FraudRiskRuleList) Next() string {
	return list.nextPagePath
}

func (list *FraudRiskRuleList) Data() []FraudRiskRule {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *FraudRiskRuleList) FetchWithContext(ctx context.Context) error {
	resources := &fraudRiskRuleList{}
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
func (list *FraudRiskRuleList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *FraudRiskRuleList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &fraudRiskRuleList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *FraudRiskRuleList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
