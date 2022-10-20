// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type DunningCampaignsBulkUpdateResponse struct {
  recurlyResponse *ResponseMetadata

  
        // Object type
        Object string `json:"object,omitempty"`

  
        // An array containing all of the `Plan` resources that have been updated.
        Plans []Plan `json:"plans,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *DunningCampaignsBulkUpdateResponse) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *DunningCampaignsBulkUpdateResponse) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type dunningCampaignsBulkUpdateResponseList struct {
	ListMetadata
  Data []DunningCampaignsBulkUpdateResponse `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *dunningCampaignsBulkUpdateResponseList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *dunningCampaignsBulkUpdateResponseList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// DunningCampaignsBulkUpdateResponseList allows you to paginate DunningCampaignsBulkUpdateResponse objects
type DunningCampaignsBulkUpdateResponseList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []DunningCampaignsBulkUpdateResponse
}

func NewDunningCampaignsBulkUpdateResponseList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *DunningCampaignsBulkUpdateResponseList {
  return &DunningCampaignsBulkUpdateResponseList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type DunningCampaignsBulkUpdateResponseLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []DunningCampaignsBulkUpdateResponse
  HasMore() bool
  Next() string
}

func (list  *DunningCampaignsBulkUpdateResponseList) HasMore() bool {
    return list.hasMore
}

func (list  *DunningCampaignsBulkUpdateResponseList) Next() string {
    return list.nextPagePath
}

func (list *DunningCampaignsBulkUpdateResponseList) Data() []DunningCampaignsBulkUpdateResponse {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *DunningCampaignsBulkUpdateResponseList) FetchWithContext(ctx context.Context) error {
	resources := &dunningCampaignsBulkUpdateResponseList{}
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
func (list *DunningCampaignsBulkUpdateResponseList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *DunningCampaignsBulkUpdateResponseList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &dunningCampaignsBulkUpdateResponseList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *DunningCampaignsBulkUpdateResponseList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
