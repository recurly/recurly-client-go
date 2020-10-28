// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

type AddOnMini struct {
	recurlyResponse *ResponseMetadata

	// Add-on ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The unique identifier for the add-on within its plan.
	Code string `json:"code,omitempty"`

	// Describes your add-on and will appear in subscribers' invoices.
	Name string `json:"name,omitempty"`

	// Whether the add-on type is fixed, or usage-based.
	AddOnType string `json:"add_on_type,omitempty"`

	// Type of usage, returns usage type if `add_on_type` is `usage`.
	UsageType string `json:"usage_type,omitempty"`

	// The percentage taken of the monetary amount of usage tracked. This can be up to 4 decimal places. A value between 0.0 and 100.0.
	UsagePercentage float64 `json:"usage_percentage,omitempty"`

	// System-generated unique identifier for an measured unit associated with the add-on.
	MeasuredUnitId string `json:"measured_unit_id,omitempty"`

	// Item ID
	ItemId string `json:"item_id,omitempty"`

	// Optional, stock keeping unit to link the item to other inventory systems.
	ExternalSku string `json:"external_sku,omitempty"`

	// Accounting code for invoice line items for this add-on. If no value is provided, it defaults to add-on's code.
	AccountingCode string `json:"accounting_code,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AddOnMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AddOnMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type addOnMiniList struct {
	ListMetadata
	Data            []AddOnMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *addOnMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *addOnMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AddOnMiniList allows you to paginate AddOnMini objects
type AddOnMiniList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []AddOnMini
}

func NewAddOnMiniList(client HttpCaller, nextPagePath string) *AddOnMiniList {
	return &AddOnMiniList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AddOnMiniList) Fetch() error {
	resources := &addOnMiniList{}
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
func (list *AddOnMiniList) Count() (*int64, error) {
	resources := &addOnMiniList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
