// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
)

type PurchaseCreate struct {
  

  
        // 3-letter ISO 4217 currency code.
        Currency *string `json:"currency,omitempty"`

  
        Account *AccountPurchase `json:"account,omitempty"`

  
        // The `billing_info_id` is the value that represents a specific billing info for an end customer. When `billing_info_id` is used to assign billing info to the subscription, all future billing events for the subscription will bill to the specified billing info. `billing_info_id` can ONLY be used for sites utilizing the Wallet feature.
        BillingInfoId *string `json:"billing_info_id,omitempty"`

  
        // Must be set to manual in order to preview a purchase for an Account that does not have payment information associated with the Billing Info.
        CollectionMethod *string `json:"collection_method,omitempty"`

  
        // For manual invoicing, this identifies the PO number associated with the subscription.
        PoNumber *string `json:"po_number,omitempty"`

  
        // Integer representing the number of days after an invoice's creation that the invoice will become past due. If an invoice's net terms are set to '0', it is due 'On Receipt' and will become past due 24 hours after it’s created. If an invoice is due net 30, it will become past due at 31 days exactly.
        NetTerms *int `json:"net_terms,omitempty"`

  
        // Terms and conditions to be put on the purchase invoice.
        TermsAndConditions *string `json:"terms_and_conditions,omitempty"`

  
        // Customer notes
        CustomerNotes *string `json:"customer_notes,omitempty"`

  
        // VAT reverse charge notes for cross border European tax settlement.
        VatReverseChargeNotes *string `json:"vat_reverse_charge_notes,omitempty"`

  
        // Notes to be put on the credit invoice resulting from credits in the purchase, if any.
        CreditCustomerNotes *string `json:"credit_customer_notes,omitempty"`

  
        // The default payment gateway identifier to be used for the purchase transaction.  This will also be applied as the default for any subscriptions included in the purchase request.
        GatewayCode *string `json:"gateway_code,omitempty"`

  
        Shipping *ShippingPurchase `json:"shipping,omitempty"`

  
        // A list of one time charges or credits to be created with the purchase.
        LineItems []LineItemCreate `json:"line_items,omitempty"`

  
        // A list of subscriptions to be created with the purchase.
        Subscriptions []SubscriptionPurchase `json:"subscriptions,omitempty"`

  
        // A list of coupon_codes to be redeemed on the subscription or account during the purchase.
        CouponCodes []string `json:"coupon_codes,omitempty"`

  
        // A gift card redemption code to be redeemed on the purchase invoice.
        GiftCardRedemptionCode *string `json:"gift_card_redemption_code,omitempty"`

  
        // An optional type designation for the payment gateway transaction created by this request. Supports 'moto' value, which is the acronym for mail order and telephone transactions.
        TransactionType *string `json:"transaction_type,omitempty"`

  
}


