package types

import "github.com/CoinSummer/solana-go-sdk/common"

type CompiledInstruction struct {
	ProgramIDIndex int
	Accounts       []int
	Data           []byte
}

type Instruction struct {
	ProgramID common.PublicKey
	Accounts  []AccountMeta
	Data      []byte
}
