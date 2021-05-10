package ptcpayclient

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"golang.org/x/crypto/ripemd160"

	"github.com/btcsuite/btcutil/base58"
)

const (
	pubkeyCompressed         byte = 0x2
	pubKeyBytesLenCompressed      = 33
)

//Crypto Crypto
type Crypto interface {
	GenerateKeyPair(ec elliptic.Curve) *ecdsa.PrivateKey
	LoadKeyPair(privateKey string, ec elliptic.Curve) *ecdsa.PrivateKey
	GetSinFromKey(kp *ecdsa.PrivateKey) string
	Sign(hash []byte, kp *ecdsa.PrivateKey) ([]byte, error)
	GetPublicKey(kp *ecdsa.PrivateKey) string
}

//Cryptography Cryptography
type Cryptography struct {
}

//New New
func (c *Cryptography) New() Crypto {
	return c
}

//GenerateKeyPair GenerateKeyPair
func (c *Cryptography) GenerateKeyPair(ec elliptic.Curve) *ecdsa.PrivateKey {
	//A public private key pair is created using elliptic curve secp256k1
	kp, _ := ecdsa.GenerateKey(ec, rand.Reader)
	return kp
}

//LoadKeyPair LoadKeyPair
func (c *Cryptography) LoadKeyPair(privateKey string, ec elliptic.Curve) *ecdsa.PrivateKey {
	//A public private key pair is created using elliptic curve secp256k1
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = ec //secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return &e
}

//GetSinFromKey GetSinFromKey
func (c *Cryptography) GetSinFromKey(kp *ecdsa.PrivateKey) string {
	//A public private key pair is created using elliptic curve secp256k1
	var sin string
	pub := kp.PublicKey
	comp := serializeCompressed(pub)
	//fmt.Println("byta: ", comp)
	key := hex.EncodeToString(comp)
	hexa := sha256ofHexString(key)
	hexa = ripemd160ofHexString(hexa)
	versionSinTypeEtc := "0F02" + hexa
	hexa = sha256ofHexString(versionSinTypeEtc)
	hexa = sha256ofHexString(hexa)
	checksum := hexa[0:8]
	hexa = versionSinTypeEtc + checksum
	byta, _ := hex.DecodeString(hexa)
	sin = base58.Encode(byta)
	//fmt.Println("sin: ", sin)
	return sin
}

//Sign Sign
func (c *Cryptography) Sign(hash []byte, kp *ecdsa.PrivateKey) ([]byte, error) {
	//A public private key pair is created using elliptic curve secp256k1
	ehash := sha256.Sum256(hash)
	sig, err := ecdsa.SignASN1(rand.Reader, kp, ehash[:])
	return sig, err
}

//GetPublicKey GetPublicKey
func (c *Cryptography) GetPublicKey(kp *ecdsa.PrivateKey) string {
	comp := serializeCompressed(kp.PublicKey)
	key := hex.EncodeToString(comp)
	return key
}

func sha256ofHexString(hexa string) string {
	//fmt.Println("hexa: ", hexa)
	byta, _ := hex.DecodeString(hexa)
	hash := sha256.New()
	hash.Write(byta)
	hashsum := hash.Sum(nil)
	hexb := hex.EncodeToString(hashsum)
	//fmt.Println("hexb: ", hexb)
	return hexb
}

func serializeCompressed(p ecdsa.PublicKey) []byte {
	b := make([]byte, 0, pubKeyBytesLenCompressed)
	format := pubkeyCompressed
	if isOdd(p.Y) {
		format |= 0x1
	}
	b = append(b, format)
	return paddedAppend(32, b, p.X.Bytes())
}

func isOdd(a *big.Int) bool {
	return a.Bit(0) == 1
}

func ripemd160ofHexString(hexa string) string {
	byta, _ := hex.DecodeString(hexa)
	hash := ripemd160.New()
	hash.Write(byta)
	hashsum := hash.Sum(nil)
	hexb := hex.EncodeToString(hashsum)
	return hexb
}

func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
