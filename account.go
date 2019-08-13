package recurly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	Code           AccountCode `json:"code"`
	Username       *string     `json:"username,omitifempty"`
	Email          *string     `json:"email,omitifempty"`
	PreferredLocal *string     `json:"preferred_locale,omitifempty"`
	CCEmails       *string     `json:"cc_emails,omitifempty"`
	FirstName      *string     `json:"first_name,omitifempty"`
	LastName       *string     `json:"last_name,omitifempty"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	DeletedAt      *time.Time  `json:"deleted_at"`
}

func (account *Account) String() string {
	data, err := json.Marshal(&account)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// GetAccountByID returns an account with the given account ID
func (c *Client) GetAccountByID(accountID int64) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodGet, fmt.Sprintf("%s/%d", accountsRoot, accountID), nil, account)
	return account, err
}

// GetAccountByAccountCode returns an account with the given account code
func (c *Client) GetAccountByAccountCode(accountCode AccountCode) (*Account, error) {
	account := &Account{}
	err := c.Call(http.MethodGet, fmt.Sprintf("%s/%s", accountsRoot, urlEncode(accountCode)), nil, account)
	return account, err
}
