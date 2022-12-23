package models

import "gorm.io/gorm"

//Wallets struct
type Wallets struct {
	gorm.Model
	// ID         uint   `gorm:"primary key;autoIncrement" json:"id"`
	UserId     uint   `json:"user_id" xml:"user_id" form:"user_id"`
	Mnemonic   string `json:"mnemonic" xml:"mnemonic" form:"mnemonic"`
	PrivateKey string `json:"private_key" xml:"private_key" form:"private_key"`
	PublicKey  string `json:"public_key" xml:"public_key" form:"public_key"`
	Address    string `json:"address" xml:"address" form:"address"`

	WalletTokens []WalletTokens `gorm:"foreignKey:WalletId"`
}
