// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "time"
)

type ExternalRefund struct {
  

  
        // Payment method used for external refund transaction.
        PaymentMethod *string `json:"payment_method,omitempty"`

  
        // Used as the refund transactions' description.
        Description *string `json:"description,omitempty"`

  
        // Date the external refund payment was made. Defaults to the current date-time.
        RefundedAt *time.Time `json:"refunded_at,omitempty"`

  
}


