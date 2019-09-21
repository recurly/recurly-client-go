package recurly

import "time"

type BillingInfo struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Company   string    `json:"company"`
	Address   Address   `json:"address,omitempty"`
	VATNumber string    `json:"vat_number"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BillingInfoAttributes for creating or updating billing information
type BillingInfoAttributes struct {
	Params                   `json:"-"`
	FirstName                string   `json:"first_name,omitempty"`
	LastName                 string   `json:"last_name,omitempty"`
	Company                  string   `json:"company,omitempty"`
	Address                  *Address `json:"address,omitempty"`
	VATNumber                string   `json:"vat_number,omitempty"`
	Number                   string   `json:"number,omitempty"`
	Month                    string   `json:"month,omitempty"`
	Year                     string   `json:"year,omitempty"`
	IPAddress                string   `json:"ip_address,omitempty"`
	AmazonBillingAgreementID string   `json:"amazon_billing_agreement_id,omitempty"`
	PaypalBillingAgreementID string   `json:"paypal_billing_agreement_id,omitempty"`
	FraudSessionID           string   `json:"fraud_session_id,omitempty"`
}
