package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
)

type WTransaction struct {
	SenderPrivetKey       *ecdsa.PrivateKey
	SenderPublicKey       *ecdsa.PublicKey
	SenderBlockchainAddr  string
	ReciverBlockchainAddr string
	Value                 float32
}

func NewTransaction(senderPrivetKey *ecdsa.PrivateKey, senderPublicKey *ecdsa.PublicKey, sender string, reciver string, value float32) *WTransaction {
	return &WTransaction{
		SenderPrivetKey:       senderPrivetKey,
		SenderPublicKey:       senderPublicKey,
		SenderBlockchainAddr:  sender,
		ReciverBlockchainAddr: reciver,
		Value:                 value,
	}
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func (t *WTransaction) GenerateSignature() *Signature {

	m, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	h := sha256.Sum256([]byte(m))
	r, s, err := ecdsa.Sign(rand.Reader, t.SenderPrivetKey, h[:])
	if err != nil {
		panic(err)
	}
	return &Signature{r, s}
}

func (s *Signature) SignatureStr() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
