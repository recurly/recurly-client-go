// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "time"
)

type ExternalTransaction struct {
  

  
        // Payment method used for external transaction.
        PaymentMethod *string `json:"payment_method,omitempty"`

  
        // Used as the transaction's description.
        Description *string `json:"description,omitempty"`

  
        // The total amount of the transcaction. Cannot excceed the invoice total.
        Amount *float64 `json:"amount,omitempty"`

  
        // Datetime that the external payment was collected. Defaults to current datetime.
        CollectedAt *time.Time `json:"collected_at,omitempty"`

  
}


