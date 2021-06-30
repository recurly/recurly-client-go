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

type ShippingMethod struct {
	recurlyResponse *ResponseMetadata

	// Shipping Method ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The internal name used identify the shipping method.
	Code string `json:"code,omitempty"`

	// The name of the shipping method displayed to customers.
	Name string `json:"name,omitempty"`

	// Accounting code for shipping method.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s built-in tax feature. The tax
	// code values are specific to each tax system. If you are using Recurly’s
	// built-in taxes the values are:
	// - `FR` – Common Carrier FOB Destination
	// - `FR022000` – Common Carrier FOB Origin
	// - `FR020400` – Non Common Carrier FOB Destination
	// - `FR020500` – Non Common Carrier FOB Origin
	// - `FR010100` – Delivery by Company Vehicle Before Passage of Title
	// - `FR010200` – Delivery by Company Vehicle After Passage of Title
	// - `NT` – Non-Taxable
	TaxCode string `json:"tax_code,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ShippingMethod) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ShippingMethod) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type shippingMethodList struct {
	ListMetadata
	Data            []ShippingMethod `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *shippingMethodList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *shippingMethodList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ShippingMethodList allows you to paginate ShippingMethod objects
type ShippingMethodList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []ShippingMethod
}

func NewShippingMethodList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ShippingMethodList {
	return &ShippingMethodList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ShippingMethodList) FetchWithContext(ctx context.Context) error {
	resources := &shippingMethodList{}
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
func (list *ShippingMethodList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ShippingMethodList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &shippingMethodList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ShippingMethodList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
