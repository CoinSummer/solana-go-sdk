package client

import (
	"errors"
	"fmt"
)

type GetProgramAccountsConfig struct {
	Filters []interface{} `json:"filters"`
}

type MemFilters struct {
	Memcmp struct{
		Offset int    `json:"offset"`
		Bytes  string `json:"bytes"`
	} `json:"memcmp"`
}

type SizeFilters struct {
	DataSize int `json:"dataSize"`
}

type GetProgramAccountsResponse struct {
	Account struct {
		Data       string `json:"data"`
		Executable bool   `json:"executable"`
		Lamports   uint64 `json:"lamports"`
		Owner      string `json:"owner"`
		RentEpoch  uint64 `json:"rentEpoch"`
	} `json:"account"`
	Pubkey string `json:"pubkey"`
}

// GetProgramAccounts https://docs.solana.com/developing/clients/jsonrpc-api#getprogramaccounts
func (s *Client) GetProgramAccounts(base58Accounts string, filters GetProgramAccountsConfig) ([]GetProgramAccountsResponse, error) {
	res := struct {
		GeneralResponse
		Result []GetProgramAccountsResponse `json:"result"`
	}{}
	fmt.Println("111111")
	err := s.request("getProgramAccounts", []interface{}{base58Accounts, filters}, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != (ErrorResponse{}) {
		return nil, errors.New(res.Error.Message)
	}
	return res.Result, nil
}
