// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type GiftCardCreate struct {

	// The product code or SKU of the gift card product.
	ProductCode *string `json:"product_code,omitempty"`

	// The amount of the gift card, which is the amount of the charge to the gifter account and the amount of credit that is applied to the recipient account upon successful redemption.
	UnitAmount *float64 `json:"unit_amount,omitempty"`

	// 3-letter ISO 4217 currency code.
	Currency *string `json:"currency,omitempty"`

	// The delivery details for the gift card.
	Delivery *GiftCardDeliveryCreate `json:"delivery,omitempty"`

	// Block of account details for the gifter. This references an existing account_code.
	GifterAccount *AccountPurchase `json:"gifter_account,omitempty"`
}
