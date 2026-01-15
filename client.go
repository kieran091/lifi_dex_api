package lifi_dex_api

import (
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	APIKey   string
	BaseURL  string
	Timeout  time.Duration
	ProxyURL string
}

type Client struct {
	config Config
	client *http.Client
}

func NewClient(c Config) (*Client, error) {
	if c.Timeout <= 0 {
		c.Timeout = 30 * time.Second
	}

	client := &http.Client{
		Timeout: c.Timeout,
	}
	if c.ProxyURL != "" {
		proxyURL, err := url.Parse(c.ProxyURL)
		if err != nil {
			return nil, err
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}
	return &Client{
		config: c,
		client: client,
	}, nil
}

func (c *Client) getBaseHeaders() map[string]string {
	return map[string]string{
		"x-lifi-api-key": c.config.APIKey,
	}
}

func (c *Client) Status(req *StatusRequest) (*StatusResponse, error) {
	endpoint := parseEndpoint(c.config.BaseURL, EndpointStatus, req.ToMap())
	return get[*StatusResponse](c, endpoint)
}

func (c *Client) Quote(req *QuoteRequest) (*QuoteResponse, error) {
	endpoint := parseEndpoint(c.config.BaseURL, EndpointQuote, req.ToMap())
	return get[*QuoteResponse](c, endpoint)
}

func (c *Client) AdvancedRoutes(req *AdvancedRoutesRequest) (*AdvancedRoutesResponse, error) {
	endpoint := parseEndpoint(c.config.BaseURL, EndpointAdvancedRoutes, nil)
	advancedRoutesResponse, err := post[*AdvancedRoutesResponse](c, endpoint, req)
	if err != nil {
		return nil, err
	}
	if len(advancedRoutesResponse.Routes) == 0 && advancedRoutesResponse.UnavailableRoutes != nil {
		return nil, &APIError{
			StatusCode: http.StatusOK,
			Method:     http.MethodPost,
			URL:        endpoint,
			Payload:    parseErrorResponse(advancedRoutesResponse.UnavailableRoutes),
		}
	}
	return advancedRoutesResponse, nil
}
