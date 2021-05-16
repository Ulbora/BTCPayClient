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
	req.TransactionSpeed = "medium"
	req.Buyer.Name = "bob willson"
	req.Buyer.Email = "bob@bob.com"

	resp := c.CreateInvoice(&req)
	aJSON, err := json.Marshal(resp)
	fmt.Println("err: ", err)

	fmt.Println("resp: ", string(aJSON))

	if resp.Data.Price != 10 {
		t.Fail()
	}

	// t.Fail()
}

func TestBTCPayClient_GetInvoice(t *testing.T) {
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

	// var req InvoiceReq
	// req.Token = testToken
	// req.Currency = "USD"
	// req.Price = 10
	// req.Buyer.Name = "bob willson"
	// req.Buyer.Email = "bob@bob.com"

	resp := c.GetInvoice("RSM815RoMU3XCqPXx3GoZX")
	aJSON, err := json.Marshal(resp)
	fmt.Println("err: ", err)

	fmt.Println("resp: ", string(aJSON))

	if resp.Data.Price != 10 {
		t.Fail()
	}

	// t.Fail()
}

func TestBTCPayClient_GetInvoices(t *testing.T) {
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
		"data":[
		   {
			  "guid":"a7515790-1a40-4083-abf4-cf4b17ab946f",
			  "id":"C39EyqDdDozG4oDSiWZ1qS",
			  "url":"https://testnet.demo.btcpayserver.org/invoice?id=C39EyqDdDozG4oDSiWZ1qS",
			  "btcPrice":"0.00021745",
			  "btcDue":"0.00000000",
			  "cryptoInfo":[
				 {
					"cryptoCode":"BTC",
					"paymentType":"BTCLike",
					"rate":1,
					"exRates":{
					   "BTC":0
					},
					"paid":"0.00021745",
					"price":"0.00021745",
					"due":"0.00000000",
					"paymentUrls":{
					   "BIP21":"bitcoin:tb1q0mkdp2q9ww6dcc4a6wcdj98rqcvfnr8suqsswz?amount=0.00",
					   "BIP72":"",
					   "BIP72b":"",
					   "BIP73":"",
					   "BOLT11":""
					},
					"address":"tb1q0mkdp2q9ww6dcc4a6wcdj98rqcvfnr8suqsswz",
					"url":"https://testnet.demo.btcpayserver.org/i/BTC/C39EyqDdDozG4oDSiWZ1qS",
					"totalDue":"0.00021745",
					"networkFee":"0.00000000",
					"txCount":1,
					"cryptoPaid":"0.00021745",
					"payments":[
					   {
						  "id":"7ef841f3b458a487fed07308887b0df1a1cd18e2a832a55d4adcc1f6fd6a56b6-0",
						  "receivedDate":"2021-05-16T19:33:25.189",
						  "value":0.00021745,
						  "fee":0,
						  "paymentType":"BTCLike",
						  "confirmed":true,
						  "completed":true,
						  "destination":"tb1q0mkdp2q9ww6dcc4a6wcdj98rqcvfnr8suqsswz"
					   }
					]
				 }
			  ],
			  "exRates":{
				 "BTC":0
			  },
			  "buyerTotalBtcAmount":"",
			  "invoiceTime":1621193482000,
			  "currentTime":1621199233675,
			  "lowFeeDetected":false,
			  "btcPaid":"0.00021745",
			  "rate":1,
			  "exceptionStatus":false,
			  "paymentUrls":{
				 "BIP21":"bitcoin:tb1q0mkdp2q9ww6dcc4a6wcdj98rqcvfnr8suqsswz?amount=0.00",
				 "BIP72":"",
				 "BIP72b":"",
				 "BIP73":"",
				 "BOLT11":""
			  },
			  "refundAddressRequestPending":false,
			  "buyerPaidBtcMinerFee":"",
			  "bitcoinAddress":"tb1q0mkdp2q9ww6dcc4a6wcdj98rqcvfnr8suqsswz",
			  "flags":{
				 "refundable":false
			  },
			  "paymentSubtotals":{
				 "BTC":21745
			  },
			  "paymentTotals":{
				 "BTC":21745
			  },
			  "amountPaid":0,
			  "minerFees":{
				 "BTC":{
					"satoshisPerByte":1,
					"totalFee":0
				 }
			  },
			  "exchangeRates":{
				 "BTC":{
					"BTC":0
				 }
			  },
			  "addresses":{
				 "BTC":"tb1q0mkdp2q9ww6dcc4a6wcdj98rqcvfnr8suqsswz"
			  },
			  "paymentCodes":{
				 "BTC":{
					"BIP21":"bitcoin:tb1q0mkdp2q9ww6dcc4a6wcdj98rqcvfnr8suqsswz?amount=0.00",
					"BIP72":"",
					"BIP72b":"",
					"BIP73":"",
					"BOLT11":""
				 }
			  },
			  "currency":"BTC",
			  "price":0.00021745,
			  "orderId":"",
			  "expirationTime":1621194382000,
			  "itemDesc":"",
			  "itemCode":"",
			  "posData":"",
			  "status":"complete",
			  "supportedTransactionCurrencies":{
				 "BTC":{
					"enabled":true,
					"reason":""
				 }
			  },
			  "taxIncluded":0,
			  "token":"FR7v1JqC7PudBZQTomUJEj",
			  "redirectAutomatically":false,
			  "notificationEmail":"",
			  "notificationURL":"",
			  "extendedNotifications":false,
			  "fullNotifications":false,
			  "buyer":{
				 "name":"",
				 "address1":"",
				 "address2":"",
				 "city":"",
				 "locality":"",
				 "region":"",
				 "postalCode":"",
				 "state":"",
				 "zip":"",
				 "country":"",
				 "phone":"",
				 "notify":false,
				 "email":"kenwwilliamson12@gmail.com"
			  }
		   },
		   {
			  "guid":"b0d7e24d-7110-4f77-a485-2964cf8e3c5d",
			  "id":"Rdq2RLuf7bdBy9SV3GATrY",
			  "url":"https://testnet.demo.btcpayserver.org/invoice?id=Rdq2RLuf7bdBy9SV3GATrY",
			  "btcPrice":"0.00021862",
			  "btcDue":"0.00000000",
			  "cryptoInfo":[
				 {
					"cryptoCode":"BTC",
					"paymentType":"BTCLike",
					"rate":45741.782,
					"exRates":{
					   "USD":0
					},
					"paid":"0.00021862",
					"price":"0.00021862",
					"due":"0.00000000",
					"paymentUrls":{
					   "BIP21":"bitcoin:tb1qutvyjc6zd98m30pk3fvxhrdcv9mhx7vlf7fnun?amount=0.00",
					   "BIP72":"",
					   "BIP72b":"",
					   "BIP73":"",
					   "BOLT11":""
					},
					"address":"tb1qutvyjc6zd98m30pk3fvxhrdcv9mhx7vlf7fnun",
					"url":"https://testnet.demo.btcpayserver.org/i/BTC/Rdq2RLuf7bdBy9SV3GATrY",
					"totalDue":"0.00021862",
					"networkFee":"0.00000000",
					"txCount":1,
					"cryptoPaid":"0.00021862",
					"payments":[
					   {
						  "id":"05e1b0687e0d2dc761fecc6b07387528e94dd3155ec9943bdb4ff5132d8775e7-0",
						  "receivedDate":"2021-05-16T19:19:02.841",
						  "value":0.00021862,
						  "fee":0,
						  "paymentType":"BTCLike",
						  "confirmed":true,
						  "completed":true,
						  "destination":"tb1qutvyjc6zd98m30pk3fvxhrdcv9mhx7vlf7fnun"
					   }
					]
				 }
			  ],
			  "exRates":{
				 "USD":0
			  },
			  "buyerTotalBtcAmount":"",
			  "invoiceTime":1621192646000,
			  "currentTime":1621199233676,
			  "lowFeeDetected":false,
			  "btcPaid":"0.00021862",
			  "rate":45741.782,
			  "exceptionStatus":false,
			  "paymentUrls":{
				 "BIP21":"bitcoin:tb1qutvyjc6zd98m30pk3fvxhrdcv9mhx7vlf7fnun?amount=0.00",
				 "BIP72":"",
				 "BIP72b":"",
				 "BIP73":"",
				 "BOLT11":""
			  },
			  "refundAddressRequestPending":false,
			  "buyerPaidBtcMinerFee":"",
			  "bitcoinAddress":"tb1qutvyjc6zd98m30pk3fvxhrdcv9mhx7vlf7fnun",
			  "flags":{
				 "refundable":false
			  },
			  "paymentSubtotals":{
				 "BTC":21862
			  },
			  "paymentTotals":{
				 "BTC":21862
			  },
			  "amountPaid":0,
			  "minerFees":{
				 "BTC":{
					"satoshisPerByte":1,
					"totalFee":0
				 }
			  },
			  "exchangeRates":{
				 "BTC":{
					"USD":0
				 }
			  },
			  "addresses":{
				 "BTC":"tb1qutvyjc6zd98m30pk3fvxhrdcv9mhx7vlf7fnun"
			  },
			  "paymentCodes":{
				 "BTC":{
					"BIP21":"bitcoin:tb1qutvyjc6zd98m30pk3fvxhrdcv9mhx7vlf7fnun?amount=0.00",
					"BIP72":"",
					"BIP72b":"",
					"BIP73":"",
					"BOLT11":""
				 }
			  },
			  "currency":"USD",
			  "price":10,
			  "orderId":"",
			  "expirationTime":1621193546000,
			  "itemDesc":"",
			  "itemCode":"",
			  "posData":"",
			  "status":"complete",
			  "supportedTransactionCurrencies":{
				 "BTC":{
					"enabled":true,
					"reason":""
				 }
			  },
			  "taxIncluded":0,
			  "token":"3SBLowzmiTRLh9a3up3auB",
			  "redirectAutomatically":false,
			  "notificationEmail":"",
			  "notificationURL":"",
			  "extendedNotifications":false,
			  "fullNotifications":false,
			  "buyer":{
				 "name":"bob willson",
				 "address1":"",
				 "address2":"",
				 "city":"",
				 "locality":"",
				 "region":"",
				 "postalCode":"",
				 "state":"",
				 "zip":"",
				 "country":"",
				 "phone":"",
				 "notify":false,
				 "email":"bob@bob.com"
			  }
		   },
		   {
			  "guid":"cb158459-76d8-46d0-a07b-2f1c709991fe",
			  "id":"6ftmsau2QJ7hUbokBcF3XT",
			  "url":"https://testnet.demo.btcpayserver.org/invoice?id=6ftmsau2QJ7hUbokBcF3XT",
			  "btcPrice":"0.00021872",
			  "btcDue":"0.00000000",
			  "cryptoInfo":[
				 {
					"cryptoCode":"BTC",
					"paymentType":"BTCLike",
					"rate":45721.859,
					"exRates":{
					   "USD":0
					},
					"paid":"0.00021872",
					"price":"0.00021872",
					"due":"0.00000000",
					"paymentUrls":{
					   "BIP21":"bitcoin:tb1q6q87s3fydk6nhj7xamhps496q5c8kng3eg26jh?amount=0.00",
					   "BIP72":"",
					   "BIP72b":"",
					   "BIP73":"",
					   "BOLT11":""
					},
					"address":"tb1q6q87s3fydk6nhj7xamhps496q5c8kng3eg26jh",
					"url":"https://testnet.demo.btcpayserver.org/i/BTC/6ftmsau2QJ7hUbokBcF3XT",
					"totalDue":"0.00021872",
					"networkFee":"0.00000000",
					"txCount":1,
					"cryptoPaid":"0.00021872",
					"payments":[
					   {
						  "id":"a49e6d075bd3b808c8dfa59bd22f2a80473e474d9a7876ffb6277df75da9e0ba-0",
						  "receivedDate":"2021-05-16T19:14:41.046",
						  "value":0.00021872,
						  "fee":0,
						  "paymentType":"BTCLike",
						  "confirmed":true,
						  "completed":true,
						  "destination":"tb1q6q87s3fydk6nhj7xamhps496q5c8kng3eg26jh"
					   }
					]
				 }
			  ],
			  "exRates":{
				 "USD":0
			  },
			  "buyerTotalBtcAmount":"",
			  "invoiceTime":1621192337000,
			  "currentTime":1621199233678,
			  "lowFeeDetected":false,
			  "btcPaid":"0.00021872",
			  "rate":45721.859,
			  "exceptionStatus":false,
			  "paymentUrls":{
				 "BIP21":"bitcoin:tb1q6q87s3fydk6nhj7xamhps496q5c8kng3eg26jh?amount=0.00",
				 "BIP72":"",
				 "BIP72b":"",
				 "BIP73":"",
				 "BOLT11":""
			  },
			  "refundAddressRequestPending":false,
			  "buyerPaidBtcMinerFee":"",
			  "bitcoinAddress":"tb1q6q87s3fydk6nhj7xamhps496q5c8kng3eg26jh",
			  "flags":{
				 "refundable":false
			  },
			  "paymentSubtotals":{
				 "BTC":21872
			  },
			  "paymentTotals":{
				 "BTC":21872
			  },
			  "amountPaid":0,
			  "minerFees":{
				 "BTC":{
					"satoshisPerByte":1,
					"totalFee":0
				 }
			  },
			  "exchangeRates":{
				 "BTC":{
					"USD":0
				 }
			  },
			  "addresses":{
				 "BTC":"tb1q6q87s3fydk6nhj7xamhps496q5c8kng3eg26jh"
			  },
			  "paymentCodes":{
				 "BTC":{
					"BIP21":"bitcoin:tb1q6q87s3fydk6nhj7xamhps496q5c8kng3eg26jh?amount=0.00",
					"BIP72":"",
					"BIP72b":"",
					"BIP73":"",
					"BOLT11":""
				 }
			  },
			  "currency":"USD",
			  "price":10,
			  "orderId":"",
			  "expirationTime":1621193237000,
			  "itemDesc":"",
			  "itemCode":"",
			  "posData":"",
			  "status":"complete",
			  "supportedTransactionCurrencies":{
				 "BTC":{
					"enabled":true,
					"reason":""
				 }
			  },
			  "taxIncluded":0,
			  "token":"LxZpddGDmygJPVWhNyc7sA",
			  "redirectAutomatically":false,
			  "notificationEmail":"",
			  "notificationURL":"",
			  "extendedNotifications":false,
			  "fullNotifications":false,
			  "buyer":{
				 "name":"bob willson",
				 "address1":"",
				 "address2":"",
				 "city":"",
				 "locality":"",
				 "region":"",
				 "postalCode":"",
				 "state":"",
				 "zip":"",
				 "country":"",
				 "phone":"",
				 "notify":false,
				 "email":"bob@bob.com"
			  }
		   },
		   {
			  "guid":"581666ae-4e2a-4e94-9b28-8246ccf34f25",
			  "id":"YPFPA2yeDAoc6YLZPLWhgF",
			  "url":"https://testnet.demo.btcpayserver.org/invoice?id=YPFPA2yeDAoc6YLZPLWhgF",
			  "btcPrice":"0.00021809",
			  "btcDue":"0.00000000",
			  "cryptoInfo":[
				 {
					"cryptoCode":"BTC",
					"paymentType":"BTCLike",
					"rate":45852.727,
					"exRates":{
					   "USD":0
					},
					"paid":"0.00021909",
					"price":"0.00021809",
					"due":"0.00000000",
					"paymentUrls":{
					   "BIP21":"bitcoin:tb1qaw4k5av85pgpggm56un3qkynq73h4sjx0vtl89?amount=0.00",
					   "BIP72":"",
					   "BIP72b":"",
					   "BIP73":"",
					   "BOLT11":""
					},
					"address":"tb1qaw4k5av85pgpggm56un3qkynq73h4sjx0vtl89",
					"url":"https://testnet.demo.btcpayserver.org/i/BTC/YPFPA2yeDAoc6YLZPLWhgF",
					"totalDue":"0.00021909",
					"networkFee":"0.00000100",
					"txCount":2,
					"cryptoPaid":"0.00021909",
					"payments":[
					   {
						  "id":"1f4bb7381eaffa59c4dfff5ee331a55f03ae28a20332af797906976642b61b6d-0",
						  "receivedDate":"2021-05-16T19:02:56.097",
						  "value":2.2e-7,
						  "fee":0,
						  "paymentType":"BTCLike",
						  "confirmed":true,
						  "completed":true,
						  "destination":"tb1qq76m59kc6hdntnz5053lf45ade93qa9a98cz68"
					   },
					   {
						  "id":"84497156933fd7593f0baf3791d9bf946264a0483069ad94c4502d43d73c5075-0",
						  "receivedDate":"2021-05-16T19:09:38.14",
						  "value":0.00021887,
						  "fee":0.000001,
						  "paymentType":"BTCLike",
						  "confirmed":true,
						  "completed":true,
						  "destination":"tb1qaw4k5av85pgpggm56un3qkynq73h4sjx0vtl89"
					   }
					]
				 }
			  ],
			  "exRates":{
				 "USD":0
			  },
			  "buyerTotalBtcAmount":"",
			  "invoiceTime":1621191554000,
			  "currentTime":1621199233679,
			  "lowFeeDetected":false,
			  "btcPaid":"0.00021909",
			  "rate":45852.727,
			  "exceptionStatus":false,
			  "paymentUrls":{
				 "BIP21":"bitcoin:tb1qaw4k5av85pgpggm56un3qkynq73h4sjx0vtl89?amount=0.00",
				 "BIP72":"",
				 "BIP72b":"",
				 "BIP73":"",
				 "BOLT11":""
			  },
			  "refundAddressRequestPending":false,
			  "buyerPaidBtcMinerFee":"",
			  "bitcoinAddress":"tb1qaw4k5av85pgpggm56un3qkynq73h4sjx0vtl89",
			  "flags":{
				 "refundable":false
			  },
			  "paymentSubtotals":{
				 "BTC":21809
			  },
			  "paymentTotals":{
				 "BTC":21909
			  },
			  "amountPaid":0,
			  "minerFees":{
				 "BTC":{
					"satoshisPerByte":1,
					"totalFee":100
				 }
			  },
			  "exchangeRates":{
				 "BTC":{
					"USD":0
				 }
			  },
			  "addresses":{
				 "BTC":"tb1qaw4k5av85pgpggm56un3qkynq73h4sjx0vtl89"
			  },
			  "paymentCodes":{
				 "BTC":{
					"BIP21":"bitcoin:tb1qaw4k5av85pgpggm56un3qkynq73h4sjx0vtl89?amount=0.00",
					"BIP72":"",
					"BIP72b":"",
					"BIP73":"",
					"BOLT11":""
				 }
			  },
			  "currency":"USD",
			  "price":10,
			  "orderId":"",
			  "expirationTime":1621192454000,
			  "itemDesc":"",
			  "itemCode":"",
			  "posData":"",
			  "status":"complete",
			  "supportedTransactionCurrencies":{
				 "BTC":{
					"enabled":true,
					"reason":""
				 }
			  },
			  "taxIncluded":0,
			  "token":"D6diehSJwVUXXjQiMAQaVo",
			  "redirectAutomatically":false,
			  "notificationEmail":"",
			  "notificationURL":"",
			  "extendedNotifications":false,
			  "fullNotifications":false,
			  "buyer":{
				 "name":"bob willson",
				 "address1":"",
				 "address2":"",
				 "city":"",
				 "locality":"",
				 "region":"",
				 "postalCode":"",
				 "state":"",
				 "zip":"",
				 "country":"",
				 "phone":"",
				 "notify":false,
				 "email":"bob@bob.com"
			  }
		   },
		   {
			  "guid":"69c6e3b7-17b1-400b-871f-6cc585a422ad",
			  "id":"9HjwUw2Qzy4PRAhXnLDjGm",
			  "url":"https://testnet.demo.btcpayserver.org/invoice?id=9HjwUw2Qzy4PRAhXnLDjGm",
			  "btcPrice":"0.00020616",
			  "btcDue":"0.00020616",
			  "cryptoInfo":[
				 {
					"cryptoCode":"BTC",
					"paymentType":"BTCLike",
					"rate":48506.026,
					"exRates":{
					   "USD":0
					},
					"paid":"0.00000000",
					"price":"0.00020616",
					"due":"0.00020616",
					"paymentUrls":{
					   "BIP21":"bitcoin:tb1qxgsa7cjhf5d8kqjta2e6kvec04j2hsrhu5laul?amount=0.00020616",
					   "BIP72":"",
					   "BIP72b":"",
					   "BIP73":"",
					   "BOLT11":""
					},
					"address":"tb1qxgsa7cjhf5d8kqjta2e6kvec04j2hsrhu5laul",
					"url":"https://testnet.demo.btcpayserver.org/i/BTC/9HjwUw2Qzy4PRAhXnLDjGm",
					"totalDue":"0.00020616",
					"networkFee":"0.00000000",
					"txCount":0,
					"cryptoPaid":"0.00000000",
					"payments":[
					   
					]
				 }
			  ],
			  "exRates":{
				 "USD":0
			  },
			  "buyerTotalBtcAmount":"",
			  "invoiceTime":1621116053000,
			  "currentTime":1621199233681,
			  "lowFeeDetected":false,
			  "btcPaid":"0.00000000",
			  "rate":48506.026,
			  "exceptionStatus":"marked",
			  "paymentUrls":{
				 "BIP21":"bitcoin:tb1qxgsa7cjhf5d8kqjta2e6kvec04j2hsrhu5laul?amount=0.00020616",
				 "BIP72":"",
				 "BIP72b":"",
				 "BIP73":"",
				 "BOLT11":""
			  },
			  "refundAddressRequestPending":false,
			  "buyerPaidBtcMinerFee":"",
			  "bitcoinAddress":"tb1qxgsa7cjhf5d8kqjta2e6kvec04j2hsrhu5laul",
			  "flags":{
				 "refundable":false
			  },
			  "paymentSubtotals":{
				 "BTC":20616
			  },
			  "paymentTotals":{
				 "BTC":20616
			  },
			  "amountPaid":0,
			  "minerFees":{
				 "BTC":{
					"satoshisPerByte":1,
					"totalFee":0
				 }
			  },
			  "exchangeRates":{
				 "BTC":{
					"USD":0
				 }
			  },
			  "addresses":{
				 "BTC":"tb1qxgsa7cjhf5d8kqjta2e6kvec04j2hsrhu5laul"
			  },
			  "paymentCodes":{
				 "BTC":{
					"BIP21":"bitcoin:tb1qxgsa7cjhf5d8kqjta2e6kvec04j2hsrhu5laul?amount=0.00020616",
					"BIP72":"",
					"BIP72b":"",
					"BIP73":"",
					"BOLT11":""
				 }
			  },
			  "currency":"USD",
			  "price":10,
			  "orderId":"",
			  "expirationTime":1621116953000,
			  "itemDesc":"",
			  "itemCode":"",
			  "posData":"",
			  "status":"complete",
			  "supportedTransactionCurrencies":{
				 "BTC":{
					"enabled":true,
					"reason":""
				 }
			  },
			  "taxIncluded":0,
			  "token":"Tx6wmK7MtCko4rG1iydXxU",
			  "redirectAutomatically":false,
			  "notificationEmail":"",
			  "notificationURL":"",
			  "extendedNotifications":false,
			  "fullNotifications":false,
			  "buyer":{
				 "name":"",
				 "address1":"",
				 "address2":"",
				 "city":"",
				 "locality":"",
				 "region":"",
				 "postalCode":"",
				 "state":"",
				 "zip":"",
				 "country":"",
				 "phone":"",
				 "notify":false,
				 "email":""
			  }
		   },
		   {
			  "guid":"7ffc3c02-5a6b-4833-b0b3-c2355538e545",
			  "id":"JCfkzcJTfhsCKidPGxFK3J",
			  "url":"https://testnet.demo.btcpayserver.org/invoice?id=JCfkzcJTfhsCKidPGxFK3J",
			  "btcPrice":"0.00020600",
			  "btcDue":"0.00020600",
			  "cryptoInfo":[
				 {
					"cryptoCode":"BTC",
					"paymentType":"BTCLike",
					"rate":48543.735,
					"exRates":{
					   "USD":0
					},
					"paid":"0.00000000",
					"price":"0.00020600",
					"due":"0.00020600",
					"paymentUrls":{
					   "BIP21":"bitcoin:tb1q7s3nz3gwknrp7zf9e3mnceq39zapv7vtud9jte?amount=0.000206",
					   "BIP72":"",
					   "BIP72b":"",
					   "BIP73":"",
					   "BOLT11":""
					},
					"address":"tb1q7s3nz3gwknrp7zf9e3mnceq39zapv7vtud9jte",
					"url":"https://testnet.demo.btcpayserver.org/i/BTC/JCfkzcJTfhsCKidPGxFK3J",
					"totalDue":"0.00020600",
					"networkFee":"0.00000000",
					"txCount":0,
					"cryptoPaid":"0.00000000",
					"payments":[
					   
					]
				 }
			  ],
			  "exRates":{
				 "USD":0
			  },
			  "buyerTotalBtcAmount":"",
			  "invoiceTime":1621115740000,
			  "currentTime":1621199233681,
			  "lowFeeDetected":false,
			  "btcPaid":"0.00000000",
			  "rate":48543.735,
			  "exceptionStatus":"marked",
			  "paymentUrls":{
				 "BIP21":"bitcoin:tb1q7s3nz3gwknrp7zf9e3mnceq39zapv7vtud9jte?amount=0.000206",
				 "BIP72":"",
				 "BIP72b":"",
				 "BIP73":"",
				 "BOLT11":""
			  },
			  "refundAddressRequestPending":false,
			  "buyerPaidBtcMinerFee":"",
			  "bitcoinAddress":"tb1q7s3nz3gwknrp7zf9e3mnceq39zapv7vtud9jte",
			  "flags":{
				 "refundable":false
			  },
			  "paymentSubtotals":{
				 "BTC":20600
			  },
			  "paymentTotals":{
				 "BTC":20600
			  },
			  "amountPaid":0,
			  "minerFees":{
				 "BTC":{
					"satoshisPerByte":1,
					"totalFee":0
				 }
			  },
			  "exchangeRates":{
				 "BTC":{
					"USD":0
				 }
			  },
			  "addresses":{
				 "BTC":"tb1q7s3nz3gwknrp7zf9e3mnceq39zapv7vtud9jte"
			  },
			  "paymentCodes":{
				 "BTC":{
					"BIP21":"bitcoin:tb1q7s3nz3gwknrp7zf9e3mnceq39zapv7vtud9jte?amount=0.000206",
					"BIP72":"",
					"BIP72b":"",
					"BIP73":"",
					"BOLT11":""
				 }
			  },
			  "currency":"USD",
			  "price":10,
			  "orderId":"",
			  "expirationTime":1621116640000,
			  "itemDesc":"",
			  "itemCode":"",
			  "posData":"",
			  "status":"complete",
			  "supportedTransactionCurrencies":{
				 "BTC":{
					"enabled":true,
					"reason":""
				 }
			  },
			  "taxIncluded":0,
			  "token":"HwkkwFGzRQKffgcQnvXYbd",
			  "redirectAutomatically":false,
			  "notificationEmail":"",
			  "notificationURL":"",
			  "extendedNotifications":false,
			  "fullNotifications":false,
			  "buyer":{
				 "name":"",
				 "address1":"",
				 "address2":"",
				 "city":"",
				 "locality":"",
				 "region":"",
				 "postalCode":"",
				 "state":"",
				 "zip":"",
				 "country":"",
				 "phone":"",
				 "notify":false,
				 "email":""
			  }
		   }
		]
	 }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 401

	c := ptc.New(testBaseURL, kp, testToken)

	ptc.OverrideProxy(&gp)

	fmt.Println("new client: ", ptc)

	// var req InvoiceReq
	// req.Token = testToken
	// req.Currency = "USD"
	// req.Price = 10
	// req.Buyer.Name = "bob willson"
	// req.Buyer.Email = "bob@bob.com"

	var req InvoiceArgs
	req.DateStart = "2021-05-01"
	req.DateEnd = "2021-06-01"
	req.Status = "complete"
	//-------------------------for mock only
	req.OrderID = "123"
	req.Limit = "100"
	req.Offset = "0"
	//-----------------------for mock only

	resp := c.GetInvoices(&req)
	aJSON, err := json.Marshal(resp)
	fmt.Println("err: ", err)

	fmt.Println("resp: ", string(aJSON))

	if resp.Data[1].Price != 10 {
		t.Fail()
	}

	// t.Fail()
}
