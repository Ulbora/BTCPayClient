package ptcpayclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	"github.com/btcsuite/btcd/btcec"
)

const (
	testStore = "TestStore"
	testToken = "AEgrxE3CwVAEWixX2gevGR58WuX9yGp9zP3BE6tJBAQHwe"
)

func TestBTCPayClient_GetRates(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())
	// kp := cc.GenerateKeyPair(btcec.S256())

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)
	ptc.crypto = cc
	//ptc.tokens.Token = testToken

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
			{
				"name": "US Dollar",
				"cryptoCode": "BTC",
				"currencyPair": "BTC_USD",
				"code": "USD",
				"rate": 51826.723
			}
		]
	}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200

	c := ptc.New(testBaseURL, kp, testToken)

	ptc.OverrideProxy(&gp)

	fmt.Println("new client: ", ptc)

	var cur = []string{"LTC_USD", "BTC_USD"}
	resp := c.GetRates(cur, testStore)
	fmt.Println("resp: ", resp)

	if resp.Data[0].CryptoCode != "BTC" {
		t.Fail()
	}

	//t.Fail()

}

func TestBTCPayClient_GetRatesFail(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())
	// kp := cc.GenerateKeyPair(btcec.S256())

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)
	//ptc.crypto = cc
	ptc.tokens.Token = testToken

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
			{
				"name": "US Dollar",
				"cryptoCode": "BTC",
				"currencyPair": "BTC_USD",
				"code": "USD",
				"rate": 51826.723
			}
		]
	}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 401

	c := ptc.New(testBaseURL, kp, testToken)

	ptc.OverrideProxy(&gp)

	fmt.Println("new client: ", ptc)

	var cur = []string{"LTC_USD", "BTC_USD"}
	resp := c.GetRates(cur, testStore)
	fmt.Println("resp: ", resp)

	if resp.Data[0].CryptoCode != "BTC" {
		t.Fail()
	}

	//t.Fail()

}
