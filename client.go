package postmord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	userAgent       = "postmord/1.0"
	baseURL         = "https://api2.postnord.com/rest/shipment/v2/trackandtrace/"
	identifierURL   = baseURL + "findByIdentifier.json?apikey=%s&id=%s&locale=%s"
	notificationURL = baseURL + "findByNotificationCode.json?apikey=%s&notificationPhoneNumber=%s&notificationCode=%s&locale=%s"
	referenceURL    = baseURL + "findByReference.json?apikey=%s&customerNumber=%s&referenceValue=%s&locale=%s"
)

var DefaultOptions = &Options{
	Locale: "en",
}

type Options struct {
	Locale string
}

type Client struct {
	APIKey     string
	Opts       *Options
	httpClient *http.Client
}

func NewClient(apiKey string, httpClient *http.Client, opts *Options) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if opts == nil {
		opts = DefaultOptions
	} else if opts.Locale == "" {
		opts.Locale = "en"
	}

	return Client{
		APIKey:     apiKey,
		Opts:       opts,
		httpClient: httpClient,
	}
}

func (c *Client) FindByIdentifier(id string) (*Response, error) {
	return c.request(fmt.Sprintf(identifierURL, c.APIKey, id, c.Opts.Locale))
}

// Not implemented yet
func (c *Client) FindByNotificationCode(notificationPhoneNumber, notificationCode string) (*Response, error) {
	return c.request(fmt.Sprintf(notificationURL, c.APIKey, notificationPhoneNumber, notificationCode, c.Opts.Locale))
}

// Not implemented yet
func (c *Client) FindByReference(customerNumber, referenceValue string) (*Response, error) {
	return c.request(fmt.Sprintf(referenceURL, c.APIKey, customerNumber, referenceValue, c.Opts.Locale))
}

func (c *Client) request(url string) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("request failed: %s", res.Status)
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}

	return r, nil
}
