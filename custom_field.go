// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type CustomField struct {
  recurlyResponse *ResponseMetadata

  
        // Fields must be created in the UI before values can be assigned to them.
        Name string `json:"name,omitempty"`

  
        // Any values that resemble a credit card number or security code (CVV/CVC) will be rejected.
        Value string `json:"value,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CustomField) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CustomField) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type customFieldList struct {
	ListMetadata
  Data []CustomField `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *customFieldList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *customFieldList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// CustomFieldList allows you to paginate CustomField objects
type CustomFieldList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []CustomField
}

func NewCustomFieldList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CustomFieldList {
  return &CustomFieldList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type CustomFieldLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []CustomField
  HasMore() bool
  Next() string
}

func (list  *CustomFieldList) HasMore() bool {
    return list.hasMore
}

func (list  *CustomFieldList) Next() string {
    return list.nextPagePath
}

func (list *CustomFieldList) Data() []CustomField {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *CustomFieldList) FetchWithContext(ctx context.Context) error {
	resources := &customFieldList{}
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
func (list *CustomFieldList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CustomFieldList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &customFieldList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CustomFieldList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
