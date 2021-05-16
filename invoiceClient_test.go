package ptcpayclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	"github.com/btcsuite/btcd/btcec"
)

func TestBTCPayClient_CreateInvoice(t *testing.T) {
	var pkh = "31eb31ecf1a640cd91e0a1105501f36235f8c7d51d67dcf74ccc968d74cb6b26"

	var cryt Cryptography
	cc := cryt.New()

	kp := cc.LoadKeyPair(pkh, btcec.S256())
	// kp := cc.GenerateKeyPair(btcec.S256())

	var ptc BTCPayClient
	var head Headers
	ptc.SetHeader(head)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"facade":"pos/invoice",
		"data":{
		   "url":"https://testnet.demo.btcpayserver.org/invoice?id=XCk3Wx3YpmY8RDKC8qi6Jc",
		   "posData":"",
		   "status":"new",
		   "btcPrice":"0.00020788",
		   "btcDue":"0.00020788",
		   "cryptoInfo":[
			  {
				 "cryptoCode":"BTC",
				 "paymentType":"BTCLike",
				 "rate":48104.782,
				 "exRates":{
					"USD":0.0
				 },
				 "paid":"0.00000000",
				 "price":"0.00020788",
				 "due":"0.00020788",
				 "paymentUrls":{
					"BIP21":"bitcoin:tb1qatye930xsacklnrgmaexl8feszq5692qnajcsu?amount=0.00020788",
					"BIP72":null,
					"BIP72b":null,
					"BIP73":null,
					"BOLT11":null
				 },
				 "address":"tb1qatye930xsacklnrgmaexl8feszq5692qnajcsu",
				 "url":"https://testnet.demo.btcpayserver.org/i/BTC/XCk3Wx3YpmY8RDKC8qi6Jc",
				 "totalDue":"0.00020788",
				 "networkFee":"0.00000000",
				 "txCount":0,
				 "cryptoPaid":"0.00000000",
				 "payments":[
					
				 ]
			  }
		   ],
		   "price":10.0,
		   "currency":"USD",
		   "exRates":{
			  "USD":0.0
		   },
		   "buyerTotalBtcAmount":null,
		   "itemDesc":"",
		   "itemCode":"",
		   "orderId":"",
		   "guid":"61a3bb79-8b78-4969-8876-f25b44cfd4d4",
		   "id":"XCk3Wx3YpmY8RDKC8qi6Jc",
		   "invoiceTime":1621119646000,
		   "expirationTime":1621120546000,
		   "currentTime":1621119646639,
		   "lowFeeDetected":false,
		   "btcPaid":"0.00000000",
		   "rate":48104.782,
		   "exceptionStatus":false,
		   "paymentUrls":{
			  "BIP21":"bitcoin:tb1qatye930xsacklnrgmaexl8feszq5692qnajcsu?amount=0.00020788",
			  "BIP72":null,
			  "BIP72b":null,
			  "BIP73":null,
			  "BOLT11":null
		   },
		   "refundAddressRequestPending":false,
		   "buyerPaidBtcMinerFee":null,
		   "bitcoinAddress":"tb1qatye930xsacklnrgmaexl8feszq5692qnajcsu",
		   "token":"C9Pa5TRzwmFtmSTZBPDoqf",
		   "flags":{
			  "refundable":false
		   },
		   "paymentSubtotals":{
			  "BTC":20788.0
		   },
		   "paymentTotals":{
			  "BTC":20788.0
		   },
		   "amountPaid":0,
		   "minerFees":{
			  "BTC":{
				 "satoshisPerByte":1.0,
				 "totalFee":0.0
			  }
		   },
		   "exchangeRates":{
			  "BTC":{
				 "USD":0.0
			  }
		   },
		   "supportedTransactionCurrencies":{
			  "BTC":{
				 "enabled":true,
				 "reason":null
			  }
		   },
		   "addresses":{
			  "BTC":"tb1qatye930xsacklnrgmaexl8feszq5692qnajcsu"
		   },
		   "paymentCodes":{
			  "BTC":{
				 "BIP21":"bitcoin:tb1qatye930xsacklnrgmaexl8feszq5692qnajcsu?amount=0.00020788",
				 "BIP72":null,
				 "BIP72b":null,
				 "BIP73":null,
				 "BOLT11":null
			  }
		   },
		   "buyer":{
			  "name":"bob willson",
			  "address1":"",
			  "address2":"",
			  "locality":"",
			  "region":"",
			  "postalCode":"",
			  "country":"",
			  "phone":"",
			  "email":"bob@bob.com"
		   }
		}
	 }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 401

	c := ptc.New(testBaseURL, kp, testToken)

	ptc.OverrideProxy(&gp)

	fmt.Println("new client: ", ptc)

	var req InvoiceReq
	req.Token = testToken
	req.Currency = "USD"
	req.Price = 10
	req.Buyer.Name = "bob willson"
	req.Buyer.Email = "bob@bob.com"

	resp := c.CreateInvoice(&req)
	aJSON, err := json.Marshal(resp)
	fmt.Println("err: ", err)

	fmt.Println("resp: ", string(aJSON))

	if resp.Data.Price != 10 {
		t.Fail()
	}

	//t.Fail()
}
