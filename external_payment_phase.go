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

type ExternalPaymentPhase struct {
	recurlyResponse *ResponseMetadata

	// System-generated unique identifier for an external payment phase ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Subscription from an external resource such as Apple App Store or Google Play Store.
	ExternalSubscription ExternalSubscription `json:"external_subscription,omitempty"`

	// Started At
	StartedAt time.Time `json:"started_at,omitempty"`

	// Ends At
	EndsAt time.Time `json:"ends_at,omitempty"`

	// Starting Billing Period Index
	StartingBillingPeriodIndex int `json:"starting_billing_period_index,omitempty"`

	// Ending Billing Period Index
	EndingBillingPeriodIndex int `json:"ending_billing_period_index,omitempty"`

	// Type of discount offer given, e.g. "FREE_TRIAL"
	OfferType string `json:"offer_type,omitempty"`

	// Name of the discount offer given, e.g. "introductory"
	OfferName string `json:"offer_name,omitempty"`

	// Number of billing periods
	PeriodCount int `json:"period_count,omitempty"`

	// Billing cycle length
	PeriodLength string `json:"period_length,omitempty"`

	// Allows up to 9 decimal places
	Amount string `json:"amount,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// When the external subscription was created in Recurly.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the external subscription was updated in Recurly.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalPaymentPhase) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalPaymentPhase) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalPaymentPhaseList struct {
	ListMetadata
	Data            []ExternalPaymentPhase `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalPaymentPhaseList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalPaymentPhaseList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalPaymentPhaseList allows you to paginate ExternalPaymentPhase objects
type ExternalPaymentPhaseList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalPaymentPhase
}

func NewExternalPaymentPhaseList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalPaymentPhaseList {
	return &ExternalPaymentPhaseList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalPaymentPhaseLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalPaymentPhase
	HasMore() bool
	Next() string
}

func (list *ExternalPaymentPhaseList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalPaymentPhaseList) Next() string {
	return list.nextPagePath
}

func (list *ExternalPaymentPhaseList) Data() []ExternalPaymentPhase {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalPaymentPhaseList) FetchWithContext(ctx context.Context) error {
	resources := &externalPaymentPhaseList{}
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
func (list *ExternalPaymentPhaseList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalPaymentPhaseList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalPaymentPhaseList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalPaymentPhaseList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
