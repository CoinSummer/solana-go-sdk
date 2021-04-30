package client

import "errors"

// sol 链的 sol 余额
func (s *Client) GetBalance(base58Addr string) (uint64, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context `json:"context"`
			Value   uint64  `json:"value"`
		} `json:"result"`
	}{}
	err := s.request("getBalance", []interface{}{base58Addr}, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}

type tokenAccountBalance struct {
	JsonRPC string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			Slot uint64 `json:"slot"`
		} `json:"context"`
		Value struct {
			Amount         string  `json:"amount"`
			Decimals       uint8   `json:"decimals"`
			UiAmount       float64 `json:"uiAmount"`
			UiAmountString string  `json:"uiAmountString"`
		} `json:"value"`
	} `json:"result"`
	ID uint64 `json:"id"`
	Error   ErrorResponse `json:"error"`
}

// solana链的非 sol 代币的余额
func (s *Client) GetTokenAccountBalance(base58Addr string) (float64, error) {
	var res tokenAccountBalance
	err := s.request("getTokenAccountBalance", []interface{}{base58Addr}, &res)
	if err != nil {
		return 0, err
	}

	if res.Error != (ErrorResponse{}) {
		return 0, errors.New(res.Error.Message)
	}
	return res.Result.Value.UiAmount, nil
}
