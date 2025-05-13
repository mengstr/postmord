package postmord

import (
	"net/http"
)

const (
	userAgent     = "postmord/1.0.0"
	DomainDefault = "https://api2.postnord.com"
	endpoint      = "/rest/shipment/v5/trackandtrace/"
)

var (
	DefaultOptions = &Options{Locale: LocaleDefault, Domain: DomainDefault}
	baseURL        string
)

// Options is used to set the options used to retrieve information about shipments.
type Options struct {
	Locale Locale
	Domain string
}

// Client is used to retrieve information about shipments.
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
	} else {
		if opts.Locale == "" {
			opts.Locale = LocaleDefault
		}
		if opts.Domain == "" {
			opts.Domain = DomainDefault
		}
	}

	baseURL = opts.Domain + endpoint

	return Client{APIKey: apiKey, Opts: opts, httpClient: httpClient}
}
