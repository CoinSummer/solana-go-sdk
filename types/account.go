package types

import (
	"crypto/ed25519"

	"github.com/CoinSummer/solana-go-sdk/common"
)

type Account struct {
	PublicKey  common.PublicKey
	PrivateKey ed25519.PrivateKey
}

func NewAccount() Account {
	_, X, _ := ed25519.GenerateKey(nil)
	return AccountFromPrivateKeyBytes(X)
}

func AccountFromPrivateKeyBytes(privateKey []byte) Account {
	sk := ed25519.PrivateKey(privateKey)
	return Account{
		PublicKey:  common.PublicKeyFromBytes(sk.Public().(ed25519.PublicKey)),
		PrivateKey: sk,
	}
}

type AccountMeta struct {
	PubKey     common.PublicKey
	IsSigner   bool
	IsWritable bool
}
