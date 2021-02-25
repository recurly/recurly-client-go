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
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []MeasuredUnit
}

func NewMeasuredUnitList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *MeasuredUnitList {
	return &MeasuredUnitList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type MeasuredUnitLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []MeasuredUnit
	HasMore() bool
	Next() string
}

func (list *MeasuredUnitList) HasMore() bool {
	return list.hasMore
}

func (list *MeasuredUnitList) Next() string {
	return list.nextPagePath
}

func (list *MeasuredUnitList) Data() []MeasuredUnit {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *MeasuredUnitList) FetchWithContext(ctx context.Context) error {
	resources := &measuredUnitList{}
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
func (list *MeasuredUnitList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *MeasuredUnitList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &measuredUnitList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *MeasuredUnitList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
