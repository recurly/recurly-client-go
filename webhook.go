package recurly

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"time"
)

const DEFAULT_WEBHOOK_TOLERANCE time.Duration = 5 * time.Minute

var (
	WebhookErrParsing           = errors.New("webhook signature header parsing failed")
	WebhookErrToleranceExceeded = errors.New("webhook tolerance exceeded")
	WebhookErrSignatureMismatch = errors.New("webhook signatures do not match")
)

func VerifyWebhookSignature(header, secret string, body []byte, tolerance time.Duration) error {
	wh, err := parseWebhookSignatureHeader(header)
	if err != nil {
		return err
	}

	err = wh.verifyTolerance(tolerance)
	if err != nil {
		return err
	}

	return wh.verifySignature([]byte(secret), body)
}

type webhookHeader struct {
	timestamp  string
	signatures [][]byte
}

func (w *webhookHeader) verifyTolerance(tolerance time.Duration) error {
	ts, err := strconv.ParseInt(w.timestamp, 10, 64)
	if err != nil {
		return err
	}

	if time.Since(unixMilli(ts)) > tolerance {
		return WebhookErrToleranceExceeded
	}

	return nil
}

func (w *webhookHeader) verifySignature(secret, body []byte) error {
	expected := getSignature(secret, []byte(w.timestamp), body)

	for _, received := range w.signatures {
		if hmac.Equal(expected, received) {
			return nil
		}
	}

	return WebhookErrSignatureMismatch
}

func parseWebhookSignatureHeader(header string) (*webhookHeader, error) {
	wh := new(webhookHeader)

	if header == "" {
		return wh, WebhookErrParsing
	}

	parsed := strings.Split(header, ",")
	if len(parsed) < 2 {
		return wh, WebhookErrParsing
	}

	wh.timestamp = parsed[0]
	for _, sig := range parsed[1:] {
		decoded, err := hex.DecodeString(sig)
		if err != nil {
			return wh, err
		}

		wh.signatures = append(wh.signatures, decoded)
	}

	return wh, nil
}

func getSignature(secret, timestamp, payload []byte) []byte {
	mac := hmac.New(sha256.New, secret)
	mac.Write(timestamp)
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}

// taken from go 1.17 source https://cs.opensource.google/go/go/+/refs/tags/go1.17.1:src/time/time.go;l=1330-1334
// UnixMilli returns the local Time corresponding to the given Unix time,
// msec milliseconds since January 1, 1970 UTC.
func unixMilli(msec int64) time.Time {
	return time.Unix(msec/1e3, (msec%1e3)*1e6)
}
