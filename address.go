package recurly

type Address struct {
	FirstName  *string `json:"first_name,omitifempty"`
	LastName   *string `json:"last_name,omitifempty"`
	Phone      *string `json:"phone,omitifempty"`
	Street1    *string `json:"street1,omitifempty"`
	Street2    *string `json:"street2,omitifempty"`
	City       *string `json:"city,omitifempty"`
	State      *string `json:"state,omitifempty"`
	Region     *string `json:"region,omitifempty"`
	PostalCode *string `json:"postal_code,omitifempty"`
	Country    *string `json:"country,omitifempty"`
}
