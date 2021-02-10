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

type Item struct {
	recurlyResponse *ResponseMetadata

	// Item ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique code to identify the item.
	Code string `json:"code,omitempty"`

	// The current state of the item.
	State string `json:"state,omitempty"`

	// This name describes your item and will appear on the invoice when it's purchased on a one time basis.
	Name string `json:"name,omitempty"`

	// Optional, description.
	Description string `json:"description,omitempty"`

	// Optional, stock keeping unit to link the item to other inventory systems.
	ExternalSku string `json:"external_sku,omitempty"`

	// Accounting code for invoice line items.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraTransactionType int `json:"avalara_transaction_type,omitempty"`

	// Used by Avalara for Communications taxes. The transaction type in combination with the service type describe how the item is taxed. Refer to [the documentation](https://help.avalara.com/AvaTax_for_Communications/Tax_Calculation/AvaTax_for_Communications_Tax_Engine/Mapping_Resources/TM_00115_AFC_Modules_Corresponding_Transaction_Types) for more available t/s types.
	AvalaraServiceType int `json:"avalara_service_type,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// `true` exempts tax on the item, `false` applies tax on the item.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Item Pricing
	Currencies []Pricing `json:"currencies,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Item) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Item) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type itemList struct {
	ListMetadata
	Data            []Item `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *itemList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *itemList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ItemList allows you to paginate Item objects
type ItemList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Item
}

type ItemLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Item
	HasMore() bool
	Next() string
}

func NewItemList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ItemList {
	return &ItemList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *ItemList) HasMore() bool {
	return list.hasMore
}

func (list *ItemList) Next() string {
	return list.nextPagePath
}

func (list *ItemList) Data() []Item {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ItemList) FetchWithContext(ctx context.Context) error {
	resources := &itemList{}
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
func (list *ItemList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ItemList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &itemList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ItemList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
