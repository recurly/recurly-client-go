// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type DunningCampaignsBulkUpdate struct {

	// List of `plan_codes` associated with the Plans for which the dunning campaign should be updated. Required unless `plan_ids` is present.
	PlanCodes []string `json:"plan_codes,omitempty"`

	// List of `plan_ids` associated with the Plans for which the dunning campaign should be updated. Required unless `plan_codes` is present.
	PlanIds []string `json:"plan_ids,omitempty"`
}
