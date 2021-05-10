package ptcpayclient

import (
	"encoding/json"
	"net/http"
)

// Copyright (c) 2018 Ulbora Labs LLC
// Copyright (c) 2018 Ken Williamson

//Token Token
func (a *BTCPayClient) Token(req *TokenRequest) *TokenResponse {
	var rtn TokenResponse

	var url = a.host + "/tokens"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(req)
	if err == nil {
		var headers Headers
		req := a.buildRequest(http.MethodPost, url, headers, aJSON)
		suc, stat := a.proxy.Do(req, &rtn)
		a.log.Debug("suc: ", suc)
		a.log.Debug("stat: ", stat)
		rtn.Code = int64(stat)
		if !suc || stat != http.StatusOK {
			a.log.Debug("proxy call failed to : ", url)
		}
	}
	a.log.Debug("rtn: ", rtn)
	a.tokens = &rtn.Data
	return &rtn

}

//PairClient PairClient
func (a *BTCPayClient) PairClient(code string) *TokenResponse {
	a.log.Debug("pairing with code: ", code)

	var tkr TokenRequest
	tkr.ID = a.clientID
	tkr.PairingCode = code
	a.log.Debug("PairClient req: ", tkr)
	return a.Token(&tkr)
}

//GetPairingCodeRequest GetPairingCodeRequest
func (a *BTCPayClient) GetPairingCodeRequest(code string) string {
	var url = a.host + "/api-access-request?pairingCode=" + code
	a.log.Debug("pair request url: ", url)
	return url
}
