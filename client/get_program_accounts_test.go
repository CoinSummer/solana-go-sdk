package client

import (
	"fmt"
	"testing"
)

//CGEVtNYVAdyoXHFj9vAiXnKxcJcUubWB7NL3VcTCoQki
//5Q544fKrFoe6tsEbD7S8EmxGTJYAKtTVhAW5Q5pge4j1

//F1BuWYNqGN5x3ZtpFD1EPyhfMFTNcRCZjeNejb7qYfMM
//EhhTKczWMGQt46ynNeRX1WfeagwwJd7ufHvCDjRxjo5Q
func TestGetProgramAccounts(t *testing.T) {
	c := NewClient(MainnetRPCEndpoint)
	accounts := `EhhTKczWMGQt46ynNeRX1WfeagwwJd7ufHvCDjRxjo5Q`
	var filter GetProgramAccountsConfig
	var memFilters MemFilters
	var size SizeFilters

	memFilters.Memcmp.Offset = 40
	memFilters.Memcmp.Bytes = `F1BuWYNqGN5x3ZtpFD1EPyhfMFTNcRCZjeNejb7qYfMM`

	size.DataSize = 88

	filter.Filters = append(filter.Filters, memFilters)
	filter.Filters = append(filter.Filters, size)

	res, err := c.GetProgramAccounts(accounts, filter)
	fmt.Println("res:", res)
	fmt.Println("err:", err)
}
