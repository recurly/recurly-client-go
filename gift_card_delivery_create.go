// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type GiftCardDeliveryCreate struct {

	// Whether the delivery method is email or postal service.
	Method *string `json:"method,omitempty"`

	// The email address of the recipient. Required if `method` is `email`.
	EmailAddress *string `json:"email_address,omitempty"`

	// When the gift card should be delivered to the recipient. If null, the gift card will be delivered immediately. If a datetime is provided, the delivery will be in an hourly window, rounding down. For example, 6:23 pm will be in the 6:00 pm hourly batch.
	DeliverAt *time.Time `json:"deliver_at,omitempty"`

	// The first name of the recipient.
	FirstName *string `json:"first_name,omitempty"`

	// The last name of the recipient.
	LastName *string `json:"last_name,omitempty"`

	// Address information for the recipient. Required if `method` is `post`.
	RecipientAddress *AddressCreate `json:"recipient_address,omitempty"`

	// The name of the gifter for the purpose of a message displayed to the recipient.
	GifterName *string `json:"gifter_name,omitempty"`

	// The personal message from the gifter to the recipient.
	PersonalMessage *string `json:"personal_message,omitempty"`
}
