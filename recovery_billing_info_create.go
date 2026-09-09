// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type RecoveryBillingInfoCreate struct {

	// First name
	FirstName *string `json:"first_name,omitempty"`

	// Last name
	LastName *string `json:"last_name,omitempty"`

	// Company name
	Company *string `json:"company,omitempty"`

	Address *RecoveryAddress `json:"address,omitempty"`

	// *STRONGLY RECOMMENDED* Customer's IP address when updating their billing information.
	IpAddress *string `json:"ip_address,omitempty"`

	// An identifier for a specific payment gateway.
	GatewayCode *string `json:"gateway_code,omitempty"`

	// The `primary_payment_method` field is used to designate the primary billing info on the account. An account can have a maximum of 1 primary. If a user sets a different payment method as a primary, then the existing primary will no longer be marked as such.
	PrimaryPaymentMethod *bool `json:"primary_payment_method,omitempty"`

	// The `backup_payment_method` field is used to designate a billing info as a backup on the account that will be tried if the initial billing info used for an invoice is declined. All payment methods, including the billing info marked `primary_payment_method` can be set as a backup. An account can have a maximum of 1 backup, if a user sets a different payment method as a backup, the existing backup will no longer be marked as such.
	BackupPaymentMethod *bool `json:"backup_payment_method,omitempty"`

	// Array of Payment Gateway References, each a reference to a third-party gateway object of varying types.
	PaymentGatewayReferences *[]PaymentGatewayReferencesCreate `json:"payment_gateway_references,omitempty"`

	// Merchant-supplied fallback payment method metadata. Recurly's own gateway-token lookup is authoritative and will override any of these fields it can determine itself; these fields are only used to fill gaps when that lookup is unavailable.
	PaymentMethod *RecoveryPaymentMethodCreate `json:"payment_method,omitempty"`

	// Network transaction ID from the previous customer-in-session subscription signup or billing info storage.
	// - 10-15 alphanumeric characters for Mastercard
	// - 14-15 alphanumeric for Visa
	// - 15 digits for all other brands
	// - 16 alphanumeric characters for Cartes Bancaires, which are processed as Visa or Mastercard
	NetworkTransactionId *string `json:"network_transaction_id,omitempty"`

	// Transactions from previous collection attempts for this payment method. Optional, unless this billing_info is the primary payment method and the account's dunning campaign skips Recurly's own retry attempts entirely -- in that case at least one entry is required.
	Transactions *[]RecoveryTransactionCreate `json:"transactions,omitempty"`
}
