package client

import (
	"fmt"
	"testing"
)

func TestGetTokenAccountsByOwner(t *testing.T) {
	c := NewClient(MainnetRPCEndpoint)

	mintAddr := `TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA`
	progAddr := `HzWpBN6ucpsA9wcfmhLAFYqEUmHjE9n2cGHwunG5avpL`

	owner, err := c.GetTokenAccountsByOwner(progAddr, mintAddr)
	fmt.Println("owner:", owner)
	fmt.Println("err:", err)
}

