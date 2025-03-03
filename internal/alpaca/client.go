package alpaca

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

// Client wraps the Alpaca client with our custom functionality
type Client struct {
	Trading    *alpaca.Client
	MarketData *marketdata.Client
}

// NewClient creates a new Alpaca client
func NewClient(apiKey, apiSecret string, paper bool) (*Client, error) {
	// Set up the base client
	baseURL := alpaca.ApiURL
	if paper {
		baseURL = alpaca.PaperApiURL
	}

	// Initialize the trading client
	tradingClient := alpaca.NewClient(alpaca.ClientOpts{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		BaseURL:   baseURL,
	})

	// Initialize the market data client
	dataClient := marketdata.NewClient(marketdata.ClientOpts{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	})

	return &Client{
		Trading:    tradingClient,
		MarketData: dataClient,
	}, nil
}

// GetAccount retrieves account information
func (c *Client) GetAccount() (*alpaca.Account, error) {
	return c.Trading.GetAccount()
}
