// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type InvoiceAddress struct {
	recurlyResponse *ResponseMetadata

	// Name on account
	NameOnAccount string `json:"name_on_account,omitempty"`

	// Company
	Company string `json:"company,omitempty"`

	// First name
	FirstName string `json:"first_name,omitempty"`

	// Last name
	LastName string `json:"last_name,omitempty"`

	// Phone number
	Phone string `json:"phone,omitempty"`

	// Street 1
	Street1 string `json:"street1,omitempty"`

	// Street 2
	Street2 string `json:"street2,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// State or province.
	Region string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO 3166-1 alpha-2 code.
	Country string `json:"country,omitempty"`

	// Code that represents a geographic entity (location or object). Only returned for Sling Vertex Integration
	GeoCode string `json:"geo_code,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *InvoiceAddress) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *InvoiceAddress) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceAddressList struct {
	ListMetadata
	Data            []InvoiceAddress `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceAddressList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceAddressList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceAddressList allows you to paginate InvoiceAddress objects
type InvoiceAddressList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []InvoiceAddress
}

func NewInvoiceAddressList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *InvoiceAddressList {
	return &InvoiceAddressList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceAddressList) FetchWithContext(ctx context.Context) error {
	resources := &invoiceAddressList{}
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
func (list *InvoiceAddressList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceAddressList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &invoiceAddressList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *InvoiceAddressList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
