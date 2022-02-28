package nomics

import (
	"net/http"
)

const (
	APIBase = "https://api.nomics.com/v1"

	CurrenciesTickerRoute   = "currencies/ticker"
	CurrenciesMetadataRoute = "currencies"

	MarketsRoute          = "markets"
	MarketCapHistoryRoute = "market-cap/history"

	GlobalVolumeHistoryRoute = "volume/history"

	FiatExchangeRatesRoute        = "exchange-rates"
	FiatExchangeRatesHistoryRoute = "exchange-rates/history"
)

// Noimcs defines the supported subset of the nomics API.
type Nomics interface {
	APIBase() string
	APIKey() string
	Client() *http.Client
	SetClient(client *http.Client)
}

// NomicsBase bundles data needed by a large number of methods in order to interact with the nomics API.
type NomicsBase struct {
	apiBase string
	apiKey  string
	client  *http.Client
}

// NewNomics  creates a new client instance.
func NewNomics(apiKey string) *NomicsBase {
	return &NomicsBase{
		apiBase: APIBase,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}
}

// APIBase returns the API Base URL configured for this client.
func (nm *NomicsBase) APIBase() string {
	return nm.apiBase
}

// ApiKey returns the API key configured for this client.
func (nm *NomicsBase) APIKey() string {
	return nm.apiKey
}

// Client returns the HTTP client configured for this client.
func (nm *NomicsBase) Client() *http.Client {
	return nm.client
}

// SetClient updates the HTTP client for this client.
func (nm *NomicsBase) SetClient(c *http.Client) {
	nm.client = c
}
