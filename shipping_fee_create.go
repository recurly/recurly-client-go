// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
)

type ShippingFeeCreate struct {
  

  
        // The id of the shipping method used to deliver the purchase. If `method_id` and `method_code` are both present, `method_id` will be used.
        MethodId *string `json:"method_id,omitempty"`

  
        // The code of the shipping method used to deliver the purchase. If `method_id` and `method_code` are both present, `method_id` will be used.
        MethodCode *string `json:"method_code,omitempty"`

  
        // This is priced in the purchase's currency.
        Amount *float64 `json:"amount,omitempty"`

  
}


