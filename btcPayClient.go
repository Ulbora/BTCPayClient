package ptcpayclient

import (
	"bytes"
	"crypto/ecdsa"
	"net/http"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

type BTCPayClient struct {
	clientID  string
	userAgent string
	host      string
	kp        *ecdsa.PrivateKey
	proxy     px.Proxy
	log       *lg.Logger
	headers   Headers
}

func (a *BTCPayClient) New(host string, kp *ecdsa.PrivateKey) Client {
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	a.log = &l

	var p px.GoProxy
	a.proxy = &p

	a.kp = kp

	var cryt Cryptography
	c := cryt.New()

	a.clientID = c.GetSinFromKey(a.kp)
	a.userAgent = userAgent
	return a
}

//OverrideProxy OverrideProxy
func (a *BTCPayClient) OverrideProxy(proxy px.Proxy) {
	a.proxy = proxy
}

//SetHeader SetHeader
func (a *BTCPayClient) SetHeader(head Headers) {
	a.headers = head
}

//SetLogLevel SetLogLevel
func (a *BTCPayClient) SetLogLevel(level int) {
	a.log.LogLevel = level
}

func (a *BTCPayClient) buildRequest(method string, url string, headers Headers, aJSON []byte) *http.Request {

	var req *http.Request
	var err error
	if method == http.MethodPost || method == http.MethodPut {
		headers.Set("Content-Type", "application/json")
		req, err = http.NewRequest(method, url, bytes.NewBuffer(aJSON))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	a.log.Debug("err in build req: ", err)
	if err == nil {
		for k, v := range headers.headers {
			a.log.Debug("header: ", k, v)
			req.Header.Set(k, v)
		}
	}
	return req
}
