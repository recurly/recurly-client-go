// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ExportFiles struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	Files []ExportFile `json:"files,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExportFiles) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExportFiles) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type exportFilesList struct {
	ListMetadata
	Data            []ExportFiles `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *exportFilesList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *exportFilesList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExportFilesList allows you to paginate ExportFiles objects
type ExportFilesList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExportFiles
}

func NewExportFilesList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExportFilesList {
	return &ExportFilesList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExportFilesLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExportFiles
	HasMore() bool
	Next() string
}

func (list *ExportFilesList) HasMore() bool {
	return list.hasMore
}

func (list *ExportFilesList) Next() string {
	return list.nextPagePath
}

func (list *ExportFilesList) Data() []ExportFiles {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExportFilesList) FetchWithContext(ctx context.Context) error {
	resources := &exportFilesList{}
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
func (list *ExportFilesList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExportFilesList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &exportFilesList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExportFilesList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
