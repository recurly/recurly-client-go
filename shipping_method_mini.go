// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type ShippingMethodMini struct {
  recurlyResponse *ResponseMetadata

  
        // Shipping Method ID
        Id string `json:"id,omitempty"`

  
        // Object type
        Object string `json:"object,omitempty"`

  
        // The internal name used identify the shipping method.
        Code string `json:"code,omitempty"`

  
        // The name of the shipping method displayed to customers.
        Name string `json:"name,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ShippingMethodMini) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ShippingMethodMini) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type shippingMethodMiniList struct {
	ListMetadata
  Data []ShippingMethodMini `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *shippingMethodMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *shippingMethodMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// ShippingMethodMiniList allows you to paginate ShippingMethodMini objects
type ShippingMethodMiniList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []ShippingMethodMini
}

func NewShippingMethodMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ShippingMethodMiniList {
  return &ShippingMethodMiniList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type ShippingMethodMiniLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []ShippingMethodMini
  HasMore() bool
  Next() string
}

func (list  *ShippingMethodMiniList) HasMore() bool {
    return list.hasMore
}

func (list  *ShippingMethodMiniList) Next() string {
    return list.nextPagePath
}

func (list *ShippingMethodMiniList) Data() []ShippingMethodMini {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *ShippingMethodMiniList) FetchWithContext(ctx context.Context) error {
	resources := &shippingMethodMiniList{}
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
func (list *ShippingMethodMiniList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ShippingMethodMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &shippingMethodMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ShippingMethodMiniList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
