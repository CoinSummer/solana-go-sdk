package client

import "errors"

// GetBalance sol 链的 sol 余额
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

type tokenAccountBalanceResp struct {
	GeneralResponse
	Result struct {
		Context struct {
			Slot uint64 `json:"slot"`
		} `json:"context"`
		Value BalanceValue `json:"value"`
	} `json:"result"`
}

type BalanceValue struct {
	Amount         string  `json:"amount"`
	Decimals       uint8   `json:"decimals"`
	UiAmount       float64 `json:"uiAmount"`
	UiAmountString string  `json:"uiAmountString"`
}

// GetTokenAccountBalance solana链的非 sol 代币的余额
// 该地址为主账户地址对应的 token 的账户地址，需要做区分
func (s *Client) GetTokenAccountBalance(base58Addr string) (*BalanceValue, error) {
	var res tokenAccountBalanceResp
	err := s.request("getTokenAccountBalance", []interface{}{base58Addr}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != (ErrorResponse{}) {
		return nil, errors.New(res.Error.Message)
	}
	return &res.Result.Value, nil
}
