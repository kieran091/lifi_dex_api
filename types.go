package lifi_dex_api

// Common structures used across multiple response types

type Token struct {
	Address  string   `json:"address"`
	ChainId  int      `json:"chainId"`
	Symbol   string   `json:"symbol"`
	Decimals int      `json:"decimals"`
	Name     string   `json:"name"`
	CoinKey  string   `json:"coinKey"`
	LogoURI  string   `json:"logoURI"`
	PriceUSD string   `json:"priceUSD,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

type ToolDetails struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	LogoURI string `json:"logoURI"`
}

type FeeCost struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Percentage  string `json:"percentage"`
	Token       Token  `json:"token"`
	Amount      string `json:"amount"`
	AmountUSD   string `json:"amountUSD"`
	Included    bool   `json:"included"`
}

type GasCost struct {
	Type      string `json:"type"`
	Price     string `json:"price,omitempty"`
	Amount    string `json:"amount"`
	Token     Token  `json:"token"`
	Estimate  string `json:"estimate,omitempty"`
	Limit     string `json:"limit,omitempty"`
	AmountUSD string `json:"amountUSD"`
}

type Action struct {
	FromChainId int     `json:"fromChainId"`
	FromAmount  string  `json:"fromAmount"`
	FromToken   Token   `json:"fromToken"`
	ToChainId   int     `json:"toChainId"`
	ToToken     Token   `json:"toToken"`
	FromAddress string  `json:"fromAddress,omitempty"`
	ToAddress   string  `json:"toAddress,omitempty"`
	Slippage    float64 `json:"slippage,omitempty"`
}

type FilteredOutError struct {
	OverallPath string `json:"overallPath"`
	Reason      string `json:"reason"`
}

type FailedError struct {
	OverallPath string               `json:"overallPath"`
	Subpaths    map[string][]Subpath `json:"subpaths"`
}

type Subpath struct {
	ErrorType string `json:"errorType"`
	Code      string `json:"code"`
	Action    Action `json:"action"`
	Tool      string `json:"tool"`
	Message   string `json:"message"`
}

type Errors struct {
	FilteredOut []FilteredOutError `json:"filteredOut"`
	Failed      []FailedError      `json:"failed"`
}

type ErrorResponse struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Errors  *Errors `json:"errors,omitempty"`
}

type StatusRequest struct {
	TxHash    string `json:"txHash"`
	Bridge    string `json:"bridge"`
	FromChain string `json:"fromChain"`
	ToChain   string `json:"toChain"`
}

type IncludedStep struct {
	Tool          string      `json:"tool"`
	ToolDetails   ToolDetails `json:"toolDetails"`
	FromAmount    string      `json:"fromAmount"`
	FromToken     Token       `json:"fromToken"`
	ToToken       Token       `json:"toToken"`
	ToAmount      string      `json:"toAmount"`
	BridgedAmount string      `json:"bridgedAmount,omitempty"`
}

type Transaction struct {
	TxHash        string        `json:"txHash"`
	TxLink        string        `json:"txLink"`
	Amount        string        `json:"amount"`
	Token         Token         `json:"token"`
	ChainId       int           `json:"chainId"`
	GasPrice      string        `json:"gasPrice"`
	GasUsed       string        `json:"gasUsed"`
	GasToken      Token         `json:"gasToken"`
	GasAmount     string        `json:"gasAmount"`
	GasAmountUSD  string        `json:"gasAmountUSD"`
	AmountUSD     string        `json:"amountUSD"`
	Value         string        `json:"value"`
	IncludedSteps []interface{} `json:"includedSteps"`
	Timestamp     int           `json:"timestamp"`
}

type SendingTransaction struct {
	TxHash        string         `json:"txHash"`
	TxLink        string         `json:"txLink"`
	Amount        string         `json:"amount"`
	Token         Token          `json:"token"`
	ChainId       int            `json:"chainId"`
	GasPrice      string         `json:"gasPrice"`
	GasUsed       string         `json:"gasUsed"`
	GasToken      Token          `json:"gasToken"`
	GasAmount     string         `json:"gasAmount"`
	GasAmountUSD  string         `json:"gasAmountUSD"`
	AmountUSD     string         `json:"amountUSD"`
	Value         string         `json:"value"`
	IncludedSteps []IncludedStep `json:"includedSteps"`
	Timestamp     int            `json:"timestamp"`
}

type StatusResponse struct {
	TransactionId    string             `json:"transactionId"`
	Sending          SendingTransaction `json:"sending"`
	Receiving        Transaction        `json:"receiving"`
	FeeCosts         []FeeCost          `json:"feeCosts"`
	LifiExplorerLink string             `json:"lifiExplorerLink"`
	FromAddress      string             `json:"fromAddress"`
	ToAddress        string             `json:"toAddress"`
	Tool             string             `json:"tool"`
	Status           string             `json:"status"`
	Substatus        string             `json:"substatus"`
	SubstatusMessage string             `json:"substatusMessage"`
	Metadata         struct {
		Integrator string `json:"integrator"`
	} `json:"metadata"`
	BridgeExplorerLink string `json:"bridgeExplorerLink"`
}

type QuoteRequest struct {
	FromChain                string   `json:"fromChain"`
	ToChain                  string   `json:"toChain"`
	FromToken                string   `json:"fromToken"`
	ToToken                  string   `json:"toToken"`
	FromAddress              string   `json:"fromAddress"`
	ToAddress                string   `json:"toAddress"`
	FromAmount               string   `json:"fromAmount"`
	Order                    string   `json:"order,omitempty"`
	Slippage                 float64  `json:"slippage,omitempty"`
	Integrator               string   `json:"integrator,omitempty"`
	Referrer                 string   `json:"referrer,omitempty"`
	Fee                      float64  `json:"fee,omitempty"`
	AllowBridges             []string `json:"allowBridges,omitempty"`
	DenyBridges              []string `json:"denyBridges,omitempty"`
	AllowExchanges           []string `json:"allowExchanges,omitempty"`
	DenyExchanges            []string `json:"denyExchanges,omitempty"`
	PreferBridges            []string `json:"preferBridges,omitempty"`
	PreferExchanges          []string `json:"preferExchanges,omitempty"`
	AllowDestinationCall     bool     `json:"allowDestinationCall,omitempty"`
	FromAmountForGas         string   `json:"fromAmountForGas,omitempty"`
	MaxPriceImpact           float64  `json:"maxPriceImpact,omitempty"`
	SwapStepTimingStrategies []string `json:"swapStepTimingStrategies,omitempty"`
	RouteTimingStrategies    []string `json:"routeTimingStrategies,omitempty"`
	SkipSimulation           bool     `json:"skipSimulation,omitempty"`
}

type Estimate struct {
	Tool              string    `json:"tool"`
	FromAmount        string    `json:"fromAmount"`
	ToAmount          string    `json:"toAmount"`
	ToAmountMin       string    `json:"toAmountMin"`
	ApprovalAddress   string    `json:"approvalAddress"`
	ExecutionDuration int       `json:"executionDuration,omitempty"`
	FromAmountUSD     string    `json:"fromAmountUSD"`
	ToAmountUSD       string    `json:"toAmountUSD"`
	FeeCosts          []FeeCost `json:"feeCosts"`
	GasCosts          []GasCost `json:"gasCosts"`
	Data              struct{}  `json:"data"`
}

type TransactionRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	ChainId  int    `json:"chainId"`
	Data     string `json:"data"`
	Value    string `json:"value"`
	GasPrice string `json:"gasPrice"`
	GasLimit string `json:"gasLimit"`
}

type QuoteIncludedStep struct {
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	Tool        string      `json:"tool"`
	ToolDetails ToolDetails `json:"toolDetails"`
	Action      struct{}    `json:"action"`
	Estimate    struct{}    `json:"estimate"`
}

type QuoteResponse struct {
	Id                 string              `json:"id"`
	Type               string              `json:"type"`
	Tool               string              `json:"tool"`
	ToolDetails        ToolDetails         `json:"toolDetails"`
	Action             Action              `json:"action"`
	Estimate           Estimate            `json:"estimate"`
	Integrator         string              `json:"integrator"`
	Referrer           string              `json:"referrer"`
	TransactionRequest TransactionRequest  `json:"transactionRequest"`
	IncludedSteps      []QuoteIncludedStep `json:"includedSteps"`
}

type AdvancedRoutesRequest struct {
	FromChainId      int          `json:"fromChainId"`
	FromAmount       string       `json:"fromAmount"`
	FromTokenAddress string       `json:"fromTokenAddress"`
	ToChainId        int          `json:"toChainId"`
	ToTokenAddress   string       `json:"toTokenAddress"`
	Options          RouteOptions `json:"options,omitempty"`
	FromAddress      string       `json:"fromAddress,omitempty"`
	ToAddress        string       `json:"toAddress,omitempty"`
	FromAmountForGas string       `json:"fromAmountForGas,omitempty"`
}

type RouteOptions struct {
	Integrator string  `json:"integrator,omitempty"`
	Referrer   string  `json:"referrer,omitempty"`
	Slippage   float64 `json:"slippage,omitempty"`
	Fee        float64 `json:"fee,omitempty"`
	Bridges    struct {
		Allow []string `json:"allow,omitempty"`
	} `json:"bridges,omitempty"`
	Exchanges struct {
		Allow []string `json:"allow,omitempty"`
	} `json:"exchanges,omitempty"`
	AllowSwitchChain bool    `json:"allowSwitchChain,omitempty"`
	Order            string  `json:"order,omitempty"`
	MaxPriceImpact   float64 `json:"maxPriceImpact,omitempty"`
}

type Protocol struct {
	Name             string `json:"name"`
	Part             int    `json:"part"`
	FromTokenAddress string `json:"fromTokenAddress"`
	ToTokenAddress   string `json:"toTokenAddress"`
}

type Bid struct {
	User                           string `json:"user"`
	Router                         string `json:"router"`
	Initiator                      string `json:"initiator"`
	SendingChainId                 int    `json:"sendingChainId"`
	SendingAssetId                 string `json:"sendingAssetId"`
	Amount                         string `json:"amount"`
	ReceivingChainId               int    `json:"receivingChainId"`
	ReceivingAssetId               string `json:"receivingAssetId"`
	AmountReceived                 string `json:"amountReceived"`
	ReceivingAddress               string `json:"receivingAddress"`
	TransactionId                  string `json:"transactionId"`
	Expiry                         int    `json:"expiry"`
	CallDataHash                   string `json:"callDataHash"`
	CallTo                         string `json:"callTo"`
	EncryptedCallData              string `json:"encryptedCallData"`
	SendingChainTxManagerAddress   string `json:"sendingChainTxManagerAddress"`
	ReceivingChainTxManagerAddress string `json:"receivingChainTxManagerAddress"`
	BidExpiry                      int    `json:"bidExpiry"`
}

type StepEstimateData struct {
	Bid                    Bid            `json:"bid,omitempty"`
	GasFeeInReceivingToken string         `json:"gasFeeInReceivingToken,omitempty"`
	TotalFee               string         `json:"totalFee,omitempty"`
	MetaTxRelayerFee       string         `json:"metaTxRelayerFee,omitempty"`
	RouterFee              string         `json:"routerFee,omitempty"`
	FromToken              Token          `json:"fromToken,omitempty"`
	ToToken                Token          `json:"toToken,omitempty"`
	ToTokenAmount          string         `json:"toTokenAmount,omitempty"`
	FromTokenAmount        string         `json:"fromTokenAmount,omitempty"`
	Protocols              [][][]Protocol `json:"protocols,omitempty"`
	EstimatedGas           int            `json:"estimatedGas,omitempty"`
}

type StepEstimate struct {
	FromAmount      string           `json:"fromAmount"`
	ToAmount        string           `json:"toAmount"`
	ToAmountMin     string           `json:"toAmountMin"`
	ApprovalAddress string           `json:"approvalAddress"`
	FeeCosts        []FeeCost        `json:"feeCosts"`
	GasCosts        []GasCost        `json:"gasCosts"`
	Data            StepEstimateData `json:"data"`
}

type RouteStep struct {
	Id         string       `json:"id"`
	Type       string       `json:"type"`
	Tool       string       `json:"tool"`
	Action     Action       `json:"action"`
	Estimate   StepEstimate `json:"estimate"`
	Integrator string       `json:"integrator"`
}

type Route struct {
	Id            string      `json:"id"`
	FromChainId   int         `json:"fromChainId"`
	FromAmountUSD string      `json:"fromAmountUSD"`
	FromAmount    string      `json:"fromAmount"`
	FromToken     Token       `json:"fromToken"`
	ToChainId     int         `json:"toChainId"`
	ToAmountUSD   string      `json:"toAmountUSD"`
	ToAmount      string      `json:"toAmount"`
	ToAmountMin   string      `json:"toAmountMin"`
	ToToken       Token       `json:"toToken"`
	GasCostUSD    string      `json:"gasCostUSD"`
	Steps         []RouteStep `json:"steps"`
}

type RouteError struct {
	ErrorType string `json:"errorType"`
	Code      string `json:"code"`
	Action    Action `json:"action"`
}

type AdvancedRoutesResponse struct {
	Routes            []Route      `json:"routes"`
	UnavailableRoutes *Errors      `json:"unavailableRoutes,omitempty"`
	Errors            []RouteError `json:"errors,omitempty"`
}

type T struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Tool        string `json:"tool"`
	ToolDetails struct {
		Key     string `json:"key"`
		Name    string `json:"name"`
		LogoURI string `json:"logoURI"`
	} `json:"toolDetails"`
	Action struct {
		FromToken struct {
			Address  string   `json:"address"`
			ChainId  int      `json:"chainId"`
			Symbol   string   `json:"symbol"`
			Decimals int      `json:"decimals"`
			Name     string   `json:"name"`
			CoinKey  string   `json:"coinKey"`
			LogoURI  string   `json:"logoURI"`
			PriceUSD string   `json:"priceUSD"`
			Tags     []string `json:"tags"`
		} `json:"fromToken"`
		FromAmount string `json:"fromAmount"`
		ToToken    struct {
			Address  string        `json:"address"`
			ChainId  int           `json:"chainId"`
			Symbol   string        `json:"symbol"`
			Decimals int           `json:"decimals"`
			Name     string        `json:"name"`
			CoinKey  string        `json:"coinKey"`
			LogoURI  string        `json:"logoURI"`
			PriceUSD string        `json:"priceUSD"`
			Tags     []interface{} `json:"tags"`
		} `json:"toToken"`
		FromChainId int     `json:"fromChainId"`
		ToChainId   int     `json:"toChainId"`
		Slippage    float64 `json:"slippage"`
		FromAddress string  `json:"fromAddress"`
		ToAddress   string  `json:"toAddress"`
	} `json:"action"`
	Estimate struct {
		Tool            string `json:"tool"`
		ApprovalAddress string `json:"approvalAddress"`
		ToAmountMin     string `json:"toAmountMin"`
		ToAmount        string `json:"toAmount"`
		FromAmount      string `json:"fromAmount"`
		FeeCosts        []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Token       struct {
				Address  string   `json:"address"`
				ChainId  int      `json:"chainId"`
				Symbol   string   `json:"symbol"`
				Decimals int      `json:"decimals"`
				Name     string   `json:"name"`
				CoinKey  string   `json:"coinKey"`
				LogoURI  string   `json:"logoURI"`
				PriceUSD string   `json:"priceUSD"`
				Tags     []string `json:"tags"`
			} `json:"token"`
			Amount     string `json:"amount"`
			AmountUSD  string `json:"amountUSD"`
			Percentage string `json:"percentage"`
			Included   bool   `json:"included"`
			FeeSplit   struct {
				IntegratorFee string `json:"integratorFee"`
				LifiFee       string `json:"lifiFee"`
			} `json:"feeSplit,omitempty"`
		} `json:"feeCosts"`
		GasCosts []struct {
			Type      string `json:"type"`
			Price     string `json:"price"`
			Estimate  string `json:"estimate"`
			Limit     string `json:"limit"`
			Amount    string `json:"amount"`
			AmountUSD string `json:"amountUSD"`
			Token     struct {
				Address  string   `json:"address"`
				ChainId  int      `json:"chainId"`
				Symbol   string   `json:"symbol"`
				Decimals int      `json:"decimals"`
				Name     string   `json:"name"`
				CoinKey  string   `json:"coinKey"`
				LogoURI  string   `json:"logoURI"`
				PriceUSD string   `json:"priceUSD"`
				Tags     []string `json:"tags"`
			} `json:"token"`
		} `json:"gasCosts"`
		ExecutionDuration int    `json:"executionDuration"`
		FromAmountUSD     string `json:"fromAmountUSD"`
		ToAmountUSD       string `json:"toAmountUSD"`
	} `json:"estimate"`
	IncludedSteps []struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Action struct {
			FromChainId int    `json:"fromChainId"`
			FromAmount  string `json:"fromAmount"`
			FromToken   struct {
				Address  string   `json:"address"`
				ChainId  int      `json:"chainId"`
				Symbol   string   `json:"symbol"`
				Decimals int      `json:"decimals"`
				Name     string   `json:"name"`
				CoinKey  string   `json:"coinKey"`
				LogoURI  string   `json:"logoURI"`
				PriceUSD string   `json:"priceUSD"`
				Tags     []string `json:"tags"`
			} `json:"fromToken"`
			ToChainId int `json:"toChainId"`
			ToToken   struct {
				Address  string   `json:"address"`
				ChainId  int      `json:"chainId"`
				Symbol   string   `json:"symbol"`
				Decimals int      `json:"decimals"`
				Name     string   `json:"name"`
				CoinKey  string   `json:"coinKey"`
				LogoURI  string   `json:"logoURI"`
				PriceUSD string   `json:"priceUSD"`
				Tags     []string `json:"tags"`
			} `json:"toToken"`
			FromAddress               string `json:"fromAddress"`
			ToAddress                 string `json:"toAddress"`
			DestinationGasConsumption string `json:"destinationGasConsumption,omitempty"`
			DestinationCallData       string `json:"destinationCallData,omitempty"`
		} `json:"action"`
		Estimate struct {
			FromAmount      string `json:"fromAmount"`
			ToAmount        string `json:"toAmount"`
			ToAmountMin     string `json:"toAmountMin"`
			Tool            string `json:"tool"`
			ApprovalAddress string `json:"approvalAddress"`
			GasCosts        []struct {
				Type      string `json:"type"`
				Price     string `json:"price"`
				Estimate  string `json:"estimate"`
				Limit     string `json:"limit"`
				Amount    string `json:"amount"`
				AmountUSD string `json:"amountUSD"`
				Token     struct {
					Address  string   `json:"address"`
					ChainId  int      `json:"chainId"`
					Symbol   string   `json:"symbol"`
					Decimals int      `json:"decimals"`
					Name     string   `json:"name"`
					CoinKey  string   `json:"coinKey"`
					LogoURI  string   `json:"logoURI"`
					PriceUSD string   `json:"priceUSD"`
					Tags     []string `json:"tags"`
				} `json:"token"`
			} `json:"gasCosts"`
			FeeCosts []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Token       struct {
					Address  string   `json:"address"`
					ChainId  int      `json:"chainId"`
					Symbol   string   `json:"symbol"`
					Decimals int      `json:"decimals"`
					Name     string   `json:"name"`
					CoinKey  string   `json:"coinKey"`
					LogoURI  string   `json:"logoURI"`
					PriceUSD string   `json:"priceUSD"`
					Tags     []string `json:"tags"`
				} `json:"token"`
				Amount     string `json:"amount"`
				AmountUSD  string `json:"amountUSD"`
				Percentage string `json:"percentage"`
				Included   bool   `json:"included"`
				FeeSplit   struct {
					IntegratorFee string `json:"integratorFee"`
					LifiFee       string `json:"lifiFee"`
				} `json:"feeSplit,omitempty"`
			} `json:"feeCosts"`
			ExecutionDuration int `json:"executionDuration"`
		} `json:"estimate"`
		Tool        string `json:"tool"`
		ToolDetails struct {
			Key     string `json:"key"`
			Name    string `json:"name"`
			LogoURI string `json:"logoURI"`
		} `json:"toolDetails"`
	} `json:"includedSteps"`
	Integrator         string `json:"integrator"`
	TransactionRequest struct {
		Value    string `json:"value"`
		To       string `json:"to"`
		Data     string `json:"data"`
		From     string `json:"from"`
		ChainId  int    `json:"chainId"`
		GasPrice string `json:"gasPrice"`
		GasLimit string `json:"gasLimit"`
	} `json:"transactionRequest"`
	TransactionId string `json:"transactionId"`
}
