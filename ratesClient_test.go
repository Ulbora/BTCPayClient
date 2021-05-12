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
	testToken = "GGsauSDZek5nffnHP9oJ6vxSRYKyhcGNN8PjTVX8aL2rwG"
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
	ptc.tokens.Token = testToken

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

	c := ptc.New(testBaseURL, kp)

	fmt.Println("new client: ", ptc)

	var cur = []string{"LTC_USD", "BTC_USD"}
	resp := c.GetRates(cur, testStore)
	fmt.Println("resp: ", resp)

	t.Fail()

}
