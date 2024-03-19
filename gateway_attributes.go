// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type GatewayAttributes struct {
	recurlyResponse *ResponseMetadata

	// Used by Adyen and Braintree gateways. For Adyen the Shopper Reference value used when the external token was created. For Braintree the PayPal PayerID is populated in the response.
	AccountReference string `json:"account_reference,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *GatewayAttributes) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *GatewayAttributes) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type gatewayAttributesList struct {
	ListMetadata
	Data            []GatewayAttributes `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *gatewayAttributesList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *gatewayAttributesList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// GatewayAttributesList allows you to paginate GatewayAttributes objects
type GatewayAttributesList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []GatewayAttributes
}

func NewGatewayAttributesList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *GatewayAttributesList {
	return &GatewayAttributesList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *GatewayAttributesList) FetchWithContext(ctx context.Context) error {
	resources := &gatewayAttributesList{}
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
func (list *GatewayAttributesList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *GatewayAttributesList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &gatewayAttributesList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *GatewayAttributesList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
