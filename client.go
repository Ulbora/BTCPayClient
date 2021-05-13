package ptcpayclient

// Copyright (c) 2018 Ulbora Labs LLC
// Copyright (c) 2018 Ken Williamson

const (
	userAgent = "Ulbora btyPayClient"
)

// go mod init github.com/Ulbora/BTCPayClient

//Headers Headers
type Headers struct {
	headers map[string]string
	//mu      sync.Mutex
}

//Set Set
func (h *Headers) Set(key string, value string) {
	//h.mu.Lock()
	//defer h.mu.Unlock()
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[key] = value
}

//PairClientResponse PairClientResponse
type PairClientResponse struct {
	Merchant string `json:"merchant"`
}

//TokenRequest TokenRequest
type TokenRequest struct {
	ID          string `json:"id"`
	Facade      string `json:"facade"`
	Label       string `json:"label"`
	PairingCode string `json:"pairingCode"`
}

//TokenResponse TokenResponse
type TokenResponse struct {
	Data []TokenData `json:"data"`
	Code int64       `json:"code"`
}

//Client Client
type Client interface {
	GetClientID() string
	Token(req *TokenRequest) *TokenResponse
	PairClient(code string) *TokenResponse
	GetPairingCodeRequest(code string) string
	GetRates(currencyPairs []string, storeID string) *RateResponse
	//CreateInvoice(inv *InvoiceReq) *Invoice
	//GetInvoice(invoiceId string, token string) *Invoice
	//GetInvoices(args IncoiceArgs, token string) *[]Invoice
}
