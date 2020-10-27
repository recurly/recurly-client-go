// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []ExportFiles
}

func NewExportFilesList(client HttpCaller, nextPagePath string) *ExportFilesList {
	return &ExportFilesList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExportFilesList) Fetch() error {
	resources := &exportFilesList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *ExportFilesList) Count() (*int64, error) {
	resources := &exportFilesList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
