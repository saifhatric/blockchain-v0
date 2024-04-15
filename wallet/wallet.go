package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Wallet struct {
	PublicKey      *ecdsa.PublicKey
	PrivetKey      *ecdsa.PrivateKey
	BlockchainAddr string
}

func NewWallet() *Wallet {
	w := &Wallet{}
	//Generate public key and privet key
	privetKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		panic(err)
	}
	w.PrivetKey = privetKey
	w.PublicKey = &w.PrivetKey.PublicKey

	// calling a function to generate short address using publicKey
	adress := w.HashedAddress()

	w.BlockchainAddr = adress
	return w
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x\n", w.PrivetKey.D.Bytes())
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes())
}
