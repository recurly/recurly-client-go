// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type DunningIntervalCreate struct {

	// Number of days before sending the next email.
	Days *int `json:"days,omitempty"`

	// The id of the custom email template to assign to this interval, from `GET /dunning_campaigns/email_templates`. `null` uses the system default template for this interval.
	EmailTemplateId *string `json:"email_template_id,omitempty"`
}
