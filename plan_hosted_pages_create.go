// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type PlanHostedPagesCreate struct {

	// URL to redirect to after signup on the hosted payment pages.
	SuccessUrl *string `json:"success_url,omitempty"`

	// URL to redirect to on canceled signup on the hosted payment pages.
	CancelUrl *string `json:"cancel_url,omitempty"`

	// If `true`, the customer will be sent directly to your `success_url` after a successful signup, bypassing Recurly's hosted confirmation page.
	BypassConfirmation *bool `json:"bypass_confirmation,omitempty"`

	// Determines if the quantity field is displayed on the hosted pages for the plan.
	DisplayQuantity *bool `json:"display_quantity,omitempty"`
}
