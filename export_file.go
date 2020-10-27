// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []ExportFile
}

func NewExportFileList(client HttpCaller, nextPagePath string) *ExportFileList {
	return &ExportFileList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExportFileList) Fetch() error {
	resources := &exportFileList{}
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
func (list *ExportFileList) Count() (*int64, error) {
	resources := &exportFileList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
