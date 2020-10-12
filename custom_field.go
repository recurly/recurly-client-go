package recurly

import (
	"net/http"
)

type CustomField struct {
	recurlyResponse *ResponseMetadata

	// Fields must be created in the UI before values can be assigned to them.
	Name string `json:"name,omitempty"`

	// Any values that resemble a credit card number or security code (CVV/CVC) will be rejected.
	Value string `json:"value,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CustomField) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CustomField) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type customFieldList struct {
	ListMetadata
	Data            []CustomField `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *customFieldList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *customFieldList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CustomFieldList allows you to paginate CustomField objects
type CustomFieldList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []CustomField
}

func NewCustomFieldList(client HttpCaller, nextPagePath string) *CustomFieldList {
	return &CustomFieldList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CustomFieldList) Fetch() error {
	resources := &customFieldList{}
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
func (list *CustomFieldList) Count() (*int64, error) {
	resources := &customFieldList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
