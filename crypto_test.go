package ptcpayclient

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec"
)

func TestCryptography_GenerateKeyPair(t *testing.T) {
	var cryt Cryptography
	c := cryt.New()
	pk := c.GenerateKeyPair(btcec.S256())
	fmt.Println("prv key: ", pk)

	fmt.Println("pubkey:", pk.PublicKey)

	fmt.Println("D:", pk.D)

	fmt.Println("Curve:", pk.Curve)
	fmt.Println("P256:", btcec.S256())
	fmt.Println("X:", pk.X)
	fmt.Println("Y:", pk.Y)

	if pk.X.String() == "" || pk.Y.String() == "" {
		t.Fail()
	}
}

func TestCryptography_Sign(t *testing.T) {
	var cryt Cryptography
	c := cryt.New()

	kp := c.GenerateKeyPair(btcec.S256())
	fmt.Println("prv key in sign: ", kp)

	msg := "hello, world"
	// hash := sha256.Sum256([]byte(msg))
	hash := []byte(msg)

	sig, err := c.Sign(hash, kp)

	nhash := sha256.Sum256([]byte(msg))
	valid := ecdsa.VerifyASN1(&kp.PublicKey, nhash[:], sig)
	fmt.Println("valid: ", valid)
	fmt.Println("err: ", err)
	if !valid || err != nil {
		t.Fail()
	}

}

func TestCryptography_SignFail(t *testing.T) {
	var cryt Cryptography
	c := cryt.New()

	kp := c.GenerateKeyPair(btcec.S256())
	//fmt.Println("priv key", kp.PrivateKey)

	kp2 := c.GenerateKeyPair(btcec.S256())
	fmt.Println("prv key in sign: ", kp)

	msg := "hello, world"
	// hash := sha256.Sum256([]byte(msg))
	hash := []byte(msg)

	sig, err := c.Sign(hash, kp)

	nhash := sha256.Sum256([]byte(msg))
	valid := ecdsa.VerifyASN1(&kp2.PublicKey, nhash[:], sig)
	fmt.Println("valid: ", valid)
	fmt.Println("err: ", err)
	if valid {
		t.Fail()
	}
}

func TestCryptography_LoadKeyPairFail(t *testing.T) {
	var pk = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"
	var cryt Cryptography
	c := cryt.New()
	kp := c.LoadKeyPair(pk, btcec.S256())

	kp2 := c.GenerateKeyPair(btcec.S256())

	fmt.Println("Curve:", kp.Curve)
	fmt.Println("P256:", btcec.S256())

	fmt.Println("prv key in sign: ", kp)

	msg := "hello, world"
	// hash := sha256.Sum256([]byte(msg))
	hash := []byte(msg)

	sig, err := c.Sign(hash, kp)

	nhash := sha256.Sum256([]byte(msg))
	valid := ecdsa.VerifyASN1(&kp2.PublicKey, nhash[:], sig)
	fmt.Println("valid: ", valid)
	fmt.Println("err: ", err)
	if valid {
		t.Fail()
	}

}

func TestCryptography_LoadKeyPair(t *testing.T) {
	var pk = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"
	var cryt Cryptography
	c := cryt.New()
	kp := c.LoadKeyPair(pk, btcec.S256())

	fmt.Println("Curve:", kp.Curve)
	fmt.Println("P256:", btcec.S256())

	fmt.Println("prv key in sign: ", kp)

	msg := "hello, world"
	// hash := sha256.Sum256([]byte(msg))
	hash := []byte(msg)

	sig, err := c.Sign(hash, kp)

	nhash := sha256.Sum256([]byte(msg))
	valid := ecdsa.VerifyASN1(&kp.PublicKey, nhash[:], sig)
	fmt.Println("valid: ", valid)
	fmt.Println("err: ", err)
	if !valid || err != nil {
		t.Fail()
	}

}

type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

func TestCryptography_GetSinFromKey(t *testing.T) {
	pm := "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEICg7E4NN53YkaWuAwpoqjfAofjzKI7Jq1f532dX+0O6QoAcGBSuBBAAK\noUQDQgAEjZcNa6Kdz6GQwXcUD9iJ+t1tJZCx7hpqBuJV2/IrQBfue8jh8H7Q/4vX\nfAArmNMaGotTpjdnymWlMfszzXJhlw==\n-----END EC PRIVATE KEY-----\n"

	clientId := "TeyN4LPrXiG5t2yuSamKqP3ynVk3F52iHrX"
	key := extractKeyFromPem(pm)

	var kp ecdsa.PrivateKey = ecdsa.PrivateKey(*key)

	var cryt Cryptography
	c := cryt.New()

	fmt.Println("prv key from pen: ", kp)

	sin := c.GetSinFromKey(&kp)
	fmt.Println("sin: ", sin)
	if clientId != sin {
		t.Fail()
	}
	//t.Fail()

}

func TestCryptography_GetSinFromKey2(t *testing.T) {

	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	clientId := "TfDnXWvj6bBhkduYiZnohg5qhtDu5VWohhw"

	var cryt Cryptography
	c := cryt.New()

	kp := c.LoadKeyPair(pkh, btcec.S256())

	fmt.Println("prv key from pen: ", kp)

	sin := c.GetSinFromKey(kp)
	fmt.Println("sin: ", sin)
	if clientId != sin {
		t.Fail()
	}
	//t.Fail()

}

func extractKeyFromPem(pm string) *btcec.PrivateKey {
	byta := []byte(pm)
	blck, _ := pem.Decode(byta)
	var ecp ecPrivateKey
	asn1.Unmarshal(blck.Bytes, &ecp)
	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), ecp.PrivateKey)
	return priv
}

func Test_paddedAppend(t *testing.T) {
	b1 := []byte{'a'}
	b2 := []byte{'b', 'c'}

	b3 := paddedAppend(2, b2, b1)
	fmt.Println("b3: ", b3)
	fmt.Println("len b3: ", len(b3))

	if len(b3) != 4 {
		t.Fail()
	}
	//t.Fail()
}
