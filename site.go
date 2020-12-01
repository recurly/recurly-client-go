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

type Site struct {
	recurlyResponse *ResponseMetadata

	// Site ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	Subdomain string `json:"subdomain,omitempty"`

	// This value is used to configure RecurlyJS to submit tokenized billing information.
	PublicApiKey string `json:"public_api_key,omitempty"`

	// Mode
	Mode string `json:"mode,omitempty"`

	Address Address `json:"address,omitempty"`

	Settings Settings `json:"settings,omitempty"`

	// A list of features enabled for the site.
	Features []string `json:"features,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Site) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Site) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type siteList struct {
	ListMetadata
	Data            []Site `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *siteList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *siteList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SiteList allows you to paginate Site objects
type SiteList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string

	HasMore bool
	Data    []Site
}

func NewSiteList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SiteList {
	return &SiteList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SiteList) FetchWithContext(ctx context.Context) error {
	resources := &siteList{}
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
func (list *SiteList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SiteList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &siteList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SiteList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
