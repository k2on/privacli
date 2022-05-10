package main

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type CardList struct {
	Cards        []Card `json:"data"`
	TotalEntries int    `json:"total_entries"`
	TotalPages   int    `json:"total_pages"`
	Page         int    `json:"page"`
}

type TransactionList struct {
	Transactions []Transaction `json:"data"`
	TotalEntries int           `json:"total_entries"`
	TotalPages   int           `json:"total_pages"`
	Page         int           `json:"page"`
}

type Card struct {
	Created            string `json:"created"`
	Token              string `json:"token"`
	LastFour           string `json:"last_four"`
	HostName           string `json:"hostname"`
	Memo               string `json:"memo"`
	Type               string `json:"type"`
	SpendLimit         string `json:"spend_limit"`
	SpendLimitDuration Duration `json:"spend_limit_duration"`
	State              State `json:"state"`
	Funding            Funding `json:"funding"`
	Pan     string `json:"pan"`
	CVV     string `json:"cvv"`
	ExpMonth     string `json:"exp_month"`
	ExpYear     string `json:"exp_year"`
}

type Type string
const (
	TypeDigitalWallet Type = "DIGITAL_WALLET"
	TypeMerchantLocked     = "MERCHANT_LOCKED"
	TypeSingleUse          = "SINGLE_USE"
	TypeUnlocked           = "UNLOCKED"
)

type Duration string
const (
	DurationAnnually Duration = "ANNUALLY"
	DurationForever           = "FOREVER"
	DurationMonthly           = "MONTHLY"
	DurationTransaction       = "TRANSACTION"
)

type State string
const (
	StateOpen State = "OPEN"
	StatePaused     = "PAUSED"
	StateClosed     = "CLOSED"
)

type Funding struct {
	Created     string `json:"created"`
	Token       string `json:"token"`
	Type        string `json:"type"`
	State       string `json:"state"`
	NickName    string `json:"nickname"`
	AccountName string `json:"account_name"`
	LastFour    string `json:"last_four"`
}

type Transaction struct {
	Card                        Card       `json:"card"`
	Amount                      int        `json:"amount"`
	AuthorizationAmount         int        `json:"authorization_amount"`
	MerchantAmount              int        `json:"merchant_amount"`
	MerchantAuthorizationAmount int        `json:"merchant_authorization_amount"`
	MerchantCurrency            string     `json:"merchant_currency"`
	AquirerFee                  int        `json:"aquirer_fee"`
	Created                     string     `json:"created"`
	Events                      []Event    `json:"events"`
	Funding                     []FundedBy `json:"funding"`
	Merchant                    Merchant   `json:"merchant"`
	Network                     string     `json:"network"`
	Result                      Result     `json:"result"`
	SettledAmount               string     `json:"settled_amount"`
	Status                      string     `json:"status"`
	Token                       string     `json:"token"`
	AuthorizationCode           string     `json:"authorization_code"`
}

type Event struct {
	Amount  int    `json:"amount"`
	Created string `json:"created"`
	Result  string `json:"result"`
	Type    string `json:"type"`
	Token   string `json:"token"`
}

type FundedBy struct {
	Amount int    `json:"amount"`
	Token  string `json:"token"`
	Type   string `json:"type"`
}

type Merchant struct {
	AcceptorId string `json:"acceptor_id"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Descriptor string `json:"descriptor"`
	MCC        string `json:"mcc"`
	State      string `json:"state"`
}

type Result string
const (
	ResultApproved Result = "APPROVED"
	ResultDeclined Result = "DECLINED"
)