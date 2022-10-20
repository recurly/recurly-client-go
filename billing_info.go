// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
        "time"
)

type BillingInfo struct {
  recurlyResponse *ResponseMetadata

  
        Id string `json:"id,omitempty"`

  
        // Object type
        Object string `json:"object,omitempty"`

  
        AccountId string `json:"account_id,omitempty"`

  
        FirstName string `json:"first_name,omitempty"`

  
        LastName string `json:"last_name,omitempty"`

  
        Company string `json:"company,omitempty"`

  
        Address Address `json:"address,omitempty"`

  
        // Customer's VAT number (to avoid having the VAT applied). This is only used for automatically collected invoices.
        VatNumber string `json:"vat_number,omitempty"`

  
        Valid bool `json:"valid,omitempty"`

  
        PaymentMethod PaymentMethod `json:"payment_method,omitempty"`

  
        // Most recent fraud result.
        Fraud FraudInfo `json:"fraud,omitempty"`

  
        // The `primary_payment_method` field is used to indicate the primary billing info on the account. The first billing info created on an account will always become primary. This payment method will be used
        PrimaryPaymentMethod bool `json:"primary_payment_method,omitempty"`

  
        // The `backup_payment_method` field is used to indicate a billing info as a backup on the account that will be tried if the initial billing info used for an invoice is declined.
        BackupPaymentMethod bool `json:"backup_payment_method,omitempty"`

  
        // When the billing information was created.
        CreatedAt time.Time `json:"created_at,omitempty"`

  
        // When the billing information was last changed.
        UpdatedAt time.Time `json:"updated_at,omitempty"`

  
        UpdatedBy BillingInfoUpdatedBy `json:"updated_by,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *BillingInfo) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *BillingInfo) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type billingInfoList struct {
	ListMetadata
  Data []BillingInfo `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *billingInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *billingInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// BillingInfoList allows you to paginate BillingInfo objects
type BillingInfoList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []BillingInfo
}

func NewBillingInfoList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *BillingInfoList {
  return &BillingInfoList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type BillingInfoLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []BillingInfo
  HasMore() bool
  Next() string
}

func (list  *BillingInfoList) HasMore() bool {
    return list.hasMore
}

func (list  *BillingInfoList) Next() string {
    return list.nextPagePath
}

func (list *BillingInfoList) Data() []BillingInfo {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *BillingInfoList) FetchWithContext(ctx context.Context) error {
	resources := &billingInfoList{}
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
func (list *BillingInfoList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *BillingInfoList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &billingInfoList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *BillingInfoList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
