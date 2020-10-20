package recurly

import (
	"net/http"
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
	Data            []AccountNote `json:"data"`
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []AccountNote
}

func NewAccountNoteList(client HttpCaller, nextPagePath string) *AccountNoteList {
	return &AccountNoteList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountNoteList) Fetch() error {
	resources := &accountNoteList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountNoteList) Count() (*int64, error) {
	resources := &accountNoteList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
