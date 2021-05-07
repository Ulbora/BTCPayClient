package ptcpayclient

// Copyright (c) 2018 Ulbora Labs LLC
// Copyright (c) 2018 Ken Williamson

import "time"

//Payload Payload
type Payload struct {
	ClientID    string `json:"id"`
	PairingCode string `json:"pairingCode"`
}

//Rate Rate
type Rate struct {
	Name         string  `json:"name"`
	CryptoCode   string  `json:"cryptoCode"`
	CurrencyPair string  `json:"currencyPair"`
	Code         string  `json:"code"`
	Rate         float64 `json:"rate"`
}

//TranCurStatus TranCurStatus
type TranCurStatus struct {
	Enabled bool
}

//Buyer Buyer
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

//BuyerFields BuyerFields
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

//CryptoCode CryptoCode
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
	URL         string    `json:"url"`
	TotalDue    string    `json:"totalDue"`
	NetworkFee  string    `json:"networkFee"`
	TxCount     int64     `json:"txCount"`
	CryptoPaid  string    `json:"cryptoPaid"`
	Payments    []string  `json:"payments"`
}

//PayURLs PayURLs
type PayURLs struct {
	BIP21  string `json:"BIP21"`
	BIP72  string `json:"BIP72"`
	BIP72b string `json:"BIP72b"`
	BIP73  string `json:"BIP73"`
	BOLT11 string `json:"BOLT11"`
}

//InvFlags InvFlags
type InvFlags struct {
	Refundable bool `json:"refundable"`
}

//InvoiceReq InvoiceReq
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
	OrderID                        string                   `json:"orderId"`
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

//IncoiceArgs IncoiceArgs
type IncoiceArgs struct {
	Status    string `json:"status"`
	OrderID   string `json:"orderId"`
	ItemCode  string `json:"itemCode"`
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
}

//Invoice Invoice
type Invoice struct {
	GUID                        string       `json:"guid"`
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
	OrderID                     string       `json:"orderId"`
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
