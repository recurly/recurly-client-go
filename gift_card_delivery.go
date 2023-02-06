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

type GiftCardDelivery struct {
	recurlyResponse *ResponseMetadata

	// Whether the delivery method is email or postal service.
	Method string `json:"method,omitempty"`

	// The email address of the recipient.
	EmailAddress string `json:"email_address,omitempty"`

	// When the gift card should be delivered to the recipient. If null, the gift card will be delivered immediately. If a datetime is provided, the delivery will be in an hourly window, rounding down. For example, 6:23 pm will be in the 6:00 pm hourly batch.
	DeliverAt time.Time `json:"deliver_at,omitempty"`

	// The first name of the recipient.
	FirstName string `json:"first_name,omitempty"`

	// The last name of the recipient.
	LastName string `json:"last_name,omitempty"`

	// Address information for the recipient.
	RecipientAddress Address `json:"recipient_address,omitempty"`

	// The name of the gifter for the purpose of a message displayed to the recipient.
	GifterName string `json:"gifter_name,omitempty"`

	// The personal message from the gifter to the recipient.
	PersonalMessage string `json:"personal_message,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *GiftCardDelivery) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *GiftCardDelivery) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type giftCardDeliveryList struct {
	ListMetadata
	Data            []GiftCardDelivery `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *giftCardDeliveryList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *giftCardDeliveryList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// GiftCardDeliveryList allows you to paginate GiftCardDelivery objects
type GiftCardDeliveryList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []GiftCardDelivery
}

func NewGiftCardDeliveryList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *GiftCardDeliveryList {
	return &GiftCardDeliveryList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type GiftCardDeliveryLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []GiftCardDelivery
	HasMore() bool
	Next() string
}

func (list *GiftCardDeliveryList) HasMore() bool {
	return list.hasMore
}

func (list *GiftCardDeliveryList) Next() string {
	return list.nextPagePath
}

func (list *GiftCardDeliveryList) Data() []GiftCardDelivery {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *GiftCardDeliveryList) FetchWithContext(ctx context.Context) error {
	resources := &giftCardDeliveryList{}
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
func (list *GiftCardDeliveryList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *GiftCardDeliveryList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &giftCardDeliveryList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *GiftCardDeliveryList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
