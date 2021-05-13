package ptcpayclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	"github.com/btcsuite/btcd/btcec"
)

const (
	//testBaseURL = "http://127.0.0.1:49392"
	testBaseURL = "https://testnet.btcpayments.net"
)

func TestBTCPayClient_Token(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
		  {
			"policies": [
			  {
				"policy": "id",
				"method": "inactive",
				"params": [
				  "Tf8UuBAFcXX6AymQpVGSyEtb4oDzXsX7yUe"
				]
			  }
			],
			"token": "6cPAzk6jdcsLQPwoB4cn8J",
			"facade": "merchant",
			"dateCreated": 1558525586681,
			"pairingExpiration": 1558611986681,
			"pairingCode": "ZHcXiqX"
		  }
		]
	  }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200

	c := ptc.New(testBaseURL, kp, "")
	ptc.SetLogLevel(lg.AllLevel)
	ptc.OverrideProxy(&gp)
	var tkr TokenRequest
	tkr.ID = c.GetClientID() //cc.GetSinFromKey(kp)
	tkr.Facade = "merchant"

	resp := c.Token(&tkr)
	if resp.Code != 200 {
		t.Fail()
	}
	//t.Fail()
}

func TestBTCPayClient_Token_fail_code(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
		  {
			"policies": [
			  {
				"policy": "id",
				"method": "inactive",
				"params": [
				  "Tf8UuBAFcXX6AymQpVGSyEtb4oDzXsX7yUe"
				]
			  }
			],
			"token": "6cPAzk6jdcsLQPwoB4cn8J",
			"facade": "merchant",
			"dateCreated": 1558525586681,
			"pairingExpiration": 1558611986681,
			"pairingCode": "ZHcXiqX"
		  }
		]
	  }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 401

	c := ptc.New(testBaseURL, kp, "")
	ptc.SetLogLevel(lg.AllLevel)
	ptc.OverrideProxy(&gp)
	var tkr TokenRequest
	tkr.ID = c.GetClientID() // cc.GetSinFromKey(kp)
	tkr.Facade = "merchant"

	resp := c.Token(&tkr)
	if resp.Code != 401 {
		t.Fail()
	}
	//t.Fail()
}

func TestBTCPayClient_PairClient(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
		  {
			"policies": [
			  {
				"policy": "id",
				"method": "inactive",
				"params": [
				  "Tf8UuBAFcXX6AymQpVGSyEtb4oDzXsX7yUe"
				]
			  }
			],
			"token": "6cPAzk6jdcsLQPwoB4cn8J",
			"facade": "merchant",
			"dateCreated": 1558525586681,
			"pairingExpiration": 1558611986681,
			"pairingCode": "ZHcXiqX"
		  }
		]
	  }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200

	c := ptc.New(testBaseURL, kp, "")

	fmt.Println("new client: ", ptc)

	ptc.SetLogLevel(lg.AllLevel)
	ptc.OverrideProxy(&gp)
	var tkr TokenRequest
	tkr.ID = c.GetClientID() // cc.GetSinFromKey(kp)
	tkr.Facade = "merchant"

	resp := c.Token(&tkr)
	if resp.Code != 200 {
		t.Fail()
	}

	var gp2 px.MockGoProxy
	var mres2 http.Response
	mres2.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
		  {
			"policies": [
			  {
				"policy": "id",
				"method": "inactive",
				"params": [
				  "Tf8UuBAFcXX6AymQpVGSyEtb4oDzXsX7yUe"
				]
			  }
			],
			"token": "6cPAzk6jdcsLQPwoB4cn8J",
			"facade": "merchant",
			"dateCreated": 1558525586681,
			"pairingExpiration": 1558611986681,
			"pairingCode": "ZHcXiqX"
		  }
		]
	  }`))
	gp2.MockResp = &mres2
	gp2.MockDoSuccess1 = true
	gp2.MockRespCode = 200

	ptc.OverrideProxy(&gp2)

	fmt.Println("pres: ", resp)
	presp := c.PairClient(resp.Data[0].ParingCode)

	fmt.Println("presp: ", presp)

	fmt.Println("new client: ", ptc)

	if presp.Code != 200 {
		t.Fail()
	}

	//t.Fail()
}

func TestBTCPayClient_PairClient2(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b25"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())
	// kp := cc.GenerateKeyPair(btcec.S256())

	pub := cc.GetPublicKey(kp)

	fmt.Println("Public key x-identity: ", pub)

	fmt.Println("invoices post x-signature: ", pub)

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
		  {
			"policies": [
			  {
				"policy": "id",
				"method": "inactive",
				"params": [
				  "Tf8UuBAFcXX6AymQpVGSyEtb4oDzXsX7yUe"
				]
			  }
			],
			"token": "6cPAzk6jdcsLQPwoB4cn8J",
			"facade": "merchant",
			"dateCreated": 1558525586681,
			"pairingExpiration": 1558611986681,
			"pairingCode": "ZHcXiqX"
		  }
		]
	  }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200

	c := ptc.New(testBaseURL, kp, "")

	fmt.Println("new client: ", ptc)

	ptc.SetLogLevel(lg.AllLevel)
	ptc.OverrideProxy(&gp) //-----------------------------------
	var tkr TokenRequest
	tkr.ID = c.GetClientID() // cc.GetSinFromKey(kp)
	tkr.Label = "Six910 access"
	tkr.Facade = "merchant"

	resp := c.Token(&tkr)
	if resp.Code != 200 {
		t.Fail()
	}

	var gp2 px.MockGoProxy
	var mres2 http.Response
	mres2.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"data": [
		  {
			"policies": [
			  {
				"policy": "id",
				"method": "inactive",
				"params": [
				  "Tf8UuBAFcXX6AymQpVGSyEtb4oDzXsX7yUe"
				]
			  }
			],
			"token": "6cPAzk6jdcsLQPwoB4cn8J",
			"facade": "merchant",
			"dateCreated": 1558525586681,
			"pairingExpiration": 1558611986681,
			"pairingCode": "ZHcXiqX"
		  }
		]
	  }`))
	gp2.MockResp = &mres2
	gp2.MockDoSuccess1 = true
	gp2.MockRespCode = 200

	ptc.OverrideProxy(&gp2) //-------------------------------------

	fmt.Println("pres: ", resp)
	presp := c.PairClient(resp.Data[0].ParingCode)

	fmt.Println("presp: ", presp)
	fmt.Println("token: ", presp.Data[0].Token)
	fmt.Println("paring code: ", presp.Data[0].ParingCode)

	pcodeURL := c.GetPairingCodeRequest(presp.Data[0].ParingCode)

	fmt.Println("new client: ", ptc)

	fmt.Println("pairing code url: ", pcodeURL)

	if presp.Code != 200 {
		t.Fail()
	}

	//t.Fail()
}
