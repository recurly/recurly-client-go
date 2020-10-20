package recurly

import (
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

	// The type of tiering used by the Add-on.
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []SubscriptionAddOn
}

func NewSubscriptionAddOnList(client HttpCaller, nextPagePath string) *SubscriptionAddOnList {
	return &SubscriptionAddOnList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionAddOnList) Fetch() error {
	resources := &subscriptionAddOnList{}
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
func (list *SubscriptionAddOnList) Count() (*int64, error) {
	resources := &subscriptionAddOnList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
