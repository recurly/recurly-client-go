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

type GeneralLedgerAccount struct {
	recurlyResponse *ResponseMetadata

	// The ID of a general ledger account. General ledger accounts are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique code to identify the ledger account. Each code must start
	// with a letter or number. The following special characters are
	// allowed: `-_.,:`
	Code string `json:"code,omitempty"`

	// Optional description.
	Description string `json:"description,omitempty"`

	AccountType string `json:"account_type,omitempty"`

	// Created at
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *GeneralLedgerAccount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *GeneralLedgerAccount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type generalLedgerAccountList struct {
	ListMetadata
	Data            []GeneralLedgerAccount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *generalLedgerAccountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *generalLedgerAccountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// GeneralLedgerAccountList allows you to paginate GeneralLedgerAccount objects
type GeneralLedgerAccountList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []GeneralLedgerAccount
}

func NewGeneralLedgerAccountList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *GeneralLedgerAccountList {
	return &GeneralLedgerAccountList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type GeneralLedgerAccountLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []GeneralLedgerAccount
	HasMore() bool
	Next() string
}

func (list *GeneralLedgerAccountList) HasMore() bool {
	return list.hasMore
}

func (list *GeneralLedgerAccountList) Next() string {
	return list.nextPagePath
}

func (list *GeneralLedgerAccountList) Data() []GeneralLedgerAccount {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *GeneralLedgerAccountList) FetchWithContext(ctx context.Context) error {
	resources := &generalLedgerAccountList{}
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
func (list *GeneralLedgerAccountList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *GeneralLedgerAccountList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &generalLedgerAccountList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *GeneralLedgerAccountList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
