// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "time"
)

type UsageCreate struct {
  

  
        // Custom field for recording the id in your own system associated with the usage, so you can provide auditable usage displays to your customers using a GET on this endpoint.
        MerchantTag *string `json:"merchant_tag,omitempty"`

  
        // The amount of usage. Can be positive, negative, or 0. No decimals allowed, we will strip them. If the usage-based add-on is billed with a percentage, your usage will be a monetary amount you will want to format in cents. (e.g., $5.00 is "500").
        Amount *float64 `json:"amount,omitempty"`

  
        // When the usage was recorded in your system.
        RecordingTimestamp *time.Time `json:"recording_timestamp,omitempty"`

  
        // When the usage actually happened. This will define the line item dates this usage is billed under and is important for revenue recognition.
        UsageTimestamp *time.Time `json:"usage_timestamp,omitempty"`

  
}


