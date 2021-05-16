package ptcpayclient

//***********************************************
//* Copyright (c) 2021 Ulbora Labs LLC
//* Copyright (c) 2021 Ken Williamson
//***********************************************

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"net/http"
)

//CreateInvoice CreateInvoice
func (a *BTCPayClient) CreateInvoice(inv *InvoiceReq) *InvoiceResponse {
	var rtn InvoiceResponse

	var url = a.host + "/invoices"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(inv)
	bodyStr := string(aJSON)
	a.log.Debug("bodyStr: ", bodyStr)
	if err == nil {
		var headers Headers
		urlb := []byte(url + bodyStr)
		signVal, _ := a.crypto.Sign(urlb, a.kp)
		headers.Set("x-identity", a.crypto.GetPublicKey(a.kp))
		headers.Set("x-signature", hex.EncodeToString(signVal))

		a.log.Debug("headers: ", headers)
		req := a.buildRequest(http.MethodPost, url, headers, aJSON)
		suc, stat := a.proxy.Do(req, &rtn) //--------------------
		// //test------------------------
		// client := &http.Client{}
		// resp, err := client.Do(req)
		// fmt.Println("client err: ", err)
		// defer resp.Body.Close()
		// stat := resp.StatusCode
		// decoder := json.NewDecoder(resp.Body)
		// bodyBytes, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// bodyString := string(bodyBytes)
		// fmt.Println("body: ", bodyString)
		// error := decoder.Decode(&rtn)
		// var suc = true
		// fmt.Println("error: ", error)
		// //-------------------

		a.log.Debug("suc: ", suc)
		a.log.Debug("stat: ", stat)
		//rtn.Code = int64(stat)
		if !suc || stat != http.StatusOK {
			a.log.Debug("proxy call failed to : ", url)
		}
	}
	a.log.Debug("rtn: ", rtn)

	return &rtn
}

//GetInvoice GetInvoice
func (a *BTCPayClient) GetInvoice(invoiceID string) *InvoiceResponse {
	var rtn InvoiceResponse
	var url = a.host + "/invoices/" + invoiceID + "?token=" + a.tokens.Token
	a.log.Debug("url: ", url)
	var headers Headers
	urlbi := []byte(url)
	signVal, _ := a.crypto.Sign(urlbi, a.kp)
	headers.Set("x-identity", a.crypto.GetPublicKey(a.kp))
	headers.Set("x-signature", hex.EncodeToString(signVal))
	fmt.Println("headers: ", headers)
	reqi := a.buildRequest(http.MethodGet, url, headers, nil)
	suc, stat := a.proxy.Do(reqi, &rtn) //--------------------
	// //test------------------------
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// fmt.Println("client err: ", err)
	// defer resp.Body.Close()
	// stat := resp.StatusCode
	// decoder := json.NewDecoder(resp.Body)
	// bodyBytes, erri := ioutil.ReadAll(resp.Body)
	// if erri != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// fmt.Println("body: ", bodyString)
	// error := decoder.Decode(&rtn)
	// var suc = true
	// fmt.Println("error: ", error)
	// //-------------------

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	if !suc || stat != http.StatusOK {
		a.log.Debug("proxy call failed to : ", url)
	}

	a.log.Debug("rtn: ", rtn)

	return &rtn
}

//GetInvoices GetInvoices
func (a *BTCPayClient) GetInvoices(args *InvoiceArgs) *InvoiceListResponse {
	var rtn InvoiceListResponse
	var url = a.host + "/invoices?token=" + a.tokens.Token
	if args.Status != "" {
		url += "&status=" + args.Status
	}
	if args.DateStart != "" {
		url += "&dateStart=" + args.DateStart
	}
	if args.DateEnd != "" {
		url += "&dateEnd=" + args.DateEnd
	}
	if args.OrderID != "" {
		url += "&orderId=" + args.OrderID
	}

	if args.Limit != "" {
		url += "&limit=" + args.Limit
	}

	if args.Offset != "" {
		url += "&offset=" + args.Offset
	}

	a.log.Debug("url: ", url)
	var headers Headers
	urlbis := []byte(url)
	signVal, _ := a.crypto.Sign(urlbis, a.kp)
	headers.Set("x-identity", a.crypto.GetPublicKey(a.kp))
	headers.Set("x-signature", hex.EncodeToString(signVal))
	fmt.Println("headers: ", headers)
	reqis := a.buildRequest(http.MethodGet, url, headers, nil)
	suc, stat := a.proxy.Do(reqis, &rtn) //--------------------
	// //test------------------------
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// fmt.Println("client err: ", err)
	// defer resp.Body.Close()
	// stat := resp.StatusCode
	// decoder := json.NewDecoder(resp.Body)
	// bodyBytes, erris := ioutil.ReadAll(resp.Body)
	// if erris != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// fmt.Println("body: ", bodyString)
	// error := decoder.Decode(&rtn)
	// var suc = true
	// fmt.Println("error: ", error)
	// //-------------------

	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)

	if !suc || stat != http.StatusOK {
		a.log.Debug("proxy call failed to : ", url)
	}

	a.log.Debug("rtn: ", rtn)
	return &rtn
}
