package recurly

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"testing"
	"time"
)

func generateTestWebhookSignatureData() (secret string, timestamp string, body []byte, signature string) {
	secret = "354fab5bd35f5a0d50845ee7e30165244468b13c2c343f104e4e730a59326d9d"
	ts := time.Now().Unix() * 1e3
	timestamp = strconv.FormatInt(ts, 10)
	body = []byte(`{"id":"rjxwmwedwqug","object_type":"account","site_id":"qc326l1hl8k9","event_type":"created","event_time":"2022-09-13T21:18:40Z","account_code":"adfas23zzz14123"}`)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "." + string(body)))
	signature = hex.EncodeToString(mac.Sum(nil))

	return
}

func TestVerifyWebhookSignature(t *testing.T) {
	secret, timestamp, body, signature := generateTestWebhookSignatureData()
	header := timestamp + "," + signature

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if err != nil {
		t.Errorf("Verify Webhook Signature failed with: %s", err.Error())
	}
}

func TestVerifyWebhookMultipleSignature(t *testing.T) {
	secret, timestamp, body, signature := generateTestWebhookSignatureData()
	header := timestamp + "," + signature

	// compute hmac again with new secret, but same "timestamp.body"
	secret = "ade46dda4474b0ae57dec22d35b9cd526b260d235f7e200c1d93f156000fe396"
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "." + string(body)))
	signature = hex.EncodeToString(mac.Sum(nil))
	header += "," + signature

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if err != nil {
		t.Errorf("Verify Webhook Signature failed with: %s", err.Error())
	}
}

func TestVerifyWebhookSignatureToleranceReached(t *testing.T) {
	secret, _, body, signature := generateTestWebhookSignatureData()
	ts := time.Now().Add(-time.Minute*10).Unix() * 1e3 // sub 10 min from timestamp sent over
	header := strconv.FormatInt(ts, 10) + "," + signature

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if !errors.Is(err, WebhookErrToleranceExceeded) {
		t.Errorf("Expected Verify Webhook Signature Tolerance Reached to be WebhookErrToleranceExceeded value, but instead got: %+v", err)
	}
}

func TestVerifyWebhookSignatureIncorrectHeader0(t *testing.T) {
	secret, _, body, _ := generateTestWebhookSignatureData()
	header := ""

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if !errors.Is(err, WebhookErrParsing) {
		t.Errorf("Expected Verify Webhook Signature Incorrect Header 0 to be WebhookErrParsing value, but instead got: %+v", err)
	}
}

func TestVerifyWebhookSignatureIncorrectHeader1(t *testing.T) {
	secret, timestamp, body, signature := generateTestWebhookSignatureData()
	header := timestamp + signature

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if !errors.Is(err, WebhookErrParsing) {
		t.Errorf("Expected Verify Webhook Signature Incorrect Header 1 to be WebhookErrParsing value, but instead got: %+v", err)
	}
}

func TestVerifyWebhookSignatureIncorrectHeader2(t *testing.T) {
	secret, timestamp, body, _ := generateTestWebhookSignatureData()
	header := timestamp

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if !errors.Is(err, WebhookErrParsing) {
		t.Errorf("Expected Verify Webhook Signature Incorrect Header 2 to be WebhookErrParsing value, but instead got: %+v", err)
	}
}

func TestVerifyWebhookSignatureIncorrectHeader3(t *testing.T) {
	secret, timestamp, body, signature := generateTestWebhookSignatureData()
	header := timestamp + "a" + "," + signature

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if err == nil {
		t.Errorf("Expected Verify Webhook Signature Incorrect Header 3 to fail with not nil err value, but instead got: nil")
	}
}

func TestVerifyWebhookSignatureMismatchBody(t *testing.T) {
	secret, timestamp, body, signature := generateTestWebhookSignatureData()
	header := timestamp + "," + signature

	err := VerifyWebhookSignature(header, secret, append(body, 'a'), DEFAULT_WEBHOOK_TOLERANCE)
	if !errors.Is(err, WebhookErrSignatureMismatch) {
		t.Errorf("Expected Verify Webhook Signature Mismatch Body to have WebhookErrSignatureMismatch value, but instead got %+v", err)
	}
}

func TestVerifyWebhookSignatureMismatchTimestamp(t *testing.T) {
	secret, timestamp, body, signature := generateTestWebhookSignatureData()
	timestamp = strconv.FormatInt(time.Now().Unix()*1e3+1, 10)
	header := timestamp + "," + signature

	err := VerifyWebhookSignature(header, secret, body, DEFAULT_WEBHOOK_TOLERANCE)
	if !errors.Is(err, WebhookErrSignatureMismatch) {
		t.Errorf("Expected Verify Webhook Signature Mismatch Body to have WebhookErrSignatureMismatch value, but instead got %+v", err)
	}
}
