package recurly

import "time"

type BillingInfo struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Company   *string   `json:"company"`
	Address   *Address  `json:"address,omitifemtpy"`
	VATNumber *string   `json:"vat_number"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
