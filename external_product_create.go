// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import ()

type ExternalProductCreate struct {

	// External product name.
	Name *string `json:"name,omitempty"`

	// Recurly plan UUID.
	PlanId *string `json:"plan_id,omitempty"`

	// List of external product references of the external product.
	ExternalProductReferences []ExternalProductReferenceBase `json:"external_product_references,omitempty"`
}
