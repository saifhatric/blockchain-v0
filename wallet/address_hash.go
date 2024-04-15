package wallet

import (
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func (w *Wallet) HashedAddress() string {
	// 1. hash public key with sha256 (32 bytes)
	h1 := sha256.New()
	h1.Write(w.PublicKey.X.Bytes())
	h1.Write(w.PublicKey.Y.Bytes())
	digest1 := h1.Sum(nil)

	// 2. preform ripmd-160 on prevHash to make it shorter (20 bytes)
	h2 := ripemd160.New()
	h2.Write(digest1)
	digest2 := h2.Sum(nil)

	// 3. add version at the bigning of the hash (0x00)
	vd3 := make([]byte, 21)
	vd3[0] = 0x00
	copy(vd3[1:], digest2[:])

	// 4. preform sha-256 on vd3
	h4 := sha256.New()
	h4.Write(vd3)
	digest4 := h4.Sum(nil)

	// 5. preform sha-256 on previous hash
	h5 := sha256.New()
	h5.Write(digest4)
	digest5 := h5.Sum(nil)

	// 6. taking first 04 bytes for checkSum
	checkSum := digest5[:4]

	// 7. adding the checkSum at the end of vd3(step 04) and make it in 25 bytes
	dc7 := make([]byte, 25)
	copy(dc7[21:], vd3[:])
	copy(dc7[:21], checkSum[:])

	// 8. convert the result into base-58
	return base58.Encode(dc7)

}
