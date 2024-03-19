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

type ShippingAddress struct {
	recurlyResponse *ResponseMetadata

	// Shipping Address ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Account ID
	AccountId string `json:"account_id,omitempty"`

	Nickname string `json:"nickname,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	Email string `json:"email,omitempty"`

	VatNumber string `json:"vat_number,omitempty"`

	Phone string `json:"phone,omitempty"`

	Street1 string `json:"street1,omitempty"`

	Street2 string `json:"street2,omitempty"`

	City string `json:"city,omitempty"`

	// State or province.
	Region string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO 3166-1 alpha-2 code.
	Country string `json:"country,omitempty"`

	// Code that represents a geographic entity (location or object). Only returned for Sling Vertex Integration
	GeoCode string `json:"geo_code,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ShippingAddress) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ShippingAddress) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type shippingAddressList struct {
	ListMetadata
	Data            []ShippingAddress `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *shippingAddressList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *shippingAddressList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ShippingAddressList allows you to paginate ShippingAddress objects
type ShippingAddressList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []ShippingAddress
}

func NewShippingAddressList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ShippingAddressList {
	return &ShippingAddressList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ShippingAddressList) FetchWithContext(ctx context.Context) error {
	resources := &shippingAddressList{}
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
func (list *ShippingAddressList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ShippingAddressList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &shippingAddressList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ShippingAddressList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
