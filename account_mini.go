// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
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
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string

	HasMore bool
	Data    []AccountMini
}

func NewAccountMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AccountMiniList {
	return &AccountMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountMiniList) FetchWithContext(ctx context.Context) error {
	resources := &accountMiniList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AccountMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &accountMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
