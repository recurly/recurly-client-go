package recurly

import ()

type ShippingAddressCreate struct {
	Params `json:"-"`

	Nickname *string `json:"nickname,omitempty"`

	FirstName *string `json:"first_name,omitempty"`

	LastName *string `json:"last_name,omitempty"`

	Company *string `json:"company,omitempty"`

	Email *string `json:"email,omitempty"`

	VatNumber *string `json:"vat_number,omitempty"`

	Phone *string `json:"phone,omitempty"`

	Street1 *string `json:"street1,omitempty"`

	Street2 *string `json:"street2,omitempty"`

	City *string `json:"city,omitempty"`

	// State or province.
	Region *string `json:"region,omitempty"`

	// Zip or postal code.
	PostalCode *string `json:"postal_code,omitempty"`

	// Country, 2-letter ISO code.
	Country *string `json:"country,omitempty"`
}

func (attr *ShippingAddressCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
