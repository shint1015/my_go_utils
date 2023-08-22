package utils

import (
	"github.com/shopspring/decimal"
)

// 共通構造体

type FormatOrder struct {
	Symbol      string
	OrderType   string
	OrderSide   string
	FromCoin    string
	ToCoin      string
	OrderRate   decimal.Decimal
	OrderFee    decimal.Decimal
	FromAmount  decimal.Decimal
	ToAmount    decimal.Decimal
	OrderStatus string
}

type DepositHistory struct {
	Symbol     string
	CreateAt   int
	Amount     string
	WalletTxId string
	Status     string
	IsInner    bool
}
