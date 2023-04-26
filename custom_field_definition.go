// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
	"time"
)

type CustomFieldDefinition struct {
	recurlyResponse *ResponseMetadata

	// Custom field definition ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Related Recurly object type
	RelatedType string `json:"related_type,omitempty"`

	// Used by the API to identify the field or reading and writing. The name can only be used once per Recurly object type.
	Name string `json:"name,omitempty"`

	// The access control applied inside Recurly's admin UI:
	// - `api_only` - No one will be able to view or edit this field's data via the admin UI.
	// - `read_only` - Users with the Customers role will be able to view this field's data via the admin UI, but
	//   editing will only be available via the API.
	// - `write` - Users with the Customers role will be able to view and edit this field's data via the admin UI.
	// - `set_only` - Users with the Customers role will be able to set this field's data via the admin console.
	UserAccess string `json:"user_access,omitempty"`

	// Used to label the field when viewing and editing the field in Recurly's admin UI.
	DisplayName string `json:"display_name,omitempty"`

	// Displayed as a tooltip when editing the field in the Recurly admin UI.
	Tooltip string `json:"tooltip,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Definitions are initially soft deleted, and once all the values are removed from the accouts or subscriptions, will be hard deleted an no longer visible.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CustomFieldDefinition) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CustomFieldDefinition) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type customFieldDefinitionList struct {
	ListMetadata
	Data            []CustomFieldDefinition `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *customFieldDefinitionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *customFieldDefinitionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CustomFieldDefinitionList allows you to paginate CustomFieldDefinition objects
type CustomFieldDefinitionList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []CustomFieldDefinition
}

func NewCustomFieldDefinitionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CustomFieldDefinitionList {
	return &CustomFieldDefinitionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *CustomFieldDefinitionList) FetchWithContext(ctx context.Context) error {
	resources := &customFieldDefinitionList{}
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
func (list *CustomFieldDefinitionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CustomFieldDefinitionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &customFieldDefinitionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CustomFieldDefinitionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
