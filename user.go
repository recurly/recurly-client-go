// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
        "time"
)

type User struct {
  recurlyResponse *ResponseMetadata

  
        Id string `json:"id,omitempty"`

  
        // Object type
        Object string `json:"object,omitempty"`

  
        Email string `json:"email,omitempty"`

  
        FirstName string `json:"first_name,omitempty"`

  
        LastName string `json:"last_name,omitempty"`

  
        TimeZone string `json:"time_zone,omitempty"`

  
        CreatedAt time.Time `json:"created_at,omitempty"`

  
        DeletedAt time.Time `json:"deleted_at,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *User) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *User) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type userList struct {
	ListMetadata
  Data []User `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *userList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *userList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// UserList allows you to paginate User objects
type UserList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []User
}

func NewUserList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *UserList {
  return &UserList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type UserLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []User
  HasMore() bool
  Next() string
}

func (list  *UserList) HasMore() bool {
    return list.hasMore
}

func (list  *UserList) Next() string {
    return list.nextPagePath
}

func (list *UserList) Data() []User {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *UserList) FetchWithContext(ctx context.Context) error {
	resources := &userList{}
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
func (list *UserList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *UserList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &userList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *UserList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
