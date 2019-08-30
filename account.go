package recurly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	accountsRoot = "/accounts"
)

// AccountCode Recurly account code
type AccountCode string

func urlEncode(accountCode AccountCode) string {
	return fmt.Sprintf("code-%s", url.PathEscape(string(accountCode)))
}

// Account account
// See https://partner-docs.recurly.com/v2018-08-09#tag/account
type Account struct {
	BaseObject
	Code                 AccountCode  `json:"code"`
	State                AccountState `json:"state"`
	HostedLoginToken     string       `json:"hosted_login_token"`
	ShippingAddresses    []Address    `json:"shipping_addresses"`
	Username             string       `json:"username,omitempty"`
	Email                string       `json:"email,omitempty"`
	PreferredLocal       string       `json:"preferred_locale,omitempty"`
	CCEmails             string       `json:"cc_emails,omitempty"`
	FirstName            string       `json:"first_name,omitempty"`
	LastName             string       `json:"last_name,omitempty"`
	Company              string       `json:"company"`
	VATNumber            string       `json:"vat_number"`
	TaxExempt            bool         `json:"tax_exempt"`
	ExceptionCertificate string       `json:"exception_certificate"`
	ParentAccountID      string       `json:"parent_account_id"`
	BillTo               string       `json:"bill_to"`
	Address              *Address     `json:"address,omitempty"`
	BillingInfo          *BillingInfo `json:"billing_info"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
	DeletedAt            *time.Time   `json:"deleted_at"`
	CustomFields         CustomFields `json:"custom_fields"`
}

// CreateAccountAttributes for creating a new account
type CreateAccountAttributes struct {
	Params `json:"-"`
	// Acquisition
	Code                 string                 `json:"code"`
	ShippingAddresses    []Address              `json:"shipping_addresses,omitempty"`
	Username             string                 `json:"username,omitempty"`
	Email                string                 `json:"email,omitempty"`
	PreferredLocal       string                 `json:"preferred_locale,omitempty"`
	CCEmails             string                 `json:"cc_emails,omitempty"`
	FirstName            string                 `json:"first_name,omitempty"`
	LastName             string                 `json:"last_name,omitempty"`
	Company              string                 `json:"company,omitempty"`
	VATNumber            string                 `json:"vat_number,omitempty"`
	TaxExempt            *bool                  `json:"tax_exempt,omitempty"`
	ExceptionCertificate string                 `json:"exception_certificate,omitempty"`
	ParentAccountID      string                 `json:"parent_account_id,omitempty"`
	ParentAccountCode    string                 `json:"parent_account_code,omitempty"`
	BillTo               string                 `json:"bill_to,omitempty"`
	Address              *Address               `json:"address,omitempty"`
	BillingInfo          *BillingInfoAttributes `json:"billing_info,omitempty"`
	CustomFields         CustomFields           `json:"custom_fields,omitempty"`
}

func (attr *CreateAccountAttributes) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

//UpdateAccountAttributes for creating a new account (omits Code)
type UpdateAccountAttributes struct {
	Params `json:"-"`
	// Acquisition
	ShippingAddresses    []Address              `json:"shipping_addresses,omitempty"`
	Username             string                 `json:"username,omitempty"`
	Email                string                 `json:"email,omitempty"`
	PreferredLocal       string                 `json:"preferred_locale,omitempty"`
	CCEmails             string                 `json:"cc_emails,omitempty"`
	FirstName            string                 `json:"first_name,omitempty"`
	LastName             string                 `json:"last_name,omitempty"`
	Company              string                 `json:"company,omitempty"`
	VATNumber            string                 `json:"vat_number,omitempty"`
	TaxExempt            *bool                  `json:"tax_exempt,omitempty"`
	ExceptionCertificate string                 `json:"exception_certificate,omitempty"`
	ParentAccountID      string                 `json:"parent_account_id,omitempty"`
	ParentAccountCode    string                 `json:"parent_account_code,omitempty"`
	BillTo               string                 `json:"bill_to,omitempty"`
	Address              *Address               `json:"address,omitempty"`
	BillingInfo          *BillingInfoAttributes `json:"billing_info,omitempty"`
	CustomFields         CustomFields           `json:"custom_fields,omitempty"`
}

func (attr *UpdateAccountAttributes) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

// AccountList allows pagination across a list of accounts
type AccountList struct {
	client       *Client
	nextPagePath string

	HasMore bool
	Data    []Account
}

// internal struct for deserializing accounts
type accountList struct {
	ListMetadata
	Data []Account `json:"data"`
}

type ListMetadata struct {
	ObjectName string `json:"object"`
	HasMore    bool   `json:"has_more"`
	Next       string `json:"next"`
}

func newAccountsList(client *Client, list *accountList) *AccountList {
	return &AccountList{
		client:       client,
		nextPagePath: list.Next,
		HasMore:      list.HasMore,
		Data:         list.Data,
	}
}

// NextPage returns the next page of accounts
func (list AccountList) NextPage() (*AccountList, error) {
	if !list.HasMore {
		return nil, nil
	}

	accounts := &accountList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, accounts)
	if err != nil {
		return nil, err
	}
	return newAccountsList(list.client, accounts), nil
}

type AccountState string

const (
	AccountStateActive = "active"
	AccountStateClosed = "closed"
)

// BaseObject includes an ID and object type for all Recurly records
type BaseObject struct {
	ID     ObjectID `json:"id"`
	Object string   `json:"object"`
}

// ObjectID base-36 encoded Recurly object ID
type ObjectID string

type AccountListParams struct {
	ListParams

	// Subscriber returns only subscribers or non-subscribers if non-nil.
	Subscriber *bool
	// PastDue returns only past-due or non-past-due if non-nil.
	PastDue *bool
}

func (list *AccountListParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
}

func (list *AccountListParams) URLParams() []KeyValue {
	options := list.ListParams.URLParams()

	if list.Subscriber != nil {
		options = append(options, KeyValue{Key: "subscriber", Value: formatBool(*list.Subscriber)})
	}
	if list.PastDue != nil {
		options = append(options, KeyValue{Key: "past_due", Value: formatBool(*list.PastDue)})
	}

	return options
}

// ToInt64 converts an ObjectID to an int64 for more efficient storage
func (id ObjectID) ToInt64() int64 {
	intID, _ := strconv.ParseInt(string(id), 36, 64)
	return intID
}

// ToObjectID converts an int64 to an ObjectID
func ToObjectID(id int64) ObjectID {
	return ObjectID(strconv.FormatInt(id, 36))
}

func (account *Account) String() string {
	data, err := json.Marshal(&account)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// GetAccount returns an account with the given account ID
func (c *Client) GetAccount(accountID ObjectID, params GenericParams) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodGet, fmt.Sprintf("%s/%s", accountsRoot, accountID), params, account)
	return account, err
}

// GetAccountByAccountCode returns an account with the given account code
func (c *Client) GetAccountByAccountCode(accountCode AccountCode, params GenericParams) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodGet, fmt.Sprintf("%s/%s", accountsRoot, urlEncode(accountCode)), params, account)
	return account, err
}

// ListAccounts returns an array of accounts
func (c *Client) ListAccounts(params *AccountListParams) (*AccountList, error) {
	accounts := &accountList{}
	err := c.Call(http.MethodGet, accountsRoot, params, accounts)
	if err != nil {
		return nil, err
	}
	return newAccountsList(c, accounts), nil
}

// CreateAccount creates a new account
func (c *Client) CreateAccount(attributes *CreateAccountAttributes) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodPost, accountsRoot, attributes, account)
	return account, err
}

// UpdateAccount updates an account
func (c *Client) UpdateAccount(accountID ObjectID, attributes *UpdateAccountAttributes) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodPut, fmt.Sprintf("%s/%s", accountsRoot, accountID), attributes, account)
	return account, err
}

// UpdateAccountByAccountCode updates an account
func (c *Client) UpdateAccountByAccountCode(accountCode AccountCode, attributes *UpdateAccountAttributes) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodPut, fmt.Sprintf("%s/%s", accountsRoot, urlEncode(accountCode)), attributes, account)
	return account, err
}

// DeactivateAccount closes an account
func (c *Client) DeactivateAccount(accountID ObjectID, params GenericParams) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodDelete, fmt.Sprintf("%s/%s", accountsRoot, accountID), params, account)
	return account, err
}

// DeactivateAccountByAccountCode closes an account
func (c *Client) DeactivateAccountByAccountCode(accountCode AccountCode, params GenericParams) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodDelete, fmt.Sprintf("%s/%s", accountsRoot, urlEncode(accountCode)), params, account)
	return account, err
}
