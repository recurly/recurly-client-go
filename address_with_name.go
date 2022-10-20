// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type AddressWithName struct {
  recurlyResponse *ResponseMetadata

  
        // First name
        FirstName string `json:"first_name,omitempty"`

  
        // Last name
        LastName string `json:"last_name,omitempty"`

  
        // Phone number
        Phone string `json:"phone,omitempty"`

  
        // Street 1
        Street1 string `json:"street1,omitempty"`

  
        // Street 2
        Street2 string `json:"street2,omitempty"`

  
        // City
        City string `json:"city,omitempty"`

  
        // State or province.
        Region string `json:"region,omitempty"`

  
        // Zip or postal code.
        PostalCode string `json:"postal_code,omitempty"`

  
        // Country, 2-letter ISO 3166-1 alpha-2 code.
        Country string `json:"country,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AddressWithName) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AddressWithName) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type addressWithNameList struct {
	ListMetadata
  Data []AddressWithName `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *addressWithNameList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *addressWithNameList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// AddressWithNameList allows you to paginate AddressWithName objects
type AddressWithNameList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []AddressWithName
}

func NewAddressWithNameList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AddressWithNameList {
  return &AddressWithNameList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type AddressWithNameLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []AddressWithName
  HasMore() bool
  Next() string
}

func (list  *AddressWithNameList) HasMore() bool {
    return list.hasMore
}

func (list  *AddressWithNameList) Next() string {
    return list.nextPagePath
}

func (list *AddressWithNameList) Data() []AddressWithName {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *AddressWithNameList) FetchWithContext(ctx context.Context) error {
	resources := &addressWithNameList{}
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
func (list *AddressWithNameList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AddressWithNameList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &addressWithNameList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AddressWithNameList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
