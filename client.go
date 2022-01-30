package recurly

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	defaultTimeout = 60 * time.Second
)

var (
	// APIHost is the base URL for Recurly API v3
	APIHost string

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

	acceptVersion = fmt.Sprintf("application/vnd.recurly.%s", APIVersion)
	userAgent     = fmt.Sprintf("Recurly/%s; go %s", clientVersion, runtime.Version())
	pathPattern   = regexp.MustCompile(`{[^}]+}`)
)

// Client submits API requests to Recurly
type Client struct {
	apiKey  string
	baseURL string

	Log        *Logger
	HTTPClient *http.Client
}

var (
	// apiHosts maps the Region to the base URL
	apiHosts = map[region]string{
		US: "https://v3.recurly.com",
		EU: "https://v3.eu.recurly.com",
	}
)

// NewClient returns a new API Client using the given APIKey
func NewClient(apiKey string) *Client {
	return NewClientWithOptions(apiKey, ClientOptions{
		Region: US,
	})
}

// NewClientWithOptions returns a new API Client using the given APIKey and options
func NewClientWithOptions(apiKey string, options ClientOptions) *Client {
	apiHost := APIHost
	if apiHost == "" {
		apiHost = apiHosts[options.Region]
	}
	return &Client{
		apiKey:     apiKey,
		baseURL:    apiHost,
		Log:        NewLogger(LevelWarn),
		HTTPClient: defaultClient,
	}
}

func validatePathParameters(params []string) error {
	invalidParams := []string{}
	for _, param := range params {
		if len(strings.TrimSpace(param)) == 0 {
			invalidParams = append(invalidParams, param)
		}
	}
	if len(invalidParams) > 0 {
		return &Error{
			Message: "Operation parameters cannot be empty strings.",
			Class:   ErrorClassClient,
			Type:    ErrorTypeBadRequest,
		}
	}
	return nil
}

// InterpolatePath takes an OpenAPI-style path such as "/accounts/{account_id}/shipping_addresses/{shipping_address_id}"
// and a list of string arguments to fill the template, and it returns the interpolated path
func (c *Client) InterpolatePath(path string, params ...string) (string, error) {
	err := validatePathParameters(params)
	if err != nil {
		return "", err
	}
	template := pathPattern.ReplaceAllString(path, "%s")
	encodedParams := make([]interface{}, len(params))
	for i, param := range params {
		encoded := url.PathEscape(param)
		encodedParams[i] = encoded
	}
	return fmt.Sprintf(template, encodedParams...), nil
}

// HTTPCaller is the generic http interface used by the Client
type HTTPCaller interface {
	Call(ctx context.Context, method string, path string, body interface{}, queryParams QueryParams, requestOptions OptionsApplier, v interface{}) error
}

// Call sends a request to Recurly API and parses the JSON response for the expected response type.
func (c *Client) Call(ctx context.Context, method string, path string, body interface{}, queryParams QueryParams, requestOptions OptionsApplier, v interface{}) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("%s/%s", c.baseURL, path)
	} else {
		path = fmt.Sprintf("%s%s", c.baseURL, path)
	}

	path = BuildURL(path, queryParams)

	req, err := c.NewHTTPRequest(ctx, method, path, body, requestOptions)
	if err != nil {
		return err
	}

	return c.Do(req, v)
}

// BuildURL will append query string parameters
func BuildURL(requestURL string, queryParams QueryParams) string {
	if queryParams != nil {
		if keyValues := queryParams.URLParams(); len(keyValues) > 0 {
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

// NewHTTPRequest generates an http.Request for the client to submit to Recurly API.
func (c *Client) NewHTTPRequest(ctx context.Context, method string, requestURL string, body interface{}, requestOptions OptionsApplier) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, requestURL, nil)
	if err != nil {
		return nil, err
	}

	req = requestOptions.applyOptions(req)

	// Set required headers after custom headers so that they aren't accidentally overwritten.
	req.Header.Set("Accept", acceptVersion)
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", userAgent)
	req.SetBasicAuth(c.apiKey, "")

	if body != nil && method != http.MethodGet {
		data, err := json.Marshal(body)
		if err != nil {
			c.Log.Errorf("Failed to marshal data to JSON payload: %v", err)
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		if c.Log.IsLevel(LevelDebug) {
			fmt.Println("Request Body: ", req.Body)
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
	// TODO this could be cleaned up to better handle
	// when the response body is empty
	bodyReader := res.Body
	contentEncoding := strings.ToLower(res.Header.Get("Content-Encoding"))
	if contentEncoding == "gzip" {
		gzipReader, err := gzip.NewReader(bodyReader)
		// If the content is encoded, we cannot tell the difference b/w
		// an empty response and one with content. The only signal we have
		// is if we receive an EOF error from `NewReader`. If we get this
		// error, we want to fall through this block and not wrap the bodyReader
		// in the gzipReader. For any other error, we want to return early.
		if err != nil {
			if err != io.EOF {
				return err
			}
		} else {
			bodyReader = gzipReader
		}
	}

	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		c.Log.Errorf("Cannot read response: %s", err)
		return err
	}

	meta := parseResponseMetadata(res)
	v.(Resource).setResponse(meta)

	if c.Log.IsLevel(LevelDebug) {
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
		if len(body) > 0 {
			if err = json.Unmarshal(body, v); err != nil {
				c.Log.Errorf("Failed to deserialize JSON:\n%s", body)
				return err
			}
		}
		return nil
	}

	return parseResponseToError(res, body)
}

func successfulStatus(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusBadRequest
}
