package recurly

import (
	"time"
)

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}

func Float(v float64) *float64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func Array(v []string) *[]string {
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}

func StringSlice(v []string) []*string {
	out := make([]*string, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

type Empty struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Empty) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Empty) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

type ListMetadata struct {
	ObjectName string `json:"object"`
	HasMore    bool   `json:"has_more"`
	Next       string `json:"next"`
}

type Resource interface {
	GetResponse() *ResponseMetadata
	setResponse(*ResponseMetadata)
}
