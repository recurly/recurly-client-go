// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type InvoiceUpdatable struct {
	Params `json:"-"`

	// This identifies the PO number associated with the invoice. Not editable for credit invoices.
	PoNumber *string `json:"po_number,omitempty"`

	// VAT Reverse Charge Notes are editable only if there was a VAT reverse charge applied to the invoice.
	VatReverseChargeNotes *string `json:"vat_reverse_charge_notes,omitempty"`

	// Terms and conditions are an optional note field. Not editable for credit invoices.
	TermsAndConditions *string `json:"terms_and_conditions,omitempty"`

	// Customer notes are an optional note field.
	CustomerNotes *string `json:"customer_notes,omitempty"`

	// Integer representing the number of days after an invoice's creation that the invoice will become past due. Changing Net terms changes due_on, and the invoice could move between past due and pending.
	NetTerms *int `json:"net_terms,omitempty"`

	Address *InvoiceAddressCreate `json:"address,omitempty"`
}

func (attr *InvoiceUpdatable) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
