package lifi_dex_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

func createClient() *Client {
	c := Config{
		APIKey:   "31371ea0-d8d5-4393-bbc7-2d6cb604dc44.cbc48dd8-cd4a-45d3-87e5-de70fe400f91",
		BaseURL:  "https://li.quest",
		Timeout:  30 * time.Second,
		ProxyURL: "http://127.0.0.1:7897",
	}

	client, err := NewClient(c)
	if err != nil {
		panic(err)
	}

	return client
}

func TestClient_Status(t *testing.T) {
	client := createClient()

	statusRequest := &StatusRequest{
		TxHash:    "0x2a95e352469b20c63b14885996badfe2984c7b5d4484c3e5e8fe38e5bffc7af6",
		Bridge:    "",
		FromChain: "",
		ToChain:   "",
	}
	statusResponse, err := client.Status(statusRequest)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Status Response: %+v", statusResponse)
}

func TestClient_AdvancedRoutes(t *testing.T) {
	client := createClient()

	advancedRoutesRequest := &AdvancedRoutesRequest{
		FromChainId:      1,
		FromAmount:       "100000000000000000000000000000000000000000000000000",
		FromTokenAddress: "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		ToChainId:        56,
		ToTokenAddress:   "0x43c934a845205f0b514417d757d7235b8f53f1b9",
	}
	advancedRoutesResponse, err := client.AdvancedRoutes(advancedRoutesRequest)
	if err != nil {
		var apiErr *APIError
		if errors.As(err, &apiErr); apiErr != nil {
			t.Logf("API Error: %+v", apiErr)
			t.Logf("API Error Payload: %+v", apiErr.Payload)
			return
		}
		t.Fatal(err)
	}
	t.Logf("Advanced Routes Response: %+v", advancedRoutesResponse)
}

func TestClient_Quote(t *testing.T) {
	client := createClient()

	quoteRequest := &QuoteRequest{
		FromChain:   "1",
		ToChain:     "56",
		FromToken:   "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		ToToken:     "0x43c934a845205f0b514417d757d7235b8f53f1b9",
		FromAddress: "0xb6789347ff8231f06f3101461d97b22f238a3019",
		ToAddress:   "0xb6789347ff8231f06f3101461d97b22f238a3019",
		FromAmount:  "1000",
		Order:       "FASTEST",
		Slippage:    0.005,
	}
	quoteResponse, err := client.Quote(quoteRequest)
	if err != nil {
		var apiErr *APIError
		if errors.As(err, &apiErr); apiErr != nil {
			t.Logf("API Error: %+v", apiErr)
			t.Logf("API Error Payload: %+v", apiErr.Payload)
			return
		}
		t.Fatal(err)
	}

	marshal, _ := json.Marshal(quoteResponse)
	fmt.Println(string(marshal))

	t.Logf("Quote Response: %+v", quoteResponse)
}
