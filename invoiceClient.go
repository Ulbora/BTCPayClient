package ptcpayclient

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
		fmt.Println("headers: ", headers)
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
	// if len(rtn.Data) > 0 {
	// 	a.tokens = rtn.Data[0]
	// }

	return &rtn
}
