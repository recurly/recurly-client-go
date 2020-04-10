package recurly

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	defaultTimeout = 60 * time.Second
)

var (
	// APIKey contains the Recurly API key used for authenticating globally.
	// A new client will use this value unless the API Key is explicitly set
	// whe ncreating the client.
	APIKey string

	// APIHost is the base URL for Recurly API v3
	APIHost = "https://v3.recurly.com"

	defaultTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: time.Minute,
			DualStack: true,
		}).DialContext,
		IdleConnTimeout:     2 * time.Minute,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	// defaultClient provides a default HTTP client for making HTTPS requests to Recurly
	defaultClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // disable redirects
		},
		Jar:       nil,
		Timeout:   defaultTimeout,
		Transport: defaultTransport,
	}

	acceptVersion  = fmt.Sprintf("application/vnd.recurly.%s", APIVersion)
	recurlyVersion = fmt.Sprintf("recurly.%s", APIVersion)
	userAgent      = fmt.Sprintf("Recurly/%s; go %s", clientVersion, runtime.Version())
	pathPattern    = regexp.MustCompile(`{[^}]+}`)
)

// Client submits API requests to Recurly
type Client struct {
	apiKey  string
	baseURL string

	Log        *Logger
	HTTPClient *http.Client
}

// DefaultClient returns the default API Client using the globally set Site and APIKey
func DefaultClient() *Client {
	return &Client{
		apiKey:     APIKey,
		baseURL:    APIHost,
		Log:        NewLogger(LevelDebug),
		HTTPClient: defaultClient,
	}
}

// NewClient creates a new Recurly API Client
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if apiKey == "" {
		apiKey = APIKey
	}
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		apiKey:     apiKey,
		baseURL:    APIHost,
		Log:        NewLogger(LevelDebug),
		HTTPClient: httpClient,
	}
}

// Takes an OpenAPI-style path such as "/accounts/{account_id}/shipping_addresses/{shipping_address_id}"
// and a list of string arguments to fill the template, and it returns the interpolated path
func (c *Client) InterpolatePath(path string, params ...string) string {
	template := pathPattern.ReplaceAllString(path, "%s")
	encodedParams := make([]interface{}, len(params))
	for i, param := range params {
		encoded := url.PathEscape(param)
		encodedParams[i] = encoded
	}
	return fmt.Sprintf(template, encodedParams...)
}

// Call sends a request to Recurly and parses the JSON response for the expected response type
func (c *Client) Call(method string, path string, genericParams GenericParams, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("%s/%s", c.baseURL, path)
	} else {
		path = fmt.Sprintf("%s%s", c.baseURL, path)
	}

	path = BuildUrl(path, genericParams)

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

// Append URL parameters
func BuildUrl(requestURL string, genericParams GenericParams) string {
	var params *Params
	if genericParams != nil && !reflect.ValueOf(genericParams).IsNil() { // test if the interface is nil
		params = genericParams.toParams()
	}

	if params != nil {
		if keyValues := params.URLParams(); len(keyValues) > 0 {
			var buf bytes.Buffer
			buf.WriteString(requestURL)
			buf.WriteByte('?')

			for i, kv := range keyValues {
				if i > 0 {
					buf.WriteByte('&')
				}
				buf.WriteString(kv.Key)
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(kv.Value))
			}
			requestURL = buf.String()
		}
	}

	return requestURL
}

// NewRequest generates an http.Request for the API client to submit to Recurly.
func (c *Client) NewRequest(method string, requestURL string, params *Params) (*http.Request, error) {
	req, err := http.NewRequest(method, requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", acceptVersion)
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
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
	c.Log.Debugf("Response: %d, %5.3f sec", res.StatusCode, requestTime.Seconds())
	defer res.Body.Close()

	// Gzip decode
	bodyReader := res.Body
	contentEncoding := strings.ToLower(res.Header.Get("Content-Encoding"))
	if contentEncoding == "gzip" {
		gzipReader, err := gzip.NewReader(bodyReader)
		if err != nil {
			return err
		}
		bodyReader = gzipReader
	}

	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		c.Log.Errorf("Cannot read response: %s", err)
		return err
	}

	if c.Log.IsLevel(LevelDebug) { // do not waste cpu unless we're at debug level
		// TODO: How do we want to expose the rate limit info?
		// Maybe use a chan(RateLimit) for something to get the data?
		meta := parseResponseMetadata(res)
		if meta.Deprecated {
			c.Log.Errorf("Endpoint %s is deprecated. Use at your own risk!", req.URL.Path)
		}
		if meta.Deprecated && meta.DeprecationDate != "" {
			c.Log.Warnf("Endpoint %s will no longer be available after %s.", req.URL.Path, meta.DeprecationDate)
		}

		bodyContentType := res.Header.Get("Content-type")
		if strings.HasPrefix(bodyContentType, "application/json") {
			c.Log.Debugf("Headers: %s\nBody:\n%s\n", meta.String(), body)
		} else {
			c.Log.Debugf("Headers: %s\nBody content type: %s", meta.String(), bodyContentType)
		}
	}

	if successfulStatus(res.StatusCode) {
		if err = json.Unmarshal(body, v); err != nil {
			c.Log.Errorf("Failed to deserialize JSON:\n%s", body)
			return err
		}
		return nil
	}

	return parseResponseToError(res, body)
}

func successfulStatus(statusCode int) bool {
	return statusCode == http.StatusOK || statusCode == http.StatusCreated
}
