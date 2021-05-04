package ptcpayclient

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"math/big"
	"sync"
	"time"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// go mod init github.com/Ulbora/BTCPayClient

//Headers Headers
type Headers struct {
	headers map[string]string
	mu      sync.Mutex
}

//Set Set
func (h *Headers) Set(key string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[key] = value
}

type BTCPayClient struct {
	clientID  string
	userAgent string
	host      string
	proxy     px.Proxy
	log       *lg.Logger
	headers   *Headers
}

type PairClientResponse struct {
	Merchant string `json:"merchant"`
}

type Payload struct {
	ClientID    string `json:"id"`
	PairingCode string `json:"pairingCode"`
}

type Rate struct {
	Name         string  `json:"name"`
	CryptoCode   string  `json:"cryptoCode"`
	CurrencyPair string  `json:"currencyPair"`
	Code         string  `json:"code"`
	Rate         float64 `json:"rate"`
}

type TranCurStatus struct {
	Enabled bool
}

type Buyer struct {
	Name       string `json:"name"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	Locality   string `json:"locality"`
	Region     string `json:"region"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
	Zip        string `json:"zip"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	Notify     bool   `json:"notify"`
	Email      string `json:"email"`
}

type BuyerFields struct {
	BuyerName     string `json:"buyerName"`
	BuyerAddress1 string `json:"buyerAddress1"`
	BuyerAddress2 string `json:"buyerAddress2"`
	BuyerCity     string `json:"buyerCity"`
	BuyerState    string `json:"buyerState"`
	BuyerZip      string `json:"buyerZip"`
	BuyerCountry  string `json:"buyerCountry"`
	BuyerPhone    string `json:"buyerPhone"`
	BuyerNotify   bool   `json:"buyerNotify"`
	BuyerEmail    string `json:"buyerEmail"`
}

type CryptoCode struct {
	CryptoCode  string    `json:"cryptoCode"`
	PaymentType string    `json:"paymentType"`
	Rate        float64   `json:"rate"`
	ExRates     []float64 `json:"exRates"`
	Paid        string    `json:"paid"`
	Price       string    `json:"price"`
	Due         string    `json:"due"`
	PaymentUrls []string  `json:"paymentUrls"`
	Address     string    `json:"address"`
	Url         string    `json:"url"`
	TotalDue    string    `json:"totalDue"`
	NetworkFee  string    `json:"networkFee"`
	TxCount     int64     `json:"txCount"`
	CryptoPaid  string    `json:"cryptoPaid"`
	Payments    []string  `json:"payments"`
}

type PayURLs struct {
	BIP21  string `json:"BIP21"`
	BIP72  string `json:"BIP72"`
	BIP72b string `json:"BIP72b"`
	BIP73  string `json:"BIP73"`
	BOLT11 string `json:"BOLT11"`
}

type InvFlags struct {
	Refundable bool `json:"refundable"`
}

type InvoiceReq struct {
	//Guid                           string                   `json:"guid"`
	//ID                             string                   `json:"id"`
	//URL                            string                   `json:"url"`
	//BtcPrice                       string                   `json:"btcPrice"`
	//BtcDue                         string                   `json:"btcDue"`
	//CryptoInfo                     []CryptoCode             `json:"cryptoInfo"`
	//ExRates                        float64                  `json:"exRates"`
	//BuyerTotalBtcAmount            string                   `json:"buyerTotalBtcAmount"`
	//InvoiceTime                    time.Time                `json:"invoiceTime"`
	//CurrentTime                    time.Time                `json:"currentTime"`
	//LowFeeDetected                 bool                     `json:"lowFeeDetected"`
	//BtcPaid                        string                   `json:"btcPaid"`
	//Rate                           float64                  `json:"rate"`
	//ExceptionStatus                string                   `json:"exceptionStatus"`
	//PaymentUrls                    PayURLs                  `json:"paymentUrls"`
	//RefundAddressRequestPending    bool                     `json:"refundAddressRequestPending"`
	//BuyerPaidBtcMinerFee           string                   `json:"buyerPaidBtcMinerFee"`
	//BitcoinAddress                 string                   `json:"bitcoinAddress"`
	//Flags                          InvFlags                 `json:"flags"`
	//PaymentSubtotals               float64                  `json:"paymentSubtotals"`
	//PaymentTotals                  float64                  `json:"paymentTotals"`
	// AmountPaid                     float64                  `json:"amountPaid"`
	// MinerFees                      float64                  `json:"minerFees"`
	// ExchangeRates                  float64                  `json:"exchangeRates"`
	// Addresses                      string                   `json:"addresses"`
	// PaymentCodes                   string                   `json:"paymentCodes"`
	Currency                       string                   `json:"currency"`
	Price                          float64                  `json:"price"`
	OrderId                        string                   `json:"orderId"`
	ExpirationTime                 time.Time                `json:"expirationTime"`
	ItemDesc                       string                   `json:"itemDesc"`
	ItemCode                       string                   `json:"itemCode"`
	PosData                        string                   `json:"posData"`
	Status                         string                   `json:"status"`
	RedirectURL                    string                   `json:"redirectURL"`
	TransactionSpeed               string                   `json:"transactionSpeed"`
	Physical                       bool                     `json:"physical"`
	SupportedTransactionCurrencies map[string]TranCurStatus `json:"supportedTransactionCurrencies"`
	Refundable                     bool                     `json:"refundable"`
	TaxIncluded                    float64                  `json:"taxIncluded"`
	Token                          string                   `json:"token"`
	RedirectAutomatically          bool                     `json:"redirectAutomatically"`
	NotificationEmail              string                   `json:"notificationEmail"`
	NotificationURL                string                   `json:"notificationURL"`
	ExtendedNotifications          bool                     `json:"extendedNotifications"`
	FullNotifications              bool                     `json:"fullNotifications"`
	BuyerFields                    BuyerFields              `json:"buyerFields"`
	//Buyer                          Buyer                    `json:"buyer"`
}

type IncoiceArgs struct {
	Status    string `json:"status"`
	OrderId   string `json:"orderId"`
	ItemCode  string `json:"itemCode"`
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
}

type Invoice struct {
	Guid                        string       `json:"guid"`
	ID                          string       `json:"id"`
	URL                         string       `json:"url"`
	BtcPrice                    string       `json:"btcPrice"`
	BtcDue                      string       `json:"btcDue"`
	CryptoInfo                  []CryptoCode `json:"cryptoInfo"`
	ExRates                     float64      `json:"exRates"`
	BuyerTotalBtcAmount         string       `json:"buyerTotalBtcAmount"`
	InvoiceTime                 time.Time    `json:"invoiceTime"`
	CurrentTime                 time.Time    `json:"currentTime"`
	LowFeeDetected              bool         `json:"lowFeeDetected"`
	BtcPaid                     string       `json:"btcPaid"`
	Rate                        float64      `json:"rate"`
	ExceptionStatus             string       `json:"exceptionStatus"`
	PaymentUrls                 PayURLs      `json:"paymentUrls"`
	RefundAddressRequestPending bool         `json:"refundAddressRequestPending"`
	BuyerPaidBtcMinerFee        string       `json:"buyerPaidBtcMinerFee"`
	BitcoinAddress              string       `json:"bitcoinAddress"`
	Flags                       InvFlags     `json:"flags"`
	PaymentSubtotals            float64      `json:"paymentSubtotals"`
	PaymentTotals               float64      `json:"paymentTotals"`
	AmountPaid                  float64      `json:"amountPaid"`
	MinerFees                   float64      `json:"minerFees"`
	ExchangeRates               float64      `json:"exchangeRates"`
	Addresses                   string       `json:"addresses"`
	PaymentCodes                string       `json:"paymentCodes"`
	Currency                    string       `json:"currency"`
	Price                       float64      `json:"price"`
	OrderId                     string       `json:"orderId"`
	ExpirationTime              time.Time    `json:"expirationTime"`
	ItemDesc                    string       `json:"itemDesc"`
	ItemCode                    string       `json:"itemCode"`
	PosData                     string       `json:"posData"`
	Status                      string       `json:"status"`
	//RedirectURL                    string                   `json:"redirectURL"`
	//TransactionSpeed               string                   `json:"transactionSpeed"`
	//Physical                       bool                     `json:"physical"`
	SupportedTransactionCurrencies map[string]TranCurStatus `json:"supportedTransactionCurrencies"`
	//Refundable                     bool                     `json:"refundable"`
	TaxIncluded           float64 `json:"taxIncluded"`
	Token                 string  `json:"token"`
	RedirectAutomatically bool    `json:"redirectAutomatically"`
	NotificationEmail     string  `json:"notificationEmail"`
	NotificationURL       string  `json:"notificationURL"`
	ExtendedNotifications bool    `json:"extendedNotifications"`
	FullNotifications     bool    `json:"fullNotifications"`
	//BuyerFields                    BuyerFields              `json:"buyerFields"`
	Buyer Buyer `json:"buyer"`
}

type Client interface {
	PairClient(code string) *PairClientResponse
	GetRates(currencyPairs []string, storeID string) *[]Rate
	CreateInvoice(inv *InvoiceReq) *Invoice
	GetInvoice(invoiceId string, token string) *Invoice
	GetInvoices(args IncoiceArgs, token string) *[]Invoice
}

type Crypto interface {
	GenerateKeyPair() *ecdsa.PrivateKey
	LoadKeyPair(privateKey string) *ecdsa.PrivateKey
	GetSinFromKey()
	Sign(hash []byte, kp *ecdsa.PrivateKey) []byte
}

type Cryptography struct {
}

func (c *Cryptography) GenerateKeyPair() *ecdsa.PrivateKey {
	kp, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//priva, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//privb, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	//puba := priva.PublicKey
	//pubb := privb.PublicKey
	//prvk :=
	return kp
}

func (c *Cryptography) LoadKeyPair(privateKey string) *ecdsa.PrivateKey {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return &e
}

func (c *Cryptography) Sign(hash []byte, kp *ecdsa.PrivateKey) []byte {
	sig, err := ecdsa.SignASN1(rand.Reader, kp, hash[:])
	if err != nil {
		log.Println(err)
	}
	return sig

}
