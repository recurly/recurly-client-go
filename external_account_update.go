// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type ExternalAccountUpdate struct {

	// Represents the account code for the external account.
	ExternalAccountCode *string `json:"external_account_code,omitempty"`

	// Represents the connection type. `AppleAppStore` or `GooglePlayStore`
	ExternalConnectionType *string `json:"external_connection_type,omitempty"`
}
