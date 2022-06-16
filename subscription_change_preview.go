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

type SubscriptionChangePreview struct {
	recurlyResponse *ResponseMetadata

	// The ID of the Subscription Change.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The ID of the subscription that is going to be changed.
	SubscriptionId string `json:"subscription_id,omitempty"`

	// Just the important parts.
	Plan PlanMini `json:"plan,omitempty"`

	// These add-ons will be used when the subscription renews.
	AddOns []SubscriptionAddOn `json:"add_ons,omitempty"`

	// Unit amount
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// This field is deprecated. Please do not use it.
	TaxInclusive bool `json:"tax_inclusive,omitempty"`

	// Subscription quantity
	Quantity int `json:"quantity,omitempty"`

	// Subscription shipping details
	Shipping SubscriptionShipping `json:"shipping,omitempty"`

	// Activated at
	ActivateAt time.Time `json:"activate_at,omitempty"`

	// Returns `true` if the subscription change is activated.
	Activated bool `json:"activated,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Setup fee revenue schedule type
	SetupFeeRevenueScheduleType string `json:"setup_fee_revenue_schedule_type,omitempty"`

	// Invoice Collection
	InvoiceCollection InvoiceCollection `json:"invoice_collection,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// Accept nested attributes for three_d_secure_action_result_token_id
	BillingInfo SubscriptionChangeBillingInfo `json:"billing_info,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionChangePreview) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionChangePreview) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionChangePreviewList struct {
	ListMetadata
	Data            []SubscriptionChangePreview `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionChangePreviewList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionChangePreviewList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionChangePreviewList allows you to paginate SubscriptionChangePreview objects
type SubscriptionChangePreviewList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []SubscriptionChangePreview
}

func NewSubscriptionChangePreviewList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SubscriptionChangePreviewList {
	return &SubscriptionChangePreviewList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionChangePreviewList) FetchWithContext(ctx context.Context) error {
	resources := &subscriptionChangePreviewList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionChangePreviewList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionChangePreviewList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &subscriptionChangePreviewList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionChangePreviewList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
