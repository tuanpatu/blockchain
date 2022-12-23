package services

import (
	"fmt"
	"log"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/go-bip39"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"gorm.io/gorm"

	"blockchain/pkg/models"
)

type WalletsServiceImpl struct {
	DB *gorm.DB
}

func NewWalletsService(DB *gorm.DB) WalletsService {
	return &WalletsServiceImpl{
		DB: DB,
	}
}
func (r *WalletsServiceImpl) GetWallets() (error, []*models.Wallets) {
	WalletsModels := []*models.Wallets{}

	err := r.DB.Find(&WalletsModels).Error

	if err != nil {
		return err, nil
	}

	return nil, WalletsModels
}

func (r *WalletsServiceImpl) GetWallet(id *string) (error, *models.Wallets) {
	wallet := models.Wallets{}
	err := r.DB.First(&wallet, "id = ?", id).Error

	if err != nil {
		return err, nil
	}

	return err, &wallet
}

func (r *WalletsServiceImpl) CreateWallet(userId *string) (error, *models.Wallets) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "")
	master, ch := hd.ComputeMastersFromSeed(seed)
	path := "m/44'/60'/0'/0/0'"
	private, _ := hd.DerivePrivateKeyForPath(master, ch, path)
	privateKey := secp256k1.PrivKey(private)

	fmt.Println("seedPhrase:", mnemonic)
	fmt.Println("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKey.Bytes()))

	publicKey := privateKey.PubKey()
	fmt.Println("Public Key:", hexutil.Encode(publicKey.Bytes()))

	address := publicKey.Address()
	fmt.Println("Address:", hexutil.Encode(address.Bytes()))

	wallet := models.Wallets{
		Mnemonic:   mnemonic,
		PrivateKey: hexutil.Encode(privateKey.Bytes()),
		PublicKey:  hexutil.Encode(publicKey.Bytes()),
		Address:    hexutil.Encode(address.Bytes()),
	}

	err = r.DB.Create(&wallet).Error

	if err != nil {
		return err, nil
	}

	return err, &wallet
}

type WalletsService interface {
	CreateWallet(*string) (error, *models.Wallets)
	GetWallet(*string) (error, *models.Wallets)
	GetWallets() (error, []*models.Wallets)
}
