package ptcpayclient

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type Crypto interface {
	GenerateKeyPair() *ecdsa.PrivateKey
	LoadKeyPair(privateKey string) *ecdsa.PrivateKey
	GetSinFromKey(kp *ecdsa.PrivateKey) string
	Sign(hash []byte, kp *ecdsa.PrivateKey) ([]byte, error)
}

type Cryptography struct {
}

func (c *Cryptography) New() Crypto {
	return c
}

func (c *Cryptography) GenerateKeyPair() *ecdsa.PrivateKey {
	kp, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return kp
}

func (c *Cryptography) LoadKeyPair(privateKey string) *ecdsa.PrivateKey {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = elliptic.P256() //secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return &e
}

func (c *Cryptography) GetSinFromKey(kp *ecdsa.PrivateKey) string {
	var sin string

	return sin
}

func (c *Cryptography) Sign(hash []byte, kp *ecdsa.PrivateKey) ([]byte, error) {
	ehash := sha256.Sum256(hash)
	sig, err := ecdsa.SignASN1(rand.Reader, kp, ehash[:])
	return sig, err
}
