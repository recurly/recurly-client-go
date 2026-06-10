// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"time"
)

type RecoveryTransactionCreate struct {

	// The error code returned by the payment gateway for the original payment collection attempt.
	GatewayErrorCode *string `json:"gateway_error_code,omitempty"`

	// The advice code returned by the payment gateway for the original payment collection attempt. This field is only applicable for certain gateways.
	MerchantAdviceCode *string `json:"merchant_advice_code,omitempty"`

	// The date the original payment collection was attempted.
	AttemptedCollectionDate *time.Time `json:"attempted_collection_date,omitempty"`
}
