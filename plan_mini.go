// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PlanMini struct {
	recurlyResponse *ResponseMetadata

	// Plan ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique code to identify the plan. This is used in Hosted Payment Page URLs and in the invoice exports.
	Code string `json:"code,omitempty"`

	// This name describes your plan and will appear on the Hosted Payment Page and the subscriber's invoice.
	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planMiniList struct {
	ListMetadata
	Data            []PlanMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanMiniList allows you to paginate PlanMini objects
type PlanMiniList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string

	HasMore bool
	Data    []PlanMini
}

func NewPlanMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanMiniList {
	return &PlanMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanMiniList) FetchWithContext(ctx context.Context) error {
	resources := &planMiniList{}
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
func (list *PlanMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
