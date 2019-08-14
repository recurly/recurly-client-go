package recurly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	// APIVersion is the current Recurly API Version
	APIVersion = "v2018-08-09"

	// APIHost is the base URL for Recurly API v3
	APIHost = "https://partner-api.recurly.com"

	userAgent = "recurly-go/0.0.1" // TODO: dynamically insert version
)

var (
	// APIKey contains the Recurly API key used for authenticating globally.
	// A new client will use this value unless the API Key is explicitly set
	// whe ncreating the client.
	APIKey string

	// SiteID is the default site identifier to use globally if not specified
	// when creating a new client.
	SiteID SiteIdentifier

	// DefaultHTTPClient provides a default HTTP client for making HTTPS requests to Recurly
	DefaultHTTPClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // disable redirects
		},
		Jar:     nil,
		Timeout: 30 * time.Second,
	}
)

// Client submits API requests to Recurly
type Client struct {
	apiKey  string
	baseURL string
	siteID  SiteIdentifier

	Log        *Logger
	HTTPClient *http.Client
}

// SiteIdentifier uniquely identifies a site making the request
type SiteIdentifier struct {
	ID        int64
	Subdomain string
}

// URLEncode returns a URL encoded string to be used in API request paths
func (id *SiteIdentifier) URLEncode() string {
	if id.ID > 0 {
		return strconv.FormatInt(id.ID, 10)
	}
	return fmt.Sprintf("subdomain-%s", url.PathEscape(id.Subdomain))
}

// DefaultClient returns the default API Client using the globally set Site and APIKey
func DefaultClient() *Client {
	return &Client{
		apiKey:     APIKey,
		baseURL:    APIHost,
		siteID:     SiteID,
		Log:        NewLogger(LevelDebug),
		HTTPClient: DefaultHTTPClient,
	}
}

// NewClient creates a new Recurly API Client
func NewClient(apiKey string, siteID SiteIdentifier, httpClient *http.Client) *Client {
	if apiKey == "" {
		apiKey = APIKey
	}
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		apiKey:     apiKey,
		baseURL:    APIHost,
		siteID:     siteID,
		Log:        NewLogger(LevelDebug),
		HTTPClient: httpClient,
	}
}

// Call sends a request to Recurly and parses the JSON response for the expected response type
func (c *Client) Call(method string, path string, genericParams GenericParams, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if strings.HasPrefix(path, sitesRoot) {
		path = fmt.Sprintf("%s%s", c.baseURL, path)
	} else {
		path = fmt.Sprintf("%s/sites/%s%s", c.baseURL, c.siteID.URLEncode(), path)
	}

	var params *Params
	if genericParams != nil && !reflect.ValueOf(genericParams).IsNil() { // test if the interface is nil
		params = genericParams.toParams()
	}

	req, err := c.NewRequest(method, path, params)
	if err != nil {
		return err
	}

	return c.Do(req, v)
}

// NewRequest generates an http.Request for the API client to submit to Recurly.
func (c *Client) NewRequest(method string, requestURL string, params *Params) (*http.Request, error) {
	if params != nil {
		// Append URL parameters
		if keyValues := params.URLParams(); len(keyValues) > 0 {
			var buf bytes.Buffer
			for i, kv := range keyValues {
				if i > 0 {
					buf.WriteByte('&')
				}
				buf.WriteString(kv.Key)
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(kv.Value))
			}
			requestURL = fmt.Sprintf("%s?%s", requestURL, buf.String())
		}
	}

	req, err := http.NewRequest(method, requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", fmt.Sprintf("Accept: application/vnd.recurly.%s", APIVersion))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Recurly-Version", fmt.Sprintf("recurly.%s", APIVersion))
	req.Header.Add("User-Agent", userAgent)
	req.SetBasicAuth(c.apiKey, "")

	if params != nil {
		// TODO: generate an idempotency key if missing?
		if params.IdempotencyKey != "" {
			req.Header.Add("Idempotency-Key", params.IdempotencyKey)
		}
		for key, v := range params.Header {
			for _, value := range v {
				req.Header.Set(key, value)
			}
		}

		if params.Data != nil && method != http.MethodGet {
			data, err := json.Marshal(params.Data)
			if err != nil {
				c.Log.Errorf("Failed to marshal data to JSON payload: %v", err)
				return nil, err
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		}

		if params.Context != nil {
			req = req.WithContext(params.Context)
		}
	}

	return req, nil
}

// Do submits the http.Request to Recurly's API and parses the JSON response
func (c *Client) Do(req *http.Request, v interface{}) error {
	c.Log.Debugf("Requesting %s %s", req.Method, req.URL.String())

	startTime := time.Now()
	res, err := c.HTTPClient.Do(req)
	requestTime := time.Since(startTime)

	if err != nil {
		c.Log.Errorf("Request failed: %v", err)
		return err
	}
	c.Log.Debugf("Response: %d, %0.3f sec", res.StatusCode, requestTime.Seconds())
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Log.Errorf("Cannot read response: %s", err)
		return err
	}

	if c.Log.IsLevel(LevelDebug) { // do not waste cpu unless we're at debug level
		// TODO: How do we want to expose the rate limit info?
		// Maybe use a chan(RateLimit) for something to get the data?
		meta := c.parseResponseMetadata(req.URL, res)

		c.Log.Debugf("Headers: %s\nBody:\n%s\n", meta.String(), body)
	}

	if err = c.unmarshalJSON(body, v); err != nil {
		// TODO: inspect the status code to determine if a 5xx error is why we can't parse
		return err
	}

	// TODO: On non-2xx response, return an error
	// 404 -- object not found
	// 422 -- invalid request
	// 429 -- rate limited

	return nil
}

func (c *Client) unmarshalJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		c.Log.Errorf("Failed to deserialize JSON:\n%s", data)
	}
	return err
}
