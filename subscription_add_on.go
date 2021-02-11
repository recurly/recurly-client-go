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

type SubscriptionAddOn struct {
	recurlyResponse *ResponseMetadata

	// Subscription Add-on ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Subscription ID
	SubscriptionId string `json:"subscription_id,omitempty"`

	// Just the important parts.
	AddOn AddOnMini `json:"add_on,omitempty"`

	// Used to determine where the associated add-on data is pulled from. If this value is set to
	// `plan_add_on` or left blank, then add-on data will be pulled from the plan's add-ons. If the associated
	// `plan` has `allow_any_item_on_subscriptions` set to `true` and this field is set to `item`, then
	// the associated add-on data will be pulled from the site's item catalog.
	AddOnSource string `json:"add_on_source,omitempty"`

	// Add-on quantity
	Quantity int `json:"quantity,omitempty"`

	// This is priced in the subscription's currency.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// The pricing model for the add-on.  For more information,
	// [click here](https://docs.recurly.com/docs/billing-models#section-quantity-based). See our
	// [Guide](https://developers.recurly.com/guides/item-addon-guide.html) for an overview of how
	// to configure quantity-based pricing models.
	TierType string `json:"tier_type,omitempty"`

	// If tiers are provided in the request, all existing tiers on the Subscription Add-on will be
	// removed and replaced by the tiers in the request.
	Tiers []SubscriptionAddOnTier `json:"tiers,omitempty"`

	// The percentage taken of the monetary amount of usage tracked. This can be up to 4 decimal places. A value between 0.0 and 100.0. Required if add_on_type is usage and usage_type is percentage.
	UsagePercentage float64 `json:"usage_percentage,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Expired at
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOn) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOn) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionAddOnList struct {
	ListMetadata
	Data            []SubscriptionAddOn `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionAddOnList allows you to paginate SubscriptionAddOn objects
type SubscriptionAddOnList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []SubscriptionAddOn
}

type SubscriptionAddOnLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []SubscriptionAddOn
	HasMore() bool
	Next() string
}

func NewSubscriptionAddOnList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SubscriptionAddOnList {
	return &SubscriptionAddOnList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *SubscriptionAddOnList) HasMore() bool {
	return list.hasMore
}

func (list *SubscriptionAddOnList) Next() string {
	return list.nextPagePath
}

func (list *SubscriptionAddOnList) Data() []SubscriptionAddOn {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionAddOnList) FetchWithContext(ctx context.Context) error {
	resources := &subscriptionAddOnList{}
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
func (list *SubscriptionAddOnList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionAddOnList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &subscriptionAddOnList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionAddOnList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
