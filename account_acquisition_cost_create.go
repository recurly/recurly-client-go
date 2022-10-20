// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
)

type AccountAcquisitionCostCreate struct {
  

  
        // 3-letter ISO 4217 currency code.
        Currency *string `json:"currency,omitempty"`

  
        // The amount of the corresponding currency used to acquire the account.
        Amount *float64 `json:"amount,omitempty"`

  
}


