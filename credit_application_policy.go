// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type CreditApplicationPolicy struct {
	recurlyResponse *ResponseMetadata

	// Determines which credit invoices are applied to invoices:
	// - `all`: All available credit invoices are applied (default)
	// - `none`: No credit invoices are applied automatically
	Mode string `json:"mode,omitempty"`

	// Optional array of credit invoice origin types to allow when mode is `all`.
	// If not specified when mode is `all`, credits from all origins are applied.
	// Only valid when mode is `all`.
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CreditApplicationPolicy) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CreditApplicationPolicy) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type creditApplicationPolicyList struct {
	ListMetadata
	Data            []CreditApplicationPolicy `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *creditApplicationPolicyList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *creditApplicationPolicyList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CreditApplicationPolicyList allows you to paginate CreditApplicationPolicy objects
type CreditApplicationPolicyList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []CreditApplicationPolicy
}

func NewCreditApplicationPolicyList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CreditApplicationPolicyList {
	return &CreditApplicationPolicyList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type CreditApplicationPolicyLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []CreditApplicationPolicy
	HasMore() bool
	Next() string
}

func (list *CreditApplicationPolicyList) HasMore() bool {
	return list.hasMore
}

func (list *CreditApplicationPolicyList) Next() string {
	return list.nextPagePath
}

func (list *CreditApplicationPolicyList) Data() []CreditApplicationPolicy {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *CreditApplicationPolicyList) FetchWithContext(ctx context.Context) error {
	resources := &creditApplicationPolicyList{}
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
func (list *CreditApplicationPolicyList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CreditApplicationPolicyList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &creditApplicationPolicyList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CreditApplicationPolicyList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
