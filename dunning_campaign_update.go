// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type DunningCampaignUpdate struct {

	// Campaign name.
	Name *string `json:"name,omitempty"`

	// Campaign code.
	Code *string `json:"code,omitempty"`

	// Campaign description.
	Description *string `json:"description,omitempty"`

	// Set to `true` to make this the default campaign for accounts or plans without an assigned dunning campaign. Cannot be set on an inactive campaign, and cannot be unset directly—assign a different campaign as the default instead.
	DefaultCampaign *bool `json:"default_campaign,omitempty"`

	// Dunning Cycle settings. One entry per collection method (`automatic`, `manual`, `trial`); each type may appear at most once. Each cycle write fully replaces that cycle's current settings version.
	DunningCycles *[]DunningCycleCreate `json:"dunning_cycles,omitempty"`
}
