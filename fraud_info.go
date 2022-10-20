// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type FraudInfo struct {
  recurlyResponse *ResponseMetadata

  
        // Kount score
        Score int `json:"score,omitempty"`

  
        // Kount decision
        Decision string `json:"decision,omitempty"`

  
        // Kount rules
        RiskRulesTriggered map[string]interface{} `json:"risk_rules_triggered,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *FraudInfo) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *FraudInfo) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type fraudInfoList struct {
	ListMetadata
  Data []FraudInfo `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *fraudInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *fraudInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// FraudInfoList allows you to paginate FraudInfo objects
type FraudInfoList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []FraudInfo
}

func NewFraudInfoList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *FraudInfoList {
  return &FraudInfoList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type FraudInfoLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []FraudInfo
  HasMore() bool
  Next() string
}

func (list  *FraudInfoList) HasMore() bool {
    return list.hasMore
}

func (list  *FraudInfoList) Next() string {
    return list.nextPagePath
}

func (list *FraudInfoList) Data() []FraudInfo {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *FraudInfoList) FetchWithContext(ctx context.Context) error {
	resources := &fraudInfoList{}
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
func (list *FraudInfoList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *FraudInfoList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &fraudInfoList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *FraudInfoList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
