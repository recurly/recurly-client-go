// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
)

type PlanRampIntervalCreate struct {
  

  
        // Represents the billing cycle where a ramp interval starts.
        StartingBillingCycle *int `json:"starting_billing_cycle,omitempty"`

  
        // Represents the price for the ramp interval.
        Currencies []PlanRampPricingCreate `json:"currencies,omitempty"`

  
}


