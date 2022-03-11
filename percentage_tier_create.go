// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type PercentageTierCreate struct {

	// Ending amount for the tier. Allows up to 2 decimal places. The last tier ending_amount is null.
	EndingAmount *float64 `json:"ending_amount,omitempty"`

	// Decimal usage percentage.
	UsagePercentage *string `json:"usage_percentage,omitempty"`
}
