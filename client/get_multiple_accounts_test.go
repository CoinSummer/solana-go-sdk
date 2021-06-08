package client

import (
	"fmt"
	"testing"
)

func TestGetMultipleAccounts(t *testing.T) {
	c := NewClient(MainnetRPCEndpoint)
	accounts := []string{`vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg`, `4fYNw3dojWmQ4dXtSGE9epjRGy9pFSx62YypT7avPYvA`}
	var cfg GetMultipleAccountsConfigDataSlice
	cfg.Offset = 0
	cfg.Length = 0

	res, err := c.GetMultipleAccounts(accounts, cfg)
	fmt.Println("res:", res)
	fmt.Println("err:", err)
}
