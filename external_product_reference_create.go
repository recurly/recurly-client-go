// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type ExternalProductReferenceCreate struct {

	// A code which associates the external product to a corresponding object or resource in an external platform like the Apple App Store or Google Play Store.
	ReferenceCode *string `json:"reference_code,omitempty"`

	ExternalConnectionType *string `json:"external_connection_type,omitempty"`
}
