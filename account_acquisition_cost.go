// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
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
	Data            []AccountAcquisitionCost `json:"data"`
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []AccountAcquisitionCost
}

func NewAccountAcquisitionCostList(client HttpCaller, nextPagePath string) *AccountAcquisitionCostList {
	return &AccountAcquisitionCostList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountAcquisitionCostList) Fetch() error {
	resources := &accountAcquisitionCostList{}
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
func (list *AccountAcquisitionCostList) Count() (*int64, error) {
	resources := &accountAcquisitionCostList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
