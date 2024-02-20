// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// APIVersion is the current Recurly API Version
	APIVersion = "v2019-10-10"
)

type ClientInterface interface {
	ListSites(params *ListSitesParams, opts ...Option) *SiteList

	GetSite(siteId string, opts ...Option) (*Site, error)
	GetSiteWithContext(ctx context.Context, siteId string, opts ...Option) (*Site, error)

	ListAccounts(params *ListAccountsParams, opts ...Option) *AccountList

	CreateAccount(body *AccountCreate, opts ...Option) (*Account, error)
	CreateAccountWithContext(ctx context.Context, body *AccountCreate, opts ...Option) (*Account, error)

	GetAccount(accountId string, opts ...Option) (*Account, error)
	GetAccountWithContext(ctx context.Context, accountId string, opts ...Option) (*Account, error)

	UpdateAccount(accountId string, body *AccountUpdate, opts ...Option) (*Account, error)
	UpdateAccountWithContext(ctx context.Context, accountId string, body *AccountUpdate, opts ...Option) (*Account, error)

	DeactivateAccount(accountId string, opts ...Option) (*Account, error)
	DeactivateAccountWithContext(ctx context.Context, accountId string, opts ...Option) (*Account, error)

	GetAccountAcquisition(accountId string, opts ...Option) (*AccountAcquisition, error)
	GetAccountAcquisitionWithContext(ctx context.Context, accountId string, opts ...Option) (*AccountAcquisition, error)

	UpdateAccountAcquisition(accountId string, body *AccountAcquisitionUpdatable, opts ...Option) (*AccountAcquisition, error)
	UpdateAccountAcquisitionWithContext(ctx context.Context, accountId string, body *AccountAcquisitionUpdatable, opts ...Option) (*AccountAcquisition, error)

	RemoveAccountAcquisition(accountId string, opts ...Option) (*Empty, error)
	RemoveAccountAcquisitionWithContext(ctx context.Context, accountId string, opts ...Option) (*Empty, error)

	ReactivateAccount(accountId string, opts ...Option) (*Account, error)
	ReactivateAccountWithContext(ctx context.Context, accountId string, opts ...Option) (*Account, error)

	GetAccountBalance(accountId string, opts ...Option) (*AccountBalance, error)
	GetAccountBalanceWithContext(ctx context.Context, accountId string, opts ...Option) (*AccountBalance, error)

	GetBillingInfo(accountId string, opts ...Option) (*BillingInfo, error)
	GetBillingInfoWithContext(ctx context.Context, accountId string, opts ...Option) (*BillingInfo, error)

	UpdateBillingInfo(accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error)
	UpdateBillingInfoWithContext(ctx context.Context, accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error)

	RemoveBillingInfo(accountId string, opts ...Option) (*Empty, error)
	RemoveBillingInfoWithContext(ctx context.Context, accountId string, opts ...Option) (*Empty, error)

	VerifyBillingInfo(accountId string, params *VerifyBillingInfoParams, opts ...Option) (*Transaction, error)
	VerifyBillingInfoWithContext(ctx context.Context, accountId string, params *VerifyBillingInfoParams, opts ...Option) (*Transaction, error)

	ListBillingInfos(accountId string, params *ListBillingInfosParams, opts ...Option) *BillingInfoList

	CreateBillingInfo(accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error)
	CreateBillingInfoWithContext(ctx context.Context, accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error)

	GetABillingInfo(accountId string, billingInfoId string, opts ...Option) (*BillingInfo, error)
	GetABillingInfoWithContext(ctx context.Context, accountId string, billingInfoId string, opts ...Option) (*BillingInfo, error)

	UpdateABillingInfo(accountId string, billingInfoId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error)
	UpdateABillingInfoWithContext(ctx context.Context, accountId string, billingInfoId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error)

	RemoveABillingInfo(accountId string, billingInfoId string, opts ...Option) (*Empty, error)
	RemoveABillingInfoWithContext(ctx context.Context, accountId string, billingInfoId string, opts ...Option) (*Empty, error)

	ListAccountCouponRedemptions(accountId string, params *ListAccountCouponRedemptionsParams, opts ...Option) *CouponRedemptionList

	GetActiveCouponRedemption(accountId string, opts ...Option) (*CouponRedemption, error)
	GetActiveCouponRedemptionWithContext(ctx context.Context, accountId string, opts ...Option) (*CouponRedemption, error)

	CreateCouponRedemption(accountId string, body *CouponRedemptionCreate, opts ...Option) (*CouponRedemption, error)
	CreateCouponRedemptionWithContext(ctx context.Context, accountId string, body *CouponRedemptionCreate, opts ...Option) (*CouponRedemption, error)

	RemoveCouponRedemption(accountId string, opts ...Option) (*CouponRedemption, error)
	RemoveCouponRedemptionWithContext(ctx context.Context, accountId string, opts ...Option) (*CouponRedemption, error)

	ListAccountCreditPayments(accountId string, params *ListAccountCreditPaymentsParams, opts ...Option) *CreditPaymentList

	ListAccountInvoices(accountId string, params *ListAccountInvoicesParams, opts ...Option) *InvoiceList

	CreateInvoice(accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error)
	CreateInvoiceWithContext(ctx context.Context, accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error)

	PreviewInvoice(accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error)
	PreviewInvoiceWithContext(ctx context.Context, accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error)

	ListAccountLineItems(accountId string, params *ListAccountLineItemsParams, opts ...Option) *LineItemList

	CreateLineItem(accountId string, body *LineItemCreate, opts ...Option) (*LineItem, error)
	CreateLineItemWithContext(ctx context.Context, accountId string, body *LineItemCreate, opts ...Option) (*LineItem, error)

	ListAccountNotes(accountId string, params *ListAccountNotesParams, opts ...Option) *AccountNoteList

	GetAccountNote(accountId string, accountNoteId string, opts ...Option) (*AccountNote, error)
	GetAccountNoteWithContext(ctx context.Context, accountId string, accountNoteId string, opts ...Option) (*AccountNote, error)

	ListShippingAddresses(accountId string, params *ListShippingAddressesParams, opts ...Option) *ShippingAddressList

	CreateShippingAddress(accountId string, body *ShippingAddressCreate, opts ...Option) (*ShippingAddress, error)
	CreateShippingAddressWithContext(ctx context.Context, accountId string, body *ShippingAddressCreate, opts ...Option) (*ShippingAddress, error)

	GetShippingAddress(accountId string, shippingAddressId string, opts ...Option) (*ShippingAddress, error)
	GetShippingAddressWithContext(ctx context.Context, accountId string, shippingAddressId string, opts ...Option) (*ShippingAddress, error)

	UpdateShippingAddress(accountId string, shippingAddressId string, body *ShippingAddressUpdate, opts ...Option) (*ShippingAddress, error)
	UpdateShippingAddressWithContext(ctx context.Context, accountId string, shippingAddressId string, body *ShippingAddressUpdate, opts ...Option) (*ShippingAddress, error)

	RemoveShippingAddress(accountId string, shippingAddressId string, opts ...Option) (*Empty, error)
	RemoveShippingAddressWithContext(ctx context.Context, accountId string, shippingAddressId string, opts ...Option) (*Empty, error)

	ListAccountSubscriptions(accountId string, params *ListAccountSubscriptionsParams, opts ...Option) *SubscriptionList

	ListAccountTransactions(accountId string, params *ListAccountTransactionsParams, opts ...Option) *TransactionList

	ListChildAccounts(accountId string, params *ListChildAccountsParams, opts ...Option) *AccountList

	ListAccountAcquisition(params *ListAccountAcquisitionParams, opts ...Option) *AccountAcquisitionList

	ListCoupons(params *ListCouponsParams, opts ...Option) *CouponList

	CreateCoupon(body *CouponCreate, opts ...Option) (*Coupon, error)
	CreateCouponWithContext(ctx context.Context, body *CouponCreate, opts ...Option) (*Coupon, error)

	GetCoupon(couponId string, opts ...Option) (*Coupon, error)
	GetCouponWithContext(ctx context.Context, couponId string, opts ...Option) (*Coupon, error)

	UpdateCoupon(couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error)
	UpdateCouponWithContext(ctx context.Context, couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error)

	DeactivateCoupon(couponId string, opts ...Option) (*Coupon, error)
	DeactivateCouponWithContext(ctx context.Context, couponId string, opts ...Option) (*Coupon, error)

	RestoreCoupon(couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error)
	RestoreCouponWithContext(ctx context.Context, couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error)

	ListUniqueCouponCodes(couponId string, params *ListUniqueCouponCodesParams, opts ...Option) *UniqueCouponCodeList

	ListCreditPayments(params *ListCreditPaymentsParams, opts ...Option) *CreditPaymentList

	GetCreditPayment(creditPaymentId string, opts ...Option) (*CreditPayment, error)
	GetCreditPaymentWithContext(ctx context.Context, creditPaymentId string, opts ...Option) (*CreditPayment, error)

	ListCustomFieldDefinitions(params *ListCustomFieldDefinitionsParams, opts ...Option) *CustomFieldDefinitionList

	GetCustomFieldDefinition(customFieldDefinitionId string, opts ...Option) (*CustomFieldDefinition, error)
	GetCustomFieldDefinitionWithContext(ctx context.Context, customFieldDefinitionId string, opts ...Option) (*CustomFieldDefinition, error)

	ListItems(params *ListItemsParams, opts ...Option) *ItemList

	CreateItem(body *ItemCreate, opts ...Option) (*Item, error)
	CreateItemWithContext(ctx context.Context, body *ItemCreate, opts ...Option) (*Item, error)

	GetItem(itemId string, opts ...Option) (*Item, error)
	GetItemWithContext(ctx context.Context, itemId string, opts ...Option) (*Item, error)

	UpdateItem(itemId string, body *ItemUpdate, opts ...Option) (*Item, error)
	UpdateItemWithContext(ctx context.Context, itemId string, body *ItemUpdate, opts ...Option) (*Item, error)

	DeactivateItem(itemId string, opts ...Option) (*Item, error)
	DeactivateItemWithContext(ctx context.Context, itemId string, opts ...Option) (*Item, error)

	ReactivateItem(itemId string, opts ...Option) (*Item, error)
	ReactivateItemWithContext(ctx context.Context, itemId string, opts ...Option) (*Item, error)

	ListMeasuredUnit(params *ListMeasuredUnitParams, opts ...Option) *MeasuredUnitList

	CreateMeasuredUnit(body *MeasuredUnitCreate, opts ...Option) (*MeasuredUnit, error)
	CreateMeasuredUnitWithContext(ctx context.Context, body *MeasuredUnitCreate, opts ...Option) (*MeasuredUnit, error)

	GetMeasuredUnit(measuredUnitId string, opts ...Option) (*MeasuredUnit, error)
	GetMeasuredUnitWithContext(ctx context.Context, measuredUnitId string, opts ...Option) (*MeasuredUnit, error)

	UpdateMeasuredUnit(measuredUnitId string, body *MeasuredUnitUpdate, opts ...Option) (*MeasuredUnit, error)
	UpdateMeasuredUnitWithContext(ctx context.Context, measuredUnitId string, body *MeasuredUnitUpdate, opts ...Option) (*MeasuredUnit, error)

	RemoveMeasuredUnit(measuredUnitId string, opts ...Option) (*MeasuredUnit, error)
	RemoveMeasuredUnitWithContext(ctx context.Context, measuredUnitId string, opts ...Option) (*MeasuredUnit, error)

	ListInvoices(params *ListInvoicesParams, opts ...Option) *InvoiceList

	GetInvoice(invoiceId string, opts ...Option) (*Invoice, error)
	GetInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error)

	PutInvoice(invoiceId string, body *InvoiceUpdatable, opts ...Option) (*Invoice, error)
	PutInvoiceWithContext(ctx context.Context, invoiceId string, body *InvoiceUpdatable, opts ...Option) (*Invoice, error)

	ApplyCreditBalance(invoiceId string, opts ...Option) (*Invoice, error)
	ApplyCreditBalanceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error)

	CollectInvoice(invoiceId string, params *CollectInvoiceParams, opts ...Option) (*Invoice, error)
	CollectInvoiceWithContext(ctx context.Context, invoiceId string, params *CollectInvoiceParams, opts ...Option) (*Invoice, error)

	FailInvoice(invoiceId string, opts ...Option) (*Invoice, error)
	FailInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error)

	MarkInvoiceSuccessful(invoiceId string, opts ...Option) (*Invoice, error)
	MarkInvoiceSuccessfulWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error)

	ReopenInvoice(invoiceId string, opts ...Option) (*Invoice, error)
	ReopenInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error)

	VoidInvoice(invoiceId string, opts ...Option) (*Invoice, error)
	VoidInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error)

	RecordExternalTransaction(invoiceId string, body *ExternalTransaction, opts ...Option) (*Transaction, error)
	RecordExternalTransactionWithContext(ctx context.Context, invoiceId string, body *ExternalTransaction, opts ...Option) (*Transaction, error)

	ListInvoiceLineItems(invoiceId string, params *ListInvoiceLineItemsParams, opts ...Option) *LineItemList

	ListInvoiceCouponRedemptions(invoiceId string, params *ListInvoiceCouponRedemptionsParams, opts ...Option) *CouponRedemptionList

	ListRelatedInvoices(invoiceId string, opts ...Option) *InvoiceList

	RefundInvoice(invoiceId string, body *InvoiceRefund, opts ...Option) (*Invoice, error)
	RefundInvoiceWithContext(ctx context.Context, invoiceId string, body *InvoiceRefund, opts ...Option) (*Invoice, error)

	ListLineItems(params *ListLineItemsParams, opts ...Option) *LineItemList

	GetLineItem(lineItemId string, opts ...Option) (*LineItem, error)
	GetLineItemWithContext(ctx context.Context, lineItemId string, opts ...Option) (*LineItem, error)

	RemoveLineItem(lineItemId string, opts ...Option) (*Empty, error)
	RemoveLineItemWithContext(ctx context.Context, lineItemId string, opts ...Option) (*Empty, error)

	ListPlans(params *ListPlansParams, opts ...Option) *PlanList

	CreatePlan(body *PlanCreate, opts ...Option) (*Plan, error)
	CreatePlanWithContext(ctx context.Context, body *PlanCreate, opts ...Option) (*Plan, error)

	GetPlan(planId string, opts ...Option) (*Plan, error)
	GetPlanWithContext(ctx context.Context, planId string, opts ...Option) (*Plan, error)

	UpdatePlan(planId string, body *PlanUpdate, opts ...Option) (*Plan, error)
	UpdatePlanWithContext(ctx context.Context, planId string, body *PlanUpdate, opts ...Option) (*Plan, error)

	RemovePlan(planId string, opts ...Option) (*Plan, error)
	RemovePlanWithContext(ctx context.Context, planId string, opts ...Option) (*Plan, error)

	ListPlanAddOns(planId string, params *ListPlanAddOnsParams, opts ...Option) *AddOnList

	CreatePlanAddOn(planId string, body *AddOnCreate, opts ...Option) (*AddOn, error)
	CreatePlanAddOnWithContext(ctx context.Context, planId string, body *AddOnCreate, opts ...Option) (*AddOn, error)

	GetPlanAddOn(planId string, addOnId string, opts ...Option) (*AddOn, error)
	GetPlanAddOnWithContext(ctx context.Context, planId string, addOnId string, opts ...Option) (*AddOn, error)

	UpdatePlanAddOn(planId string, addOnId string, body *AddOnUpdate, opts ...Option) (*AddOn, error)
	UpdatePlanAddOnWithContext(ctx context.Context, planId string, addOnId string, body *AddOnUpdate, opts ...Option) (*AddOn, error)

	RemovePlanAddOn(planId string, addOnId string, opts ...Option) (*AddOn, error)
	RemovePlanAddOnWithContext(ctx context.Context, planId string, addOnId string, opts ...Option) (*AddOn, error)

	ListAddOns(params *ListAddOnsParams, opts ...Option) *AddOnList

	GetAddOn(addOnId string, opts ...Option) (*AddOn, error)
	GetAddOnWithContext(ctx context.Context, addOnId string, opts ...Option) (*AddOn, error)

	ListShippingMethods(params *ListShippingMethodsParams, opts ...Option) *ShippingMethodList

	CreateShippingMethod(body *ShippingMethodCreate, opts ...Option) (*ShippingMethod, error)
	CreateShippingMethodWithContext(ctx context.Context, body *ShippingMethodCreate, opts ...Option) (*ShippingMethod, error)

	GetShippingMethod(id string, opts ...Option) (*ShippingMethod, error)
	GetShippingMethodWithContext(ctx context.Context, id string, opts ...Option) (*ShippingMethod, error)

	UpdateShippingMethod(shippingMethodId string, body *ShippingMethodUpdate, opts ...Option) (*ShippingMethod, error)
	UpdateShippingMethodWithContext(ctx context.Context, shippingMethodId string, body *ShippingMethodUpdate, opts ...Option) (*ShippingMethod, error)

	DeactivateShippingMethod(shippingMethodId string, opts ...Option) (*ShippingMethod, error)
	DeactivateShippingMethodWithContext(ctx context.Context, shippingMethodId string, opts ...Option) (*ShippingMethod, error)

	ListSubscriptions(params *ListSubscriptionsParams, opts ...Option) *SubscriptionList

	CreateSubscription(body *SubscriptionCreate, opts ...Option) (*Subscription, error)
	CreateSubscriptionWithContext(ctx context.Context, body *SubscriptionCreate, opts ...Option) (*Subscription, error)

	GetSubscription(subscriptionId string, opts ...Option) (*Subscription, error)
	GetSubscriptionWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error)

	ModifySubscription(subscriptionId string, body *SubscriptionUpdate, opts ...Option) (*Subscription, error)
	ModifySubscriptionWithContext(ctx context.Context, subscriptionId string, body *SubscriptionUpdate, opts ...Option) (*Subscription, error)

	TerminateSubscription(subscriptionId string, params *TerminateSubscriptionParams, opts ...Option) (*Subscription, error)
	TerminateSubscriptionWithContext(ctx context.Context, subscriptionId string, params *TerminateSubscriptionParams, opts ...Option) (*Subscription, error)

	CancelSubscription(subscriptionId string, params *CancelSubscriptionParams, opts ...Option) (*Subscription, error)
	CancelSubscriptionWithContext(ctx context.Context, subscriptionId string, params *CancelSubscriptionParams, opts ...Option) (*Subscription, error)

	ReactivateSubscription(subscriptionId string, opts ...Option) (*Subscription, error)
	ReactivateSubscriptionWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error)

	PauseSubscription(subscriptionId string, body *SubscriptionPause, opts ...Option) (*Subscription, error)
	PauseSubscriptionWithContext(ctx context.Context, subscriptionId string, body *SubscriptionPause, opts ...Option) (*Subscription, error)

	ResumeSubscription(subscriptionId string, opts ...Option) (*Subscription, error)
	ResumeSubscriptionWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error)

	ConvertTrial(subscriptionId string, opts ...Option) (*Subscription, error)
	ConvertTrialWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error)

	GetPreviewRenewal(subscriptionId string, opts ...Option) (*InvoiceCollection, error)
	GetPreviewRenewalWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*InvoiceCollection, error)

	GetSubscriptionChange(subscriptionId string, opts ...Option) (*SubscriptionChange, error)
	GetSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*SubscriptionChange, error)

	CreateSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChange, error)
	CreateSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChange, error)

	RemoveSubscriptionChange(subscriptionId string, opts ...Option) (*Empty, error)
	RemoveSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Empty, error)

	PreviewSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChangePreview, error)
	PreviewSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChangePreview, error)

	ListSubscriptionInvoices(subscriptionId string, params *ListSubscriptionInvoicesParams, opts ...Option) *InvoiceList

	ListSubscriptionLineItems(subscriptionId string, params *ListSubscriptionLineItemsParams, opts ...Option) *LineItemList

	ListSubscriptionCouponRedemptions(subscriptionId string, params *ListSubscriptionCouponRedemptionsParams, opts ...Option) *CouponRedemptionList

	ListUsage(subscriptionId string, addOnId string, params *ListUsageParams, opts ...Option) *UsageList

	CreateUsage(subscriptionId string, addOnId string, body *UsageCreate, opts ...Option) (*Usage, error)
	CreateUsageWithContext(ctx context.Context, subscriptionId string, addOnId string, body *UsageCreate, opts ...Option) (*Usage, error)

	GetUsage(usageId string, opts ...Option) (*Usage, error)
	GetUsageWithContext(ctx context.Context, usageId string, opts ...Option) (*Usage, error)

	UpdateUsage(usageId string, body *UsageCreate, opts ...Option) (*Usage, error)
	UpdateUsageWithContext(ctx context.Context, usageId string, body *UsageCreate, opts ...Option) (*Usage, error)

	RemoveUsage(usageId string, opts ...Option) (*Empty, error)
	RemoveUsageWithContext(ctx context.Context, usageId string, opts ...Option) (*Empty, error)

	ListTransactions(params *ListTransactionsParams, opts ...Option) *TransactionList

	GetTransaction(transactionId string, opts ...Option) (*Transaction, error)
	GetTransactionWithContext(ctx context.Context, transactionId string, opts ...Option) (*Transaction, error)

	GetUniqueCouponCode(uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error)
	GetUniqueCouponCodeWithContext(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error)

	DeactivateUniqueCouponCode(uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error)
	DeactivateUniqueCouponCodeWithContext(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error)

	ReactivateUniqueCouponCode(uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error)
	ReactivateUniqueCouponCodeWithContext(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error)

	CreatePurchase(body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error)
	CreatePurchaseWithContext(ctx context.Context, body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error)

	PreviewPurchase(body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error)
	PreviewPurchaseWithContext(ctx context.Context, body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error)

	GetExportDates(opts ...Option) (*ExportDates, error)
	GetExportDatesWithContext(ctx context.Context, opts ...Option) (*ExportDates, error)

	GetExportFiles(exportDate string, opts ...Option) (*ExportFiles, error)
	GetExportFilesWithContext(ctx context.Context, exportDate string, opts ...Option) (*ExportFiles, error)

	ListDunningCampaigns(params *ListDunningCampaignsParams, opts ...Option) *DunningCampaignList

	GetDunningCampaign(dunningCampaignId string, opts ...Option) (*DunningCampaign, error)
	GetDunningCampaignWithContext(ctx context.Context, dunningCampaignId string, opts ...Option) (*DunningCampaign, error)

	PutDunningCampaignBulkUpdate(dunningCampaignId string, body *DunningCampaignsBulkUpdate, opts ...Option) (*DunningCampaignsBulkUpdateResponse, error)
	PutDunningCampaignBulkUpdateWithContext(ctx context.Context, dunningCampaignId string, body *DunningCampaignsBulkUpdate, opts ...Option) (*DunningCampaignsBulkUpdateResponse, error)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_sites
//
// Returns: A list of sites.
func (c *Client) ListSites(params *ListSitesParams, opts ...Option) *SiteList {
	path, err := c.InterpolatePath("/sites")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewSiteList(c, path, requestOptions)
}

// GetSite wraps GetSiteWithContext using the background context
func (c *Client) GetSite(siteId string, opts ...Option) (*Site, error) {
	ctx := context.Background()
	return c.getSite(ctx, siteId, opts...)
}

// GetSiteWithContext Fetch a site
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_site
//
// Returns: A site.
func (c *Client) GetSiteWithContext(ctx context.Context, siteId string, opts ...Option) (*Site, error) {
	return c.getSite(ctx, siteId, opts...)
}

func (c *Client) getSite(ctx context.Context, siteId string, opts ...Option) (*Site, error) {
	path, err := c.InterpolatePath("/sites/{site_id}", siteId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Site{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_accounts
//
// Returns: A list of the site's accounts.
func (c *Client) ListAccounts(params *ListAccountsParams, opts ...Option) *AccountList {
	path, err := c.InterpolatePath("/accounts")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewAccountList(c, path, requestOptions)
}

// CreateAccount wraps CreateAccountWithContext using the background context
func (c *Client) CreateAccount(body *AccountCreate, opts ...Option) (*Account, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createAccount(ctx, body, opts...)
}

// CreateAccountWithContext Create an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_account
//
// Returns: An account.
func (c *Client) CreateAccountWithContext(ctx context.Context, body *AccountCreate, opts ...Option) (*Account, error) {
	return c.createAccount(ctx, body, opts...)
}

func (c *Client) createAccount(ctx context.Context, body *AccountCreate, opts ...Option) (*Account, error) {
	path, err := c.InterpolatePath("/accounts")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Account{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetAccount wraps GetAccountWithContext using the background context
func (c *Client) GetAccount(accountId string, opts ...Option) (*Account, error) {
	ctx := context.Background()
	return c.getAccount(ctx, accountId, opts...)
}

// GetAccountWithContext Fetch an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_account
//
// Returns: An account.
func (c *Client) GetAccountWithContext(ctx context.Context, accountId string, opts ...Option) (*Account, error) {
	return c.getAccount(ctx, accountId, opts...)
}

func (c *Client) getAccount(ctx context.Context, accountId string, opts ...Option) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Account{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateAccount wraps UpdateAccountWithContext using the background context
func (c *Client) UpdateAccount(accountId string, body *AccountUpdate, opts ...Option) (*Account, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateAccount(ctx, accountId, body, opts...)
}

// UpdateAccountWithContext Modify an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_account
//
// Returns: An account.
func (c *Client) UpdateAccountWithContext(ctx context.Context, accountId string, body *AccountUpdate, opts ...Option) (*Account, error) {
	return c.updateAccount(ctx, accountId, body, opts...)
}

func (c *Client) updateAccount(ctx context.Context, accountId string, body *AccountUpdate, opts ...Option) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Account{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateAccount wraps DeactivateAccountWithContext using the background context
func (c *Client) DeactivateAccount(accountId string, opts ...Option) (*Account, error) {
	ctx := context.Background()
	return c.deactivateAccount(ctx, accountId, opts...)
}

// DeactivateAccountWithContext Deactivate an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/deactivate_account
//
// Returns: An account.
func (c *Client) DeactivateAccountWithContext(ctx context.Context, accountId string, opts ...Option) (*Account, error) {
	return c.deactivateAccount(ctx, accountId, opts...)
}

func (c *Client) deactivateAccount(ctx context.Context, accountId string, opts ...Option) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Account{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetAccountAcquisition wraps GetAccountAcquisitionWithContext using the background context
func (c *Client) GetAccountAcquisition(accountId string, opts ...Option) (*AccountAcquisition, error) {
	ctx := context.Background()
	return c.getAccountAcquisition(ctx, accountId, opts...)
}

// GetAccountAcquisitionWithContext Fetch an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_account_acquisition
//
// Returns: An account's acquisition data.
func (c *Client) GetAccountAcquisitionWithContext(ctx context.Context, accountId string, opts ...Option) (*AccountAcquisition, error) {
	return c.getAccountAcquisition(ctx, accountId, opts...)
}

func (c *Client) getAccountAcquisition(ctx context.Context, accountId string, opts ...Option) (*AccountAcquisition, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &AccountAcquisition{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateAccountAcquisition wraps UpdateAccountAcquisitionWithContext using the background context
func (c *Client) UpdateAccountAcquisition(accountId string, body *AccountAcquisitionUpdatable, opts ...Option) (*AccountAcquisition, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateAccountAcquisition(ctx, accountId, body, opts...)
}

// UpdateAccountAcquisitionWithContext Update an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_account_acquisition
//
// Returns: An account's updated acquisition data.
func (c *Client) UpdateAccountAcquisitionWithContext(ctx context.Context, accountId string, body *AccountAcquisitionUpdatable, opts ...Option) (*AccountAcquisition, error) {
	return c.updateAccountAcquisition(ctx, accountId, body, opts...)
}

func (c *Client) updateAccountAcquisition(ctx context.Context, accountId string, body *AccountAcquisitionUpdatable, opts ...Option) (*AccountAcquisition, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &AccountAcquisition{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveAccountAcquisition wraps RemoveAccountAcquisitionWithContext using the background context
func (c *Client) RemoveAccountAcquisition(accountId string, opts ...Option) (*Empty, error) {
	ctx := context.Background()
	return c.removeAccountAcquisition(ctx, accountId, opts...)
}

// RemoveAccountAcquisitionWithContext Remove an account's acquisition data
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_account_acquisition
//
// Returns: Acquisition data was succesfully deleted.
func (c *Client) RemoveAccountAcquisitionWithContext(ctx context.Context, accountId string, opts ...Option) (*Empty, error) {
	return c.removeAccountAcquisition(ctx, accountId, opts...)
}

func (c *Client) removeAccountAcquisition(ctx context.Context, accountId string, opts ...Option) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/acquisition", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Empty{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateAccount wraps ReactivateAccountWithContext using the background context
func (c *Client) ReactivateAccount(accountId string, opts ...Option) (*Account, error) {
	ctx := context.Background()
	return c.reactivateAccount(ctx, accountId, opts...)
}

// ReactivateAccountWithContext Reactivate an inactive account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/reactivate_account
//
// Returns: An account.
func (c *Client) ReactivateAccountWithContext(ctx context.Context, accountId string, opts ...Option) (*Account, error) {
	return c.reactivateAccount(ctx, accountId, opts...)
}

func (c *Client) reactivateAccount(ctx context.Context, accountId string, opts ...Option) (*Account, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/reactivate", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Account{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetAccountBalance wraps GetAccountBalanceWithContext using the background context
func (c *Client) GetAccountBalance(accountId string, opts ...Option) (*AccountBalance, error) {
	ctx := context.Background()
	return c.getAccountBalance(ctx, accountId, opts...)
}

// GetAccountBalanceWithContext Fetch an account's balance and past due status
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_account_balance
//
// Returns: An account's balance.
func (c *Client) GetAccountBalanceWithContext(ctx context.Context, accountId string, opts ...Option) (*AccountBalance, error) {
	return c.getAccountBalance(ctx, accountId, opts...)
}

func (c *Client) getAccountBalance(ctx context.Context, accountId string, opts ...Option) (*AccountBalance, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/balance", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &AccountBalance{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetBillingInfo wraps GetBillingInfoWithContext using the background context
func (c *Client) GetBillingInfo(accountId string, opts ...Option) (*BillingInfo, error) {
	ctx := context.Background()
	return c.getBillingInfo(ctx, accountId, opts...)
}

// GetBillingInfoWithContext Fetch an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_billing_info
//
// Returns: An account's billing information.
func (c *Client) GetBillingInfoWithContext(ctx context.Context, accountId string, opts ...Option) (*BillingInfo, error) {
	return c.getBillingInfo(ctx, accountId, opts...)
}

func (c *Client) getBillingInfo(ctx context.Context, accountId string, opts ...Option) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &BillingInfo{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateBillingInfo wraps UpdateBillingInfoWithContext using the background context
func (c *Client) UpdateBillingInfo(accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateBillingInfo(ctx, accountId, body, opts...)
}

// UpdateBillingInfoWithContext Set an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_billing_info
//
// Returns: Updated billing information.
func (c *Client) UpdateBillingInfoWithContext(ctx context.Context, accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	return c.updateBillingInfo(ctx, accountId, body, opts...)
}

func (c *Client) updateBillingInfo(ctx context.Context, accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &BillingInfo{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveBillingInfo wraps RemoveBillingInfoWithContext using the background context
func (c *Client) RemoveBillingInfo(accountId string, opts ...Option) (*Empty, error) {
	ctx := context.Background()
	return c.removeBillingInfo(ctx, accountId, opts...)
}

// RemoveBillingInfoWithContext Remove an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_billing_info
//
// Returns: Billing information deleted
func (c *Client) RemoveBillingInfoWithContext(ctx context.Context, accountId string, opts ...Option) (*Empty, error) {
	return c.removeBillingInfo(ctx, accountId, opts...)
}

func (c *Client) removeBillingInfo(ctx context.Context, accountId string, opts ...Option) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Empty{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type VerifyBillingInfoParams struct {
	Params

	// Body - The body of the request.
	Body *BillingInfoVerify
}

func (list *VerifyBillingInfoParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *VerifyBillingInfoParams) URLParams() []KeyValue {
	var options []KeyValue

	return options
}

// VerifyBillingInfo wraps VerifyBillingInfoWithContext using the background context
func (c *Client) VerifyBillingInfo(accountId string, params *VerifyBillingInfoParams, opts ...Option) (*Transaction, error) {
	ctx := context.Background()
	return c.verifyBillingInfo(ctx, accountId, params, opts...)
}

// VerifyBillingInfoWithContext Verify an account's credit card billing information
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/verify_billing_info
//
// Returns: Transaction information from verify.
func (c *Client) VerifyBillingInfoWithContext(ctx context.Context, accountId string, params *VerifyBillingInfoParams, opts ...Option) (*Transaction, error) {
	return c.verifyBillingInfo(ctx, accountId, params, opts...)
}

func (c *Client) verifyBillingInfo(ctx context.Context, accountId string, params *VerifyBillingInfoParams, opts ...Option) (*Transaction, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_info/verify", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Transaction{}
	err = c.Call(ctx, http.MethodPost, path, nil, params, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_billing_infos
//
// Returns: A list of the the billing information for an account's
func (c *Client) ListBillingInfos(accountId string, params *ListBillingInfosParams, opts ...Option) *BillingInfoList {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewBillingInfoList(c, path, requestOptions)
}

// CreateBillingInfo wraps CreateBillingInfoWithContext using the background context
func (c *Client) CreateBillingInfo(accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createBillingInfo(ctx, accountId, body, opts...)
}

// CreateBillingInfoWithContext Add new billing information on an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_billing_info
//
// Returns: Updated billing information.
func (c *Client) CreateBillingInfoWithContext(ctx context.Context, accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	return c.createBillingInfo(ctx, accountId, body, opts...)
}

func (c *Client) createBillingInfo(ctx context.Context, accountId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &BillingInfo{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetABillingInfo wraps GetABillingInfoWithContext using the background context
func (c *Client) GetABillingInfo(accountId string, billingInfoId string, opts ...Option) (*BillingInfo, error) {
	ctx := context.Background()
	return c.getABillingInfo(ctx, accountId, billingInfoId, opts...)
}

// GetABillingInfoWithContext Fetch a billing info
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_a_billing_info
//
// Returns: A billing info.
func (c *Client) GetABillingInfoWithContext(ctx context.Context, accountId string, billingInfoId string, opts ...Option) (*BillingInfo, error) {
	return c.getABillingInfo(ctx, accountId, billingInfoId, opts...)
}

func (c *Client) getABillingInfo(ctx context.Context, accountId string, billingInfoId string, opts ...Option) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &BillingInfo{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateABillingInfo wraps UpdateABillingInfoWithContext using the background context
func (c *Client) UpdateABillingInfo(accountId string, billingInfoId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateABillingInfo(ctx, accountId, billingInfoId, body, opts...)
}

// UpdateABillingInfoWithContext Update an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_a_billing_info
//
// Returns: Updated billing information.
func (c *Client) UpdateABillingInfoWithContext(ctx context.Context, accountId string, billingInfoId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	return c.updateABillingInfo(ctx, accountId, billingInfoId, body, opts...)
}

func (c *Client) updateABillingInfo(ctx context.Context, accountId string, billingInfoId string, body *BillingInfoCreate, opts ...Option) (*BillingInfo, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &BillingInfo{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveABillingInfo wraps RemoveABillingInfoWithContext using the background context
func (c *Client) RemoveABillingInfo(accountId string, billingInfoId string, opts ...Option) (*Empty, error) {
	ctx := context.Background()
	return c.removeABillingInfo(ctx, accountId, billingInfoId, opts...)
}

// RemoveABillingInfoWithContext Remove an account's billing information
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_a_billing_info
//
// Returns: Billing information deleted
func (c *Client) RemoveABillingInfoWithContext(ctx context.Context, accountId string, billingInfoId string, opts ...Option) (*Empty, error) {
	return c.removeABillingInfo(ctx, accountId, billingInfoId, opts...)
}

func (c *Client) removeABillingInfo(ctx context.Context, accountId string, billingInfoId string, opts ...Option) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/billing_infos/{billing_info_id}", accountId, billingInfoId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Empty{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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

	// State - Filter by state.
	State *string
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

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
	}

	return options
}

// ListAccountCouponRedemptions List the coupon redemptions for an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_coupon_redemptions
//
// Returns: A list of the the coupon redemptions on an account.
func (c *Client) ListAccountCouponRedemptions(accountId string, params *ListAccountCouponRedemptionsParams, opts ...Option) *CouponRedemptionList {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewCouponRedemptionList(c, path, requestOptions)
}

// GetActiveCouponRedemption wraps GetActiveCouponRedemptionWithContext using the background context
func (c *Client) GetActiveCouponRedemption(accountId string, opts ...Option) (*CouponRedemption, error) {
	ctx := context.Background()
	return c.getActiveCouponRedemption(ctx, accountId, opts...)
}

// GetActiveCouponRedemptionWithContext Fetch the coupon redemption that is active on an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_active_coupon_redemption
//
// Returns: An active coupon redemption on an account.
func (c *Client) GetActiveCouponRedemptionWithContext(ctx context.Context, accountId string, opts ...Option) (*CouponRedemption, error) {
	return c.getActiveCouponRedemption(ctx, accountId, opts...)
}

func (c *Client) getActiveCouponRedemption(ctx context.Context, accountId string, opts ...Option) (*CouponRedemption, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &CouponRedemption{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// CreateCouponRedemption wraps CreateCouponRedemptionWithContext using the background context
func (c *Client) CreateCouponRedemption(accountId string, body *CouponRedemptionCreate, opts ...Option) (*CouponRedemption, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createCouponRedemption(ctx, accountId, body, opts...)
}

// CreateCouponRedemptionWithContext Generate an active coupon redemption on an account or subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_coupon_redemption
//
// Returns: Returns the new coupon redemption.
func (c *Client) CreateCouponRedemptionWithContext(ctx context.Context, accountId string, body *CouponRedemptionCreate, opts ...Option) (*CouponRedemption, error) {
	return c.createCouponRedemption(ctx, accountId, body, opts...)
}

func (c *Client) createCouponRedemption(ctx context.Context, accountId string, body *CouponRedemptionCreate, opts ...Option) (*CouponRedemption, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &CouponRedemption{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveCouponRedemption wraps RemoveCouponRedemptionWithContext using the background context
func (c *Client) RemoveCouponRedemption(accountId string, opts ...Option) (*CouponRedemption, error) {
	ctx := context.Background()
	return c.removeCouponRedemption(ctx, accountId, opts...)
}

// RemoveCouponRedemptionWithContext Delete the active coupon redemption from an account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_coupon_redemption
//
// Returns: Coupon redemption deleted.
func (c *Client) RemoveCouponRedemptionWithContext(ctx context.Context, accountId string, opts ...Option) (*CouponRedemption, error) {
	return c.removeCouponRedemption(ctx, accountId, opts...)
}

func (c *Client) removeCouponRedemption(ctx context.Context, accountId string, opts ...Option) (*CouponRedemption, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/coupon_redemptions/active", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &CouponRedemption{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_credit_payments
//
// Returns: A list of the account's credit payments.
func (c *Client) ListAccountCreditPayments(accountId string, params *ListAccountCreditPaymentsParams, opts ...Option) *CreditPaymentList {
	path, err := c.InterpolatePath("/accounts/{account_id}/credit_payments", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewCreditPaymentList(c, path, requestOptions)
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

	// State - Invoice state.
	State *string

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

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_invoices
//
// Returns: A list of the account's invoices.
func (c *Client) ListAccountInvoices(accountId string, params *ListAccountInvoicesParams, opts ...Option) *InvoiceList {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewInvoiceList(c, path, requestOptions)
}

// CreateInvoice wraps CreateInvoiceWithContext using the background context
func (c *Client) CreateInvoice(accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createInvoice(ctx, accountId, body, opts...)
}

// CreateInvoiceWithContext Create an invoice for pending line items
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_invoice
//
// Returns: Returns the new invoices.
func (c *Client) CreateInvoiceWithContext(ctx context.Context, accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error) {
	return c.createInvoice(ctx, accountId, body, opts...)
}

func (c *Client) createInvoice(ctx context.Context, accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &InvoiceCollection{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PreviewInvoice wraps PreviewInvoiceWithContext using the background context
func (c *Client) PreviewInvoice(accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.previewInvoice(ctx, accountId, body, opts...)
}

// PreviewInvoiceWithContext Preview new invoice for pending line items
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/preview_invoice
//
// Returns: Returns the invoice previews.
func (c *Client) PreviewInvoiceWithContext(ctx context.Context, accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error) {
	return c.previewInvoice(ctx, accountId, body, opts...)
}

func (c *Client) previewInvoice(ctx context.Context, accountId string, body *InvoiceCreate, opts ...Option) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/invoices/preview", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &InvoiceCollection{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_line_items
//
// Returns: A list of the account's line items.
func (c *Client) ListAccountLineItems(accountId string, params *ListAccountLineItemsParams, opts ...Option) *LineItemList {
	path, err := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewLineItemList(c, path, requestOptions)
}

// CreateLineItem wraps CreateLineItemWithContext using the background context
func (c *Client) CreateLineItem(accountId string, body *LineItemCreate, opts ...Option) (*LineItem, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createLineItem(ctx, accountId, body, opts...)
}

// CreateLineItemWithContext Create a new line item for the account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_line_item
//
// Returns: Returns the new line item.
func (c *Client) CreateLineItemWithContext(ctx context.Context, accountId string, body *LineItemCreate, opts ...Option) (*LineItem, error) {
	return c.createLineItem(ctx, accountId, body, opts...)
}

func (c *Client) createLineItem(ctx context.Context, accountId string, body *LineItemCreate, opts ...Option) (*LineItem, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/line_items", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &LineItem{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
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

// ListAccountNotes List an account's notes
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_notes
//
// Returns: A list of an account's notes.
func (c *Client) ListAccountNotes(accountId string, params *ListAccountNotesParams, opts ...Option) *AccountNoteList {
	path, err := c.InterpolatePath("/accounts/{account_id}/notes", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewAccountNoteList(c, path, requestOptions)
}

// GetAccountNote wraps GetAccountNoteWithContext using the background context
func (c *Client) GetAccountNote(accountId string, accountNoteId string, opts ...Option) (*AccountNote, error) {
	ctx := context.Background()
	return c.getAccountNote(ctx, accountId, accountNoteId, opts...)
}

// GetAccountNoteWithContext Fetch an account note
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_account_note
//
// Returns: An account note.
func (c *Client) GetAccountNoteWithContext(ctx context.Context, accountId string, accountNoteId string, opts ...Option) (*AccountNote, error) {
	return c.getAccountNote(ctx, accountId, accountNoteId, opts...)
}

func (c *Client) getAccountNote(ctx context.Context, accountId string, accountNoteId string, opts ...Option) (*AccountNote, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/notes/{account_note_id}", accountId, accountNoteId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &AccountNote{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_shipping_addresses
//
// Returns: A list of an account's shipping addresses.
func (c *Client) ListShippingAddresses(accountId string, params *ListShippingAddressesParams, opts ...Option) *ShippingAddressList {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewShippingAddressList(c, path, requestOptions)
}

// CreateShippingAddress wraps CreateShippingAddressWithContext using the background context
func (c *Client) CreateShippingAddress(accountId string, body *ShippingAddressCreate, opts ...Option) (*ShippingAddress, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createShippingAddress(ctx, accountId, body, opts...)
}

// CreateShippingAddressWithContext Create a new shipping address for the account
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_shipping_address
//
// Returns: Returns the new shipping address.
func (c *Client) CreateShippingAddressWithContext(ctx context.Context, accountId string, body *ShippingAddressCreate, opts ...Option) (*ShippingAddress, error) {
	return c.createShippingAddress(ctx, accountId, body, opts...)
}

func (c *Client) createShippingAddress(ctx context.Context, accountId string, body *ShippingAddressCreate, opts ...Option) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &ShippingAddress{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetShippingAddress wraps GetShippingAddressWithContext using the background context
func (c *Client) GetShippingAddress(accountId string, shippingAddressId string, opts ...Option) (*ShippingAddress, error) {
	ctx := context.Background()
	return c.getShippingAddress(ctx, accountId, shippingAddressId, opts...)
}

// GetShippingAddressWithContext Fetch an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_shipping_address
//
// Returns: A shipping address.
func (c *Client) GetShippingAddressWithContext(ctx context.Context, accountId string, shippingAddressId string, opts ...Option) (*ShippingAddress, error) {
	return c.getShippingAddress(ctx, accountId, shippingAddressId, opts...)
}

func (c *Client) getShippingAddress(ctx context.Context, accountId string, shippingAddressId string, opts ...Option) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &ShippingAddress{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateShippingAddress wraps UpdateShippingAddressWithContext using the background context
func (c *Client) UpdateShippingAddress(accountId string, shippingAddressId string, body *ShippingAddressUpdate, opts ...Option) (*ShippingAddress, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateShippingAddress(ctx, accountId, shippingAddressId, body, opts...)
}

// UpdateShippingAddressWithContext Update an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_shipping_address
//
// Returns: The updated shipping address.
func (c *Client) UpdateShippingAddressWithContext(ctx context.Context, accountId string, shippingAddressId string, body *ShippingAddressUpdate, opts ...Option) (*ShippingAddress, error) {
	return c.updateShippingAddress(ctx, accountId, shippingAddressId, body, opts...)
}

func (c *Client) updateShippingAddress(ctx context.Context, accountId string, shippingAddressId string, body *ShippingAddressUpdate, opts ...Option) (*ShippingAddress, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &ShippingAddress{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveShippingAddress wraps RemoveShippingAddressWithContext using the background context
func (c *Client) RemoveShippingAddress(accountId string, shippingAddressId string, opts ...Option) (*Empty, error) {
	ctx := context.Background()
	return c.removeShippingAddress(ctx, accountId, shippingAddressId, opts...)
}

// RemoveShippingAddressWithContext Remove an account's shipping address
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_shipping_address
//
// Returns: Shipping address deleted.
func (c *Client) RemoveShippingAddressWithContext(ctx context.Context, accountId string, shippingAddressId string, opts ...Option) (*Empty, error) {
	return c.removeShippingAddress(ctx, accountId, shippingAddressId, opts...)
}

func (c *Client) removeShippingAddress(ctx context.Context, accountId string, shippingAddressId string, opts ...Option) (*Empty, error) {
	path, err := c.InterpolatePath("/accounts/{account_id}/shipping_addresses/{shipping_address_id}", accountId, shippingAddressId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Empty{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_subscriptions
//
// Returns: A list of the account's subscriptions.
func (c *Client) ListAccountSubscriptions(accountId string, params *ListAccountSubscriptionsParams, opts ...Option) *SubscriptionList {
	path, err := c.InterpolatePath("/accounts/{account_id}/subscriptions", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewSubscriptionList(c, path, requestOptions)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_transactions
//
// Returns: A list of the account's transactions.
func (c *Client) ListAccountTransactions(accountId string, params *ListAccountTransactionsParams, opts ...Option) *TransactionList {
	path, err := c.InterpolatePath("/accounts/{account_id}/transactions", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewTransactionList(c, path, requestOptions)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_child_accounts
//
// Returns: A list of an account's child accounts.
func (c *Client) ListChildAccounts(accountId string, params *ListChildAccountsParams, opts ...Option) *AccountList {
	path, err := c.InterpolatePath("/accounts/{account_id}/accounts", accountId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewAccountList(c, path, requestOptions)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_account_acquisition
//
// Returns: A list of the site's account acquisition data.
func (c *Client) ListAccountAcquisition(params *ListAccountAcquisitionParams, opts ...Option) *AccountAcquisitionList {
	path, err := c.InterpolatePath("/acquisitions")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewAccountAcquisitionList(c, path, requestOptions)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_coupons
//
// Returns: A list of the site's coupons.
func (c *Client) ListCoupons(params *ListCouponsParams, opts ...Option) *CouponList {
	path, err := c.InterpolatePath("/coupons")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewCouponList(c, path, requestOptions)
}

// CreateCoupon wraps CreateCouponWithContext using the background context
func (c *Client) CreateCoupon(body *CouponCreate, opts ...Option) (*Coupon, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createCoupon(ctx, body, opts...)
}

// CreateCouponWithContext Create a new coupon
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_coupon
//
// Returns: A new coupon.
func (c *Client) CreateCouponWithContext(ctx context.Context, body *CouponCreate, opts ...Option) (*Coupon, error) {
	return c.createCoupon(ctx, body, opts...)
}

func (c *Client) createCoupon(ctx context.Context, body *CouponCreate, opts ...Option) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Coupon{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetCoupon wraps GetCouponWithContext using the background context
func (c *Client) GetCoupon(couponId string, opts ...Option) (*Coupon, error) {
	ctx := context.Background()
	return c.getCoupon(ctx, couponId, opts...)
}

// GetCouponWithContext Fetch a coupon
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_coupon
//
// Returns: A coupon.
func (c *Client) GetCouponWithContext(ctx context.Context, couponId string, opts ...Option) (*Coupon, error) {
	return c.getCoupon(ctx, couponId, opts...)
}

func (c *Client) getCoupon(ctx context.Context, couponId string, opts ...Option) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Coupon{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateCoupon wraps UpdateCouponWithContext using the background context
func (c *Client) UpdateCoupon(couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateCoupon(ctx, couponId, body, opts...)
}

// UpdateCouponWithContext Update an active coupon
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_coupon
//
// Returns: The updated coupon.
func (c *Client) UpdateCouponWithContext(ctx context.Context, couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error) {
	return c.updateCoupon(ctx, couponId, body, opts...)
}

func (c *Client) updateCoupon(ctx context.Context, couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Coupon{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateCoupon wraps DeactivateCouponWithContext using the background context
func (c *Client) DeactivateCoupon(couponId string, opts ...Option) (*Coupon, error) {
	ctx := context.Background()
	return c.deactivateCoupon(ctx, couponId, opts...)
}

// DeactivateCouponWithContext Expire a coupon
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/deactivate_coupon
//
// Returns: The expired Coupon
func (c *Client) DeactivateCouponWithContext(ctx context.Context, couponId string, opts ...Option) (*Coupon, error) {
	return c.deactivateCoupon(ctx, couponId, opts...)
}

func (c *Client) deactivateCoupon(ctx context.Context, couponId string, opts ...Option) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}", couponId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Coupon{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RestoreCoupon wraps RestoreCouponWithContext using the background context
func (c *Client) RestoreCoupon(couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.restoreCoupon(ctx, couponId, body, opts...)
}

// RestoreCouponWithContext Restore an inactive coupon
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/restore_coupon
//
// Returns: The restored coupon.
func (c *Client) RestoreCouponWithContext(ctx context.Context, couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error) {
	return c.restoreCoupon(ctx, couponId, body, opts...)
}

func (c *Client) restoreCoupon(ctx context.Context, couponId string, body *CouponUpdate, opts ...Option) (*Coupon, error) {
	path, err := c.InterpolatePath("/coupons/{coupon_id}/restore", couponId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Coupon{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_unique_coupon_codes
//
// Returns: A list of unique coupon codes that were generated
func (c *Client) ListUniqueCouponCodes(couponId string, params *ListUniqueCouponCodesParams, opts ...Option) *UniqueCouponCodeList {
	path, err := c.InterpolatePath("/coupons/{coupon_id}/unique_coupon_codes", couponId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewUniqueCouponCodeList(c, path, requestOptions)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_credit_payments
//
// Returns: A list of the site's credit payments.
func (c *Client) ListCreditPayments(params *ListCreditPaymentsParams, opts ...Option) *CreditPaymentList {
	path, err := c.InterpolatePath("/credit_payments")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewCreditPaymentList(c, path, requestOptions)
}

// GetCreditPayment wraps GetCreditPaymentWithContext using the background context
func (c *Client) GetCreditPayment(creditPaymentId string, opts ...Option) (*CreditPayment, error) {
	ctx := context.Background()
	return c.getCreditPayment(ctx, creditPaymentId, opts...)
}

// GetCreditPaymentWithContext Fetch a credit payment
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_credit_payment
//
// Returns: A credit payment.
func (c *Client) GetCreditPaymentWithContext(ctx context.Context, creditPaymentId string, opts ...Option) (*CreditPayment, error) {
	return c.getCreditPayment(ctx, creditPaymentId, opts...)
}

func (c *Client) getCreditPayment(ctx context.Context, creditPaymentId string, opts ...Option) (*CreditPayment, error) {
	path, err := c.InterpolatePath("/credit_payments/{credit_payment_id}", creditPaymentId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &CreditPayment{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_custom_field_definitions
//
// Returns: A list of the site's custom field definitions.
func (c *Client) ListCustomFieldDefinitions(params *ListCustomFieldDefinitionsParams, opts ...Option) *CustomFieldDefinitionList {
	path, err := c.InterpolatePath("/custom_field_definitions")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewCustomFieldDefinitionList(c, path, requestOptions)
}

// GetCustomFieldDefinition wraps GetCustomFieldDefinitionWithContext using the background context
func (c *Client) GetCustomFieldDefinition(customFieldDefinitionId string, opts ...Option) (*CustomFieldDefinition, error) {
	ctx := context.Background()
	return c.getCustomFieldDefinition(ctx, customFieldDefinitionId, opts...)
}

// GetCustomFieldDefinitionWithContext Fetch an custom field definition
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_custom_field_definition
//
// Returns: An custom field definition.
func (c *Client) GetCustomFieldDefinitionWithContext(ctx context.Context, customFieldDefinitionId string, opts ...Option) (*CustomFieldDefinition, error) {
	return c.getCustomFieldDefinition(ctx, customFieldDefinitionId, opts...)
}

func (c *Client) getCustomFieldDefinition(ctx context.Context, customFieldDefinitionId string, opts ...Option) (*CustomFieldDefinition, error) {
	path, err := c.InterpolatePath("/custom_field_definitions/{custom_field_definition_id}", customFieldDefinitionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &CustomFieldDefinition{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_items
//
// Returns: A list of the site's items.
func (c *Client) ListItems(params *ListItemsParams, opts ...Option) *ItemList {
	path, err := c.InterpolatePath("/items")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewItemList(c, path, requestOptions)
}

// CreateItem wraps CreateItemWithContext using the background context
func (c *Client) CreateItem(body *ItemCreate, opts ...Option) (*Item, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createItem(ctx, body, opts...)
}

// CreateItemWithContext Create a new item
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_item
//
// Returns: A new item.
func (c *Client) CreateItemWithContext(ctx context.Context, body *ItemCreate, opts ...Option) (*Item, error) {
	return c.createItem(ctx, body, opts...)
}

func (c *Client) createItem(ctx context.Context, body *ItemCreate, opts ...Option) (*Item, error) {
	path, err := c.InterpolatePath("/items")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Item{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetItem wraps GetItemWithContext using the background context
func (c *Client) GetItem(itemId string, opts ...Option) (*Item, error) {
	ctx := context.Background()
	return c.getItem(ctx, itemId, opts...)
}

// GetItemWithContext Fetch an item
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_item
//
// Returns: An item.
func (c *Client) GetItemWithContext(ctx context.Context, itemId string, opts ...Option) (*Item, error) {
	return c.getItem(ctx, itemId, opts...)
}

func (c *Client) getItem(ctx context.Context, itemId string, opts ...Option) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Item{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateItem wraps UpdateItemWithContext using the background context
func (c *Client) UpdateItem(itemId string, body *ItemUpdate, opts ...Option) (*Item, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateItem(ctx, itemId, body, opts...)
}

// UpdateItemWithContext Update an active item
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_item
//
// Returns: The updated item.
func (c *Client) UpdateItemWithContext(ctx context.Context, itemId string, body *ItemUpdate, opts ...Option) (*Item, error) {
	return c.updateItem(ctx, itemId, body, opts...)
}

func (c *Client) updateItem(ctx context.Context, itemId string, body *ItemUpdate, opts ...Option) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Item{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateItem wraps DeactivateItemWithContext using the background context
func (c *Client) DeactivateItem(itemId string, opts ...Option) (*Item, error) {
	ctx := context.Background()
	return c.deactivateItem(ctx, itemId, opts...)
}

// DeactivateItemWithContext Deactivate an item
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/deactivate_item
//
// Returns: An item.
func (c *Client) DeactivateItemWithContext(ctx context.Context, itemId string, opts ...Option) (*Item, error) {
	return c.deactivateItem(ctx, itemId, opts...)
}

func (c *Client) deactivateItem(ctx context.Context, itemId string, opts ...Option) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}", itemId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Item{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateItem wraps ReactivateItemWithContext using the background context
func (c *Client) ReactivateItem(itemId string, opts ...Option) (*Item, error) {
	ctx := context.Background()
	return c.reactivateItem(ctx, itemId, opts...)
}

// ReactivateItemWithContext Reactivate an inactive item
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/reactivate_item
//
// Returns: An item.
func (c *Client) ReactivateItemWithContext(ctx context.Context, itemId string, opts ...Option) (*Item, error) {
	return c.reactivateItem(ctx, itemId, opts...)
}

func (c *Client) reactivateItem(ctx context.Context, itemId string, opts ...Option) (*Item, error) {
	path, err := c.InterpolatePath("/items/{item_id}/reactivate", itemId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Item{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_measured_unit
//
// Returns: A list of the site's measured units.
func (c *Client) ListMeasuredUnit(params *ListMeasuredUnitParams, opts ...Option) *MeasuredUnitList {
	path, err := c.InterpolatePath("/measured_units")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewMeasuredUnitList(c, path, requestOptions)
}

// CreateMeasuredUnit wraps CreateMeasuredUnitWithContext using the background context
func (c *Client) CreateMeasuredUnit(body *MeasuredUnitCreate, opts ...Option) (*MeasuredUnit, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createMeasuredUnit(ctx, body, opts...)
}

// CreateMeasuredUnitWithContext Create a new measured unit
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_measured_unit
//
// Returns: A new measured unit.
func (c *Client) CreateMeasuredUnitWithContext(ctx context.Context, body *MeasuredUnitCreate, opts ...Option) (*MeasuredUnit, error) {
	return c.createMeasuredUnit(ctx, body, opts...)
}

func (c *Client) createMeasuredUnit(ctx context.Context, body *MeasuredUnitCreate, opts ...Option) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &MeasuredUnit{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetMeasuredUnit wraps GetMeasuredUnitWithContext using the background context
func (c *Client) GetMeasuredUnit(measuredUnitId string, opts ...Option) (*MeasuredUnit, error) {
	ctx := context.Background()
	return c.getMeasuredUnit(ctx, measuredUnitId, opts...)
}

// GetMeasuredUnitWithContext Fetch a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_measured_unit
//
// Returns: An item.
func (c *Client) GetMeasuredUnitWithContext(ctx context.Context, measuredUnitId string, opts ...Option) (*MeasuredUnit, error) {
	return c.getMeasuredUnit(ctx, measuredUnitId, opts...)
}

func (c *Client) getMeasuredUnit(ctx context.Context, measuredUnitId string, opts ...Option) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &MeasuredUnit{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateMeasuredUnit wraps UpdateMeasuredUnitWithContext using the background context
func (c *Client) UpdateMeasuredUnit(measuredUnitId string, body *MeasuredUnitUpdate, opts ...Option) (*MeasuredUnit, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateMeasuredUnit(ctx, measuredUnitId, body, opts...)
}

// UpdateMeasuredUnitWithContext Update a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_measured_unit
//
// Returns: The updated measured_unit.
func (c *Client) UpdateMeasuredUnitWithContext(ctx context.Context, measuredUnitId string, body *MeasuredUnitUpdate, opts ...Option) (*MeasuredUnit, error) {
	return c.updateMeasuredUnit(ctx, measuredUnitId, body, opts...)
}

func (c *Client) updateMeasuredUnit(ctx context.Context, measuredUnitId string, body *MeasuredUnitUpdate, opts ...Option) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &MeasuredUnit{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveMeasuredUnit wraps RemoveMeasuredUnitWithContext using the background context
func (c *Client) RemoveMeasuredUnit(measuredUnitId string, opts ...Option) (*MeasuredUnit, error) {
	ctx := context.Background()
	return c.removeMeasuredUnit(ctx, measuredUnitId, opts...)
}

// RemoveMeasuredUnitWithContext Remove a measured unit
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_measured_unit
//
// Returns: A measured unit.
func (c *Client) RemoveMeasuredUnitWithContext(ctx context.Context, measuredUnitId string, opts ...Option) (*MeasuredUnit, error) {
	return c.removeMeasuredUnit(ctx, measuredUnitId, opts...)
}

func (c *Client) removeMeasuredUnit(ctx context.Context, measuredUnitId string, opts ...Option) (*MeasuredUnit, error) {
	path, err := c.InterpolatePath("/measured_units/{measured_unit_id}", measuredUnitId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &MeasuredUnit{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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

	// State - Invoice state.
	State *string

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

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_invoices
//
// Returns: A list of the site's invoices.
func (c *Client) ListInvoices(params *ListInvoicesParams, opts ...Option) *InvoiceList {
	path, err := c.InterpolatePath("/invoices")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewInvoiceList(c, path, requestOptions)
}

// GetInvoice wraps GetInvoiceWithContext using the background context
func (c *Client) GetInvoice(invoiceId string, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	return c.getInvoice(ctx, invoiceId, opts...)
}

// GetInvoiceWithContext Fetch an invoice
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_invoice
//
// Returns: An invoice.
func (c *Client) GetInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	return c.getInvoice(ctx, invoiceId, opts...)
}

func (c *Client) getInvoice(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PutInvoice wraps PutInvoiceWithContext using the background context
func (c *Client) PutInvoice(invoiceId string, body *InvoiceUpdatable, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.putInvoice(ctx, invoiceId, body, opts...)
}

// PutInvoiceWithContext Update an invoice
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/put_invoice
//
// Returns: An invoice.
func (c *Client) PutInvoiceWithContext(ctx context.Context, invoiceId string, body *InvoiceUpdatable, opts ...Option) (*Invoice, error) {
	return c.putInvoice(ctx, invoiceId, body, opts...)
}

func (c *Client) putInvoice(ctx context.Context, invoiceId string, body *InvoiceUpdatable, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ApplyCreditBalance wraps ApplyCreditBalanceWithContext using the background context
func (c *Client) ApplyCreditBalance(invoiceId string, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	return c.applyCreditBalance(ctx, invoiceId, opts...)
}

// ApplyCreditBalanceWithContext Apply available credit to a pending or past due charge invoice
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/apply_credit_balance
//
// Returns: The updated invoice.
func (c *Client) ApplyCreditBalanceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	return c.applyCreditBalance(ctx, invoiceId, opts...)
}

func (c *Client) applyCreditBalance(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/apply_credit_balance", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
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

func (list *CollectInvoiceParams) URLParams() []KeyValue {
	var options []KeyValue

	return options
}

// CollectInvoice wraps CollectInvoiceWithContext using the background context
func (c *Client) CollectInvoice(invoiceId string, params *CollectInvoiceParams, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	return c.collectInvoice(ctx, invoiceId, params, opts...)
}

// CollectInvoiceWithContext Collect a pending or past due, automatic invoice
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/collect_invoice
//
// Returns: The updated invoice.
func (c *Client) CollectInvoiceWithContext(ctx context.Context, invoiceId string, params *CollectInvoiceParams, opts ...Option) (*Invoice, error) {
	return c.collectInvoice(ctx, invoiceId, params, opts...)
}

func (c *Client) collectInvoice(ctx context.Context, invoiceId string, params *CollectInvoiceParams, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/collect", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPut, path, nil, params, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// FailInvoice wraps FailInvoiceWithContext using the background context
func (c *Client) FailInvoice(invoiceId string, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	return c.failInvoice(ctx, invoiceId, opts...)
}

// FailInvoiceWithContext Mark an open invoice as failed
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/fail_invoice
//
// Returns: The updated invoice.
func (c *Client) FailInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	return c.failInvoice(ctx, invoiceId, opts...)
}

func (c *Client) failInvoice(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/mark_failed", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// MarkInvoiceSuccessful wraps MarkInvoiceSuccessfulWithContext using the background context
func (c *Client) MarkInvoiceSuccessful(invoiceId string, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	return c.markInvoiceSuccessful(ctx, invoiceId, opts...)
}

// MarkInvoiceSuccessfulWithContext Mark an open invoice as successful
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/mark_invoice_successful
//
// Returns: The updated invoice.
func (c *Client) MarkInvoiceSuccessfulWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	return c.markInvoiceSuccessful(ctx, invoiceId, opts...)
}

func (c *Client) markInvoiceSuccessful(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/mark_successful", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReopenInvoice wraps ReopenInvoiceWithContext using the background context
func (c *Client) ReopenInvoice(invoiceId string, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	return c.reopenInvoice(ctx, invoiceId, opts...)
}

// ReopenInvoiceWithContext Reopen a closed, manual invoice
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/reopen_invoice
//
// Returns: The updated invoice.
func (c *Client) ReopenInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	return c.reopenInvoice(ctx, invoiceId, opts...)
}

func (c *Client) reopenInvoice(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/reopen", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// VoidInvoice wraps VoidInvoiceWithContext using the background context
func (c *Client) VoidInvoice(invoiceId string, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	return c.voidInvoice(ctx, invoiceId, opts...)
}

// VoidInvoiceWithContext Void a credit invoice.
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/void_invoice
//
// Returns: The updated invoice.
func (c *Client) VoidInvoiceWithContext(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	return c.voidInvoice(ctx, invoiceId, opts...)
}

func (c *Client) voidInvoice(ctx context.Context, invoiceId string, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/void", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RecordExternalTransaction wraps RecordExternalTransactionWithContext using the background context
func (c *Client) RecordExternalTransaction(invoiceId string, body *ExternalTransaction, opts ...Option) (*Transaction, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.recordExternalTransaction(ctx, invoiceId, body, opts...)
}

// RecordExternalTransactionWithContext Record an external payment for a manual invoices.
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/record_external_transaction
//
// Returns: The recorded transaction.
func (c *Client) RecordExternalTransactionWithContext(ctx context.Context, invoiceId string, body *ExternalTransaction, opts ...Option) (*Transaction, error) {
	return c.recordExternalTransaction(ctx, invoiceId, body, opts...)
}

func (c *Client) recordExternalTransaction(ctx context.Context, invoiceId string, body *ExternalTransaction, opts ...Option) (*Transaction, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/transactions", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Transaction{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_invoice_line_items
//
// Returns: A list of the invoice's line items.
func (c *Client) ListInvoiceLineItems(invoiceId string, params *ListInvoiceLineItemsParams, opts ...Option) *LineItemList {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/line_items", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewLineItemList(c, path, requestOptions)
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

// ListInvoiceCouponRedemptions List the coupon redemptions applied to an invoice
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_invoice_coupon_redemptions
//
// Returns: A list of the the coupon redemptions associated with the invoice.
func (c *Client) ListInvoiceCouponRedemptions(invoiceId string, params *ListInvoiceCouponRedemptionsParams, opts ...Option) *CouponRedemptionList {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/coupon_redemptions", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewCouponRedemptionList(c, path, requestOptions)
}

// ListRelatedInvoices List an invoice's related credit or charge invoices
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_related_invoices
//
// Returns: A list of the credit or charge invoices associated with the invoice.
func (c *Client) ListRelatedInvoices(invoiceId string, opts ...Option) *InvoiceList {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/related_invoices", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	return NewInvoiceList(c, path, requestOptions)
}

// RefundInvoice wraps RefundInvoiceWithContext using the background context
func (c *Client) RefundInvoice(invoiceId string, body *InvoiceRefund, opts ...Option) (*Invoice, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.refundInvoice(ctx, invoiceId, body, opts...)
}

// RefundInvoiceWithContext Refund an invoice
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/refund_invoice
//
// Returns: Returns the new credit invoice.
func (c *Client) RefundInvoiceWithContext(ctx context.Context, invoiceId string, body *InvoiceRefund, opts ...Option) (*Invoice, error) {
	return c.refundInvoice(ctx, invoiceId, body, opts...)
}

func (c *Client) refundInvoice(ctx context.Context, invoiceId string, body *InvoiceRefund, opts ...Option) (*Invoice, error) {
	path, err := c.InterpolatePath("/invoices/{invoice_id}/refund", invoiceId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Invoice{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_line_items
//
// Returns: A list of the site's line items.
func (c *Client) ListLineItems(params *ListLineItemsParams, opts ...Option) *LineItemList {
	path, err := c.InterpolatePath("/line_items")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewLineItemList(c, path, requestOptions)
}

// GetLineItem wraps GetLineItemWithContext using the background context
func (c *Client) GetLineItem(lineItemId string, opts ...Option) (*LineItem, error) {
	ctx := context.Background()
	return c.getLineItem(ctx, lineItemId, opts...)
}

// GetLineItemWithContext Fetch a line item
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_line_item
//
// Returns: A line item.
func (c *Client) GetLineItemWithContext(ctx context.Context, lineItemId string, opts ...Option) (*LineItem, error) {
	return c.getLineItem(ctx, lineItemId, opts...)
}

func (c *Client) getLineItem(ctx context.Context, lineItemId string, opts ...Option) (*LineItem, error) {
	path, err := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &LineItem{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveLineItem wraps RemoveLineItemWithContext using the background context
func (c *Client) RemoveLineItem(lineItemId string, opts ...Option) (*Empty, error) {
	ctx := context.Background()
	return c.removeLineItem(ctx, lineItemId, opts...)
}

// RemoveLineItemWithContext Delete an uninvoiced line item
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_line_item
//
// Returns: Line item deleted.
func (c *Client) RemoveLineItemWithContext(ctx context.Context, lineItemId string, opts ...Option) (*Empty, error) {
	return c.removeLineItem(ctx, lineItemId, opts...)
}

func (c *Client) removeLineItem(ctx context.Context, lineItemId string, opts ...Option) (*Empty, error) {
	path, err := c.InterpolatePath("/line_items/{line_item_id}", lineItemId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Empty{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_plans
//
// Returns: A list of plans.
func (c *Client) ListPlans(params *ListPlansParams, opts ...Option) *PlanList {
	path, err := c.InterpolatePath("/plans")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewPlanList(c, path, requestOptions)
}

// CreatePlan wraps CreatePlanWithContext using the background context
func (c *Client) CreatePlan(body *PlanCreate, opts ...Option) (*Plan, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createPlan(ctx, body, opts...)
}

// CreatePlanWithContext Create a plan
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_plan
//
// Returns: A plan.
func (c *Client) CreatePlanWithContext(ctx context.Context, body *PlanCreate, opts ...Option) (*Plan, error) {
	return c.createPlan(ctx, body, opts...)
}

func (c *Client) createPlan(ctx context.Context, body *PlanCreate, opts ...Option) (*Plan, error) {
	path, err := c.InterpolatePath("/plans")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Plan{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetPlan wraps GetPlanWithContext using the background context
func (c *Client) GetPlan(planId string, opts ...Option) (*Plan, error) {
	ctx := context.Background()
	return c.getPlan(ctx, planId, opts...)
}

// GetPlanWithContext Fetch a plan
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_plan
//
// Returns: A plan.
func (c *Client) GetPlanWithContext(ctx context.Context, planId string, opts ...Option) (*Plan, error) {
	return c.getPlan(ctx, planId, opts...)
}

func (c *Client) getPlan(ctx context.Context, planId string, opts ...Option) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Plan{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdatePlan wraps UpdatePlanWithContext using the background context
func (c *Client) UpdatePlan(planId string, body *PlanUpdate, opts ...Option) (*Plan, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updatePlan(ctx, planId, body, opts...)
}

// UpdatePlanWithContext Update a plan
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_plan
//
// Returns: A plan.
func (c *Client) UpdatePlanWithContext(ctx context.Context, planId string, body *PlanUpdate, opts ...Option) (*Plan, error) {
	return c.updatePlan(ctx, planId, body, opts...)
}

func (c *Client) updatePlan(ctx context.Context, planId string, body *PlanUpdate, opts ...Option) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Plan{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemovePlan wraps RemovePlanWithContext using the background context
func (c *Client) RemovePlan(planId string, opts ...Option) (*Plan, error) {
	ctx := context.Background()
	return c.removePlan(ctx, planId, opts...)
}

// RemovePlanWithContext Remove a plan
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_plan
//
// Returns: Plan deleted
func (c *Client) RemovePlanWithContext(ctx context.Context, planId string, opts ...Option) (*Plan, error) {
	return c.removePlan(ctx, planId, opts...)
}

func (c *Client) removePlan(ctx context.Context, planId string, opts ...Option) (*Plan, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}", planId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Plan{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_plan_add_ons
//
// Returns: A list of add-ons.
func (c *Client) ListPlanAddOns(planId string, params *ListPlanAddOnsParams, opts ...Option) *AddOnList {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewAddOnList(c, path, requestOptions)
}

// CreatePlanAddOn wraps CreatePlanAddOnWithContext using the background context
func (c *Client) CreatePlanAddOn(planId string, body *AddOnCreate, opts ...Option) (*AddOn, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createPlanAddOn(ctx, planId, body, opts...)
}

// CreatePlanAddOnWithContext Create an add-on
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_plan_add_on
//
// Returns: An add-on.
func (c *Client) CreatePlanAddOnWithContext(ctx context.Context, planId string, body *AddOnCreate, opts ...Option) (*AddOn, error) {
	return c.createPlanAddOn(ctx, planId, body, opts...)
}

func (c *Client) createPlanAddOn(ctx context.Context, planId string, body *AddOnCreate, opts ...Option) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons", planId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &AddOn{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetPlanAddOn wraps GetPlanAddOnWithContext using the background context
func (c *Client) GetPlanAddOn(planId string, addOnId string, opts ...Option) (*AddOn, error) {
	ctx := context.Background()
	return c.getPlanAddOn(ctx, planId, addOnId, opts...)
}

// GetPlanAddOnWithContext Fetch a plan's add-on
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_plan_add_on
//
// Returns: An add-on.
func (c *Client) GetPlanAddOnWithContext(ctx context.Context, planId string, addOnId string, opts ...Option) (*AddOn, error) {
	return c.getPlanAddOn(ctx, planId, addOnId, opts...)
}

func (c *Client) getPlanAddOn(ctx context.Context, planId string, addOnId string, opts ...Option) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &AddOn{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdatePlanAddOn wraps UpdatePlanAddOnWithContext using the background context
func (c *Client) UpdatePlanAddOn(planId string, addOnId string, body *AddOnUpdate, opts ...Option) (*AddOn, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updatePlanAddOn(ctx, planId, addOnId, body, opts...)
}

// UpdatePlanAddOnWithContext Update an add-on
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_plan_add_on
//
// Returns: An add-on.
func (c *Client) UpdatePlanAddOnWithContext(ctx context.Context, planId string, addOnId string, body *AddOnUpdate, opts ...Option) (*AddOn, error) {
	return c.updatePlanAddOn(ctx, planId, addOnId, body, opts...)
}

func (c *Client) updatePlanAddOn(ctx context.Context, planId string, addOnId string, body *AddOnUpdate, opts ...Option) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &AddOn{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemovePlanAddOn wraps RemovePlanAddOnWithContext using the background context
func (c *Client) RemovePlanAddOn(planId string, addOnId string, opts ...Option) (*AddOn, error) {
	ctx := context.Background()
	return c.removePlanAddOn(ctx, planId, addOnId, opts...)
}

// RemovePlanAddOnWithContext Remove an add-on
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_plan_add_on
//
// Returns: Add-on deleted
func (c *Client) RemovePlanAddOnWithContext(ctx context.Context, planId string, addOnId string, opts ...Option) (*AddOn, error) {
	return c.removePlanAddOn(ctx, planId, addOnId, opts...)
}

func (c *Client) removePlanAddOn(ctx context.Context, planId string, addOnId string, opts ...Option) (*AddOn, error) {
	path, err := c.InterpolatePath("/plans/{plan_id}/add_ons/{add_on_id}", planId, addOnId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &AddOn{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_add_ons
//
// Returns: A list of add-ons.
func (c *Client) ListAddOns(params *ListAddOnsParams, opts ...Option) *AddOnList {
	path, err := c.InterpolatePath("/add_ons")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewAddOnList(c, path, requestOptions)
}

// GetAddOn wraps GetAddOnWithContext using the background context
func (c *Client) GetAddOn(addOnId string, opts ...Option) (*AddOn, error) {
	ctx := context.Background()
	return c.getAddOn(ctx, addOnId, opts...)
}

// GetAddOnWithContext Fetch an add-on
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_add_on
//
// Returns: An add-on.
func (c *Client) GetAddOnWithContext(ctx context.Context, addOnId string, opts ...Option) (*AddOn, error) {
	return c.getAddOn(ctx, addOnId, opts...)
}

func (c *Client) getAddOn(ctx context.Context, addOnId string, opts ...Option) (*AddOn, error) {
	path, err := c.InterpolatePath("/add_ons/{add_on_id}", addOnId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &AddOn{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_shipping_methods
//
// Returns: A list of the site's shipping methods.
func (c *Client) ListShippingMethods(params *ListShippingMethodsParams, opts ...Option) *ShippingMethodList {
	path, err := c.InterpolatePath("/shipping_methods")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewShippingMethodList(c, path, requestOptions)
}

// CreateShippingMethod wraps CreateShippingMethodWithContext using the background context
func (c *Client) CreateShippingMethod(body *ShippingMethodCreate, opts ...Option) (*ShippingMethod, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createShippingMethod(ctx, body, opts...)
}

// CreateShippingMethodWithContext Create a new shipping method
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_shipping_method
//
// Returns: A new shipping method.
func (c *Client) CreateShippingMethodWithContext(ctx context.Context, body *ShippingMethodCreate, opts ...Option) (*ShippingMethod, error) {
	return c.createShippingMethod(ctx, body, opts...)
}

func (c *Client) createShippingMethod(ctx context.Context, body *ShippingMethodCreate, opts ...Option) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &ShippingMethod{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetShippingMethod wraps GetShippingMethodWithContext using the background context
func (c *Client) GetShippingMethod(id string, opts ...Option) (*ShippingMethod, error) {
	ctx := context.Background()
	return c.getShippingMethod(ctx, id, opts...)
}

// GetShippingMethodWithContext Fetch a shipping method
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_shipping_method
//
// Returns: A shipping method.
func (c *Client) GetShippingMethodWithContext(ctx context.Context, id string, opts ...Option) (*ShippingMethod, error) {
	return c.getShippingMethod(ctx, id, opts...)
}

func (c *Client) getShippingMethod(ctx context.Context, id string, opts ...Option) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{id}", id)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &ShippingMethod{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateShippingMethod wraps UpdateShippingMethodWithContext using the background context
func (c *Client) UpdateShippingMethod(shippingMethodId string, body *ShippingMethodUpdate, opts ...Option) (*ShippingMethod, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateShippingMethod(ctx, shippingMethodId, body, opts...)
}

// UpdateShippingMethodWithContext Update an active Shipping Method
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_shipping_method
//
// Returns: The updated shipping method.
func (c *Client) UpdateShippingMethodWithContext(ctx context.Context, shippingMethodId string, body *ShippingMethodUpdate, opts ...Option) (*ShippingMethod, error) {
	return c.updateShippingMethod(ctx, shippingMethodId, body, opts...)
}

func (c *Client) updateShippingMethod(ctx context.Context, shippingMethodId string, body *ShippingMethodUpdate, opts ...Option) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &ShippingMethod{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateShippingMethod wraps DeactivateShippingMethodWithContext using the background context
func (c *Client) DeactivateShippingMethod(shippingMethodId string, opts ...Option) (*ShippingMethod, error) {
	ctx := context.Background()
	return c.deactivateShippingMethod(ctx, shippingMethodId, opts...)
}

// DeactivateShippingMethodWithContext Deactivate a shipping method
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/deactivate_shipping_method
//
// Returns: A shipping method.
func (c *Client) DeactivateShippingMethodWithContext(ctx context.Context, shippingMethodId string, opts ...Option) (*ShippingMethod, error) {
	return c.deactivateShippingMethod(ctx, shippingMethodId, opts...)
}

func (c *Client) deactivateShippingMethod(ctx context.Context, shippingMethodId string, opts ...Option) (*ShippingMethod, error) {
	path, err := c.InterpolatePath("/shipping_methods/{shipping_method_id}", shippingMethodId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &ShippingMethod{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_subscriptions
//
// Returns: A list of the site's subscriptions.
func (c *Client) ListSubscriptions(params *ListSubscriptionsParams, opts ...Option) *SubscriptionList {
	path, err := c.InterpolatePath("/subscriptions")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewSubscriptionList(c, path, requestOptions)
}

// CreateSubscription wraps CreateSubscriptionWithContext using the background context
func (c *Client) CreateSubscription(body *SubscriptionCreate, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createSubscription(ctx, body, opts...)
}

// CreateSubscriptionWithContext Create a new subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_subscription
//
// Returns: A subscription.
func (c *Client) CreateSubscriptionWithContext(ctx context.Context, body *SubscriptionCreate, opts ...Option) (*Subscription, error) {
	return c.createSubscription(ctx, body, opts...)
}

func (c *Client) createSubscription(ctx context.Context, body *SubscriptionCreate, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetSubscription wraps GetSubscriptionWithContext using the background context
func (c *Client) GetSubscription(subscriptionId string, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	return c.getSubscription(ctx, subscriptionId, opts...)
}

// GetSubscriptionWithContext Fetch a subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_subscription
//
// Returns: A subscription.
func (c *Client) GetSubscriptionWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	return c.getSubscription(ctx, subscriptionId, opts...)
}

func (c *Client) getSubscription(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ModifySubscription wraps ModifySubscriptionWithContext using the background context
func (c *Client) ModifySubscription(subscriptionId string, body *SubscriptionUpdate, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.modifySubscription(ctx, subscriptionId, body, opts...)
}

// ModifySubscriptionWithContext Modify a subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/modify_subscription
//
// Returns: A subscription.
func (c *Client) ModifySubscriptionWithContext(ctx context.Context, subscriptionId string, body *SubscriptionUpdate, opts ...Option) (*Subscription, error) {
	return c.modifySubscription(ctx, subscriptionId, body, opts...)
}

func (c *Client) modifySubscription(ctx context.Context, subscriptionId string, body *SubscriptionUpdate, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
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

	// Charge - Applicable only if the subscription has usage based add-ons and unbilled usage logged for the current billing cycle. If true, current billing cycle unbilled usage is billed on the final invoice. If false, Recurly will create a negative usage record for current billing cycle usage that will zero out the final invoice line items.
	Charge *bool
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

	if list.Charge != nil {
		options = append(options, KeyValue{Key: "charge", Value: strconv.FormatBool(*list.Charge)})
	}

	return options
}

// TerminateSubscription wraps TerminateSubscriptionWithContext using the background context
func (c *Client) TerminateSubscription(subscriptionId string, params *TerminateSubscriptionParams, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	return c.terminateSubscription(ctx, subscriptionId, params, opts...)
}

// TerminateSubscriptionWithContext Terminate a subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/terminate_subscription
//
// Returns: An expired subscription.
func (c *Client) TerminateSubscriptionWithContext(ctx context.Context, subscriptionId string, params *TerminateSubscriptionParams, opts ...Option) (*Subscription, error) {
	return c.terminateSubscription(ctx, subscriptionId, params, opts...)
}

func (c *Client) terminateSubscription(ctx context.Context, subscriptionId string, params *TerminateSubscriptionParams, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodDelete, path, nil, params, requestOptions, result)
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

func (list *CancelSubscriptionParams) URLParams() []KeyValue {
	var options []KeyValue

	return options
}

// CancelSubscription wraps CancelSubscriptionWithContext using the background context
func (c *Client) CancelSubscription(subscriptionId string, params *CancelSubscriptionParams, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	return c.cancelSubscription(ctx, subscriptionId, params, opts...)
}

// CancelSubscriptionWithContext Cancel a subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/cancel_subscription
//
// Returns: A canceled or failed subscription.
func (c *Client) CancelSubscriptionWithContext(ctx context.Context, subscriptionId string, params *CancelSubscriptionParams, opts ...Option) (*Subscription, error) {
	return c.cancelSubscription(ctx, subscriptionId, params, opts...)
}

func (c *Client) cancelSubscription(ctx context.Context, subscriptionId string, params *CancelSubscriptionParams, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/cancel", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodPut, path, nil, params, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateSubscription wraps ReactivateSubscriptionWithContext using the background context
func (c *Client) ReactivateSubscription(subscriptionId string, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	return c.reactivateSubscription(ctx, subscriptionId, opts...)
}

// ReactivateSubscriptionWithContext Reactivate a canceled subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/reactivate_subscription
//
// Returns: An active subscription.
func (c *Client) ReactivateSubscriptionWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	return c.reactivateSubscription(ctx, subscriptionId, opts...)
}

func (c *Client) reactivateSubscription(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/reactivate", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PauseSubscription wraps PauseSubscriptionWithContext using the background context
func (c *Client) PauseSubscription(subscriptionId string, body *SubscriptionPause, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.pauseSubscription(ctx, subscriptionId, body, opts...)
}

// PauseSubscriptionWithContext Pause subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/pause_subscription
//
// Returns: A subscription.
func (c *Client) PauseSubscriptionWithContext(ctx context.Context, subscriptionId string, body *SubscriptionPause, opts ...Option) (*Subscription, error) {
	return c.pauseSubscription(ctx, subscriptionId, body, opts...)
}

func (c *Client) pauseSubscription(ctx context.Context, subscriptionId string, body *SubscriptionPause, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/pause", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ResumeSubscription wraps ResumeSubscriptionWithContext using the background context
func (c *Client) ResumeSubscription(subscriptionId string, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	return c.resumeSubscription(ctx, subscriptionId, opts...)
}

// ResumeSubscriptionWithContext Resume subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/resume_subscription
//
// Returns: A subscription.
func (c *Client) ResumeSubscriptionWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	return c.resumeSubscription(ctx, subscriptionId, opts...)
}

func (c *Client) resumeSubscription(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/resume", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ConvertTrial wraps ConvertTrialWithContext using the background context
func (c *Client) ConvertTrial(subscriptionId string, opts ...Option) (*Subscription, error) {
	ctx := context.Background()
	return c.convertTrial(ctx, subscriptionId, opts...)
}

// ConvertTrialWithContext Convert trial subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/convert_trial
//
// Returns: A subscription.
func (c *Client) ConvertTrialWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	return c.convertTrial(ctx, subscriptionId, opts...)
}

func (c *Client) convertTrial(ctx context.Context, subscriptionId string, opts ...Option) (*Subscription, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/convert_trial", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Subscription{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetPreviewRenewal wraps GetPreviewRenewalWithContext using the background context
func (c *Client) GetPreviewRenewal(subscriptionId string, opts ...Option) (*InvoiceCollection, error) {
	ctx := context.Background()
	return c.getPreviewRenewal(ctx, subscriptionId, opts...)
}

// GetPreviewRenewalWithContext Fetch a preview of a subscription's renewal invoice(s)
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_preview_renewal
//
// Returns: A preview of the subscription's renewal invoice(s).
func (c *Client) GetPreviewRenewalWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*InvoiceCollection, error) {
	return c.getPreviewRenewal(ctx, subscriptionId, opts...)
}

func (c *Client) getPreviewRenewal(ctx context.Context, subscriptionId string, opts ...Option) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/preview_renewal", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &InvoiceCollection{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetSubscriptionChange wraps GetSubscriptionChangeWithContext using the background context
func (c *Client) GetSubscriptionChange(subscriptionId string, opts ...Option) (*SubscriptionChange, error) {
	ctx := context.Background()
	return c.getSubscriptionChange(ctx, subscriptionId, opts...)
}

// GetSubscriptionChangeWithContext Fetch a subscription's pending change
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_subscription_change
//
// Returns: A subscription's pending change.
func (c *Client) GetSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*SubscriptionChange, error) {
	return c.getSubscriptionChange(ctx, subscriptionId, opts...)
}

func (c *Client) getSubscriptionChange(ctx context.Context, subscriptionId string, opts ...Option) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &SubscriptionChange{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// CreateSubscriptionChange wraps CreateSubscriptionChangeWithContext using the background context
func (c *Client) CreateSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChange, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createSubscriptionChange(ctx, subscriptionId, body, opts...)
}

// CreateSubscriptionChangeWithContext Create a new subscription change
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_subscription_change
//
// Returns: A subscription change.
func (c *Client) CreateSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChange, error) {
	return c.createSubscriptionChange(ctx, subscriptionId, body, opts...)
}

func (c *Client) createSubscriptionChange(ctx context.Context, subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChange, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &SubscriptionChange{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveSubscriptionChange wraps RemoveSubscriptionChangeWithContext using the background context
func (c *Client) RemoveSubscriptionChange(subscriptionId string, opts ...Option) (*Empty, error) {
	ctx := context.Background()
	return c.removeSubscriptionChange(ctx, subscriptionId, opts...)
}

// RemoveSubscriptionChangeWithContext Delete the pending subscription change
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_subscription_change
//
// Returns: Subscription change was deleted.
func (c *Client) RemoveSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, opts ...Option) (*Empty, error) {
	return c.removeSubscriptionChange(ctx, subscriptionId, opts...)
}

func (c *Client) removeSubscriptionChange(ctx context.Context, subscriptionId string, opts ...Option) (*Empty, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Empty{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PreviewSubscriptionChange wraps PreviewSubscriptionChangeWithContext using the background context
func (c *Client) PreviewSubscriptionChange(subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChangePreview, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.previewSubscriptionChange(ctx, subscriptionId, body, opts...)
}

// PreviewSubscriptionChangeWithContext Preview a new subscription change
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/preview_subscription_change
//
// Returns: A subscription change.
func (c *Client) PreviewSubscriptionChangeWithContext(ctx context.Context, subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChangePreview, error) {
	return c.previewSubscriptionChange(ctx, subscriptionId, body, opts...)
}

func (c *Client) previewSubscriptionChange(ctx context.Context, subscriptionId string, body *SubscriptionChangeCreate, opts ...Option) (*SubscriptionChangePreview, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/change/preview", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &SubscriptionChangePreview{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
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

	// State - Invoice state.
	State *string

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

	if list.State != nil {
		options = append(options, KeyValue{Key: "state", Value: *list.State})
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_subscription_invoices
//
// Returns: A list of the subscription's invoices.
func (c *Client) ListSubscriptionInvoices(subscriptionId string, params *ListSubscriptionInvoicesParams, opts ...Option) *InvoiceList {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/invoices", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewInvoiceList(c, path, requestOptions)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_subscription_line_items
//
// Returns: A list of the subscription's line items.
func (c *Client) ListSubscriptionLineItems(subscriptionId string, params *ListSubscriptionLineItemsParams, opts ...Option) *LineItemList {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/line_items", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewLineItemList(c, path, requestOptions)
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

// ListSubscriptionCouponRedemptions List the coupon redemptions for a subscription
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_subscription_coupon_redemptions
//
// Returns: A list of the the coupon redemptions on a subscription.
func (c *Client) ListSubscriptionCouponRedemptions(subscriptionId string, params *ListSubscriptionCouponRedemptionsParams, opts ...Option) *CouponRedemptionList {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/coupon_redemptions", subscriptionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewCouponRedemptionList(c, path, requestOptions)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_usage
//
// Returns: A list of the subscription add-on's usage records.
func (c *Client) ListUsage(subscriptionId string, addOnId string, params *ListUsageParams, opts ...Option) *UsageList {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/add_ons/{add_on_id}/usage", subscriptionId, addOnId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewUsageList(c, path, requestOptions)
}

// CreateUsage wraps CreateUsageWithContext using the background context
func (c *Client) CreateUsage(subscriptionId string, addOnId string, body *UsageCreate, opts ...Option) (*Usage, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createUsage(ctx, subscriptionId, addOnId, body, opts...)
}

// CreateUsageWithContext Log a usage record on this subscription add-on
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_usage
//
// Returns: The created usage record.
func (c *Client) CreateUsageWithContext(ctx context.Context, subscriptionId string, addOnId string, body *UsageCreate, opts ...Option) (*Usage, error) {
	return c.createUsage(ctx, subscriptionId, addOnId, body, opts...)
}

func (c *Client) createUsage(ctx context.Context, subscriptionId string, addOnId string, body *UsageCreate, opts ...Option) (*Usage, error) {
	path, err := c.InterpolatePath("/subscriptions/{subscription_id}/add_ons/{add_on_id}/usage", subscriptionId, addOnId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Usage{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetUsage wraps GetUsageWithContext using the background context
func (c *Client) GetUsage(usageId string, opts ...Option) (*Usage, error) {
	ctx := context.Background()
	return c.getUsage(ctx, usageId, opts...)
}

// GetUsageWithContext Get a usage record
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_usage
//
// Returns: The usage record.
func (c *Client) GetUsageWithContext(ctx context.Context, usageId string, opts ...Option) (*Usage, error) {
	return c.getUsage(ctx, usageId, opts...)
}

func (c *Client) getUsage(ctx context.Context, usageId string, opts ...Option) (*Usage, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Usage{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateUsage wraps UpdateUsageWithContext using the background context
func (c *Client) UpdateUsage(usageId string, body *UsageCreate, opts ...Option) (*Usage, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.updateUsage(ctx, usageId, body, opts...)
}

// UpdateUsageWithContext Update a usage record
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/update_usage
//
// Returns: The updated usage record.
func (c *Client) UpdateUsageWithContext(ctx context.Context, usageId string, body *UsageCreate, opts ...Option) (*Usage, error) {
	return c.updateUsage(ctx, usageId, body, opts...)
}

func (c *Client) updateUsage(ctx context.Context, usageId string, body *UsageCreate, opts ...Option) (*Usage, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &Usage{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// RemoveUsage wraps RemoveUsageWithContext using the background context
func (c *Client) RemoveUsage(usageId string, opts ...Option) (*Empty, error) {
	ctx := context.Background()
	return c.removeUsage(ctx, usageId, opts...)
}

// RemoveUsageWithContext Delete a usage record.
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/remove_usage
//
// Returns: Usage was successfully deleted.
func (c *Client) RemoveUsageWithContext(ctx context.Context, usageId string, opts ...Option) (*Empty, error) {
	return c.removeUsage(ctx, usageId, opts...)
}

func (c *Client) removeUsage(ctx context.Context, usageId string, opts ...Option) (*Empty, error) {
	path, err := c.InterpolatePath("/usage/{usage_id}", usageId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Empty{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
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
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_transactions
//
// Returns: A list of the site's transactions.
func (c *Client) ListTransactions(params *ListTransactionsParams, opts ...Option) *TransactionList {
	path, err := c.InterpolatePath("/transactions")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewTransactionList(c, path, requestOptions)
}

// GetTransaction wraps GetTransactionWithContext using the background context
func (c *Client) GetTransaction(transactionId string, opts ...Option) (*Transaction, error) {
	ctx := context.Background()
	return c.getTransaction(ctx, transactionId, opts...)
}

// GetTransactionWithContext Fetch a transaction
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_transaction
//
// Returns: A transaction.
func (c *Client) GetTransactionWithContext(ctx context.Context, transactionId string, opts ...Option) (*Transaction, error) {
	return c.getTransaction(ctx, transactionId, opts...)
}

func (c *Client) getTransaction(ctx context.Context, transactionId string, opts ...Option) (*Transaction, error) {
	path, err := c.InterpolatePath("/transactions/{transaction_id}", transactionId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &Transaction{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetUniqueCouponCode wraps GetUniqueCouponCodeWithContext using the background context
func (c *Client) GetUniqueCouponCode(uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	ctx := context.Background()
	return c.getUniqueCouponCode(ctx, uniqueCouponCodeId, opts...)
}

// GetUniqueCouponCodeWithContext Fetch a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) GetUniqueCouponCodeWithContext(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	return c.getUniqueCouponCode(ctx, uniqueCouponCodeId, opts...)
}

func (c *Client) getUniqueCouponCode(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &UniqueCouponCode{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// DeactivateUniqueCouponCode wraps DeactivateUniqueCouponCodeWithContext using the background context
func (c *Client) DeactivateUniqueCouponCode(uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	ctx := context.Background()
	return c.deactivateUniqueCouponCode(ctx, uniqueCouponCodeId, opts...)
}

// DeactivateUniqueCouponCodeWithContext Deactivate a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/deactivate_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) DeactivateUniqueCouponCodeWithContext(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	return c.deactivateUniqueCouponCode(ctx, uniqueCouponCodeId, opts...)
}

func (c *Client) deactivateUniqueCouponCode(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}", uniqueCouponCodeId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &UniqueCouponCode{}
	err = c.Call(ctx, http.MethodDelete, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ReactivateUniqueCouponCode wraps ReactivateUniqueCouponCodeWithContext using the background context
func (c *Client) ReactivateUniqueCouponCode(uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	ctx := context.Background()
	return c.reactivateUniqueCouponCode(ctx, uniqueCouponCodeId, opts...)
}

// ReactivateUniqueCouponCodeWithContext Restore a unique coupon code
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/reactivate_unique_coupon_code
//
// Returns: A unique coupon code.
func (c *Client) ReactivateUniqueCouponCodeWithContext(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	return c.reactivateUniqueCouponCode(ctx, uniqueCouponCodeId, opts...)
}

func (c *Client) reactivateUniqueCouponCode(ctx context.Context, uniqueCouponCodeId string, opts ...Option) (*UniqueCouponCode, error) {
	path, err := c.InterpolatePath("/unique_coupon_codes/{unique_coupon_code_id}/restore", uniqueCouponCodeId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &UniqueCouponCode{}
	err = c.Call(ctx, http.MethodPut, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// CreatePurchase wraps CreatePurchaseWithContext using the background context
func (c *Client) CreatePurchase(body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.createPurchase(ctx, body, opts...)
}

// CreatePurchaseWithContext Create a new purchase
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/create_purchase
//
// Returns: Returns the new invoices
func (c *Client) CreatePurchaseWithContext(ctx context.Context, body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error) {
	return c.createPurchase(ctx, body, opts...)
}

func (c *Client) createPurchase(ctx context.Context, body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/purchases")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &InvoiceCollection{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PreviewPurchase wraps PreviewPurchaseWithContext using the background context
func (c *Client) PreviewPurchase(body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.previewPurchase(ctx, body, opts...)
}

// PreviewPurchaseWithContext Preview a new purchase
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/preview_purchase
//
// Returns: Returns preview of the new invoices
func (c *Client) PreviewPurchaseWithContext(ctx context.Context, body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error) {
	return c.previewPurchase(ctx, body, opts...)
}

func (c *Client) previewPurchase(ctx context.Context, body *PurchaseCreate, opts ...Option) (*InvoiceCollection, error) {
	path, err := c.InterpolatePath("/purchases/preview")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &InvoiceCollection{}
	err = c.Call(ctx, http.MethodPost, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetExportDates wraps GetExportDatesWithContext using the background context
func (c *Client) GetExportDates(opts ...Option) (*ExportDates, error) {
	ctx := context.Background()
	return c.getExportDates(ctx, opts...)
}

// GetExportDatesWithContext List the dates that have an available export to download.
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_export_dates
//
// Returns: Returns a list of dates.
func (c *Client) GetExportDatesWithContext(ctx context.Context, opts ...Option) (*ExportDates, error) {
	return c.getExportDates(ctx, opts...)
}

func (c *Client) getExportDates(ctx context.Context, opts ...Option) (*ExportDates, error) {
	path, err := c.InterpolatePath("/export_dates")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &ExportDates{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// GetExportFiles wraps GetExportFilesWithContext using the background context
func (c *Client) GetExportFiles(exportDate string, opts ...Option) (*ExportFiles, error) {
	ctx := context.Background()
	return c.getExportFiles(ctx, exportDate, opts...)
}

// GetExportFilesWithContext List of the export files that are available to download.
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_export_files
//
// Returns: Returns a list of export files to download.
func (c *Client) GetExportFilesWithContext(ctx context.Context, exportDate string, opts ...Option) (*ExportFiles, error) {
	return c.getExportFiles(ctx, exportDate, opts...)
}

func (c *Client) getExportFiles(ctx context.Context, exportDate string, opts ...Option) (*ExportFiles, error) {
	path, err := c.InterpolatePath("/export_dates/{export_date}/export_files", exportDate)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &ExportFiles{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

type ListDunningCampaignsParams struct {
	Params

	// Sort - Sort field. You *really* only want to sort by `updated_at` in ascending
	// order. In descending order updated records will move behind the cursor and could
	// prevent some records from being returned.
	Sort *string
}

func (list *ListDunningCampaignsParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *ListDunningCampaignsParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.Sort != nil {
		options = append(options, KeyValue{Key: "sort", Value: *list.Sort})
	}

	return options
}

// ListDunningCampaigns List the dunning campaigns for a site
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/list_dunning_campaigns
//
// Returns: A list of the the dunning_campaigns on an account.
func (c *Client) ListDunningCampaigns(params *ListDunningCampaignsParams, opts ...Option) *DunningCampaignList {
	path, err := c.InterpolatePath("/dunning_campaigns")
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(params.Params, opts...)
	path = BuildURL(path, params)
	return NewDunningCampaignList(c, path, requestOptions)
}

// GetDunningCampaign wraps GetDunningCampaignWithContext using the background context
func (c *Client) GetDunningCampaign(dunningCampaignId string, opts ...Option) (*DunningCampaign, error) {
	ctx := context.Background()
	return c.getDunningCampaign(ctx, dunningCampaignId, opts...)
}

// GetDunningCampaignWithContext Fetch a dunning campaign
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/get_dunning_campaign
//
// Returns: Settings for a dunning campaign.
func (c *Client) GetDunningCampaignWithContext(ctx context.Context, dunningCampaignId string, opts ...Option) (*DunningCampaign, error) {
	return c.getDunningCampaign(ctx, dunningCampaignId, opts...)
}

func (c *Client) getDunningCampaign(ctx context.Context, dunningCampaignId string, opts ...Option) (*DunningCampaign, error) {
	path, err := c.InterpolatePath("/dunning_campaigns/{dunning_campaign_id}", dunningCampaignId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := NewRequestOptions(opts...)
	result := &DunningCampaign{}
	err = c.Call(ctx, http.MethodGet, path, nil, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// PutDunningCampaignBulkUpdate wraps PutDunningCampaignBulkUpdateWithContext using the background context
func (c *Client) PutDunningCampaignBulkUpdate(dunningCampaignId string, body *DunningCampaignsBulkUpdate, opts ...Option) (*DunningCampaignsBulkUpdateResponse, error) {
	ctx := context.Background()
	if body.Params.Context != nil {
		ctx = body.Params.Context
	}
	return c.putDunningCampaignBulkUpdate(ctx, dunningCampaignId, body, opts...)
}

// PutDunningCampaignBulkUpdateWithContext Assign a dunning campaign to multiple plans
//
// API Documentation: https://developers.recurly.com/api/v2019-10-10#operation/put_dunning_campaign_bulk_update
//
// Returns: A list of updated plans.
func (c *Client) PutDunningCampaignBulkUpdateWithContext(ctx context.Context, dunningCampaignId string, body *DunningCampaignsBulkUpdate, opts ...Option) (*DunningCampaignsBulkUpdateResponse, error) {
	return c.putDunningCampaignBulkUpdate(ctx, dunningCampaignId, body, opts...)
}

func (c *Client) putDunningCampaignBulkUpdate(ctx context.Context, dunningCampaignId string, body *DunningCampaignsBulkUpdate, opts ...Option) (*DunningCampaignsBulkUpdateResponse, error) {
	path, err := c.InterpolatePath("/dunning_campaigns/{dunning_campaign_id}/bulk_update", dunningCampaignId)
	if err != nil {
		// NOOP in 3.x client
	}
	requestOptions := RequestOptionsFromParams(body.Params, opts...)
	result := &DunningCampaignsBulkUpdateResponse{}
	err = c.Call(ctx, http.MethodPut, path, body, nil, requestOptions, result)
	if err != nil {
		return nil, err
	}
	return result, err
}
