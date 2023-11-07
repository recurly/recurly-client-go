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

type Account struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Accounts can be either active or inactive.
	State string `json:"state,omitempty"`

	// The unique token for automatically logging the account in to the hosted management pages. You may automatically log the user into their hosted management pages by directing the user to: `https://{subdomain}.recurly.com/account/{hosted_login_token}`.
	HostedLoginToken string `json:"hosted_login_token,omitempty"`

	// The shipping addresses on the account.
	ShippingAddresses []ShippingAddress `json:"shipping_addresses,omitempty"`

	// Indicates if the account has a subscription that is either active, canceled, future, or paused.
	HasLiveSubscription bool `json:"has_live_subscription,omitempty"`

	// Indicates if the account has an active subscription.
	HasActiveSubscription bool `json:"has_active_subscription,omitempty"`

	// Indicates if the account has a future subscription.
	HasFutureSubscription bool `json:"has_future_subscription,omitempty"`

	// Indicates if the account has a canceled subscription.
	HasCanceledSubscription bool `json:"has_canceled_subscription,omitempty"`

	// Indicates if the account has a paused subscription.
	HasPausedSubscription bool `json:"has_paused_subscription,omitempty"`

	// Indicates if the account has a past due invoice.
	HasPastDueInvoice bool `json:"has_past_due_invoice,omitempty"`

	// When the account was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the account was last changed.
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// If present, when the account was last marked inactive.
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// The unique identifier of the account. This cannot be changed once the account is created.
	Code string `json:"code,omitempty"`

	// A secondary value for the account.
	Username string `json:"username,omitempty"`

	// The email address used for communicating with this customer. The customer will also use this email address to log into your hosted account management pages. This value does not need to be unique.
	Email string `json:"email,omitempty"`

	// Unique ID to identify the business entity assigned to the account. Available when the `Multiple Business Entities` feature is enabled.
	OverrideBusinessEntityId string `json:"override_business_entity_id,omitempty"`

	// Used to determine the language and locale of emails sent on behalf of the merchant to the customer.
	PreferredLocale string `json:"preferred_locale,omitempty"`

	// The [IANA time zone name](https://docs.recurly.com/docs/email-time-zones-and-time-stamps#supported-api-iana-time-zone-names) used to determine the time zone of emails sent on behalf of the merchant to the customer.
	PreferredTimeZone string `json:"preferred_time_zone,omitempty"`

	// Additional email address that should receive account correspondence. These should be separated only by commas. These CC emails will receive all emails that the `email` field also receives.
	CcEmails string `json:"cc_emails,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	// The VAT number of the account (to avoid having the VAT applied). This is only used for manually collected invoices.
	VatNumber string `json:"vat_number,omitempty"`

	// The tax status of the account. `true` exempts tax on the account, `false` applies tax on the account.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// The tax exemption certificate number for the account. If the merchant has an integration for the Vertex tax provider, this optional value will be sent in any tax calculation requests for the account.
	ExemptionCertificate string `json:"exemption_certificate,omitempty"`

	// The external accounts belonging to this account
	ExternalAccounts []ExternalAccount `json:"external_accounts,omitempty"`

	// The UUID of the parent account associated with this account.
	ParentAccountId string `json:"parent_account_id,omitempty"`

	// An enumerable describing the billing behavior of the account, specifically whether the account is self-paying or will rely on the parent account to pay.
	BillTo string `json:"bill_to,omitempty"`

	// Unique ID to identify a dunning campaign. Used to specify if a non-default dunning campaign should be assigned to this account. For sites without multiple dunning campaigns enabled, the default dunning campaign will always be used.
	DunningCampaignId string `json:"dunning_campaign_id,omitempty"`

	// Unique ID to identify an invoice template. Available when the site is on a Pro or Elite plan. Used to specify if a non-default invoice template will be used to generate invoices for the account. For sites without multiple invoice templates enabled, the default template will always be used.
	InvoiceTemplateId string `json:"invoice_template_id,omitempty"`

	Address Address `json:"address,omitempty"`

	BillingInfo BillingInfo `json:"billing_info,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Account) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Account) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountList struct {
	ListMetadata
	Data            []Account `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountList allows you to paginate Account objects
type AccountList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Account
}

func NewAccountList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AccountList {
	return &AccountList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type AccountLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Account
	HasMore() bool
	Next() string
}

func (list *AccountList) HasMore() bool {
	return list.hasMore
}

func (list *AccountList) Next() string {
	return list.nextPagePath
}

func (list *AccountList) Data() []Account {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountList) FetchWithContext(ctx context.Context) error {
	resources := &accountList{}
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
func (list *AccountList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AccountList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &accountList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
