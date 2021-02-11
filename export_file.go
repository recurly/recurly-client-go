// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ExportFile struct {
	recurlyResponse *ResponseMetadata

	// Name of the export file.
	Name string `json:"name,omitempty"`

	// MD5 hash of the export file.
	Md5sum string `json:"md5sum,omitempty"`

	// A presigned link to download the export file.
	Href string `json:"href,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExportFile) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExportFile) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type exportFileList struct {
	ListMetadata
	Data            []ExportFile `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *exportFileList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *exportFileList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExportFileList allows you to paginate ExportFile objects
type ExportFileList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExportFile
}

type ExportFileLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExportFile
	HasMore() bool
	Next() string
}

func NewExportFileList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExportFileList {
	return &ExportFileList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *ExportFileList) HasMore() bool {
	return list.hasMore
}

func (list *ExportFileList) Next() string {
	return list.nextPagePath
}

func (list *ExportFileList) Data() []ExportFile {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExportFileList) FetchWithContext(ctx context.Context) error {
	resources := &exportFileList{}
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
func (list *ExportFileList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExportFileList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &exportFileList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExportFileList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
