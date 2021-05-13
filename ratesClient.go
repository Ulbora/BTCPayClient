package ptcpayclient

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

//GetRates GetRates
func (a *BTCPayClient) GetRates(currencyPairs []string, storeID string) *RateResponse {
	var rtn RateResponse
	args := strings.Join(currencyPairs[:], ",")

	var url = a.host + "/rates?currencyPairs=" + args + "&storeID=" + storeID + "&token=" + a.tokens.Token
	a.log.Debug("url: ", url)
	//aJSON, err := json.Marshal(req)
	//if err == nil {
	var headers Headers
	//create signed headers
	urlb := []byte(url)
	signVal, _ := a.crypto.Sign(urlb, a.kp)
	headers.Set("X-Identity", a.crypto.GetPublicKey(a.kp))
	headers.Set("X-Signature", hex.EncodeToString(signVal))
	fmt.Println("headers: ", headers)
	req := a.buildRequest(http.MethodGet, url, headers, nil)
	suc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", suc)
	a.log.Debug("stat: ", stat)
	//rtn.Code = int64(stat)
	if !suc || stat != http.StatusOK {
		a.log.Debug("proxy call failed to : ", url)
	}
	//}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}
