package ptcpayclient

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestCryptography_GenerateKeyPair(t *testing.T) {
	var cryt Cryptography
	c := cryt.New()
	pk := c.GenerateKeyPair()
	fmt.Println("prv key: ", pk)

	fmt.Println("pubkey:", pk.PublicKey)

	fmt.Println("D:", pk.D)

	fmt.Println("Curve:", pk.Curve)
	fmt.Println("P256:", elliptic.P256())
	fmt.Println("X:", pk.X)
	fmt.Println("Y:", pk.Y)

	if pk.X.String() == "" || pk.Y.String() == "" {
		t.Fail()
	}
}

func TestCryptography_Sign(t *testing.T) {
	var cryt Cryptography
	c := cryt.New()

	kp := c.GenerateKeyPair()
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

	kp := c.GenerateKeyPair()
	//fmt.Println("priv key", kp.PrivateKey)

	kp2 := c.GenerateKeyPair()
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
	kp := c.LoadKeyPair(pk)

	kp2 := c.GenerateKeyPair()

	fmt.Println("Curve:", kp.Curve)
	fmt.Println("P256:", elliptic.P256())

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
	kp := c.LoadKeyPair(pk)

	fmt.Println("Curve:", kp.Curve)
	fmt.Println("P256:", elliptic.P256())

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
