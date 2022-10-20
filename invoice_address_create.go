// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
)

type InvoiceAddressCreate struct {
  

  
        // Name on account
        NameOnAccount *string `json:"name_on_account,omitempty"`

  
        // Company
        Company *string `json:"company,omitempty"`

  
        // Phone number
        Phone *string `json:"phone,omitempty"`

  
        // Street 1
        Street1 *string `json:"street1,omitempty"`

  
        // Street 2
        Street2 *string `json:"street2,omitempty"`

  
        // City
        City *string `json:"city,omitempty"`

  
        // State or province.
        Region *string `json:"region,omitempty"`

  
        // Zip or postal code.
        PostalCode *string `json:"postal_code,omitempty"`

  
        // Country, 2-letter ISO 3166-1 alpha-2 code.
        Country *string `json:"country,omitempty"`

  
        // First name
        FirstName *string `json:"first_name,omitempty"`

  
        // Last name
        LastName *string `json:"last_name,omitempty"`

  
}


