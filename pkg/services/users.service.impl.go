package services

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/go-bip39"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"blockchain/pkg/models"
	"blockchain/pkg/utils/token"
)

type UsersServiceImpl struct {
	DB *gorm.DB
}

func NewUsersService(DB *gorm.DB) UsersService {
	return &UsersServiceImpl{
		DB: DB,
	}
}

func (r *UsersServiceImpl) GetUsers() (error, []*models.Users) {
	userModels := []*models.Users{}

	err := r.DB.Preload(clause.Associations).Find(&userModels).Error

	if err != nil {
		return err, nil
	}

	return nil, userModels
}

func (r *UsersServiceImpl) GetUser(id *string) (error, *models.Users) {
	user := models.Users{}
	err := r.DB.Preload("Wallet").First(&user, "id = ?", id).Error

	if err != nil {
		return err, nil
	}

	user.PrepareResponse()

	return err, &user
}

func (r *UsersServiceImpl) GetUserByAddress(address *string) (error, *models.Users) {
	user := models.Users{}
	wallet := models.Wallets{}
	err := r.DB.First(&wallet, "address = ?", address).Error

	err = r.DB.Preload("Wallet").First(&user, "id = ?", wallet.UserId).Error

	if err != nil {
		return err, nil
	}

	user.PrepareResponse()

	return err, &user
}

func (r *UsersServiceImpl) CreateUser(User *models.Users) (error, *models.Users) {
	err := r.DB.Create(&User).Error

	if err != nil {
		return err, nil
	}

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
	publicKey := privateKey.PubKey()
	address := publicKey.Address()

	wallet := models.Wallets{
		UserId:     User.ID,
		Mnemonic:   mnemonic,
		PrivateKey: hexutil.Encode(privateKey.Bytes()),
		PublicKey:  hexutil.Encode(publicKey.Bytes()),
		Address:    hexutil.Encode(address.Bytes()),
	}

	err = r.DB.Create(&wallet).Error

	if err != nil {
		return err, nil
	}

	err = r.DB.Preload("Wallet").First(&User, "id = ?", User.ID).Error
	User.PrepareResponse()

	return err, User
}

func (r *UsersServiceImpl) DeleteUser(id *string) error {
	User := models.Users{}

	err := r.DB.Delete(&User, id).Error

	if err != nil {
		return errors.New("could not delete User")
	}

	return err
}

func (r *UsersServiceImpl) UpdateUser(User *models.Users, id *string) (error, *models.Users) {
	err := r.DB.Model(&User).Where("id = ?", id).Update("email", User.Email).Error

	if err != nil {
		return errors.New("could not update User"), nil
	}

	return err, User
}

func (r *UsersServiceImpl) CurrentUser(ctx *gin.Context) {

	user_id, err := token.ExtractTokenID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	str_user_id := strconv.Itoa(int(user_id))

	err, u := r.GetUser(&str_user_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type UsersService interface {
	CreateUser(*models.Users) (error, *models.Users)
	GetUser(*string) (error, *models.Users)
	GetUsers() (error, []*models.Users)
	UpdateUser(*models.Users, *string) (error, *models.Users)
	DeleteUser(*string) error
	GetUserByAddress(*string) (error, *models.Users)
}
