package recurly

import (
	"net/http"
)

type AccountMini struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The unique identifier of the account.
	Code string `json:"code,omitempty"`

	// The email address used for communicating with this customer.
	Email string `json:"email,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	ParentAccountId string `json:"parent_account_id,omitempty"`

	BillTo string `json:"bill_to,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountMiniList struct {
	ListMetadata
	Data            []AccountMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountMiniList allows you to paginate AccountMini objects
type AccountMiniList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []AccountMini
}

func NewAccountMiniList(client HttpCaller, nextPagePath string) *AccountMiniList {
	return &AccountMiniList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountMiniList) Fetch() error {
	resources := &accountMiniList{}
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
func (list *AccountMiniList) Count() (*int64, error) {
	resources := &accountMiniList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
