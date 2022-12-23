package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Users struct
type Users struct {
	gorm.Model
	// ID      uint      `gorm:"primary key;autoIncrement" json:"id"`
	Email    string  `gorm:"not null;unique" json:"email" xml:"email" form:"email"`
	Username string  `gorm:"not null;unique" json:"username" xml:"username" form:"username"`
	Password string  `gorm:"size:255;not null;" json:"password" xml:"password" form:"password"`
	Wallet   Wallets `gorm:"foreignKey:UserId"`
}

func (u *Users) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("Password", hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func (u *Users) PrepareResponse() {
	u.Password = "******"
}
