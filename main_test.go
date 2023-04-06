package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type transportFunc func(*http.Request) (*http.Response, error)

func (f transportFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// TestFetchRate tests the fetchRate function.
func TestFetchRate(t *testing.T) {
	currency := "USD"
	rateResponse, err := fetchRate(currency)

	assert.NoError(t, err)
	assert.Equal(t, currency, rateResponse.Currency)
	assert.NotZero(t, rateResponse.Price)
}

// TestRateRocketHandler tests the rateRocketHandler function.
func TestRateRocketHandler(t *testing.T) {
	// Create a test HTTP server that returns a dummy RateResponse.
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rateResponse := &RateResponse{
			Currency:  "USD",
			Price:     12345.67,
			Timestamp: "2023-04-05T09:12:54.884Z",
			Date:      "4/5/2023",
			Time:      "9:12:54 AM",
		}
		json.NewEncoder(w).Encode(rateResponse)
	}))
	defer testServer.Close()

	// Replace the http.DefaultTransport to intercept and modify the request URL.
	oldTransport := http.DefaultTransport
	http.DefaultTransport = transportFunc(func(req *http.Request) (*http.Response, error) {
		parsedURL, _ := url.Parse(testServer.URL)
		req.URL = parsedURL
		return oldTransport.RoundTrip(req)
	})
	defer func() { http.DefaultTransport = oldTransport }()

	// Create a test HTTP request to the rateRocketHandler.
	form := url.Values{}
	form.Add("token", "test-token")
	form.Add("text", "USD")
	req := httptest.NewRequest("POST", "/raterocket", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Create a test HTTP response recorder.
	rr := httptest.NewRecorder()

	// Set the SLACK_VERIFICATION_TOKEN environment variable.
	oldToken := os.Getenv("SLACK_VERIFICATION_TOKEN")
	os.Setenv("SLACK_VERIFICATION_TOKEN", "test-token")
	defer func() { os.Setenv("SLACK_VERIFICATION_TOKEN", oldToken) }()

	// Call the rateRocketHandler function.
	rateRocketHandler(rr, req)

	// Check the response status code and body.
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "1 BTC to USD rate: 12345.67 (4/5/2023 9:12:54 AM)")
}