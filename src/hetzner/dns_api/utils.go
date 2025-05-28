package dns_api

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"
)

type HetznerTime struct {
	time.Time
}

func (ht *HetznerTime) UnmarshalJSON(b []byte) error {
	s := string(b)

	if len(s) > 1 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}

	if s == "" {
		ht.Time = time.Time{}
		return nil
	}

	layouts := []string{
		"2006-01-02 15:04:05.000 -0700 MST",
		"2006-01-02 15:04:05.00 -0700 MST",
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02T15:04:05.000Z",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			ht.Time = t
			return nil
		}
	}

	return fmt.Errorf("failed to parse time string %s", s)
}

// Creates a http client to use for requests.
func CreateClient() *http.Client {
	return &http.Client{
		Timeout: 20 * time.Second,
	}
}

// Create a new request to the hetzner dns api endpoint.
func NewHetznerRequest(method string, url string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(url, "/") {
		url = fmt.Sprintf("/%s", url)
	}

	return http.NewRequest(method, fmt.Sprintf("https://dns.hetzner.com/api/v1%s", url), body)
}

// Sends a http request via the client and returns the response.
func SendRequest(client *http.Client, request *http.Request, accepted_status_codes []int) ([]byte, error) {
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if !slices.Contains(accepted_status_codes, resp.StatusCode) {
		return nil, fmt.Errorf("API request failed with response status of %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func AddContentTypeHeader(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
}

func AddAcceptHeader(request *http.Request) {
	request.Header.Add("Accept", "application/json")
}

func AddTokenHeader(request *http.Request, token string) {
	request.Header.Add("Auth-API-Token", token)
}
