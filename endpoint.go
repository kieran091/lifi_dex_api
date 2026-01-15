package lifi_dex_api

import (
	"fmt"
	"net/url"
	"strings"
)

type Endpoint string

var (
	// EndpointStatus represents the API endpoint for transaction status checks.
	EndpointStatus Endpoint = "/v1/status"

	// EndpointQuote represents the API endpoint for obtaining quotes.
	EndpointQuote Endpoint = "/v1/quote"

	// EndpointAdvancedRoutes represents the API endpoint for advanced routing options.
	EndpointAdvancedRoutes Endpoint = "/v1/advanced/routes"
)

func parseEndpoint(host string, endpoint Endpoint, params map[string]any) string {
	sb := strings.Builder{}
	host = strings.TrimSuffix(host, "/")
	sb.WriteString(host)
	if !strings.HasPrefix(string(endpoint), "/") {
		sb.WriteString("/")
	}
	sb.WriteString(string(endpoint))

	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, fmt.Sprintf("%v", v))
		}

		sb.WriteString("?")
		sb.WriteString(values.Encode())
	}

	return sb.String()
}
