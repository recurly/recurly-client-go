package recurly

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}

func Bool(v bool) *bool {
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
}

type ListMetadata struct {
	ObjectName string `json:"object"`
	HasMore    bool   `json:"has_more"`
	Next       string `json:"next"`
}
