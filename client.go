package ptcpayclient

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

type PairClientResponse struct {
	Merchant string `json:"merchant"`
}

type Client interface {
	//Token(sin string)
	PairClient(code string) *PairClientResponse
	//GetRates(currencyPairs []string, storeID string) *[]Rate
	//CreateInvoice(inv *InvoiceReq) *Invoice
	//GetInvoice(invoiceId string, token string) *Invoice
	//GetInvoices(args IncoiceArgs, token string) *[]Invoice
}
