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

type ExternalSubscription struct {
	recurlyResponse *ResponseMetadata

	// System-generated unique identifier for an external subscription ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// External Product Reference details
	ExternalProductReference ExternalProductReferenceMini `json:"external_product_reference,omitempty"`

	// The id of the subscription in the external systems., I.e. Apple App Store or Google Play Store.
	ExternalId string `json:"external_id,omitempty"`

	// When a new billing event occurred on the external subscription in conjunction with a recent billing period, reactivation or upgrade/downgrade.
	LastPurchased time.Time `json:"last_purchased,omitempty"`

	// An indication of whether or not the external subscription will auto-renew at the expiration date.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// An indication of whether or not the external subscription is in a grace period.
	InGracePeriod bool `json:"in_grace_period,omitempty"`

	// Identifier of the app that generated the external subscription.
	AppIdentifier string `json:"app_identifier,omitempty"`

	// An indication of the quantity of a subscribed item's quantity.
	Quantity int `json:"quantity,omitempty"`

	// External subscriptions can be active, canceled, expired, or past_due.
	State string `json:"state,omitempty"`

	// When the external subscription was activated in the external platform.
	ActivatedAt time.Time `json:"activated_at,omitempty"`

	// When the external subscription was canceled in the external platform.
	CanceledAt time.Time `json:"canceled_at,omitempty"`

	// When the external subscription expires in the external platform.
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	// When the external subscription trial period started in the external platform.
	TrialStartedAt time.Time `json:"trial_started_at,omitempty"`

	// When the external subscription trial period ends in the external platform.
	TrialEndsAt time.Time `json:"trial_ends_at,omitempty"`

	// An indication of whether or not the external subscription was purchased in a sandbox environment.
	Test bool `json:"test,omitempty"`

	// When the external subscription was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the external subscription was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalSubscription) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalSubscription) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalSubscriptionList struct {
	ListMetadata
	Data            []ExternalSubscription `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalSubscriptionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalSubscriptionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalSubscriptionList allows you to paginate ExternalSubscription objects
type ExternalSubscriptionList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalSubscription
}

func NewExternalSubscriptionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalSubscriptionList {
	return &ExternalSubscriptionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalSubscriptionLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalSubscription
	HasMore() bool
	Next() string
}

func (list *ExternalSubscriptionList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalSubscriptionList) Next() string {
	return list.nextPagePath
}

func (list *ExternalSubscriptionList) Data() []ExternalSubscription {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalSubscriptionList) FetchWithContext(ctx context.Context) error {
	resources := &externalSubscriptionList{}
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
func (list *ExternalSubscriptionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalSubscriptionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalSubscriptionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalSubscriptionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
