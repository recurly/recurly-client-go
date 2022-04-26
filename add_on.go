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

type AddOn struct {
	recurlyResponse *ResponseMetadata

	// Add-on ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Plan ID
	PlanId string `json:"plan_id,omitempty"`

	// The unique identifier for the add-on within its plan.
	Code string `json:"code,omitempty"`

	// Add-ons can be either active or inactive.
	State string `json:"state,omitempty"`

	// Describes your add-on and will appear in subscribers' invoices.
	Name string `json:"name,omitempty"`

	// Whether the add-on type is fixed, or usage-based.
	AddOnType string `json:"add_on_type,omitempty"`

	// Type of usage, returns usage type if `add_on_type` is `usage`.
	UsageType string `json:"usage_type,omitempty"`

	// The percentage taken of the monetary amount of usage tracked. This can be up to 4 decimal places. A value between 0.0 and 100.0, represented as a string.
	UsagePercentage string `json:"usage_percentage,omitempty"`

	// System-generated unique identifier for an measured unit associated with the add-on.
	MeasuredUnitId string `json:"measured_unit_id,omitempty"`

	// Accounting code for invoice line items for this add-on. If no value is provided, it defaults to add-on's code.
	AccountingCode string `json:"accounting_code,omitempty"`

	// When this add-on is invoiced, the line item will use this revenue schedule. If `item_code`/`item_id` is part of the request then `revenue_schedule_type` must be absent in the request as the value will be set from the item.
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the add-on is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraTransactionType int `json:"avalara_transaction_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the add-on is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraServiceType int `json:"avalara_service_type,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// Determines if the quantity field is displayed on the hosted pages for the add-on.
	DisplayQuantity bool `json:"display_quantity,omitempty"`

	// Default quantity for the hosted pages.
	DefaultQuantity int `json:"default_quantity,omitempty"`

	// Whether the add-on is optional for the customer to include in their purchase on the hosted payment page. If false, the add-on will be included when a subscription is created through the Recurly UI. However, the add-on will not be included when a subscription is created through the API.
	Optional bool `json:"optional,omitempty"`

	// Add-on pricing
	Currencies []AddOnPricing `json:"currencies,omitempty"`

	// Just the important parts.
	Item ItemMini `json:"item,omitempty"`

	// The pricing model for the add-on.  For more information,
	// [click here](https://docs.recurly.com/docs/billing-models#section-quantity-based). See our
	// [Guide](https://developers.recurly.com/guides/item-addon-guide.html) for an overview of how
	// to configure quantity-based pricing models.
	TierType string `json:"tier_type,omitempty"`

	// The time at which usage totals are reset for billing purposes.
	UsageTimeframe string `json:"usage_timeframe,omitempty"`

	// Tiers
	Tiers []Tier `json:"tiers,omitempty"`

	// Percentage Tiers
	PercentageTiers []PercentageTiersByCurrency `json:"percentage_tiers,omitempty"`

	// Optional, stock keeping unit to link the item to other inventory systems.
	ExternalSku string `json:"external_sku,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AddOn) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AddOn) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type addOnList struct {
	ListMetadata
	Data            []AddOn `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *addOnList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *addOnList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AddOnList allows you to paginate AddOn objects
type AddOnList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []AddOn
}

func NewAddOnList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AddOnList {
	return &AddOnList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type AddOnLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []AddOn
	HasMore() bool
	Next() string
}

func (list *AddOnList) HasMore() bool {
	return list.hasMore
}

func (list *AddOnList) Next() string {
	return list.nextPagePath
}

func (list *AddOnList) Data() []AddOn {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *AddOnList) FetchWithContext(ctx context.Context) error {
	resources := &addOnList{}
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
func (list *AddOnList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AddOnList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &addOnList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AddOnList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
