package recurly

import (
	"net/http"
	"time"
)

type Site struct {
	recurlyResponse *ResponseMetadata

	// Site ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	Subdomain string `json:"subdomain,omitempty"`

	// This value is used to configure RecurlyJS to submit tokenized billing information.
	PublicApiKey string `json:"public_api_key,omitempty"`

	// Mode
	Mode string `json:"mode,omitempty"`

	Address Address `json:"address,omitempty"`

	Settings Settings `json:"settings,omitempty"`

	// A list of features enabled for the site.
	Features []string `json:"features,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Site) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Site) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type siteList struct {
	ListMetadata
	Data            []Site `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *siteList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *siteList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SiteList allows you to paginate Site objects
type SiteList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Site
}

// Fetch fetches the next page of data into the `Data` property
func (list *SiteList) Fetch() error {
	resources := &siteList{}
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

type Address struct {
	recurlyResponse *ResponseMetadata

	// First name
	FirstName string `json:"first_name,omitempty"`

	// Last name
	LastName string `json:"last_name,omitempty"`

	// Phone number
	Phone string `json:"phone,omitempty"`

	// Street 1
	Street1 string `json:"street1,omitempty"`

	// Street 2
	Street2 string `json:"street2,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// State or province.
	Region string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO code.
	Country string `json:"country,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Address) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Address) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type addressList struct {
	ListMetadata
	Data            []Address `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *addressList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *addressList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AddressList allows you to paginate Address objects
type AddressList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Address
}

// Fetch fetches the next page of data into the `Data` property
func (list *AddressList) Fetch() error {
	resources := &addressList{}
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

type Settings struct {
	recurlyResponse *ResponseMetadata

	// - full:      Full Address (Street, City, State, Postal Code and Country)
	// - streetzip: Street and Postal Code only
	// - zip:       Postal Code only
	// - none:      No Address
	BillingAddressRequirement string `json:"billing_address_requirement,omitempty"`

	AcceptedCurrencies []string `json:"accepted_currencies,omitempty"`

	// The default 3-letter ISO 4217 currency code.
	DefaultCurrency string `json:"default_currency,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Settings) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Settings) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type settingsList struct {
	ListMetadata
	Data            []Settings `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *settingsList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *settingsList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SettingsList allows you to paginate Settings objects
type SettingsList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Settings
}

// Fetch fetches the next page of data into the `Data` property
func (list *SettingsList) Fetch() error {
	resources := &settingsList{}
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

	// Used to determine the language and locale of emails sent on behalf of the merchant to the customer.
	PreferredLocale string `json:"preferred_locale,omitempty"`

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

	// The UUID of the parent account associated with this account.
	ParentAccountId string `json:"parent_account_id,omitempty"`

	// An enumerable describing the billing behavior of the account, specifically whether the account is self-paying or will rely on the parent account to pay.
	BillTo string `json:"bill_to,omitempty"`

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
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Account
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountList) Fetch() error {
	resources := &accountList{}
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

type ShippingAddress struct {
	recurlyResponse *ResponseMetadata

	// Shipping Address ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Account ID
	AccountId string `json:"account_id,omitempty"`

	Nickname string `json:"nickname,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	Email string `json:"email,omitempty"`

	VatNumber string `json:"vat_number,omitempty"`

	Phone string `json:"phone,omitempty"`

	Street1 string `json:"street1,omitempty"`

	Street2 string `json:"street2,omitempty"`

	City string `json:"city,omitempty"`

	// State or province.
	Region string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO code.
	Country string `json:"country,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ShippingAddress) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ShippingAddress) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type shippingAddressList struct {
	ListMetadata
	Data            []ShippingAddress `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *shippingAddressList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *shippingAddressList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ShippingAddressList allows you to paginate ShippingAddress objects
type ShippingAddressList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []ShippingAddress
}

// Fetch fetches the next page of data into the `Data` property
func (list *ShippingAddressList) Fetch() error {
	resources := &shippingAddressList{}
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
	Data            []BillingInfo `json:"data"`
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
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []BillingInfo
}

// Fetch fetches the next page of data into the `Data` property
func (list *BillingInfoList) Fetch() error {
	resources := &billingInfoList{}
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

type PaymentMethod struct {
	recurlyResponse *ResponseMetadata

	Object string `json:"object,omitempty"`

	// Visa, MasterCard, American Express, Discover, JCB, etc.
	CardType string `json:"card_type,omitempty"`

	// Credit card number's first six digits.
	FirstSix string `json:"first_six,omitempty"`

	// Credit card number's last four digits. Will refer to bank account if payment method is ACH.
	LastFour string `json:"last_four,omitempty"`

	// The IBAN bank account's last two digits.
	LastTwo string `json:"last_two,omitempty"`

	// Expiration month.
	ExpMonth int `json:"exp_month,omitempty"`

	// Expiration year.
	ExpYear int `json:"exp_year,omitempty"`

	// A token used in place of a credit card in order to perform transactions.
	GatewayToken string `json:"gateway_token,omitempty"`

	// An identifier for a specific payment gateway.
	GatewayCode string `json:"gateway_code,omitempty"`

	// Billing Agreement identifier. Only present for Amazon or Paypal payment methods.
	BillingAgreementId string `json:"billing_agreement_id,omitempty"`

	// The bank account type. Only present for ACH payment methods.
	AccountType string `json:"account_type,omitempty"`

	// The bank account's routing number. Only present for ACH payment methods.
	RoutingNumber string `json:"routing_number,omitempty"`

	// The bank name of this routing number.
	RoutingNumberBank string `json:"routing_number_bank,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PaymentMethod) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PaymentMethod) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type paymentMethodList struct {
	ListMetadata
	Data            []PaymentMethod `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *paymentMethodList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *paymentMethodList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PaymentMethodList allows you to paginate PaymentMethod objects
type PaymentMethodList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []PaymentMethod
}

// Fetch fetches the next page of data into the `Data` property
func (list *PaymentMethodList) Fetch() error {
	resources := &paymentMethodList{}
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

type FraudInfo struct {
	recurlyResponse *ResponseMetadata

	// Kount score
	Score int `json:"score,omitempty"`

	// Kount decision
	Decision string `json:"decision,omitempty"`

	// Kount rules
	RiskRulesTriggered map[string]interface{} `json:"risk_rules_triggered,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *FraudInfo) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *FraudInfo) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type fraudInfoList struct {
	ListMetadata
	Data            []FraudInfo `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *fraudInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *fraudInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// FraudInfoList allows you to paginate FraudInfo objects
type FraudInfoList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []FraudInfo
}

// Fetch fetches the next page of data into the `Data` property
func (list *FraudInfoList) Fetch() error {
	resources := &fraudInfoList{}
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

type BillingInfoUpdatedBy struct {
	recurlyResponse *ResponseMetadata

	// Customer's IP address when updating their billing information.
	Ip string `json:"ip,omitempty"`

	// Country of IP address, if known by Recurly.
	Country string `json:"country,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *BillingInfoUpdatedBy) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *BillingInfoUpdatedBy) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type billingInfoUpdatedByList struct {
	ListMetadata
	Data            []BillingInfoUpdatedBy `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *billingInfoUpdatedByList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *billingInfoUpdatedByList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// BillingInfoUpdatedByList allows you to paginate BillingInfoUpdatedBy objects
type BillingInfoUpdatedByList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []BillingInfoUpdatedBy
}

// Fetch fetches the next page of data into the `Data` property
func (list *BillingInfoUpdatedByList) Fetch() error {
	resources := &billingInfoUpdatedByList{}
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

type CustomField struct {
	recurlyResponse *ResponseMetadata

	// Fields must be created in the UI before values can be assigned to them.
	Name string `json:"name,omitempty"`

	// Any values that resemble a credit card number or security code (CVV/CVC) will be rejected.
	Value string `json:"value,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CustomField) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CustomField) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type customFieldList struct {
	ListMetadata
	Data            []CustomField `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *customFieldList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *customFieldList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CustomFieldList allows you to paginate CustomField objects
type CustomFieldList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CustomField
}

// Fetch fetches the next page of data into the `Data` property
func (list *CustomFieldList) Fetch() error {
	resources := &customFieldList{}
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

type ErrorMayHaveTransaction struct {
	recurlyResponse *ResponseMetadata

	// Type
	Type string `json:"type,omitempty"`

	// Message
	Message string `json:"message,omitempty"`

	// Parameter specific errors
	Params []map[string]interface{} `json:"params,omitempty"`

	// This is only included on errors with `type=transaction`.
	TransactionError TransactionError `json:"transaction_error,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ErrorMayHaveTransaction) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ErrorMayHaveTransaction) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type errorMayHaveTransactionList struct {
	ListMetadata
	Data            []ErrorMayHaveTransaction `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *errorMayHaveTransactionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *errorMayHaveTransactionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ErrorMayHaveTransactionList allows you to paginate ErrorMayHaveTransaction objects
type ErrorMayHaveTransactionList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []ErrorMayHaveTransaction
}

// Fetch fetches the next page of data into the `Data` property
func (list *ErrorMayHaveTransactionList) Fetch() error {
	resources := &errorMayHaveTransactionList{}
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

type AccountAcquisition struct {
	recurlyResponse *ResponseMetadata

	// Account balance
	Cost AccountAcquisitionCost `json:"cost,omitempty"`

	// The channel through which the account was acquired.
	Channel string `json:"channel,omitempty"`

	// An arbitrary subchannel string representing a distinction/subcategory within a broader channel.
	Subchannel string `json:"subchannel,omitempty"`

	// An arbitrary identifier for the marketing campaign that led to the acquisition of this account.
	Campaign string `json:"campaign,omitempty"`

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// When the account acquisition data was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the account acquisition data was last changed.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountAcquisition) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountAcquisition) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountAcquisitionList struct {
	ListMetadata
	Data            []AccountAcquisition `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountAcquisitionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountAcquisitionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountAcquisitionList allows you to paginate AccountAcquisition objects
type AccountAcquisitionList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AccountAcquisition
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountAcquisitionList) Fetch() error {
	resources := &accountAcquisitionList{}
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

type AccountAcquisitionCost struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// The amount of the corresponding currency used to acquire the account.
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountAcquisitionCost) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountAcquisitionCost) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountAcquisitionCostList struct {
	ListMetadata
	Data            []AccountAcquisitionCost `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountAcquisitionCostList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountAcquisitionCostList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountAcquisitionCostList allows you to paginate AccountAcquisitionCost objects
type AccountAcquisitionCostList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AccountAcquisitionCost
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountAcquisitionCostList) Fetch() error {
	resources := &accountAcquisitionCostList{}
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

type AccountMini struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The unique identifier of the account.
	Code string `json:"code,omitempty"`

	// The email address used for communicating with this customer.
	Email string `json:"email,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Company string `json:"company,omitempty"`

	ParentAccountId string `json:"parent_account_id,omitempty"`

	BillTo string `json:"bill_to,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountMiniList struct {
	ListMetadata
	Data            []AccountMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountMiniList allows you to paginate AccountMini objects
type AccountMiniList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AccountMini
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountMiniList) Fetch() error {
	resources := &accountMiniList{}
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

type AccountBalance struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	PastDue bool `json:"past_due,omitempty"`

	Balances []AccountBalanceAmount `json:"balances,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountBalance) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountBalance) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountBalanceList struct {
	ListMetadata
	Data            []AccountBalance `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountBalanceList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountBalanceList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountBalanceList allows you to paginate AccountBalance objects
type AccountBalanceList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AccountBalance
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountBalanceList) Fetch() error {
	resources := &accountBalanceList{}
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

type AccountBalanceAmount struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total amount the account is past due.
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountBalanceAmount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountBalanceAmount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountBalanceAmountList struct {
	ListMetadata
	Data            []AccountBalanceAmount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountBalanceAmountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountBalanceAmountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountBalanceAmountList allows you to paginate AccountBalanceAmount objects
type AccountBalanceAmountList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AccountBalanceAmount
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountBalanceAmountList) Fetch() error {
	resources := &accountBalanceAmountList{}
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

type CouponRedemption struct {
	recurlyResponse *ResponseMetadata

	// Coupon Redemption ID
	Id string `json:"id,omitempty"`

	// Will always be `coupon`.
	Object string `json:"object,omitempty"`

	// The Account on which the coupon was applied.
	Account AccountMini `json:"account,omitempty"`

	Coupon Coupon `json:"coupon,omitempty"`

	// Coupon Redemption state
	State string `json:"state,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// The amount that was discounted upon the application of the coupon, formatted with the currency.
	Discounted float64 `json:"discounted,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date and time the redemption was removed from the account (un-redeemed).
	RemovedAt time.Time `json:"removed_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponRedemption) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponRedemption) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponRedemptionList struct {
	ListMetadata
	Data            []CouponRedemption `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponRedemptionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponRedemptionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponRedemptionList allows you to paginate CouponRedemption objects
type CouponRedemptionList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CouponRedemption
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponRedemptionList) Fetch() error {
	resources := &couponRedemptionList{}
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

type Coupon struct {
	recurlyResponse *ResponseMetadata

	// Coupon ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code string `json:"code,omitempty"`

	// The internal name for the coupon.
	Name string `json:"name,omitempty"`

	// Indicates if the coupon is redeemable, and if it is not, why.
	State string `json:"state,omitempty"`

	// A maximum number of redemptions for the coupon. The coupon will expire when it hits its maximum redemptions.
	MaxRedemptions int `json:"max_redemptions,omitempty"`

	// Redemptions per account is the number of times a specific account can redeem the coupon. Set redemptions per account to `1` if you want to keep customers from gaming the system and getting more than one discount from the coupon campaign.
	MaxRedemptionsPerAccount int `json:"max_redemptions_per_account,omitempty"`

	// When this number reaches `max_redemptions` the coupon will no longer be redeemable.
	UniqueCouponCodesCount int `json:"unique_coupon_codes_count,omitempty"`

	// On a bulk coupon, the template from which unique coupon codes are generated.
	UniqueCodeTemplate string `json:"unique_code_template,omitempty"`

	// - "single_use" coupons applies to the first invoice only.
	// - "temporal" coupons will apply to invoices for the duration determined by the `temporal_unit` and `temporal_amount` attributes.
	Duration string `json:"duration,omitempty"`

	// If `duration` is "temporal" than `temporal_amount` is an integer which is multiplied by `temporal_unit` to define the duration that the coupon will be applied to invoices for.
	TemporalAmount int `json:"temporal_amount,omitempty"`

	// If `duration` is "temporal" than `temporal_unit` is multiplied by `temporal_amount` to define the duration that the coupon will be applied to invoices for.
	TemporalUnit string `json:"temporal_unit,omitempty"`

	// Description of the unit of time the coupon is for. Used with `free_trial_amount` to determine the duration of time the coupon is for.
	FreeTrialUnit string `json:"free_trial_unit,omitempty"`

	// Sets the duration of time the `free_trial_unit` is for.
	FreeTrialAmount int `json:"free_trial_amount,omitempty"`

	// The coupon is valid for all plans if true. If false then `plans` and `plans_names` will list the applicable plans.
	AppliesToAllPlans bool `json:"applies_to_all_plans,omitempty"`

	// The coupon is valid for one-time, non-plan charges if true.
	AppliesToNonPlanCharges bool `json:"applies_to_non_plan_charges,omitempty"`

	// TODO
	PlansNames []string `json:"plans_names,omitempty"`

	// A list of plans for which this coupon applies. This will be `null` if `applies_to_all_plans=true`.
	Plans []PlanMini `json:"plans,omitempty"`

	// Whether the discount is for all eligible charges on the account, or only a specific subscription.
	RedemptionResource string `json:"redemption_resource,omitempty"`

	// Details of the discount a coupon applies. Will contain a `type`
	// property and one of the following properties: `percent`, `fixed`, `trial`.
	Discount CouponDiscount `json:"discount,omitempty"`

	// Whether the coupon is "single_code" or "bulk". Bulk coupons will require a `unique_code_template` and will generate unique codes through the `/generate` endpoint.
	CouponType string `json:"coupon_type,omitempty"`

	// This description will show up when a customer redeems a coupon on your Hosted Payment Pages, or if you choose to show the description on your own checkout page.
	HostedPageDescription string `json:"hosted_page_description,omitempty"`

	// Description of the coupon on the invoice.
	InvoiceDescription string `json:"invoice_description,omitempty"`

	// The date and time the coupon will expire and can no longer be redeemed. Time is always 11:59:59, the end-of-day Pacific time.
	RedeemBy time.Time `json:"redeem_by,omitempty"`

	// The date and time the unique coupon code was redeemed. This is only present for bulk coupons.
	RedeemedAt time.Time `json:"redeemed_at,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date and time the coupon was expired early or reached its `max_redemptions`.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Coupon) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Coupon) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponList struct {
	ListMetadata
	Data            []Coupon `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponList allows you to paginate Coupon objects
type CouponList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Coupon
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponList) Fetch() error {
	resources := &couponList{}
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

type PlanMini struct {
	recurlyResponse *ResponseMetadata

	// Plan ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique code to identify the plan. This is used in Hosted Payment Page URLs and in the invoice exports.
	Code string `json:"code,omitempty"`

	// This name describes your plan and will appear on the Hosted Payment Page and the subscriber's invoice.
	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planMiniList struct {
	ListMetadata
	Data            []PlanMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanMiniList allows you to paginate PlanMini objects
type PlanMiniList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []PlanMini
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanMiniList) Fetch() error {
	resources := &planMiniList{}
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

type CouponDiscount struct {
	recurlyResponse *ResponseMetadata

	Type string `json:"type,omitempty"`

	// This is only present when `type=percent`.
	Percent int `json:"percent,omitempty"`

	// This is only present when `type=fixed`.
	Currencies []CouponDiscountPricing `json:"currencies,omitempty"`

	// This is only present when `type=free_trial`.
	Trial CouponDiscountTrial `json:"trial,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponDiscountList struct {
	ListMetadata
	Data            []CouponDiscount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponDiscountList allows you to paginate CouponDiscount objects
type CouponDiscountList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CouponDiscount
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountList) Fetch() error {
	resources := &couponDiscountList{}
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

type CouponDiscountPricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Value of the fixed discount that this coupon applies.
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscountPricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscountPricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponDiscountPricingList struct {
	ListMetadata
	Data            []CouponDiscountPricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponDiscountPricingList allows you to paginate CouponDiscountPricing objects
type CouponDiscountPricingList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CouponDiscountPricing
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountPricingList) Fetch() error {
	resources := &couponDiscountPricingList{}
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

type CouponDiscountTrial struct {
	recurlyResponse *ResponseMetadata

	// Temporal unit of the free trial
	Unit string `json:"unit,omitempty"`

	// Trial length measured in the units specified by the sibling `unit` property
	Length int `json:"length,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponDiscountTrial) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponDiscountTrial) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponDiscountTrialList struct {
	ListMetadata
	Data            []CouponDiscountTrial `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponDiscountTrialList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponDiscountTrialList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponDiscountTrialList allows you to paginate CouponDiscountTrial objects
type CouponDiscountTrialList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CouponDiscountTrial
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponDiscountTrialList) Fetch() error {
	resources := &couponDiscountTrialList{}
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

type CreditPayment struct {
	recurlyResponse *ResponseMetadata

	// Credit Payment ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
	Uuid string `json:"uuid,omitempty"`

	// The action for which the credit was created.
	Action string `json:"action,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// Invoice mini details
	AppliedToInvoice InvoiceMini `json:"applied_to_invoice,omitempty"`

	// Invoice mini details
	OriginalInvoice InvoiceMini `json:"original_invoice,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total credit payment amount applied to the charge invoice.
	Amount float64 `json:"amount,omitempty"`

	// For credit payments with action `refund`, this is the credit payment that was refunded.
	OriginalCreditPaymentId string `json:"original_credit_payment_id,omitempty"`

	RefundTransaction Transaction `json:"refund_transaction,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Voided at
	VoidedAt time.Time `json:"voided_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CreditPayment) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CreditPayment) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type creditPaymentList struct {
	ListMetadata
	Data            []CreditPayment `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *creditPaymentList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *creditPaymentList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CreditPaymentList allows you to paginate CreditPayment objects
type CreditPaymentList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CreditPayment
}

// Fetch fetches the next page of data into the `Data` property
func (list *CreditPaymentList) Fetch() error {
	resources := &creditPaymentList{}
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

type InvoiceMini struct {
	recurlyResponse *ResponseMetadata

	// Invoice ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Invoice number
	Number string `json:"number,omitempty"`

	// Invoice type
	Type string `json:"type,omitempty"`

	// Invoice state
	State string `json:"state,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *InvoiceMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *InvoiceMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceMiniList struct {
	ListMetadata
	Data            []InvoiceMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceMiniList allows you to paginate InvoiceMini objects
type InvoiceMiniList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []InvoiceMini
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceMiniList) Fetch() error {
	resources := &invoiceMiniList{}
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

type Transaction struct {
	recurlyResponse *ResponseMetadata

	// Transaction ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
	Uuid string `json:"uuid,omitempty"`

	// If this transaction is a refund (`type=refund`), this will be the ID of the original transaction on the invoice being refunded.
	OriginalTransactionId string `json:"original_transaction_id,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// Invoice mini details
	Invoice InvoiceMini `json:"invoice,omitempty"`

	// Invoice mini details
	VoidedByInvoice InvoiceMini `json:"voided_by_invoice,omitempty"`

	// If the transaction is charging or refunding for one or more subscriptions, these are their IDs.
	SubscriptionIds []string `json:"subscription_ids,omitempty"`

	// - `authorization` – verifies billing information and places a hold on money in the customer's account.
	// - `capture` – captures funds held by an authorization and completes a purchase.
	// - `purchase` – combines the authorization and capture in one transaction.
	// - `refund` – returns all or a portion of the money collected in a previous transaction to the customer.
	// - `verify` – a $0 or $1 transaction used to verify billing information which is immediately voided.
	Type string `json:"type,omitempty"`

	// Describes how the transaction was triggered.
	Origin string `json:"origin,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total transaction amount sent to the payment gateway.
	Amount float64 `json:"amount,omitempty"`

	// The current transaction status. Note that the status may change, e.g. a `pending` transaction may become `declined` or `success` may later become `void`.
	Status string `json:"status,omitempty"`

	// Did this transaction complete successfully?
	Success bool `json:"success,omitempty"`

	// Indicates if part or all of this transaction was refunded.
	Refunded bool `json:"refunded,omitempty"`

	BillingAddress Address `json:"billing_address,omitempty"`

	// The method by which the payment was collected.
	CollectionMethod string `json:"collection_method,omitempty"`

	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`

	// IP address provided when the billing information was collected:
	// - When the customer enters billing information into the Recurly.js or Hosted Payment Pages, Recurly records the IP address.
	// - When the merchant enters billing information using the API, the merchant may provide an IP address.
	// - When the merchant enters billing information using the UI, no IP address is recorded.
	IpAddressV4 string `json:"ip_address_v4,omitempty"`

	// IP address's country
	IpAddressCountry string `json:"ip_address_country,omitempty"`

	// Status code
	StatusCode string `json:"status_code,omitempty"`

	// For declined (`success=false`) transactions, the message displayed to the merchant.
	StatusMessage string `json:"status_message,omitempty"`

	// For declined (`success=false`) transactions, the message displayed to the customer.
	CustomerMessage string `json:"customer_message,omitempty"`

	// Language code for the message
	CustomerMessageLocale string `json:"customer_message_locale,omitempty"`

	PaymentGateway TransactionPaymentGateway `json:"payment_gateway,omitempty"`

	// Transaction message from the payment gateway.
	GatewayMessage string `json:"gateway_message,omitempty"`

	// Transaction reference number from the payment gateway.
	GatewayReference string `json:"gateway_reference,omitempty"`

	// Transaction approval code from the payment gateway.
	GatewayApprovalCode string `json:"gateway_approval_code,omitempty"`

	// For declined transactions (`success=false`), this field lists the gateway error code.
	GatewayResponseCode string `json:"gateway_response_code,omitempty"`

	// Time, in seconds, for gateway to process the transaction.
	GatewayResponseTime float64 `json:"gateway_response_time,omitempty"`

	// The values in this field will vary from gateway to gateway.
	GatewayResponseValues map[string]interface{} `json:"gateway_response_values,omitempty"`

	// When processed, result from checking the CVV/CVC value on the transaction.
	CvvCheck string `json:"cvv_check,omitempty"`

	// When processed, result from checking the overall AVS on the transaction.
	AvsCheck string `json:"avs_check,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Voided at
	VoidedAt time.Time `json:"voided_at,omitempty"`

	// Collected at, or if not collected yet, the time the transaction was created.
	CollectedAt time.Time `json:"collected_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Transaction) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Transaction) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type transactionList struct {
	ListMetadata
	Data            []Transaction `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *transactionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *transactionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TransactionList allows you to paginate Transaction objects
type TransactionList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Transaction
}

// Fetch fetches the next page of data into the `Data` property
func (list *TransactionList) Fetch() error {
	resources := &transactionList{}
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

type TransactionPaymentGateway struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TransactionPaymentGateway) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TransactionPaymentGateway) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type transactionPaymentGatewayList struct {
	ListMetadata
	Data            []TransactionPaymentGateway `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *transactionPaymentGatewayList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *transactionPaymentGatewayList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TransactionPaymentGatewayList allows you to paginate TransactionPaymentGateway objects
type TransactionPaymentGatewayList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []TransactionPaymentGateway
}

// Fetch fetches the next page of data into the `Data` property
func (list *TransactionPaymentGatewayList) Fetch() error {
	resources := &transactionPaymentGatewayList{}
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

type Invoice struct {
	recurlyResponse *ResponseMetadata

	// Invoice ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Invoices are either charge, credit, or legacy invoices.
	Type string `json:"type,omitempty"`

	// The event that created the invoice.
	Origin string `json:"origin,omitempty"`

	// Invoice state
	State string `json:"state,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// If the invoice is charging or refunding for one or more subscriptions, these are their IDs.
	SubscriptionIds []string `json:"subscription_ids,omitempty"`

	// On refund invoices, this value will exist and show the invoice ID of the purchase invoice the refund was created from.
	PreviousInvoiceId string `json:"previous_invoice_id,omitempty"`

	// If VAT taxation and the Country Invoice Sequencing feature are enabled, invoices will have country-specific invoice numbers for invoices billed to EU countries (ex: FR1001). Non-EU invoices will continue to use the site-level invoice number sequence.
	Number string `json:"number,omitempty"`

	// An automatic invoice means a corresponding transaction is run using the account's billing information at the same time the invoice is created. Manual invoices are created without a corresponding transaction. The merchant must enter a manual payment transaction or have the customer pay the invoice with an automatic method, like credit card, PayPal, Amazon, or ACH bank payment.
	CollectionMethod string `json:"collection_method,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	Address InvoiceAddress `json:"address,omitempty"`

	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total discounts applied to this invoice.
	Discount float64 `json:"discount,omitempty"`

	// The summation of charges, discounts, and credits, before tax.
	Subtotal float64 `json:"subtotal,omitempty"`

	// The total tax on this invoice.
	Tax float64 `json:"tax,omitempty"`

	// The final total on this invoice. The summation of invoice charges, discounts, credits, and tax.
	Total float64 `json:"total,omitempty"`

	// The refundable amount on a charge invoice. It will be null for all other invoices.
	RefundableAmount float64 `json:"refundable_amount,omitempty"`

	// The total amount of successful payments transaction on this invoice.
	Paid float64 `json:"paid,omitempty"`

	// The outstanding balance remaining on this invoice.
	Balance float64 `json:"balance,omitempty"`

	// Tax info
	TaxInfo TaxInfo `json:"tax_info,omitempty"`

	// VAT registration number for the customer on this invoice. This will come from the VAT Number field in the Billing Info or the Account Info depending on your tax settings and the invoice collection method.
	VatNumber string `json:"vat_number,omitempty"`

	// VAT Reverse Charge Notes only appear if you have EU VAT enabled or are using your own Avalara AvaTax account and the customer is in the EU, has a VAT number, and is in a different country than your own. This will default to the VAT Reverse Charge Notes text specified on the Tax Settings page in your Recurly admin, unless custom notes were created with the original subscription.
	VatReverseChargeNotes string `json:"vat_reverse_charge_notes,omitempty"`

	// This will default to the Terms and Conditions text specified on the Invoice Settings page in your Recurly admin. Specify custom notes to add or override Terms and Conditions.
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// This will default to the Customer Notes text specified on the Invoice Settings. Specify custom notes to add or override Customer Notes.
	CustomerNotes string `json:"customer_notes,omitempty"`

	LineItems LineItemList `json:"line_items,omitempty"`

	// Transactions
	Transactions []Transaction `json:"transactions,omitempty"`

	// Credit payments
	CreditPayments []CreditPayment `json:"credit_payments,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Date invoice is due. This is the date the net terms are reached.
	DueAt time.Time `json:"due_at,omitempty"`

	// Date invoice was marked paid or failed.
	ClosedAt time.Time `json:"closed_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Invoice) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Invoice) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceList struct {
	ListMetadata
	Data            []Invoice `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceList allows you to paginate Invoice objects
type InvoiceList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Invoice
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceList) Fetch() error {
	resources := &invoiceList{}
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

type InvoiceAddress struct {
	recurlyResponse *ResponseMetadata

	// Name on account
	NameOnAccount string `json:"name_on_account,omitempty"`

	// Company
	Company string `json:"company,omitempty"`

	// First name
	FirstName string `json:"first_name,omitempty"`

	// Last name
	LastName string `json:"last_name,omitempty"`

	// Phone number
	Phone string `json:"phone,omitempty"`

	// Street 1
	Street1 string `json:"street1,omitempty"`

	// Street 2
	Street2 string `json:"street2,omitempty"`

	// City
	City string `json:"city,omitempty"`

	// State or province.
	Region string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO code.
	Country string `json:"country,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *InvoiceAddress) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *InvoiceAddress) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceAddressList struct {
	ListMetadata
	Data            []InvoiceAddress `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceAddressList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceAddressList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceAddressList allows you to paginate InvoiceAddress objects
type InvoiceAddressList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []InvoiceAddress
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceAddressList) Fetch() error {
	resources := &invoiceAddressList{}
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

type TaxInfo struct {
	recurlyResponse *ResponseMetadata

	// Provides the tax type as "vat" for EU VAT, "usst" for U.S. Sales Tax, or the 2 letter country code for country level tax types like Canada, Australia, New Zealand, Israel, and all non-EU European countries.
	Type string `json:"type,omitempty"`

	// Provides the tax region applied on an invoice. For U.S. Sales Tax, this will be the 2 letter state code. For EU VAT this will be the 2 letter country code. For all country level tax types, this will display the regional tax, like VAT, GST, or PST.
	Region string `json:"region,omitempty"`

	// Rate
	Rate float64 `json:"rate,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TaxInfo) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TaxInfo) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type taxInfoList struct {
	ListMetadata
	Data            []TaxInfo `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *taxInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *taxInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TaxInfoList allows you to paginate TaxInfo objects
type TaxInfoList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []TaxInfo
}

// Fetch fetches the next page of data into the `Data` property
func (list *TaxInfoList) Fetch() error {
	resources := &taxInfoList{}
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

type LineItem struct {
	recurlyResponse *ResponseMetadata

	// Line item ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
	Uuid string `json:"uuid,omitempty"`

	// Charges are positive line items that debit the account. Credits are negative line items that credit the account.
	Type string `json:"type,omitempty"`

	// Unique code to identify an item. Available when the Credit Invoices and Subscription Billing Terms features are enabled.
	ItemCode string `json:"item_code,omitempty"`

	// System-generated unique identifier for an item. Available when the Credit Invoices and Subscription Billing Terms features are enabled.
	ItemId string `json:"item_id,omitempty"`

	// Optional Stock Keeping Unit assigned to an item. Available when the Credit Invoices and Subscription Billing Terms features are enabled.
	ExternalSku string `json:"external_sku,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Pending line items are charges or credits on an account that have not been applied to an invoice yet. Invoiced line items will always have an `invoice_id` value.
	State string `json:"state,omitempty"`

	// Category to describe the role of a line item on a legacy invoice:
	// - "charges" refers to charges being billed for on this invoice.
	// - "credits" refers to refund or proration credits. This portion of the invoice can be considered a credit memo.
	// - "applied_credits" refers to previous credits applied to this invoice. See their original_line_item_id to determine where the credit first originated.
	// - "carryforwards" can be ignored. They exist to consume any remaining credit balance. A new credit with the same amount will be created and placed back on the account.
	LegacyCategory string `json:"legacy_category,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// If the line item is a charge or credit for a subscription, this is its ID.
	SubscriptionId string `json:"subscription_id,omitempty"`

	// If the line item is a charge or credit for a plan or add-on, this is the plan's ID.
	PlanId string `json:"plan_id,omitempty"`

	// If the line item is a charge or credit for a plan or add-on, this is the plan's code.
	PlanCode string `json:"plan_code,omitempty"`

	// If the line item is a charge or credit for an add-on this is its ID.
	AddOnId string `json:"add_on_id,omitempty"`

	// If the line item is a charge or credit for an add-on, this is its code.
	AddOnCode string `json:"add_on_code,omitempty"`

	// Once the line item has been invoiced this will be the invoice's ID.
	InvoiceId string `json:"invoice_id,omitempty"`

	// Once the line item has been invoiced this will be the invoice's number. If VAT taxation and the Country Invoice Sequencing feature are enabled, invoices will have country-specific invoice numbers for invoices billed to EU countries (ex: FR1001). Non-EU invoices will continue to use the site-level invoice number sequence.
	InvoiceNumber string `json:"invoice_number,omitempty"`

	// Will only have a value if the line item is a credit created from a previous credit, or if the credit was created from a charge refund.
	PreviousLineItemId string `json:"previous_line_item_id,omitempty"`

	// The invoice where the credit originated. Will only have a value if the line item is a credit created from a previous credit, or if the credit was created from a charge refund.
	OriginalLineItemInvoiceId string `json:"original_line_item_invoice_id,omitempty"`

	// A credit created from an original charge will have the value of the charge's origin.
	Origin string `json:"origin,omitempty"`

	// Internal accounting code to help you reconcile your revenue to the correct ledger. Line items created as part of a subscription invoice will use the plan or add-on's accounting code, otherwise the value will only be present if you define an accounting code when creating the line item.
	AccountingCode string `json:"accounting_code,omitempty"`

	// For plan-related line items this will be the plan's code, for add-on related line items it will be the add-on's code. For item-related line itmes it will be the item's `external_sku`.
	ProductCode string `json:"product_code,omitempty"`

	// The reason the credit was given when line item is `type=credit`.
	CreditReasonCode string `json:"credit_reason_code,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// `(quantity * unit_amount) - (discount + tax)`
	Amount float64 `json:"amount,omitempty"`

	// Description that appears on the invoice. For subscription related items this will be filled in automatically.
	Description string `json:"description,omitempty"`

	// This number will be multiplied by the unit amount to compute the subtotal before any discounts or taxes.
	Quantity int `json:"quantity,omitempty"`

	// Positive amount for a charge, negative amount for a credit.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// `quantity * unit_amount`
	Subtotal float64 `json:"subtotal,omitempty"`

	// The discount applied to the line item.
	Discount float64 `json:"discount,omitempty"`

	// The tax amount for the line item.
	Tax float64 `json:"tax,omitempty"`

	// `true` if the line item is taxable, `false` if it is not.
	Taxable bool `json:"taxable,omitempty"`

	// `true` exempts tax on charges, `false` applies tax on charges. If not defined, then defaults to the Plan and Site settings. This attribute does not work for credits (negative line items). Credits are always applied post-tax. Pre-tax discounts should use the Coupons feature.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// Tax info
	TaxInfo TaxInfo `json:"tax_info,omitempty"`

	// When a line item has been prorated, this is the rate of the proration. Proration rates were made available for line items created after March 30, 2017. For line items created prior to that date, the proration rate will be `null`, even if the line item was prorated.
	ProrationRate float64 `json:"proration_rate,omitempty"`

	// Refund?
	Refund bool `json:"refund,omitempty"`

	// For refund charges, the quantity being refunded. For non-refund charges, the total quantity refunded (possibly over multiple refunds).
	RefundedQuantity int `json:"refunded_quantity,omitempty"`

	// The amount of credit from this line item that was applied to the invoice.
	CreditApplied float64 `json:"credit_applied,omitempty"`

	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`

	// If an end date is present, this is value indicates the beginning of a billing time range. If no end date is present it indicates billing for a specific date.
	StartDate time.Time `json:"start_date,omitempty"`

	// If this date is provided, it indicates the end of a time range.
	EndDate time.Time `json:"end_date,omitempty"`

	// When the line item was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// When the line item was last changed.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *LineItem) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *LineItem) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type lineItemList struct {
	ListMetadata
	Data            []LineItem `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *lineItemList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *lineItemList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// LineItemList allows you to paginate LineItem objects
type LineItemList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []LineItem
}

// Fetch fetches the next page of data into the `Data` property
func (list *LineItemList) Fetch() error {
	resources := &lineItemList{}
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

type InvoiceCollection struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	ChargeInvoice Invoice `json:"charge_invoice,omitempty"`

	// Credit invoices
	CreditInvoices []Invoice `json:"credit_invoices,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *InvoiceCollection) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *InvoiceCollection) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type invoiceCollectionList struct {
	ListMetadata
	Data            []InvoiceCollection `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *invoiceCollectionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *invoiceCollectionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// InvoiceCollectionList allows you to paginate InvoiceCollection objects
type InvoiceCollectionList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []InvoiceCollection
}

// Fetch fetches the next page of data into the `Data` property
func (list *InvoiceCollectionList) Fetch() error {
	resources := &invoiceCollectionList{}
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

type AccountNote struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	AccountId string `json:"account_id,omitempty"`

	User User `json:"user,omitempty"`

	Message string `json:"message,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountNote) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountNote) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountNoteList struct {
	ListMetadata
	Data            []AccountNote `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountNoteList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountNoteList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountNoteList allows you to paginate AccountNote objects
type AccountNoteList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AccountNote
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountNoteList) Fetch() error {
	resources := &accountNoteList{}
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

type User struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	TimeZone string `json:"time_zone,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *User) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *User) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type userList struct {
	ListMetadata
	Data            []User `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *userList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *userList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// UserList allows you to paginate User objects
type UserList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []User
}

// Fetch fetches the next page of data into the `Data` property
func (list *UserList) Fetch() error {
	resources := &userList{}
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

type Subscription struct {
	recurlyResponse *ResponseMetadata

	// Subscription ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The UUID is useful for matching data with the CSV exports and building URLs into Recurly's UI.
	Uuid string `json:"uuid,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	// Just the important parts.
	Plan PlanMini `json:"plan,omitempty"`

	// State
	State string `json:"state,omitempty"`

	// Subscription shipping details
	Shipping SubscriptionShipping `json:"shipping,omitempty"`

	// Coupon redemptions
	CouponRedemptions []CouponRedemptionMini `json:"coupon_redemptions,omitempty"`

	// Subscription Change
	PendingChange SubscriptionChange `json:"pending_change,omitempty"`

	// Current billing period started at
	CurrentPeriodStartedAt time.Time `json:"current_period_started_at,omitempty"`

	// Current billing period ends at
	CurrentPeriodEndsAt time.Time `json:"current_period_ends_at,omitempty"`

	// The start date of the term when the first billing period starts. The subscription term is the length of time that a customer will be committed to a subscription. A term can span multiple billing periods.
	CurrentTermStartedAt time.Time `json:"current_term_started_at,omitempty"`

	// When the term ends. This is calculated by a plan's interval and `total_billing_cycles` in a term. Subscription changes with a `timeframe=renewal` will be applied on this date.
	CurrentTermEndsAt time.Time `json:"current_term_ends_at,omitempty"`

	// Trial period started at
	TrialStartedAt time.Time `json:"trial_started_at,omitempty"`

	// Trial period ends at
	TrialEndsAt time.Time `json:"trial_ends_at,omitempty"`

	// The remaining billing cycles in the current term.
	RemainingBillingCycles int `json:"remaining_billing_cycles,omitempty"`

	// The number of cycles/billing periods in a term. When `remaining_billing_cycles=0`, if `auto_renew=true` the subscription will renew and a new term will begin, otherwise the subscription will expire.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// If `auto_renew=true`, when a term completes, `total_billing_cycles` takes this value as the length of subsequent terms. Defaults to the plan's `total_billing_cycles`.
	RenewalBillingCycles int `json:"renewal_billing_cycles,omitempty"`

	// Whether the subscription renews at the end of its term.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Null unless subscription is paused or will pause at the end of the current billing period.
	PausedAt time.Time `json:"paused_at,omitempty"`

	// Null unless subscription is paused or will pause at the end of the current billing period.
	RemainingPauseCycles int `json:"remaining_pause_cycles,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Subscription unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Subscription quantity
	Quantity int `json:"quantity,omitempty"`

	// Add-ons
	AddOns []SubscriptionAddOn `json:"add_ons,omitempty"`

	// Total price of add-ons
	AddOnsTotal float64 `json:"add_ons_total,omitempty"`

	// Estimated total, before tax.
	Subtotal float64 `json:"subtotal,omitempty"`

	// Collection method
	CollectionMethod string `json:"collection_method,omitempty"`

	// For manual invoicing, this identifies the PO number associated with the subscription.
	PoNumber string `json:"po_number,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
	NetTerms int `json:"net_terms,omitempty"`

	// Terms and conditions
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`

	// Customer notes
	CustomerNotes string `json:"customer_notes,omitempty"`

	// Expiration reason
	ExpirationReason string `json:"expiration_reason,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields []CustomField `json:"custom_fields,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Activated at
	ActivatedAt time.Time `json:"activated_at,omitempty"`

	// Canceled at
	CanceledAt time.Time `json:"canceled_at,omitempty"`

	// Expires at
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	// Recurring subscriptions paid with ACH will have this attribute set. This timestamp is used for alerting customers to reauthorize in 3 years in accordance with NACHA rules. If a subscription becomes inactive or the billing info is no longer a bank account, this timestamp is cleared.
	BankAccountAuthorizedAt time.Time `json:"bank_account_authorized_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Subscription) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Subscription) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionList struct {
	ListMetadata
	Data            []Subscription `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionList allows you to paginate Subscription objects
type SubscriptionList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Subscription
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionList) Fetch() error {
	resources := &subscriptionList{}
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

type SubscriptionShipping struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	Address ShippingAddress `json:"address,omitempty"`

	Method ShippingMethodMini `json:"method,omitempty"`

	// Subscription's shipping cost
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionShipping) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionShipping) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionShippingList struct {
	ListMetadata
	Data            []SubscriptionShipping `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionShippingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionShippingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionShippingList allows you to paginate SubscriptionShipping objects
type SubscriptionShippingList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []SubscriptionShipping
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionShippingList) Fetch() error {
	resources := &subscriptionShippingList{}
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

type ShippingMethodMini struct {
	recurlyResponse *ResponseMetadata

	// Shipping Method ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The internal name used identify the shipping method.
	Code string `json:"code,omitempty"`

	// The name of the shipping method displayed to customers.
	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ShippingMethodMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ShippingMethodMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type shippingMethodMiniList struct {
	ListMetadata
	Data            []ShippingMethodMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *shippingMethodMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *shippingMethodMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ShippingMethodMiniList allows you to paginate ShippingMethodMini objects
type ShippingMethodMiniList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []ShippingMethodMini
}

// Fetch fetches the next page of data into the `Data` property
func (list *ShippingMethodMiniList) Fetch() error {
	resources := &shippingMethodMiniList{}
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

type CouponRedemptionMini struct {
	recurlyResponse *ResponseMetadata

	// Coupon Redemption ID
	Id string `json:"id,omitempty"`

	// Will always be `coupon`.
	Object string `json:"object,omitempty"`

	Coupon CouponMini `json:"coupon,omitempty"`

	// Invoice state
	State string `json:"state,omitempty"`

	// The amount that was discounted upon the application of the coupon, formatted with the currency.
	Discounted float64 `json:"discounted,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponRedemptionMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponRedemptionMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponRedemptionMiniList struct {
	ListMetadata
	Data            []CouponRedemptionMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponRedemptionMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponRedemptionMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponRedemptionMiniList allows you to paginate CouponRedemptionMini objects
type CouponRedemptionMiniList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CouponRedemptionMini
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponRedemptionMiniList) Fetch() error {
	resources := &couponRedemptionMiniList{}
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

type CouponMini struct {
	recurlyResponse *ResponseMetadata

	// Coupon ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code string `json:"code,omitempty"`

	// The internal name for the coupon.
	Name string `json:"name,omitempty"`

	// Indicates if the coupon is redeemable, and if it is not, why.
	State string `json:"state,omitempty"`

	// Details of the discount a coupon applies. Will contain a `type`
	// property and one of the following properties: `percent`, `fixed`, `trial`.
	Discount CouponDiscount `json:"discount,omitempty"`

	// Whether the coupon is "single_code" or "bulk". Bulk coupons will require a `unique_code_template` and will generate unique codes through the `/generate` endpoint.
	CouponType string `json:"coupon_type,omitempty"`

	// The date and time the coupon was expired early or reached its `max_redemptions`.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CouponMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CouponMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type couponMiniList struct {
	ListMetadata
	Data            []CouponMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *couponMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *couponMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CouponMiniList allows you to paginate CouponMini objects
type CouponMiniList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CouponMini
}

// Fetch fetches the next page of data into the `Data` property
func (list *CouponMiniList) Fetch() error {
	resources := &couponMiniList{}
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

type SubscriptionChange struct {
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

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionChange) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionChange) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionChangeList struct {
	ListMetadata
	Data            []SubscriptionChange `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionChangeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionChangeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionChangeList allows you to paginate SubscriptionChange objects
type SubscriptionChangeList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []SubscriptionChange
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionChangeList) Fetch() error {
	resources := &subscriptionChangeList{}
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

type SubscriptionAddOn struct {
	recurlyResponse *ResponseMetadata

	// Subscription Add-on ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Subscription ID
	SubscriptionId string `json:"subscription_id,omitempty"`

	// Just the important parts.
	AddOn AddOnMini `json:"add_on,omitempty"`

	// Add-on quantity
	Quantity int `json:"quantity,omitempty"`

	// This is priced in the subscription's currency.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Expired at
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOn) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOn) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionAddOnList struct {
	ListMetadata
	Data            []SubscriptionAddOn `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionAddOnList allows you to paginate SubscriptionAddOn objects
type SubscriptionAddOnList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []SubscriptionAddOn
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionAddOnList) Fetch() error {
	resources := &subscriptionAddOnList{}
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
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AddOnMini
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

type UniqueCouponCode struct {
	recurlyResponse *ResponseMetadata

	// Unique Coupon Code ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The code the customer enters to redeem the coupon.
	Code string `json:"code,omitempty"`

	// Indicates if the unique coupon code is redeemable or why not.
	State string `json:"state,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// The date and time the unique coupon code was redeemed.
	RedeemedAt time.Time `json:"redeemed_at,omitempty"`

	// The date and time the coupon was expired early or reached its `max_redemptions`.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *UniqueCouponCode) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *UniqueCouponCode) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type uniqueCouponCodeList struct {
	ListMetadata
	Data            []UniqueCouponCode `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *uniqueCouponCodeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// UniqueCouponCodeList allows you to paginate UniqueCouponCode objects
type UniqueCouponCodeList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []UniqueCouponCode
}

// Fetch fetches the next page of data into the `Data` property
func (list *UniqueCouponCodeList) Fetch() error {
	resources := &uniqueCouponCodeList{}
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

type CustomFieldDefinition struct {
	recurlyResponse *ResponseMetadata

	// Custom field definition ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Related Recurly object type
	RelatedType string `json:"related_type,omitempty"`

	// Used by the API to identify the field or reading and writing. The name can only be used once per Recurly object type.
	Name string `json:"name,omitempty"`

	// The access control applied inside Recurly's admin UI:
	// - `api_only` - No one will be able to view or edit this field's data via the admin UI.
	// - `read_only` - Users with the Customers role will be able to view this field's data via the admin UI, but
	//   editing will only be available via the API.
	// - `write` - Users with the Customers role will be able to view and edit this field's data via the admin UI.
	UserAccess string `json:"user_access,omitempty"`

	// Used to label the field when viewing and editing the field in Recurly's admin UI.
	DisplayName string `json:"display_name,omitempty"`

	// Displayed as a tooltip when editing the field in the Recurly admin UI.
	Tooltip string `json:"tooltip,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Definitions are initially soft deleted, and once all the values are removed from the accouts or subscriptions, will be hard deleted an no longer visible.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CustomFieldDefinition) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CustomFieldDefinition) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type customFieldDefinitionList struct {
	ListMetadata
	Data            []CustomFieldDefinition `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *customFieldDefinitionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *customFieldDefinitionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CustomFieldDefinitionList allows you to paginate CustomFieldDefinition objects
type CustomFieldDefinitionList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []CustomFieldDefinition
}

// Fetch fetches the next page of data into the `Data` property
func (list *CustomFieldDefinitionList) Fetch() error {
	resources := &customFieldDefinitionList{}
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
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Item
}

// Fetch fetches the next page of data into the `Data` property
func (list *ItemList) Fetch() error {
	resources := &itemList{}
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

type Pricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Pricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Pricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type pricingList struct {
	ListMetadata
	Data            []Pricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *pricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *pricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PricingList allows you to paginate Pricing objects
type PricingList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Pricing
}

// Fetch fetches the next page of data into the `Data` property
func (list *PricingList) Fetch() error {
	resources := &pricingList{}
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

type BinaryFile struct {
	recurlyResponse *ResponseMetadata

	Data string `json:"data,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *BinaryFile) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *BinaryFile) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type binaryFileList struct {
	ListMetadata
	Data            []BinaryFile `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *binaryFileList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *binaryFileList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// BinaryFileList allows you to paginate BinaryFile objects
type BinaryFileList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []BinaryFile
}

// Fetch fetches the next page of data into the `Data` property
func (list *BinaryFileList) Fetch() error {
	resources := &binaryFileList{}
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

type Plan struct {
	recurlyResponse *ResponseMetadata

	// Plan ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Unique code to identify the plan. This is used in Hosted Payment Page URLs and in the invoice exports.
	Code string `json:"code,omitempty"`

	// The current state of the plan.
	State string `json:"state,omitempty"`

	// This name describes your plan and will appear on the Hosted Payment Page and the subscriber's invoice.
	Name string `json:"name,omitempty"`

	// Optional description, not displayed.
	Description string `json:"description,omitempty"`

	// Unit for the plan's billing interval.
	IntervalUnit string `json:"interval_unit,omitempty"`

	// Length of the plan's billing interval in `interval_unit`.
	IntervalLength int `json:"interval_length,omitempty"`

	// Units for the plan's trial period.
	TrialUnit string `json:"trial_unit,omitempty"`

	// Length of plan's trial period in `trial_units`. `0` means `no trial`.
	TrialLength int `json:"trial_length,omitempty"`

	// Automatically terminate subscriptions after a defined number of billing cycles. Number of billing cycles before the plan automatically stops renewing, defaults to `null` for continuous, automatic renewal.
	TotalBillingCycles int `json:"total_billing_cycles,omitempty"`

	// Subscriptions will automatically inherit this value once they are active. If `auto_renew` is `true`, then a subscription will automatically renew its term at renewal. If `auto_renew` is `false`, then a subscription will expire at the end of its term. `auto_renew` can be overridden on the subscription record itself.
	AutoRenew bool `json:"auto_renew,omitempty"`

	// Accounting code for invoice line items for the plan. If no value is provided, it defaults to plan's code.
	AccountingCode string `json:"accounting_code,omitempty"`

	// Revenue schedule type
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

	// Setup fee revenue schedule type
	SetupFeeRevenueScheduleType string `json:"setup_fee_revenue_schedule_type,omitempty"`

	// Accounting code for invoice line items for the plan's setup fee. If no value is provided, it defaults to plan's accounting code.
	SetupFeeAccountingCode string `json:"setup_fee_accounting_code,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s EU VAT tax feature. The tax code values are specific to each tax system. If you are using Recurly’s EU VAT feature you can use `unknown`, `physical`, or `digital`.
	TaxCode string `json:"tax_code,omitempty"`

	// `true` exempts tax on the plan, `false` applies tax on the plan.
	TaxExempt bool `json:"tax_exempt,omitempty"`

	// Pricing
	Currencies []PlanPricing `json:"currencies,omitempty"`

	// Hosted pages settings
	HostedPages PlanHostedPages `json:"hosted_pages,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Plan) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Plan) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planList struct {
	ListMetadata
	Data            []Plan `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanList allows you to paginate Plan objects
type PlanList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Plan
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanList) Fetch() error {
	resources := &planList{}
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

type PlanPricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Amount of one-time setup fee automatically charged at the beginning of a subscription billing cycle. For subscription plans with a trial, the setup fee will be charged at the time of signup. Setup fees do not increase with the quantity of a subscription plan.
	SetupFee float64 `json:"setup_fee,omitempty"`

	// Unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanPricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanPricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planPricingList struct {
	ListMetadata
	Data            []PlanPricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanPricingList allows you to paginate PlanPricing objects
type PlanPricingList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []PlanPricing
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanPricingList) Fetch() error {
	resources := &planPricingList{}
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

type PlanHostedPages struct {
	recurlyResponse *ResponseMetadata

	// URL to redirect to after signup on the hosted payment pages.
	SuccessUrl string `json:"success_url,omitempty"`

	// URL to redirect to on canceled signup on the hosted payment pages.
	CancelUrl string `json:"cancel_url,omitempty"`

	// If `true`, the customer will be sent directly to your `success_url` after a successful signup, bypassing Recurly's hosted confirmation page.
	BypassConfirmation bool `json:"bypass_confirmation,omitempty"`

	// Determines if the quantity field is displayed on the hosted pages for the plan.
	DisplayQuantity bool `json:"display_quantity,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanHostedPages) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanHostedPages) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planHostedPagesList struct {
	ListMetadata
	Data            []PlanHostedPages `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planHostedPagesList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planHostedPagesList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanHostedPagesList allows you to paginate PlanHostedPages objects
type PlanHostedPagesList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []PlanHostedPages
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanHostedPagesList) Fetch() error {
	resources := &planHostedPagesList{}
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

	// Accounting code for invoice line items for this add-on. If no value is provided, it defaults to add-on's code.
	AccountingCode string `json:"accounting_code,omitempty"`

	// When this add-on is invoiced, the line item will use this revenue schedule. If `item_code`/`item_id` is part of the request then `revenue_schedule_type` must be absent in the request as the value will be set from the item.
	RevenueScheduleType string `json:"revenue_schedule_type,omitempty"`

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
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AddOn
}

// Fetch fetches the next page of data into the `Data` property
func (list *AddOnList) Fetch() error {
	resources := &addOnList{}
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

type AddOnPricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AddOnPricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AddOnPricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type addOnPricingList struct {
	ListMetadata
	Data            []AddOnPricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *addOnPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *addOnPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AddOnPricingList allows you to paginate AddOnPricing objects
type AddOnPricingList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []AddOnPricing
}

// Fetch fetches the next page of data into the `Data` property
func (list *AddOnPricingList) Fetch() error {
	resources := &addOnPricingList{}
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

type ItemMini struct {
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
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ItemMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ItemMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type itemMiniList struct {
	ListMetadata
	Data            []ItemMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *itemMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *itemMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ItemMiniList allows you to paginate ItemMini objects
type ItemMiniList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []ItemMini
}

// Fetch fetches the next page of data into the `Data` property
func (list *ItemMiniList) Fetch() error {
	resources := &itemMiniList{}
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

type ShippingMethod struct {
	recurlyResponse *ResponseMetadata

	// Shipping Method ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The internal name used identify the shipping method.
	Code string `json:"code,omitempty"`

	// The name of the shipping method displayed to customers.
	Name string `json:"name,omitempty"`

	// Used by Avalara, Vertex, and Recurly’s built-in tax feature. The tax
	// code values are specific to each tax system. If you are using Recurly’s
	// built-in taxes the values are:
	// - `FR` – Common Carrier FOB Destination
	// - `FR022000` – Common Carrier FOB Origin
	// - `FR020400` – Non Common Carrier FOB Destination
	// - `FR020500` – Non Common Carrier FOB Origin
	// - `FR010100` – Delivery by Company Vehicle Before Passage of Title
	// - `FR010200` – Delivery by Company Vehicle After Passage of Title
	// - `NT` – Non-Taxable
	TaxCode string `json:"tax_code,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Deleted at
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ShippingMethod) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ShippingMethod) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type shippingMethodList struct {
	ListMetadata
	Data            []ShippingMethod `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *shippingMethodList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *shippingMethodList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ShippingMethodList allows you to paginate ShippingMethod objects
type ShippingMethodList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []ShippingMethod
}

// Fetch fetches the next page of data into the `Data` property
func (list *ShippingMethodList) Fetch() error {
	resources := &shippingMethodList{}
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
