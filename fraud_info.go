package recurly

import (
	"net/http"
)

type FraudInfo struct {
	recurlyResponse *ResponseMetadata

	// Kount score
	Score int `json:"score,omitempty"`

	// Kount decision
	Decision string `json:"decision,omitempty"`

	// Kount rules
	RiskRulesTriggered map[string]interface{} `json:"risk_rules_triggered,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *FraudInfo) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *FraudInfo) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type fraudInfoList struct {
	ListMetadata
	Data            []FraudInfo `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *fraudInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *fraudInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// FraudInfoList allows you to paginate FraudInfo objects
type FraudInfoList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []FraudInfo
}

func NewFraudInfoList(client HttpCaller, nextPagePath string) *FraudInfoList {
	return &FraudInfoList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *FraudInfoList) Fetch() error {
	resources := &fraudInfoList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *FraudInfoList) Count() (*int64, error) {
	resources := &fraudInfoList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
