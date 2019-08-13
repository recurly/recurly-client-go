package recurly

type CustomFields []CustomFields

type CustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
