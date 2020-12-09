// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type MeasuredUnitCreate struct {

	// Unique internal name of the measured unit on your site.
	Name *string `json:"name,omitempty"`

	// Display name for the measured unit.
	DisplayName *string `json:"display_name,omitempty"`

	// Optional internal description.
	Description *string `json:"description,omitempty"`
}
