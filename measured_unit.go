package recurly

import (
	"net/http"
	"time"
)

type MeasuredUnit struct {
	recurlyResponse *ResponseMetadata

	// Item ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique internal name of the measured unit on your site.
	Name string `json:"name,omitempty"`

	// Display name for the measured unit. Can only contain spaces, underscores and must be alphanumeric.
	DisplayName string `json:"display_name,omitempty"`

	// The current state of the measured unit.
	State string `json:"state,omitempty"`

	// Optional internal description.
	Description string `json:"description,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *MeasuredUnit) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *MeasuredUnit) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type measuredUnitList struct {
	ListMetadata
	Data            []MeasuredUnit `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *measuredUnitList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *measuredUnitList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// MeasuredUnitList allows you to paginate MeasuredUnit objects
type MeasuredUnitList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []MeasuredUnit
}

func NewMeasuredUnitList(client HttpCaller, nextPagePath string) *MeasuredUnitList {
	return &MeasuredUnitList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *MeasuredUnitList) Fetch() error {
	resources := &measuredUnitList{}
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
func (list *MeasuredUnitList) Count() (*int64, error) {
	resources := &measuredUnitList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
