package common

import (
	"crypto/sha256"
	"errors"

	"github.com/mr-tron/base58"
	"github.com/teserakt-io/golang-ed25519/edwards25519"
)

const (
	PublicKeyLength = 32
	MaxSeedLength   = 32
	MaxSeed         = 16
)

type PublicKey [PublicKeyLength]byte

func PublicKeyFromString(s string) PublicKey {
	d, _ := base58.Decode(s)
	return PublicKeyFromBytes(d)
}

func PublicKeyFromBytes(b []byte) PublicKey {
	var pubkey PublicKey
	if len(b) > PublicKeyLength {
		b = b[:PublicKeyLength]
	}
	copy(pubkey[PublicKeyLength-len(b):], b)
	return pubkey
}

func CreateProgramAddress(seeds [][]byte, programId PublicKey) (PublicKey, error) {
	if len(seeds) > MaxSeed {
		return PublicKey{}, errors.New("Length of the seed is too long for address generation")
	}

	buf := []byte{}
	for _, seed := range seeds {
		if len(seed) > MaxSeedLength {
			return PublicKey{}, errors.New("Length of the seed is too long for address generation")
		}
		buf = append(buf, seed...)
	}
	buf = append(buf, programId[:]...)
	buf = append(buf, []byte("ProgramDerivedAddress")...)
	h := sha256.Sum256(buf)
	pubKey := PublicKeyFromBytes(h[:])

	// public key is on curve
	var A edwards25519.ExtendedGroupElement
	if A.FromBytes((*[32]byte)(&pubKey)) {
		return PublicKey{}, errors.New("Invalid seeds, address must fall off the curve")
	}
	return pubKey, nil
}

func (p PublicKey) ToBase58() string {
	return base58.Encode(p[:])
}

func (p PublicKey) Bytes() []byte {
	return p[:]
}

func CreateWithSeed(from PublicKey, seed string, programID PublicKey) PublicKey {
	b := make([]byte, 0, 64+len(seed))
	b = append(b, from[:]...)
	b = append(b, seed[:]...)
	b = append(b, programID[:]...)
	hash := sha256.Sum256(b)
	return PublicKeyFromBytes(hash[:])
}

func FindAssociatedTokenAddress(walletAddress, tokenMintAddress PublicKey) (PublicKey, int, error) {
	seeds := [][]byte{}
	seeds = append(seeds, walletAddress.Bytes())
	seeds = append(seeds, TokenProgramID.Bytes())
	seeds = append(seeds, tokenMintAddress.Bytes())

	return FindProgramAddress(seeds, SPLAssociatedTokenAccountProgramID)
}

func FindProgramAddress(seed [][]byte, programID PublicKey) (PublicKey, int, error) {
	var pubKey PublicKey
	var err error
	nonce := 0xff
	for nonce != 0x0 {
		pubKey, err = CreateProgramAddress(append(seed, []byte{byte(nonce)}), programID)
		if err == nil {
			return pubKey, nonce, nil
		}
		nonce--
	}
	return PublicKey{}, nonce, errors.New("unable to find a viable program address")
}
