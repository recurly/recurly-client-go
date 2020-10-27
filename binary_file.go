// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type BinaryFile struct {
	recurlyResponse *ResponseMetadata

	Data string `json:"data,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *BinaryFile) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *BinaryFile) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type binaryFileList struct {
	ListMetadata
	Data            []BinaryFile `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *binaryFileList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *binaryFileList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// BinaryFileList allows you to paginate BinaryFile objects
type BinaryFileList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []BinaryFile
}

func NewBinaryFileList(client HttpCaller, nextPagePath string) *BinaryFileList {
	return &BinaryFileList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *BinaryFileList) Fetch() error {
	resources := &binaryFileList{}
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
func (list *BinaryFileList) Count() (*int64, error) {
	resources := &binaryFileList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
