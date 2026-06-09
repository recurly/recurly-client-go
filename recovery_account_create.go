// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type RecoveryAccountCreate struct {
	Address *RecoveryAddress `json:"address,omitempty"`

	// If the premium Wallet feature is enabled, more than one payment method can be associated with an account, and one can be designated as a primary and one as a backup. Without the Wallet feature, only one payment method will be accepted.
	BillingInfos *[]RecoveryBillingInfoCreate `json:"billing_infos,omitempty"`

	// The unique identifier of the account. This cannot be changed once the account is created.
	Code *string `json:"code,omitempty"`

	// The email address used for communicating with this customer.
	Email *string `json:"email,omitempty"`

	// The custom fields will only be altered when they are included in a request. Sending an empty array will not remove any existing values. To remove a field send the name with a null or empty value.
	CustomFields *[]CustomFieldCreate `json:"custom_fields,omitempty"`

	// Unique ID to identify a dunning campaign. Used to specify if a non-default dunning campaign should be assigned to this account. For sites without multiple dunning campaigns enabled, the default dunning campaign will always be used.
	DunningCampaignId *string `json:"dunning_campaign_id,omitempty"`
}
