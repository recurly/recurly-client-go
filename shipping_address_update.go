// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type ShippingAddressUpdate struct {
	Params `json:"-"`

	// Shipping Address ID
	Id *string `json:"id,omitempty"`

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

	// Country, 2-letter ISO 3166-1 alpha-2 code.
	Country *string `json:"country,omitempty"`

	// Code that represents a geographic entity (location or object). Only returned for Sling Vertex Integration
	GeoCode *string `json:"geo_code,omitempty"`
}

func (attr *ShippingAddressUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
