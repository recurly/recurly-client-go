// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ItemMini struct {
	recurlyResponse *ResponseMetadata

	// Item ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique code to identify the item.
	Code string `json:"code,omitempty"`

	// The current state of the item.
	State string `json:"state,omitempty"`

	// This name describes your item and will appear on the invoice when it's purchased on a one time basis.
	Name string `json:"name,omitempty"`

	// Optional, description.
	Description string `json:"description,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ItemMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ItemMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type itemMiniList struct {
	ListMetadata
	Data            []ItemMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *itemMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *itemMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ItemMiniList allows you to paginate ItemMini objects
type ItemMiniList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []ItemMini
}

func NewItemMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ItemMiniList {
	return &ItemMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ItemMiniList) FetchWithContext(ctx context.Context) error {
	resources := &itemMiniList{}
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
func (list *ItemMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ItemMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &itemMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ItemMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
