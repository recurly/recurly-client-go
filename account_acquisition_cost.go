// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type AccountAcquisitionCost struct {
  recurlyResponse *ResponseMetadata

  
        // 3-letter ISO 4217 currency code.
        Currency string `json:"currency,omitempty"`

  
        // The amount of the corresponding currency used to acquire the account.
        Amount float64 `json:"amount,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountAcquisitionCost) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountAcquisitionCost) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type accountAcquisitionCostList struct {
	ListMetadata
  Data []AccountAcquisitionCost `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountAcquisitionCostList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountAcquisitionCostList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// AccountAcquisitionCostList allows you to paginate AccountAcquisitionCost objects
type AccountAcquisitionCostList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []AccountAcquisitionCost
}

func NewAccountAcquisitionCostList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AccountAcquisitionCostList {
  return &AccountAcquisitionCostList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type AccountAcquisitionCostLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []AccountAcquisitionCost
  HasMore() bool
  Next() string
}

func (list  *AccountAcquisitionCostList) HasMore() bool {
    return list.hasMore
}

func (list  *AccountAcquisitionCostList) Next() string {
    return list.nextPagePath
}

func (list *AccountAcquisitionCostList) Data() []AccountAcquisitionCost {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *AccountAcquisitionCostList) FetchWithContext(ctx context.Context) error {
	resources := &accountAcquisitionCostList{}
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
func (list *AccountAcquisitionCostList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AccountAcquisitionCostList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &accountAcquisitionCostList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountAcquisitionCostList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
