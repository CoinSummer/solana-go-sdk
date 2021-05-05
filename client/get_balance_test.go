/**
 * Created by Goland.
 * Description:
 * User: 礼凯
 * Date: 2021/5/5 11:34 PM
 */
package client

import (
	"fmt"
	"github.com/CoinSummer/solana-go-sdk/common"
	"testing"
)

func TestGetTokenAccountBalance(t *testing.T) {
	//FcviTJpspTP4PTpq933rMNdfDdmKjfpwhUc1VmniXLcG
	c := NewClient(MainnetRPCEndpoint)
	publicKey := common.PublicKeyFromString(`ABbtLMJrWHvxoewEGxLF7NXurcBLawZhdsnSEFS1hiZa`)
	balance, err := c.GetTokenAccountBalance(publicKey.ToBase58())
	fmt.Println("balance:", balance)
	fmt.Println("err:", err)
}

func TestGetBalance(t *testing.T) {
	c := NewClient(MainnetRPCEndpoint)
	publicKey := common.PublicKeyFromString(`FcviTJpspTP4PTpq933rMNdfDdmKjfpwhUc1VmniXLcG`)
	balance, err := c.GetBalance(publicKey.ToBase58())
	fmt.Println("balance:", balance)
	fmt.Println("err:", err)
}
