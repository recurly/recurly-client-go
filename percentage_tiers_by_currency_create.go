// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
)

type PercentageTiersByCurrencyCreate struct {
  

  
        // 3-letter ISO 4217 currency code.
        Currency *string `json:"currency,omitempty"`

  
        // Tiers
        Tiers []PercentageTierCreate `json:"tiers,omitempty"`

  
}


