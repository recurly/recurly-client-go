package recurly

import (
	"net/http"
	"time"
)

const (
	sitesRoot = "/sites"
)

type Site struct {
	BaseObject
	Subdomain    string       `json:"subdomain"`
	PublicAPIKey *string      `json:"public_api_key,omitifempty"`
	Mode         SiteMode     `json:"site_mode"`
	Address      *Address     `json:"address,omitifempty"`
	SiteSettings SiteSettings `json:"site_settings"`
	Features     []string     `json:"features,omitifempty"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at"`
}

type SiteSettings struct {
	BillingAddressRequirement string   `json:"billing_address_requirement"`
	AcceptedCurrencies        []string `json:"accepted_currencies`
	DefaultCurrency           string   `json:"default_currency"`
}

type SiteMode string

const (
	SiteModeProduction SiteMode = "production"
	SiteModeSandbox    SiteMode = "sandbox"
)

type SiteList struct {
	ListMetadata
	Data []*Site `json:"data"`
}

// ListSites returns an array of sites
func (c *Client) ListSites(params *ListParams) (*SiteList, error) {
	sites := &SiteList{}
	err := c.Call(http.MethodGet, sitesRoot, params, sites)
	return sites, err
}
