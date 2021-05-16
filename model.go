package ptcpayclient

//***********************************************
//* Copyright (c) 2021 Ulbora Labs LLC
//* Copyright (c) 2021 Ken Williamson
//***********************************************

//Policy Policy
type Policy struct {
	Policy string   `json:"policy"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

//TokenData TokenData
type TokenData struct {
	Policies          []Policy `json:"policies"`
	Token             string   `json:"token"`
	Facade            string   `json:"facade"`
	CreateDate        int64    `json:"dateCreated"`
	PairingExpiration int64    `json:"pairingExpiration"`
	ParingCode        string   `json:"pairingCode"`
}

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

//RateResponse RateResponse
type RateResponse struct {
	Data []Rate `json:"data"`
}

//TranCurStatus TranCurStatus
type TranCurStatus struct {
	Enabled bool   `json:"enabled"`
	Reason  string `json:"reason"`
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

//Erate Erate
type Erate struct {
}

//CryptoCode CryptoCode
type CryptoCode struct {
	CryptoCode  string             `json:"cryptoCode"`
	PaymentType string             `json:"paymentType"`
	Rate        float64            `json:"rate"`
	ExRates     map[string]float64 `json:"exRates"`
	Paid        string             `json:"paid"`
	Price       string             `json:"price"`
	Due         string             `json:"due"`
	PaymentUrls map[string]string  `json:"paymentUrls"`
	Address     string             `json:"address"`
	URL         string             `json:"url"`
	TotalDue    string             `json:"totalDue"`
	NetworkFee  string             `json:"networkFee"`
	TxCount     int64              `json:"txCount"`
	CryptoPaid  string             `json:"cryptoPaid"`
	Payments    []Payment          `json:"payments"`
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
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
	OrderID  string  `json:"orderId"`
	//ExpirationTime                 int64                    `json:"expirationTime"`
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
	Buyer                          Buyer                    `json:"buyer"`
}

//InvoiceArgs InvoiceArgs
type InvoiceArgs struct {
	Status    string `json:"status"`
	OrderID   string `json:"orderId"`
	ItemCode  string `json:"itemCode"`
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
}

//MinerFee MinerFee
type MinerFee struct {
	SatoshisPerByte float64 `json:"satoshisPerByte"`
	TotalFee        float64 `json:"totalFee"`
}

//Payment Payment
type Payment struct {
	ID           string  `json:"id"`
	ReceivedDate string  `json:"receivedDate"`
	Value        float64 `json:"value"`
	Fee          float64 `json:"fee"`
	PaymentType  string  `json:"paymentType"`
	Confirmed    bool    `json:"confirmed"`
	Completed    bool    `json:"completed"`
	Destination  string  `json:"destination"`
}

//Invoice Invoice
type Invoice struct {
	GUID                        string                        `json:"guid"`
	ID                          string                        `json:"id"`
	URL                         string                        `json:"url"`
	BtcPrice                    string                        `json:"btcPrice"`
	BtcDue                      string                        `json:"btcDue"`
	CryptoInfo                  []CryptoCode                  `json:"cryptoInfo"`
	ExRates                     map[string]float64            `json:"exRates"`
	BuyerTotalBtcAmount         string                        `json:"buyerTotalBtcAmount"`
	InvoiceTime                 int64                         `json:"invoiceTime"`
	CurrentTime                 int64                         `json:"currentTime"`
	LowFeeDetected              bool                          `json:"lowFeeDetected"`
	BtcPaid                     string                        `json:"btcPaid"`
	Rate                        float64                       `json:"rate"`
	ExceptionStatus             interface{}                   `json:"exceptionStatus"`
	PaymentUrls                 PayURLs                       `json:"paymentUrls"`
	RefundAddressRequestPending bool                          `json:"refundAddressRequestPending"`
	BuyerPaidBtcMinerFee        string                        `json:"buyerPaidBtcMinerFee"`
	BitcoinAddress              string                        `json:"bitcoinAddress"`
	Flags                       InvFlags                      `json:"flags"`
	PaymentSubtotals            map[string]float64            `json:"paymentSubtotals"`
	PaymentTotals               map[string]float64            `json:"paymentTotals"`
	AmountPaid                  float64                       `json:"amountPaid"`
	MinerFees                   map[string]MinerFee           `json:"minerFees"`
	ExchangeRates               map[string]map[string]float64 `json:"exchangeRates"`
	Addresses                   map[string]string             `json:"addresses"`
	PaymentCodes                map[string]map[string]string  `json:"paymentCodes"`
	Currency                    string                        `json:"currency"`
	Price                       float64                       `json:"price"`
	OrderID                     string                        `json:"orderId"`
	ExpirationTime              int64                         `json:"expirationTime"`
	ItemDesc                    string                        `json:"itemDesc"`
	ItemCode                    string                        `json:"itemCode"`
	PosData                     string                        `json:"posData"`
	Status                      string                        `json:"status"`
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

//InvoiceResponse InvoiceResponse
type InvoiceResponse struct {
	Data Invoice `json:"data"`
}

//InvoiceListResponse InvoiceListResponse
type InvoiceListResponse struct {
	Data []Invoice `json:"data"`
}
