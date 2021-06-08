package client

import "errors"

type GetMultipleAccountsConfig struct {
	DataSlice GetMultipleAccountsConfigDataSlice `json:"dataSlice"`
}

type GetMultipleAccountsConfigDataSlice struct {
	Offset uint64 `json:"offset"`
	Length uint64 `json:"length"`
}

type GetMultipleAccountsResponse struct {
	Data       []string `json:"data"`
	Executable bool     `json:"executable"`
	Lamports   uint64   `json:"lamports"`
	Owner      string   `json:"owner"`
	RentEpoch  uint64   `json:"rentEpoch"`
}

// GetMultipleAccounts https://docs.solana.com/developing/clients/jsonrpc-api#getmultipleaccounts
// GetMultipleAccounts s
// account 要查询的公钥数组
func (s *Client) GetMultipleAccounts(account []string, cfg GetMultipleAccountsConfigDataSlice) ([]GetMultipleAccountsResponse, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context                       `json:"context"`
			Value   []GetMultipleAccountsResponse `json:"value"`
		} `json:"result"`
	}{}
	err := s.request("getMultipleAccounts", []interface{}{account, cfg}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != (ErrorResponse{}) {
		return nil, errors.New(res.Error.Message)
	}
	return res.Result.Value, nil
}
