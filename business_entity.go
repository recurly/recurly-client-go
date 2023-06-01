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

type BusinessEntity struct {
	recurlyResponse *ResponseMetadata

	// Business entity ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The entity code of the business entity.
	Code string `json:"code,omitempty"`

	// This name describes your business entity and will appear on the invoice.
	Name string `json:"name,omitempty"`

	// Address information for the business entity that will appear on the invoice.
	InvoiceDisplayAddress Address `json:"invoice_display_address,omitempty"`

	// Address information for the business entity that will be used for calculating taxes.
	TaxAddress Address `json:"tax_address,omitempty"`

	// VAT number for the customer used on the invoice.
	DefaultVatNumber string `json:"default_vat_number,omitempty"`

	// Registration number for the customer used on the invoice.
	DefaultRegistrationNumber string `json:"default_registration_number,omitempty"`

	// List of countries for which the business entity will be used.
	SubscriberLocationCountries []string `json:"subscriber_location_countries,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *BusinessEntity) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *BusinessEntity) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type businessEntityList struct {
	ListMetadata
	Data            []BusinessEntity `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *businessEntityList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *businessEntityList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// BusinessEntityList allows you to paginate BusinessEntity objects
type BusinessEntityList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []BusinessEntity
}

func NewBusinessEntityList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *BusinessEntityList {
	return &BusinessEntityList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type BusinessEntityLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []BusinessEntity
	HasMore() bool
	Next() string
}

func (list *BusinessEntityList) HasMore() bool {
	return list.hasMore
}

func (list *BusinessEntityList) Next() string {
	return list.nextPagePath
}

func (list *BusinessEntityList) Data() []BusinessEntity {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *BusinessEntityList) FetchWithContext(ctx context.Context) error {
	resources := &businessEntityList{}
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
func (list *BusinessEntityList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *BusinessEntityList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &businessEntityList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *BusinessEntityList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
