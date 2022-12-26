package models

import (
	"gorm.io/gorm"
)

type WalletTokens struct {
	gorm.Model
	// ID      uint      `gorm:"primary key;autoIncrement" json:"id"`
	WalletId uint   `json:"wallet_id" xml:"wallet_id" form:"wallet_id"`
	Name     string `json:"name" xml:"name" form:"name"`
	Symbol   string `json:"symbol" xml:"symbol" form:"symbol"`
	Amount   uint   `json:"amount" xml:"amount" form:"amount"`
}
