package lifi_dex_api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func overloadHeaders(method string, headers map[string]string) map[string]string {
	if headers == nil {
		headers = make(map[string]string)
	}

	headers["User-Agent"] = "go-lifiapi-client/1.0.0"
	headers["Accept"] = "*/*"
	headers["Connection"] = "keep-alive"
	headers["Content-Type"] = "application/json"

	// If we set Accept-Encoding ourselves, net/http will NOT auto-decompress.
	// So we must handle Content-Encoding in readResponseBody.
	if method == http.MethodGet {
		headers["Accept-Encoding"] = "gzip"
	}
	return headers
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	if resp == nil || resp.Body == nil {
		return nil, nil
	}

	body := resp.Body
	ce := strings.TrimSpace(strings.ToLower(resp.Header.Get("Content-Encoding")))
	switch ce {
	case "", "identity":
		// no-op
	case "gzip":
		gr, err := gzip.NewReader(body)
		if err != nil {
			return nil, fmt.Errorf("failed to init gzip reader: %w", err)
		}
		defer func() { _ = gr.Close() }()
		body = gr
	default:
		return nil, fmt.Errorf("unsupported content encoding: %q", ce)
	}

	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func request[T any](c *Client, endpoint string, method string, data any) (T, error) {
	var zero T
	var bodyBytes []byte
	var bodyReader io.Reader
	var err error

	if data != nil {
		switch v := data.(type) {
		case string:
			bodyBytes = []byte(v)
			bodyReader = bytes.NewReader(bodyBytes)
		default:
			bodyBytes, err = json.Marshal(data)
			if err != nil {
				return zero, fmt.Errorf("failed to marshal request body: %w", err)
			}
			bodyReader = bytes.NewReader(bodyBytes)
		}
	}

	baseHeaders := c.getBaseHeaders()
	baseHeaders = overloadHeaders(method, baseHeaders)

	req, err := http.NewRequest(method, endpoint, bodyReader)
	if err != nil {
		return zero, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range baseHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return zero, fmt.Errorf("request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := readResponseBody(resp)
	if err != nil {
		return zero, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		_ = json.Unmarshal(respBody, &errResp)

		if errResp.Errors == nil {
			return zero, &APIError{
				StatusCode: resp.StatusCode,
				BizCode:    errResp.Code,
				Method:     method,
				URL:        endpoint,
				Payload: &APIErrorPayload{
					Type:    OtherError,
					Message: errResp.Message,
				},
			}
		}
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			BizCode:    errResp.Code,
			Method:     method,
			URL:        endpoint,
			Payload:    parseErrorResponse(errResp.Errors),
		}
		return zero, apiErr
	}

	var baseResp T
	err = json.Unmarshal(respBody, &baseResp)
	if err != nil {
		return zero, fmt.Errorf("failed to decode response json: %w; body=%s", err, string(respBody))
	}

	return baseResp, nil
}

func post[T any](c *Client, endpoint string, data any) (T, error) {
	return request[T](c, endpoint, http.MethodPost, data)
}

func get[T any](c *Client, endpoint string) (T, error) {
	return request[T](c, endpoint, http.MethodGet, nil)
}
