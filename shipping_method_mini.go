package recurly

import (
	"net/http"
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
	Data            []ShippingMethodMini `json:"data"`
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []ShippingMethodMini
}

func NewShippingMethodMiniList(client HttpCaller, nextPagePath string) *ShippingMethodMiniList {
	return &ShippingMethodMiniList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ShippingMethodMiniList) Fetch() error {
	resources := &shippingMethodMiniList{}
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
func (list *ShippingMethodMiniList) Count() (*int64, error) {
	resources := &shippingMethodMiniList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
