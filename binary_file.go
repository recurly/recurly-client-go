// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
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
  Data []BinaryFile `json:"data"`
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
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []BinaryFile
}

func NewBinaryFileList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *BinaryFileList {
  return &BinaryFileList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type BinaryFileLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []BinaryFile
  HasMore() bool
  Next() string
}

func (list  *BinaryFileList) HasMore() bool {
    return list.hasMore
}

func (list  *BinaryFileList) Next() string {
    return list.nextPagePath
}

func (list *BinaryFileList) Data() []BinaryFile {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *BinaryFileList) FetchWithContext(ctx context.Context) error {
	resources := &binaryFileList{}
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
func (list *BinaryFileList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *BinaryFileList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &binaryFileList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *BinaryFileList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
