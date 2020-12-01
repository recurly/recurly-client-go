// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type AccountAcquisitionUpdatable struct {

	// Account balance
	Cost *AccountAcquisitionCostCreate `json:"cost,omitempty"`

	// The channel through which the account was acquired.
	Channel *string `json:"channel,omitempty"`

	// An arbitrary subchannel string representing a distinction/subcategory within a broader channel.
	Subchannel *string `json:"subchannel,omitempty"`

	// An arbitrary identifier for the marketing campaign that led to the acquisition of this account.
	Campaign *string `json:"campaign,omitempty"`
}
