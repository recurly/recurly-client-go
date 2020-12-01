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

type Usage struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Custom field for recording the id in your own system associated with the usage, so you can provide auditable usage displays to your customers using a GET on this endpoint.
	MerchantTag string `json:"merchant_tag,omitempty"`

	// The amount of usage. Can be positive, negative, or 0. No decimals allowed, we will strip them. If the usage-based add-on is billed with a percentage, your usage will be a monetary amount you will want to format in cents. (e.g., $5.00 is "500").
	Amount float64 `json:"amount,omitempty"`

	// Type of usage, returns usage type if `add_on_type` is `usage`.
	UsageType string `json:"usage_type,omitempty"`

	// The pricing model for the add-on.  For more information,
	// [click here](https://docs.recurly.com/docs/billing-models#section-quantity-based).
	TierType string `json:"tier_type,omitempty"`

	// The tiers and prices of the subscription based on the usage_timestamp. If tier_type = flat, tiers = null
	Tiers []SubscriptionAddOnTier `json:"tiers,omitempty"`

	// The ID of the measured unit associated with the add-on the usage record is for.
	MeasuredUnitId string `json:"measured_unit_id,omitempty"`

	// When the usage was recorded in your system.
	RecordingTimestamp time.Time `json:"recording_timestamp,omitempty"`

	// When the usage actually happened. This will define the line item dates this usage is billed under and is important for revenue recognition.
	UsageTimestamp time.Time `json:"usage_timestamp,omitempty"`

	// The percentage taken of the monetary amount of usage tracked. This can be up to 4 decimal places. A value between 0.0 and 100.0.
	UsagePercentage float64 `json:"usage_percentage,omitempty"`

	// Unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// When the usage record was billed on an invoice.
	BilledAt time.Time `json:"billed_at,omitempty"`

	// When the usage record was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the usage record was billed on an invoice.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Usage) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Usage) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type usageList struct {
	ListMetadata
	Data            []Usage `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *usageList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *usageList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// UsageList allows you to paginate Usage objects
type UsageList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string

	HasMore bool
	Data    []Usage
}

func NewUsageList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *UsageList {
	return &UsageList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *UsageList) FetchWithContext(ctx context.Context) error {
	resources := &usageList{}
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
func (list *UsageList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *UsageList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &usageList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *UsageList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
