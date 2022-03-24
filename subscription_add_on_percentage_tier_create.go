// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type SubscriptionAddOnPercentageTierCreate struct {

	// Ending amount for the tier. Allows up to 2 decimal places. Must be left empty if it is the final tier.
	EndingAmount *float64 `json:"ending_amount,omitempty"`

	// The percentage taken of the monetary amount of usage tracked.
	// This can be up to 4 decimal places represented as a string.
	UsagePercentage *string `json:"usage_percentage,omitempty"`
}
