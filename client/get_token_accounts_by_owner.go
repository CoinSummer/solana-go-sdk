package client

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GetTokenAccountsByOwnerMint struct {
	Mint string `json:"mint"`
}

type GetTokenAccountsByOwnerEncoding struct {
	Encoding string `json:"encoding"`
}

type GetTokenAccountsByOwnerValue struct {
	Data struct {
		Program string `json:"program"`
		Parsed  struct {
			AccountType string `json:"accountType"`
			Info        struct {
				TokenAmount struct {
					Amount         string  `json:"amount"`
					Decimals       int     `json:"decimals"`
					UiAmount       float64 `json:"uiAmount"`
					UiAmountString string  `json:"uiAmountString"`
				} `json:"tokenAmount"`
				Delegate        interface{} `json:"delegate"`
				DelegatedAmount uint64      `json:"delegatedAmount"`
				IsInitialized   bool        `json:"isInitialized"`
				IsNative        bool        `json:"IsNative"`
				Mint            string      `json:"mint"`
				Owner           string      `json:"owner"`
			}
		} `json:"parsed"`
	} `json:"data"`
	Executable bool   `json:"executable"`
	Lamports   uint64 `json:"lamports"`
	Owner      string `json:"owner"`
	RentEpoch  int    `json:"rentEpoch"`
}

// GetTokenAccountsByOwner https://docs.solana.com/developing/clients/jsonrpc-api#gettokenaccountsbyowner
// todo: no test data
func (s *Client) GetTokenAccountsByOwner(progAddr, mintAddr string) ([]GetTokenAccountsByOwnerValue, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                        `json:"context"`
			Value   []GetTokenAccountsByOwnerValue `json:"value"`
		} `json:"result"`
	}{}
	var mint GetTokenAccountsByOwnerMint
	var encoding GetTokenAccountsByOwnerEncoding
	mint.Mint = mintAddr
	encoding.Encoding = "jsonParsed"

	tt := []interface{}{progAddr, mint, encoding}
	marshal, _ := json.Marshal(tt)
	fmt.Println("sssss:", string(marshal))
	err := s.request("getTokenAccountsByOwner", []interface{}{progAddr, mint, encoding}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != (ErrorResponse{}) {
		return nil, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
