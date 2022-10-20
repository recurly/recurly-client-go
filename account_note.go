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

type AccountNote struct {
  recurlyResponse *ResponseMetadata

  
        Id string `json:"id,omitempty"`

  
        // Object type
        Object string `json:"object,omitempty"`

  
        AccountId string `json:"account_id,omitempty"`

  
        User User `json:"user,omitempty"`

  
        Message string `json:"message,omitempty"`

  
        CreatedAt time.Time `json:"created_at,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountNote) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountNote) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type accountNoteList struct {
	ListMetadata
  Data []AccountNote `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountNoteList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountNoteList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// AccountNoteList allows you to paginate AccountNote objects
type AccountNoteList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []AccountNote
}

func NewAccountNoteList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AccountNoteList {
  return &AccountNoteList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type AccountNoteLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []AccountNote
  HasMore() bool
  Next() string
}

func (list  *AccountNoteList) HasMore() bool {
    return list.hasMore
}

func (list  *AccountNoteList) Next() string {
    return list.nextPagePath
}

func (list *AccountNoteList) Data() []AccountNote {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *AccountNoteList) FetchWithContext(ctx context.Context) error {
	resources := &accountNoteList{}
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
func (list *AccountNoteList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AccountNoteList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &accountNoteList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountNoteList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
