// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type PaymentGatewayReferencesCreate struct {

	// Reference value used when the external token was created. If Stripe gateway is used, this value will need to be accompanied by its reference_type.
	Token *string `json:"token,omitempty"`

	// The type of reference token. Required if token is passed in for Stripe Gateway.
	ReferenceType *string `json:"reference_type,omitempty"`
}
