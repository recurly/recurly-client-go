// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// APIVersion is the current Recurly API Version
	APIVersion = "v2020-01-01"
)

type ClientInterface interface {
	ListSites(params *ListSitesParams) (*SiteList, error)

	GetSite(siteId string) (*Site, error)

	ListAccounts(params *ListAccountsParams) (*AccountList, error)

	CreateAccount(body *AccountCreate) (*Account, error)

	GetAccount(accountId string) (*Account, error)

	UpdateAccount(accountId string, body *AccountUpdate) (*Account, error)

	DeactivateAccount(accountId string) (*Account, error)

	GetAccountAcquisition(accountId string) (*AccountAcquisition, error)

	UpdateAccountAcquisition(accountId string, body *AccountAcquisitionUpdatable) (*AccountAcquisition, error)

	RemoveAccountAcquisition(accountId string) (*Empty, error)

	ReactivateAccount(accountId string) (*Account, error)

	GetAccountBalance(accountId string) (*AccountBalance, error)

	GetBillingInfo(accountId string) (*BillingInfo, error)

	UpdateBillingInfo(accountId string, body *BillingInfoCreate) (*BillingInfo, error)

	RemoveBillingInfo(accountId string) (*Empty, error)

	ListBillingInfos(accountId string, params *ListBillingInfosParams) (*BillingInfoList, error)

	CreateBillingInfo(accountId string, body *BillingInfoCreate) (*BillingInfo, error)

	GetABillingInfo(accountId string, billingInfoId string) (*BillingInfo, error)

	UpdateABillingInfo(accountId string, billingInfoId string, body *BillingInfoCreate) (*BillingInfo, error)

	RemoveABillingInfo(accountId string, billingInfoId string) (*Empty, error)

	ListAccountCouponRedemptions(accountId string, params *ListAccountCouponRedemptionsParams) (*CouponRedemptionList, error)

	ListActiveCouponRedemptions(accountId string) (*CouponRedemptionList, error)

	CreateCouponRedemption(accountId string, body *CouponRedemptionCreate) (*CouponRedemption, error)

	RemoveCouponRedemption(accountId string) (*CouponRedemption, error)

	ListAccountCreditPayments(accountId string, params *ListAccountCreditPaymentsParams) (*CreditPaymentList, error)

	ListAccountInvoices(accountId string, params *ListAccountInvoicesParams) (*InvoiceList, error)

	CreateInvoice(accountId string, body *InvoiceCreate) (*InvoiceCollection, error)

	PreviewInvoice(accountId string, body *InvoiceCreate) (*InvoiceCollection, error)

	ListAccountLineItems(accountId string, params *ListAccountLineItemsParams) (*LineItemList, error)

	CreateLineItem(accountId string, body *LineItemCreate) (*LineItem, error)

	ListAccountNotes(accountId string, params *ListAccountNotesParams) (*AccountNoteList, error)

	GetAccountNote(accountId string, accountNoteId string) (*AccountNote, error)

	ListShippingAddresses(accountId string, params *ListShippingAddressesParams) (*ShippingAddressList, error)

	CreateShippingAddress(accountId string, body *ShippingAddressCreate) (*ShippingAddress, error)

	GetShippingAddress(accountId string, shippingAddressId string) (*ShippingAddress, error)

	UpdateShippingAddress(accountId string, shippingAddressId string, body *ShippingAddressUpdate) (*ShippingAddress, error)

	RemoveShippingAddress(accountId string, shippingAddressId string) (*Empty, error)

	ListAccountSubscriptions(accountId string, params *ListAccountSubscriptionsParams) (*SubscriptionList, error)

	ListAccountTransactions(accountId string, params *ListAccountTransactionsParams) (*TransactionList, error)

	ListChildAccounts(accountId string, params *ListChildAccountsParams) (*AccountList, error)

	ListAccountAcquisition(params *ListAccountAcquisitionParams) (*AccountAcquisitionList, error)

	ListCoupons(params *ListCouponsParams) (*CouponList, error)

	CreateCoupon(body *CouponCreate) (*Coupon, error)

	GetCoupon(couponId string) (*Coupon, error)

	UpdateCoupon(couponId string, body *CouponUpdate) (*Coupon, error)

	DeactivateCoupon(couponId string) (*Coupon, error)

	RestoreCoupon(couponId string, body *CouponUpdate) (*Coupon, error)

	ListUniqueCouponCodes(couponId string, params *ListUniqueCouponCodesParams) (*UniqueCouponCodeList, error)

	ListCreditPayments(params *ListCreditPaymentsParams) (*CreditPaymentList, error)

	GetCreditPayment(creditPaymentId string) (*CreditPayment, error)

	ListCustomFieldDefinitions(params *ListCustomFieldDefinitionsParams) (*CustomFieldDefinitionList, error)

	GetCustomFieldDefinition(customFieldDefinitionId string) (*CustomFieldDefinition, error)

	ListItems(params *ListItemsParams) (*ItemList, error)

	CreateItem(body *ItemCreate) (*Item, error)

	GetItem(itemId string) (*Item, error)

	UpdateItem(itemId string, body *ItemUpdate) (*Item, error)

	DeactivateItem(itemId string) (*Item, error)

	ReactivateItem(itemId string) (*Item, error)

	ListMeasuredUnit(params *ListMeasuredUnitParams) (*MeasuredUnitList, error)

	CreateMeasuredUnit(body *MeasuredUnitCreate) (*MeasuredUnit, error)

	GetMeasuredUnit(measuredUnitId string) (*MeasuredUnit, error)

	UpdateMeasuredUnit(measuredUnitId string, body *MeasuredUnitUpdate) (*MeasuredUnit, error)

	RemoveMeasuredUnit(measuredUnitId string) (*MeasuredUnit, error)

	ListInvoices(params *ListInvoicesParams) (*InvoiceList, error)

	GetInvoice(invoiceId string) (*Invoice, error)

	PutInvoice(invoiceId string, body *InvoiceUpdatable) (*Invoice, error)

	CollectInvoice(invoiceId string, params *CollectInvoiceParams) (*Invoice, error)

	FailInvoice(invoiceId string) (*Invoice, error)

	MarkInvoiceSuccessful(invoiceId string) (*Invoice, error)

	ReopenInvoice(invoiceId string) (*Invoice, error)

	VoidInvoice(invoiceId string) (*Invoice, error)

	RecordExternalTransaction(invoiceId string, body *ExternalTransaction) (*Transaction, error)

	ListInvoiceLineItems(invoiceId string, params *ListInvoiceLineItemsParams) (*LineItemList, error)

	ListInvoiceCouponRedemptions(invoiceId string, params *ListInvoiceCouponRedemptionsParams) (*CouponRedemptionList, error)

	ListRelatedInvoices(invoiceId string) (*InvoiceList, error)

	RefundInvoice(invoiceId string, body *InvoiceRefund) (*Invoice, error)

	ListLineItems(params *ListLineItemsParams) (*LineItemList, error)

	GetLineItem(lineItemId string) (*LineItem, error)

	RemoveLineItem(lineItemId string) (*Empty, error)

	ListPlans(params *ListPlansParams) (*PlanList, error)

	CreatePlan(body *PlanCreate) (*Plan, error)

	GetPlan(planId string) (*Plan, error)

	UpdatePlan(planId string, body *PlanUpdate) (*Plan, error)

	RemovePlan(planId string) (*Plan, error)

	ListPlanAddOns(planId string, params *ListPlanAddOnsParams) (*AddOnList, error)

	CreatePlanAddOn(planId string, body *AddOnCreate) (*AddOn, error)

	GetPlanAddOn(planId string, addOnId string) (*AddOn, error)

	UpdatePlanAddOn(planId string, addOnId string, body *AddOnUpdate) (*AddOn, error)

	RemovePlanAddOn(planId string, addOnId string) (*AddOn, error)

	ListAddOns(params *ListAddOnsParams) (*AddOnList, error)

	GetAddOn(addOnId string) (*AddOn, error)

	ListShippingMethods(params *ListShippingMethodsParams) (*ShippingMethodList, error)

	CreateShippingMethod(body *ShippingMethodCreate) (*ShippingMethod, error)

	GetShippingMethod(shippingMethodId string) (*ShippingMethod, error)

	UpdateShippingMethod(shippingMethodId string, body *ShippingMethodUpdate) (*ShippingMethod, error)

	DeactivateShippingMethod(shippingMethodId string) (*ShippingMethod, error)

	ListSubscriptions(params *ListSubscriptionsParams) (*SubscriptionList, error)

	CreateSubscription(body *SubscriptionCreate) (*Subscription, error)

	GetSubscription(subscriptionId string) (*Subscription, error)

	ModifySubscription(subscriptionId string, body *SubscriptionUpdate) (*Subscription, error)

	TerminateSubscription(subscriptionId string, params *TerminateSubscriptionParams) (*Subscription, error)

	CancelSubscription(subscriptionId string, params *CancelSubscriptionParams) (*Subscription, error)

	ReactivateSubscription(subscriptionId string) (*Subscription, error)

	PauseSubscription(subscriptionId string, body *SubscriptionPause) (*Subscription, error)

	ResumeSubscription(subscriptionId string) (*Subscription, error)

	ConvertTrial(subscriptionId string) (*Subscription, error)

	GetSubscriptionChange(subscriptionId string) (*SubscriptionChange, error)

	CreateSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate) (*SubscriptionChange, error)

	RemoveSubscriptionChange(subscriptionId string) (*Empty, error)

	PreviewSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate) (*SubscriptionChange, error)

	ListSubscriptionInvoices(subscriptionId string, params *ListSubscriptionInvoicesParams) (*InvoiceList, error)

	ListSubscriptionLineItems(subscriptionId string, params *ListSubscriptionLineItemsParams) (*LineItemList, error)

	ListSubscriptionCouponRedemptions(subscriptionId string, params *ListSubscriptionCouponRedemptionsParams) (*CouponRedemptionList, error)

	ListUsage(subscriptionId string, addOnId string, params *ListUsageParams) (*UsageList, error)

	CreateUsage(subscriptionId string, addOnId string, body *UsageCreate) (*Usage, error)

	GetUsage(usageId string) (*Usage, error)

	UpdateUsage(usageId string, body *UsageCreate) (*Usage, error)

	RemoveUsage(usageId string) (*Empty, error)

	ListTransactions(params *ListTransactionsParams) (*TransactionList, error)

	GetTransaction(transactionId string) (*Transaction, error)

	GetUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error)

	DeactivateUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error)

	ReactivateUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error)

	CreatePurchase(body *PurchaseCreate) (*InvoiceCollection, error)

	PreviewPurchase(body *PurchaseCreate) (*InvoiceCollection, error)

	GetExportDates() (*ExportDates, error)

	GetExportFiles(exportDate string) (*ExportFiles, error)
}

type ListSitesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// State - Filter by state.
	State *string
}

func (list *ListSitesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListSitesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListSites List sites
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_sites
//
// Returns: A list of sites.
func (c *Client) ListSites(params *ListSitesParams) (*SiteList, error) {
	path, err := c.InterpolatePath("/sites")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewSiteList(c, path), nil
}

// GetSite Fetch a site
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_site
//
// Returns: A site.
func (c *Client) GetSite(siteId string) (*Site, error) {
	path, err := c.InterpolatePath("/sites/{site_id}", siteId)
	if err != nil {
		return nil, err
	}
	result := &Site{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Email - Filter for accounts with this exact email address. A blank value will return accounts with both `null` and `""` email addresses. Note that multiple accounts can share one email address.
	Email *string

	// Subscriber - Filter for accounts with or without a subscription in the `active`,
	// `canceled`, or `future` state.
	Subscriber *bool

	// PastDue - Filter for accounts with an invoice in the `past_due` state.
	PastDue *string
}

func (list *ListAccountsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Email != nil {
		options = append(options, KeyValue{Key: "email", Value: *list.Email})
	}

	if list.Subscriber != nil {
		options = append(options, KeyValue{Key: "subscriber", Value: strconv.FormatBool(*list.Subscriber)})
	}

	if list.PastDue != nil {
		options = append(options, KeyValue{Key: "past_due", Value: *list.PastDue})
	}

	return options
}

// ListAccounts List a site's accounts
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_accounts
//
// Returns: A list of the site's accounts.
func (c *Client) ListAccounts(params *ListAccountsParams) (*AccountList, error) {
	path, err := c.InterpolatePath("/accounts")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewAccountList(c, path), nil
}

// CreateAccount Create an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_account
//
// Returns: An account.
func (c *Client) CreateAccount(body *AccountCreate) (*Account, error) {
	path, err := c.InterpolatePath("/accounts")
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetAccount Fetch an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account
//
// Returns: An account.
func (c *Client) GetAccount(accountId string) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateAccount Modify an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_account
//
// Returns: An account.
func (c *Client) UpdateAccount(accountId string, body *AccountUpdate) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateAccount Deactivate an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_account
//
// Returns: An account.
func (c *Client) DeactivateAccount(accountId string) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetAccountAcquisition Fetch an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account_acquisition
//
// Returns: An account's acquisition data.
func (c *Client) GetAccountAcquisition(accountId string) (*AccountAcquisition, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		return nil, err
	}
	result := &AccountAcquisition{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateAccountAcquisition Update an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_account_acquisition
//
// Returns: An account's updated acquisition data.
func (c *Client) UpdateAccountAcquisition(accountId string, body *AccountAcquisitionUpdatable) (*AccountAcquisition, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		return nil, err
	}
	result := &AccountAcquisition{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveAccountAcquisition Remove an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_account_acquisition
//
// Returns: Acquisition data was succesfully deleted.
func (c *Client) RemoveAccountAcquisition(accountId string) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateAccount Reactivate an inactive account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_account
//
// Returns: An account.
func (c *Client) ReactivateAccount(accountId string) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/reactivate", accountId)
	if err != nil {
		return nil, err
	}
	result := &Account{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetAccountBalance Fetch an account's balance and past due status
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account_balance
//
// Returns: An account's balance.
func (c *Client) GetAccountBalance(accountId string) (*AccountBalance, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/balance", accountId)
	if err != nil {
		return nil, err
	}
	result := &AccountBalance{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetBillingInfo Fetch an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_billing_info
//
// Returns: An account's billing information.
func (c *Client) GetBillingInfo(accountId string) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateBillingInfo Set an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_billing_info
//
// Returns: Updated billing information.
func (c *Client) UpdateBillingInfo(accountId string, body *BillingInfoCreate) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveBillingInfo Remove an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_billing_info
//
// Returns: Billing information deleted
func (c *Client) RemoveBillingInfo(accountId string) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListBillingInfosParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListBillingInfosParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListBillingInfosParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListBillingInfos Get the list of billing information associated with an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_billing_infos
//
// Returns: A list of the the billing information for an account's
func (c *Client) ListBillingInfos(accountId string, params *ListBillingInfosParams) (*BillingInfoList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewBillingInfoList(c, path), nil
}

// CreateBillingInfo Set an account's billing information when the wallet feature is enabled
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_billing_info
//
// Returns: Updated billing information.
func (c *Client) CreateBillingInfo(accountId string, body *BillingInfoCreate) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos", accountId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetABillingInfo Fetch a billing info
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_a_billing_info
//
// Returns: A billing info.
func (c *Client) GetABillingInfo(accountId string, billingInfoId string) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateABillingInfo Update an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_a_billing_info
//
// Returns: Updated billing information.
func (c *Client) UpdateABillingInfo(accountId string, billingInfoId string, body *BillingInfoCreate) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		return nil, err
	}
	result := &BillingInfo{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveABillingInfo Remove an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_a_billing_info
//
// Returns: Billing information deleted
func (c *Client) RemoveABillingInfo(accountId string, billingInfoId string) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountCouponRedemptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListAccountCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountCouponRedemptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListAccountCouponRedemptions Show the coupon redemptions for an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_coupon_redemptions
//
// Returns: A list of the the coupon redemptions on an account.
func (c *Client) ListAccountCouponRedemptions(accountId string, params *ListAccountCouponRedemptionsParams) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewCouponRedemptionList(c, path), nil
}

// ListActiveCouponRedemptions Show the coupon redemptions that are active on an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_active_coupon_redemptions
//
// Returns: Active coupon redemptions on an account.
func (c *Client) ListActiveCouponRedemptions(accountId string) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		return nil, err
	}
	return NewCouponRedemptionList(c, path), nil
}

// CreateCouponRedemption Generate an active coupon redemption on an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_coupon_redemption
//
// Returns: Returns the new coupon redemption.
func (c *Client) CreateCouponRedemption(accountId string, body *CouponRedemptionCreate) (*CouponRedemption, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		return nil, err
	}
	result := &CouponRedemption{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveCouponRedemption Delete the active coupon redemption from an account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_coupon_redemption
//
// Returns: Coupon redemption deleted.
func (c *Client) RemoveCouponRedemption(accountId string) (*CouponRedemption, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		return nil, err
	}
	result := &CouponRedemption{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountCreditPaymentsParams struct {
	Params

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListAccountCreditPaymentsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountCreditPaymentsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListAccountCreditPayments List an account's credit payments
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_credit_payments
//
// Returns: A list of the account's credit payments.
func (c *Client) ListAccountCreditPayments(accountId string, params *ListAccountCreditPaymentsParams) (*CreditPaymentList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/credit_payments", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewCreditPaymentList(c, path), nil
}

type ListAccountInvoicesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type when:
	// - `type=charge`, only charge invoices will be returned.
	// - `type=credit`, only credit invoices will be returned.
	// - `type=non-legacy`, only charge and credit invoices will be returned.
	// - `type=legacy`, only legacy invoices will be returned.
	Type *string
}

func (list *ListAccountInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountInvoicesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListAccountInvoices List an account's invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_invoices
//
// Returns: A list of the account's invoices.
func (c *Client) ListAccountInvoices(accountId string, params *ListAccountInvoicesParams) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewInvoiceList(c, path), nil
}

// CreateInvoice Create an invoice for pending line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_invoice
//
// Returns: Returns the new invoices.
func (c *Client) CreateInvoice(accountId string, body *InvoiceCreate) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PreviewInvoice Preview new invoice for pending line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/preview_invoice
//
// Returns: Returns the invoice previews.
func (c *Client) PreviewInvoice(accountId string, body *InvoiceCreate) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices/preview", accountId)
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListAccountLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListAccountLineItems List an account's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_line_items
//
// Returns: A list of the account's line items.
func (c *Client) ListAccountLineItems(accountId string, params *ListAccountLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewLineItemList(c, path), nil
}

// CreateLineItem Create a new line item for the account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_line_item
//
// Returns: Returns the new line item.
func (c *Client) CreateLineItem(accountId string, body *LineItemCreate) (*LineItem, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	if err != nil {
		return nil, err
	}
	result := &LineItem{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountNotesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string
}

func (list *ListAccountNotesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountNotesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	return options
}

// ListAccountNotes Fetch a list of an account's notes
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_notes
//
// Returns: A list of an account's notes.
func (c *Client) ListAccountNotes(accountId string, params *ListAccountNotesParams) (*AccountNoteList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/notes", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewAccountNoteList(c, path), nil
}

// GetAccountNote Fetch an account note
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_account_note
//
// Returns: An account note.
func (c *Client) GetAccountNote(accountId string, accountNoteId string) (*AccountNote, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/notes/{account_note_id}", accountId, accountNoteId)
	if err != nil {
		return nil, err
	}
	result := &AccountNote{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListShippingAddressesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListShippingAddressesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListShippingAddressesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListShippingAddresses Fetch a list of an account's shipping addresses
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_shipping_addresses
//
// Returns: A list of an account's shipping addresses.
func (c *Client) ListShippingAddresses(accountId string, params *ListShippingAddressesParams) (*ShippingAddressList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewShippingAddressList(c, path), nil
}

// CreateShippingAddress Create a new shipping address for the account
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_shipping_address
//
// Returns: Returns the new shipping address.
func (c *Client) CreateShippingAddress(accountId string, body *ShippingAddressCreate) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	if err != nil {
		return nil, err
	}
	result := &ShippingAddress{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetShippingAddress Fetch an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_shipping_address
//
// Returns: A shipping address.
func (c *Client) GetShippingAddress(accountId string, shippingAddressId string) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		return nil, err
	}
	result := &ShippingAddress{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateShippingAddress Update an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_shipping_address
//
// Returns: The updated shipping address.
func (c *Client) UpdateShippingAddress(accountId string, shippingAddressId string, body *ShippingAddressUpdate) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		return nil, err
	}
	result := &ShippingAddress{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveShippingAddress Remove an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_shipping_address
//
// Returns: Shipping address deleted.
func (c *Client) RemoveShippingAddress(accountId string, shippingAddressId string) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAccountSubscriptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	// - When `state=active`, `state=canceled`, `state=expired`, or `state=future`, subscriptions with states that match the query and only those subscriptions will be returned.
	// - When `state=in_trial`, only subscriptions that have a trial_started_at date earlier than now and a trial_ends_at date later than now will be returned.
	// - When `state=live`, only subscriptions that are in an active, canceled, or future state or are in trial will be returned.
	State *string
}

func (list *ListAccountSubscriptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountSubscriptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListAccountSubscriptions List an account's subscriptions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_subscriptions
//
// Returns: A list of the account's subscriptions.
func (c *Client) ListAccountSubscriptions(accountId string, params *ListAccountSubscriptionsParams) (*SubscriptionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/subscriptions", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewSubscriptionList(c, path), nil
}

type ListAccountTransactionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type field. The value `payment` will return both `purchase` and `capture` transactions.
	Type *string

	// Success - Filter by success field.
	Success *string
}

func (list *ListAccountTransactionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountTransactionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	if list.Success != nil {
		options = append(options, KeyValue{Key: "success", Value: *list.Success})
	}

	return options
}

// ListAccountTransactions List an account's transactions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_transactions
//
// Returns: A list of the account's transactions.
func (c *Client) ListAccountTransactions(accountId string, params *ListAccountTransactionsParams) (*TransactionList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/transactions", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewTransactionList(c, path), nil
}

type ListChildAccountsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Email - Filter for accounts with this exact email address. A blank value will return accounts with both `null` and `""` email addresses. Note that multiple accounts can share one email address.
	Email *string

	// Subscriber - Filter for accounts with or without a subscription in the `active`,
	// `canceled`, or `future` state.
	Subscriber *bool

	// PastDue - Filter for accounts with an invoice in the `past_due` state.
	PastDue *string
}

func (list *ListChildAccountsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListChildAccountsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Email != nil {
		options = append(options, KeyValue{Key: "email", Value: *list.Email})
	}

	if list.Subscriber != nil {
		options = append(options, KeyValue{Key: "subscriber", Value: strconv.FormatBool(*list.Subscriber)})
	}

	if list.PastDue != nil {
		options = append(options, KeyValue{Key: "past_due", Value: *list.PastDue})
	}

	return options
}

// ListChildAccounts List an account's child accounts
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_child_accounts
//
// Returns: A list of an account's child accounts.
func (c *Client) ListChildAccounts(accountId string, params *ListChildAccountsParams) (*AccountList, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/accounts", accountId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewAccountList(c, path), nil
}

type ListAccountAcquisitionParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListAccountAcquisitionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAccountAcquisitionParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListAccountAcquisition List a site's account acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_account_acquisition
//
// Returns: A list of the site's account acquisition data.
func (c *Client) ListAccountAcquisition(params *ListAccountAcquisitionParams) (*AccountAcquisitionList, error) {
	path, err := c.InterpolatePath("/acquisitions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewAccountAcquisitionList(c, path), nil
}

type ListCouponsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListCouponsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListCouponsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListCoupons List a site's coupons
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_coupons
//
// Returns: A list of the site's coupons.
func (c *Client) ListCoupons(params *ListCouponsParams) (*CouponList, error) {
	path, err := c.InterpolatePath("/coupons")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewCouponList(c, path), nil
}

// CreateCoupon Create a new coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_coupon
//
// Returns: A new coupon.
func (c *Client) CreateCoupon(body *CouponCreate) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons")
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetCoupon Fetch a coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_coupon
//
// Returns: A coupon.
func (c *Client) GetCoupon(couponId string) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateCoupon Update an active coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_coupon
//
// Returns: The updated coupon.
func (c *Client) UpdateCoupon(couponId string, body *CouponUpdate) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateCoupon Expire a coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_coupon
//
// Returns: The expired Coupon
func (c *Client) DeactivateCoupon(couponId string) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RestoreCoupon Restore an inactive coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/restore_coupon
//
// Returns: The restored coupon.
func (c *Client) RestoreCoupon(couponId string, body *CouponUpdate) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}/restore", couponId)
	if err != nil {
		return nil, err
	}
	result := &Coupon{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListUniqueCouponCodesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListUniqueCouponCodesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListUniqueCouponCodesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListUniqueCouponCodes List unique coupon codes associated with a bulk coupon
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_unique_coupon_codes
//
// Returns: A list of unique coupon codes that were generated
func (c *Client) ListUniqueCouponCodes(couponId string, params *ListUniqueCouponCodesParams) (*UniqueCouponCodeList, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}/unique_coupon_codes", couponId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewUniqueCouponCodeList(c, path), nil
}

type ListCreditPaymentsParams struct {
	Params

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListCreditPaymentsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListCreditPaymentsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListCreditPayments List a site's credit payments
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_credit_payments
//
// Returns: A list of the site's credit payments.
func (c *Client) ListCreditPayments(params *ListCreditPaymentsParams) (*CreditPaymentList, error) {
	path, err := c.InterpolatePath("/credit_payments")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewCreditPaymentList(c, path), nil
}

// GetCreditPayment Fetch a credit payment
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_credit_payment
//
// Returns: A credit payment.
func (c *Client) GetCreditPayment(creditPaymentId string) (*CreditPayment, error) {
	path, err := c.InterpolatePath("/credit_payments/{credit_payment_id}", creditPaymentId)
	if err != nil {
		return nil, err
	}
	result := &CreditPayment{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListCustomFieldDefinitionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// RelatedType - Filter by related type.
	RelatedType *string
}

func (list *ListCustomFieldDefinitionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListCustomFieldDefinitionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.RelatedType != nil {
		options = append(options, KeyValue{Key: "related_type", Value: *list.RelatedType})
	}

	return options
}

// ListCustomFieldDefinitions List a site's custom field definitions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_custom_field_definitions
//
// Returns: A list of the site's custom field definitions.
func (c *Client) ListCustomFieldDefinitions(params *ListCustomFieldDefinitionsParams) (*CustomFieldDefinitionList, error) {
	path, err := c.InterpolatePath("/custom_field_definitions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewCustomFieldDefinitionList(c, path), nil
}

// GetCustomFieldDefinition Fetch an custom field definition
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_custom_field_definition
//
// Returns: An custom field definition.
func (c *Client) GetCustomFieldDefinition(customFieldDefinitionId string) (*CustomFieldDefinition, error) {
	path, err := c.InterpolatePath("/custom_field_definitions/{custom_field_definition_id}", customFieldDefinitionId)
	if err != nil {
		return nil, err
	}
	result := &CustomFieldDefinition{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListItems List a site's items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_items
//
// Returns: A list of the site's items.
func (c *Client) ListItems(params *ListItemsParams) (*ItemList, error) {
	path, err := c.InterpolatePath("/items")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewItemList(c, path), nil
}

// CreateItem Create a new item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_item
//
// Returns: A new item.
func (c *Client) CreateItem(body *ItemCreate) (*Item, error) {
	path, err := c.InterpolatePath("/items")
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetItem Fetch an item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_item
//
// Returns: An item.
func (c *Client) GetItem(itemId string) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateItem Update an active item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_item
//
// Returns: The updated item.
func (c *Client) UpdateItem(itemId string, body *ItemUpdate) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateItem Deactivate an item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_item
//
// Returns: An item.
func (c *Client) DeactivateItem(itemId string) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateItem Reactivate an inactive item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_item
//
// Returns: An item.
func (c *Client) ReactivateItem(itemId string) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}/reactivate", itemId)
	if err != nil {
		return nil, err
	}
	result := &Item{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListMeasuredUnitParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListMeasuredUnitParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListMeasuredUnitParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListMeasuredUnit List a site's measured units
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_measured_unit
//
// Returns: A list of the site's measured units.
func (c *Client) ListMeasuredUnit(params *ListMeasuredUnitParams) (*MeasuredUnitList, error) {
	path, err := c.InterpolatePath("/measured_units")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewMeasuredUnitList(c, path), nil
}

// CreateMeasuredUnit Create a new measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_measured_unit
//
// Returns: A new measured unit.
func (c *Client) CreateMeasuredUnit(body *MeasuredUnitCreate) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units")
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetMeasuredUnit Fetch a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_measured_unit
//
// Returns: An item.
func (c *Client) GetMeasuredUnit(measuredUnitId string) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateMeasuredUnit Update a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_measured_unit
//
// Returns: The updated measured_unit.
func (c *Client) UpdateMeasuredUnit(measuredUnitId string, body *MeasuredUnitUpdate) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveMeasuredUnit Remove a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_measured_unit
//
// Returns: A measured unit.
func (c *Client) RemoveMeasuredUnit(measuredUnitId string) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		return nil, err
	}
	result := &MeasuredUnit{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListInvoicesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type when:
	// - `type=charge`, only charge invoices will be returned.
	// - `type=credit`, only credit invoices will be returned.
	// - `type=non-legacy`, only charge and credit invoices will be returned.
	// - `type=legacy`, only legacy invoices will be returned.
	Type *string
}

func (list *ListInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListInvoicesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListInvoices List a site's invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_invoices
//
// Returns: A list of the site's invoices.
func (c *Client) ListInvoices(params *ListInvoicesParams) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/invoices")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewInvoiceList(c, path), nil
}

// GetInvoice Fetch an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_invoice
//
// Returns: An invoice.
func (c *Client) GetInvoice(invoiceId string) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PutInvoice Update an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/put_invoice
//
// Returns: An invoice.
func (c *Client) PutInvoice(invoiceId string, body *InvoiceUpdatable) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type CollectInvoiceParams struct {
	Params

	// Body - The body of the request.
	Body *InvoiceCollect
}

func (list *CollectInvoiceParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// CollectInvoice Collect a pending or past due, automatic invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/collect_invoice
//
// Returns: The updated invoice.
func (c *Client) CollectInvoice(invoiceId string, params *CollectInvoiceParams) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/collect", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, params, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// FailInvoice Mark an open invoice as failed
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/fail_invoice
//
// Returns: The updated invoice.
func (c *Client) FailInvoice(invoiceId string) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/mark_failed", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// MarkInvoiceSuccessful Mark an open invoice as successful
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/mark_invoice_successful
//
// Returns: The updated invoice.
func (c *Client) MarkInvoiceSuccessful(invoiceId string) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/mark_successful", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReopenInvoice Reopen a closed, manual invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reopen_invoice
//
// Returns: The updated invoice.
func (c *Client) ReopenInvoice(invoiceId string) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/reopen", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// VoidInvoice Void a credit invoice.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/void_invoice
//
// Returns: The updated invoice.
func (c *Client) VoidInvoice(invoiceId string) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/void", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RecordExternalTransaction Record an external payment for a manual invoices.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/record_external_transaction
//
// Returns: The recorded transaction.
func (c *Client) RecordExternalTransaction(invoiceId string, body *ExternalTransaction) (*Transaction, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/transactions", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Transaction{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListInvoiceLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListInvoiceLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListInvoiceLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListInvoiceLineItems List an invoice's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_invoice_line_items
//
// Returns: A list of the invoice's line items.
func (c *Client) ListInvoiceLineItems(invoiceId string, params *ListInvoiceLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/line_items", invoiceId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewLineItemList(c, path), nil
}

type ListInvoiceCouponRedemptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListInvoiceCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListInvoiceCouponRedemptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListInvoiceCouponRedemptions Show the coupon redemptions applied to an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_invoice_coupon_redemptions
//
// Returns: A list of the the coupon redemptions associated with the invoice.
func (c *Client) ListInvoiceCouponRedemptions(invoiceId string, params *ListInvoiceCouponRedemptionsParams) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/coupon_redemptions", invoiceId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewCouponRedemptionList(c, path), nil
}

// ListRelatedInvoices List an invoice's related credit or charge invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_related_invoices
//
// Returns: A list of the credit or charge invoices associated with the invoice.
func (c *Client) ListRelatedInvoices(invoiceId string) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/related_invoices", invoiceId)
	if err != nil {
		return nil, err
	}
	return NewInvoiceList(c, path), nil
}

// RefundInvoice Refund an invoice
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/refund_invoice
//
// Returns: Returns the new credit invoice.
func (c *Client) RefundInvoice(invoiceId string, body *InvoiceRefund) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/refund", invoiceId)
	if err != nil {
		return nil, err
	}
	result := &Invoice{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListLineItems List a site's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_line_items
//
// Returns: A list of the site's line items.
func (c *Client) ListLineItems(params *ListLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/line_items")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewLineItemList(c, path), nil
}

// GetLineItem Fetch a line item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_line_item
//
// Returns: A line item.
func (c *Client) GetLineItem(lineItemId string) (*LineItem, error) {
	path, err := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	if err != nil {
		return nil, err
	}
	result := &LineItem{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveLineItem Delete an uninvoiced line item
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_line_item
//
// Returns: Line item deleted.
func (c *Client) RemoveLineItem(lineItemId string) (*Empty, error) {
	path, err := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListPlansParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListPlansParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListPlansParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListPlans List a site's plans
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_plans
//
// Returns: A list of plans.
func (c *Client) ListPlans(params *ListPlansParams) (*PlanList, error) {
	path, err := c.InterpolatePath("/plans")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewPlanList(c, path), nil
}

// CreatePlan Create a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_plan
//
// Returns: A plan.
func (c *Client) CreatePlan(body *PlanCreate) (*Plan, error) {
	path, err := c.InterpolatePath("/plans")
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetPlan Fetch a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_plan
//
// Returns: A plan.
func (c *Client) GetPlan(planId string) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdatePlan Update a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_plan
//
// Returns: A plan.
func (c *Client) UpdatePlan(planId string, body *PlanUpdate) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemovePlan Remove a plan
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_plan
//
// Returns: Plan deleted
func (c *Client) RemovePlan(planId string) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		return nil, err
	}
	result := &Plan{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListPlanAddOnsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListPlanAddOnsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListPlanAddOnsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListPlanAddOns List a plan's add-ons
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_plan_add_ons
//
// Returns: A list of add-ons.
func (c *Client) ListPlanAddOns(planId string, params *ListPlanAddOnsParams) (*AddOnList, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewAddOnList(c, path), nil
}

// CreatePlanAddOn Create an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_plan_add_on
//
// Returns: An add-on.
func (c *Client) CreatePlanAddOn(planId string, body *AddOnCreate) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetPlanAddOn Fetch a plan's add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_plan_add_on
//
// Returns: An add-on.
func (c *Client) GetPlanAddOn(planId string, addOnId string) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdatePlanAddOn Update an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_plan_add_on
//
// Returns: An add-on.
func (c *Client) UpdatePlanAddOn(planId string, addOnId string, body *AddOnUpdate) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemovePlanAddOn Remove an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_plan_add_on
//
// Returns: Add-on deleted
func (c *Client) RemovePlanAddOn(planId string, addOnId string) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListAddOnsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	State *string
}

func (list *ListAddOnsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListAddOnsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListAddOns List a site's add-ons
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_add_ons
//
// Returns: A list of add-ons.
func (c *Client) ListAddOns(params *ListAddOnsParams) (*AddOnList, error) {
	path, err := c.InterpolatePath("/add_ons")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewAddOnList(c, path), nil
}

// GetAddOn Fetch an add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_add_on
//
// Returns: An add-on.
func (c *Client) GetAddOn(addOnId string) (*AddOn, error) {
	path, err := c.InterpolatePath("/add_ons/{add_on_id}", addOnId)
	if err != nil {
		return nil, err
	}
	result := &AddOn{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListShippingMethodsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListShippingMethodsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListShippingMethodsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListShippingMethods List a site's shipping methods
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_shipping_methods
//
// Returns: A list of the site's shipping methods.
func (c *Client) ListShippingMethods(params *ListShippingMethodsParams) (*ShippingMethodList, error) {
	path, err := c.InterpolatePath("/shipping_methods")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewShippingMethodList(c, path), nil
}

// CreateShippingMethod Create a new shipping method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_shipping_method
//
// Returns: A new shipping method.
func (c *Client) CreateShippingMethod(body *ShippingMethodCreate) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods")
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetShippingMethod Fetch a shipping method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_shipping_method
//
// Returns: A shipping method.
func (c *Client) GetShippingMethod(shippingMethodId string) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateShippingMethod Update an active Shipping Method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_shipping_method
//
// Returns: The updated shipping method.
func (c *Client) UpdateShippingMethod(shippingMethodId string, body *ShippingMethodUpdate) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateShippingMethod Deactivate a shipping method
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_shipping_method
//
// Returns: A shipping method.
func (c *Client) DeactivateShippingMethod(shippingMethodId string) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		return nil, err
	}
	result := &ShippingMethod{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListSubscriptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// State - Filter by state.
	// - When `state=active`, `state=canceled`, `state=expired`, or `state=future`, subscriptions with states that match the query and only those subscriptions will be returned.
	// - When `state=in_trial`, only subscriptions that have a trial_started_at date earlier than now and a trial_ends_at date later than now will be returned.
	// - When `state=live`, only subscriptions that are in an active, canceled, or future state or are in trial will be returned.
	State *string
}

func (list *ListSubscriptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListSubscriptions List a site's subscriptions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscriptions
//
// Returns: A list of the site's subscriptions.
func (c *Client) ListSubscriptions(params *ListSubscriptionsParams) (*SubscriptionList, error) {
	path, err := c.InterpolatePath("/subscriptions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewSubscriptionList(c, path), nil
}

// CreateSubscription Create a new subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_subscription
//
// Returns: A subscription.
func (c *Client) CreateSubscription(body *SubscriptionCreate) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions")
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetSubscription Fetch a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_subscription
//
// Returns: A subscription.
func (c *Client) GetSubscription(subscriptionId string) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ModifySubscription Modify a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/modify_subscription
//
// Returns: A subscription.
func (c *Client) ModifySubscription(subscriptionId string, body *SubscriptionUpdate) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type TerminateSubscriptionParams struct {
	Params

	// Refund - The type of refund to perform:
	// * `full` - Performs a full refund of the last invoice for the current subscription term.
	// * `partial` - Prorates a refund based on the amount of time remaining in the current bill cycle.
	// * `none` - Terminates the subscription without a refund.
	// In the event that the most recent invoice is a $0 invoice paid entirely by credit, Recurly will apply the credit back to the customer’s account.
	// You may also terminate a subscription with no refund and then manually refund specific invoices.
	Refund *string
}

func (list *TerminateSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *TerminateSubscriptionParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Refund != nil {
		options = append(options, KeyValue{Key: "refund", Value: *list.Refund})
	}

	return options
}

// TerminateSubscription Terminate a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/terminate_subscription
//
// Returns: An expired subscription.
func (c *Client) TerminateSubscription(subscriptionId string, params *TerminateSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodDelete, path, params, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type CancelSubscriptionParams struct {
	Params

	// Body - The body of the request.
	Body *SubscriptionCancel
}

func (list *CancelSubscriptionParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

// CancelSubscription Cancel a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/cancel_subscription
//
// Returns: A canceled or failed subscription.
func (c *Client) CancelSubscription(subscriptionId string, params *CancelSubscriptionParams) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/cancel", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, params, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateSubscription Reactivate a canceled subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_subscription
//
// Returns: An active subscription.
func (c *Client) ReactivateSubscription(subscriptionId string) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/reactivate", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PauseSubscription Pause subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/pause_subscription
//
// Returns: A subscription.
func (c *Client) PauseSubscription(subscriptionId string, body *SubscriptionPause) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/pause", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ResumeSubscription Resume subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/resume_subscription
//
// Returns: A subscription.
func (c *Client) ResumeSubscription(subscriptionId string) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/resume", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ConvertTrial Convert trial subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/convert_trial
//
// Returns: A subscription.
func (c *Client) ConvertTrial(subscriptionId string) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/convert_trial", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Subscription{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetSubscriptionChange Fetch a subscription's pending change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_subscription_change
//
// Returns: A subscription's pending change.
func (c *Client) GetSubscriptionChange(subscriptionId string) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &SubscriptionChange{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// CreateSubscriptionChange Create a new subscription change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_subscription_change
//
// Returns: A subscription change.
func (c *Client) CreateSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &SubscriptionChange{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveSubscriptionChange Delete the pending subscription change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_subscription_change
//
// Returns: Subscription change was deleted.
func (c *Client) RemoveSubscriptionChange(subscriptionId string) (*Empty, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PreviewSubscriptionChange Preview a new subscription change
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/preview_subscription_change
//
// Returns: A subscription change.
func (c *Client) PreviewSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change/preview", subscriptionId)
	if err != nil {
		return nil, err
	}
	result := &SubscriptionChange{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListSubscriptionInvoicesParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type when:
	// - `type=charge`, only charge invoices will be returned.
	// - `type=credit`, only credit invoices will be returned.
	// - `type=non-legacy`, only charge and credit invoices will be returned.
	// - `type=legacy`, only legacy invoices will be returned.
	Type *string
}

func (list *ListSubscriptionInvoicesParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionInvoicesParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListSubscriptionInvoices List a subscription's invoices
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscription_invoices
//
// Returns: A list of the subscription's invoices.
func (c *Client) ListSubscriptionInvoices(subscriptionId string, params *ListSubscriptionInvoicesParams) (*InvoiceList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/invoices", subscriptionId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewInvoiceList(c, path), nil
}

type ListSubscriptionLineItemsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Original - Filter by original field.
	Original *string

	// State - Filter by state field.
	State *string

	// Type - Filter by type field.
	Type *string
}

func (list *ListSubscriptionLineItemsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionLineItemsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Original != nil {
		options = append(options, KeyValue{Key: "original", Value: *list.Original})
	}

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	return options
}

// ListSubscriptionLineItems List a subscription's line items
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscription_line_items
//
// Returns: A list of the subscription's line items.
func (c *Client) ListSubscriptionLineItems(subscriptionId string, params *ListSubscriptionLineItemsParams) (*LineItemList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/line_items", subscriptionId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewLineItemList(c, path), nil
}

type ListSubscriptionCouponRedemptionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time
}

func (list *ListSubscriptionCouponRedemptionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListSubscriptionCouponRedemptionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

// ListSubscriptionCouponRedemptions Show the coupon redemptions for a subscription
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_subscription_coupon_redemptions
//
// Returns: A list of the the coupon redemptions on a subscription.
func (c *Client) ListSubscriptionCouponRedemptions(subscriptionId string, params *ListSubscriptionCouponRedemptionsParams) (*CouponRedemptionList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/coupon_redemptions", subscriptionId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewCouponRedemptionList(c, path), nil
}

type ListUsageParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `usage_timestamp` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=usage_timestamp` or `sort=recorded_timestamp`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=usage_timestamp` or `sort=recorded_timestamp`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// BillingStatus - Filter by usage record's billing status
	BillingStatus *string
}

func (list *ListUsageParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListUsageParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.BillingStatus != nil {
		options = append(options, KeyValue{Key: "billing_status", Value: *list.BillingStatus})
	}

	return options
}

// ListUsage List a subscription add-on's usage records
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_usage
//
// Returns: A list of the subscription add-on's usage records.
func (c *Client) ListUsage(subscriptionId string, addOnId string, params *ListUsageParams) (*UsageList, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/add_ons/{add_on_id}/usage", subscriptionId, addOnId)
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewUsageList(c, path), nil
}

// CreateUsage Log a usage record on this subscription add-on
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_usage
//
// Returns: The created usage record.
func (c *Client) CreateUsage(subscriptionId string, addOnId string, body *UsageCreate) (*Usage, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/add_ons/{add_on_id}/usage", subscriptionId, addOnId)
	if err != nil {
		return nil, err
	}
	result := &Usage{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetUsage Get a usage record
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_usage
//
// Returns: The usage record.
func (c *Client) GetUsage(usageId string) (*Usage, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		return nil, err
	}
	result := &Usage{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateUsage Update a usage record
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/update_usage
//
// Returns: The updated usage record.
func (c *Client) UpdateUsage(usageId string, body *UsageCreate) (*Usage, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		return nil, err
	}
	result := &Usage{}
	err = c.Call(http.MethodPut, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveUsage Delete a usage record.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/remove_usage
//
// Returns: Usage was successfully deleted.
func (c *Client) RemoveUsage(usageId string) (*Empty, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListTransactionsParams struct {
	Params

	// Ids - Filter results by their IDs. Up to 200 IDs can be passed at once using
	// commas as separators, e.g. `ids=h1at4d57xlmy,gyqgg0d3v9n1,jrsm5b4yefg6`.
	// **Important notes:**
	// * The `ids` parameter cannot be used with any other ordering or filtering
	//   parameters (`limit`, `order`, `sort`, `begin_time`, `end_time`, etc)
	// * Invalid or unknown IDs will be ignored, so you should check that the
	//   results correspond to your request.
	// * Records are returned in an arbitrary order. Since results are all
	//   returned at once you can sort the records yourself.
	Ids []string

	// Limit - Limit number of records 1-200.
	Limit *int

	// Order - Sort order.
	Order *string

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string

	// BeginTime - Inclusively filter by begin_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	BeginTime *time.Time

	// EndTime - Inclusively filter by end_time when `sort=created_at` or `sort=updated_at`.
	// **Note:** this value is an ISO8601 timestamp. A partial timestamp that does not include a time zone will default to UTC.
	EndTime *time.Time

	// Type - Filter by type field. The value `payment` will return both `purchase` and `capture` transactions.
	Type *string

	// Success - Filter by success field.
	Success *string
}

func (list *ListTransactionsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListTransactionsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Ids != nil {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.Ids, ",")})
	}

	if list.Limit != nil {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(*list.Limit)})
	}

	if list.Order != nil {
		options = append(options, KeyValue{Key: "order", Value: *list.Order})
	}

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}

	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	if list.Type != nil {
		options = append(options, KeyValue{Key: "type", Value: *list.Type})
	}

	if list.Success != nil {
		options = append(options, KeyValue{Key: "success", Value: *list.Success})
	}

	return options
}

// ListTransactions List a site's transactions
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/list_transactions
//
// Returns: A list of the site's transactions.
func (c *Client) ListTransactions(params *ListTransactionsParams) (*TransactionList, error) {
	path, err := c.InterpolatePath("/transactions")
	if err != nil {
		return nil, err
	}
	path = BuildUrl(path, params)
	return NewTransactionList(c, path), nil
}

// GetTransaction Fetch a transaction
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_transaction
//
// Returns: A transaction.
func (c *Client) GetTransaction(transactionId string) (*Transaction, error) {
	path, err := c.InterpolatePath("/transactions/{transaction_id}", transactionId)
	if err != nil {
		return nil, err
	}
	result := &Transaction{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetUniqueCouponCode Fetch a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) GetUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	if err != nil {
		return nil, err
	}
	result := &UniqueCouponCode{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateUniqueCouponCode Deactivate a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/deactivate_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) DeactivateUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	if err != nil {
		return nil, err
	}
	result := &UniqueCouponCode{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateUniqueCouponCode Restore a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/reactivate_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) ReactivateUniqueCouponCode(uniqueCouponCodeId string) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}/restore", uniqueCouponCodeId)
	if err != nil {
		return nil, err
	}
	result := &UniqueCouponCode{}
	err = c.Call(http.MethodPut, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// CreatePurchase Create a new purchase
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/create_purchase
//
// Returns: Returns the new invoices
func (c *Client) CreatePurchase(body *PurchaseCreate) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/purchases")
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PreviewPurchase Preview a new purchase
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/preview_purchase
//
// Returns: Returns preview of the new invoices
func (c *Client) PreviewPurchase(body *PurchaseCreate) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/purchases/preview")
	if err != nil {
		return nil, err
	}
	result := &InvoiceCollection{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetExportDates List the dates that have an available export to download.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_export_dates
//
// Returns: Returns a list of dates.
func (c *Client) GetExportDates() (*ExportDates, error) {
	path, err := c.InterpolatePath("/export_dates")
	if err != nil {
		return nil, err
	}
	result := &ExportDates{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetExportFiles List of the export files that are available to download.
//
// API Documentation: https://developers.recurly.com/api/v2020-01-01#operation/get_export_files
//
// Returns: Returns a list of export files to download.
func (c *Client) GetExportFiles(exportDate string) (*ExportFiles, error) {
	path, err := c.InterpolatePath("/export_dates/{export_date}/export_files", exportDate)
	if err != nil {
		return nil, err
	}
	result := &ExportFiles{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}
