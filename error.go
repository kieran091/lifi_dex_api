package lifi_dex_api

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

//const (
//	DefaultError = 1000 + iota
//	FailedToBuildTransactionError
//	NoQuoteError
//	NotFoundError
//	NotProcessableError
//	RateLimitError
//	ServerError
//	SlippageError
//	ThirdPartyError
//	TimeoutError
//	UnauthorizedError
//	ValidationError
//	RpcFailure
//	MalformedSchema
//)
//
//var ToolErrorMessage = map[string]string{
//	"NO_POSSIBLE_ROUTE":                 "no route was found for this action",
//	"INSUFFICIENT_LIQUIDITY":            "not enough liquidity for this action",
//	"TOOL_TIMEOUT":                      "the tool took too long to respond",
//	"UNKNOWN_ERROR":                     "an unknown error occurred in the tool",
//	"RPC_ERROR":                         "the chain rpc call failed",
//	"AMOUNT_TOO_LOW":                    "the specified amount is too low",
//	"AMOUNT_TOO_HIGH":                   "the specified amount is too high",
//	"FEES_HGHER_THAN_AMOUNT":            "the fees are higher than the amount",
//	"DIFFERENT_RECIPIENT_NOT_SUPPORTED": "different recipient is not supported",
//	"TOOL_SPECIFIC_ERROR":               "a tool specific error occurred",
//	"CANNOT_GUARANTEE_MIN_AMOUNT":       "cannot guarantee minimum amount due to market volatility",
//}

const (
	AmountTooLow       = "AMOUNT_TOO_LOW"
	AmountTooHigh      = "AMOUNT_TOO_HIGH"
	PriceImpactTooHigh = "PRICE_IMPACT_TOO_HIGH"
	OtherError         = "OTHER_ERROR"
	UnknownError       = "UNKNOWN_ERROR"
)

const (
	USD         = "USD"
	TokenAmount = "TOKEN_AMOUNT"
	PriceImpact = "PRICE_IMPACT"
)

const multiSignatureErrorMessage = "Bridge glacis does not support destination call and the request options do not allow multiple signatures"

type APIErrorPayload struct {
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	Amount  string `json:"amount"`
	Message string `json:"message"`
}

// APIError represents an error returned by the OKX API.
type APIError struct {
	// StatusCode is the HTTP status code returned by the API.
	StatusCode int
	// BizCode is the business code returned by the API.
	BizCode int
	// Method is the HTTP method used for the request.
	Method string
	// URL is the URL of the request.
	URL string
	// Body is the response body returned by the API.
	Payload *APIErrorPayload
}

// Error returns a string representation of the error.
func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("api error: %s %s: status=%d, biz_code=%d, message: %s", e.Method, e.URL, e.StatusCode, e.BizCode, e.Payload.Message)
}

// Unwrap returns the error itself.
// implementation of the Wrapper interface.
func (e *APIError) Unwrap() error {
	return e
}

// Cause returns the error itself.
// implementation of the Causer interface.
func (e *APIError) Cause() error {
	return e
}

func parseErrorResponse(errs *Errors) *APIErrorPayload {
	if errs == nil || len(errs.FilteredOut) == 0 {
		return &APIErrorPayload{
			Type: UnknownError,
		}
	}
	filteredOutReason := errs.FilteredOut[0].Reason

	if filteredOutReason == multiSignatureErrorMessage {
		return &APIErrorPayload{
			Type:    OtherError,
			Message: filteredOutReason,
		}
	}

	// check USD amount errors
	if strings.Contains(filteredOutReason, "USD") {
		re := regexp.MustCompile(`\b\d+\b`)
		matches := re.FindAllString(filteredOutReason, -1)
		if len(matches) > 0 {
			return &APIErrorPayload{
				Type:    AmountTooLow,
				Subtype: USD,
				Amount:  matches[len(matches)-1],
				Message: filteredOutReason,
			}
		}
	}

	// check Token amount errors
	// format: Transferred amount (X) out of acceptable range (min: Y, max: Z)
	re := regexp.MustCompile(`Transferred amount \((\d+)\) out of acceptable range \(min: (\d+|Infinity), max: (\d+|Infinity)\)`)
	matches := re.FindStringSubmatch(filteredOutReason)
	if len(matches) == 4 {
		transferredAmount := new(big.Int)
		_, ok := transferredAmount.SetString(matches[1], 10)
		if ok {
			var minAmount, maxAmount *big.Int

			// 解析最小值
			if strings.ToLower(matches[2]) != "infinity" {
				minAmount = new(big.Int)
				_, _ = minAmount.SetString(matches[2], 10)
			}

			// 解析最大值
			if strings.ToLower(matches[3]) != "infinity" {
				maxAmount = new(big.Int)
				_, _ = maxAmount.SetString(matches[3], 10)
			}

			amount := big.NewInt(0)
			errorType := UnknownError
			if minAmount != nil && transferredAmount.Cmp(minAmount) < 0 {
				errorType = AmountTooLow
				amount = minAmount
			} else if maxAmount != nil && transferredAmount.Cmp(maxAmount) > 0 {
				errorType = AmountTooHigh
				amount = maxAmount
			}

			return &APIErrorPayload{
				Type:    errorType,
				Subtype: TokenAmount,
				Amount:  amount.String(),
				Message: filteredOutReason,
			}
		}
	}

	// check Price impact errors
	// format: Price impact of (X)% is higher than the max allowed (Y)%
	re = regexp.MustCompile(`Price impact of (\d+(\.\d+)?)% is higher than the max allowed (\d+(\.\d+)?)%`)
	matches = re.FindStringSubmatch(filteredOutReason)
	if len(matches) > 0 {
		return &APIErrorPayload{
			Type:    PriceImpactTooHigh,
			Subtype: PriceImpact,
			Amount:  matches[1] + "%",
			Message: filteredOutReason,
		}
	}

	if strings.Contains(filteredOutReason, "insufficient") {
		return &APIErrorPayload{
			Type:    AmountTooLow,
			Subtype: USD,
			Amount:  "1",
			Message: filteredOutReason,
		}
	}

	return &APIErrorPayload{
		Type:    UnknownError,
		Message: filteredOutReason,
	}
}
