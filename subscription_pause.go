// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type SubscriptionPause struct {

	// Number of billing cycles to pause the subscriptions.
	RemainingPauseCycles *int `json:"remaining_pause_cycles,omitempty"`
}
