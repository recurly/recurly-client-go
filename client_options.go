package recurly

// Region specific to Recurly
type region string

const (
	// US Datacenter region
	US = region("us")
	// EU Datacenter region
	EU = region("eu")
)

// ClientOptions for a new API Client
type ClientOptions struct {
	Region region
}
