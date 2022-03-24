// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PercentageTier struct {
	recurlyResponse *ResponseMetadata

	// Ending amount for the tier. Allows up to 2 decimal places. Must be left empty if it is the final tier.
	EndingAmount float64 `json:"ending_amount,omitempty"`

	// The percentage taken of the monetary amount of usage tracked.
	// This can be up to 4 decimal places represented as a string.
	UsagePercentage string `json:"usage_percentage,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PercentageTier) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PercentageTier) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type percentageTierList struct {
	ListMetadata
	Data            []PercentageTier `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *percentageTierList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *percentageTierList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PercentageTierList allows you to paginate PercentageTier objects
type PercentageTierList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PercentageTier
}

func NewPercentageTierList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PercentageTierList {
	return &PercentageTierList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PercentageTierLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PercentageTier
	HasMore() bool
	Next() string
}

func (list *PercentageTierList) HasMore() bool {
	return list.hasMore
}

func (list *PercentageTierList) Next() string {
	return list.nextPagePath
}

func (list *PercentageTierList) Data() []PercentageTier {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PercentageTierList) FetchWithContext(ctx context.Context) error {
	resources := &percentageTierList{}
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
func (list *PercentageTierList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PercentageTierList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &percentageTierList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PercentageTierList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
