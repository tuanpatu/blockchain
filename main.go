package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	tptoken "blockchain/contracts"
	"blockchain/pkg/models"
	"blockchain/pkg/routes"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

var (
	db          *gorm.DB
	ctx         context.Context
	url         = "https://goerli.infura.io/v3/1d9d0d19d4df45a99d2f4d162a7d830e"
	client, err = ethclient.DialContext(ctx, url)
)

func HomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func serveWsTPT() {

}

func main() {

	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
	fmt.Println(client.BlockNumber(context.Background()))
	defer models.CloseDatabaseConnection(db)
	app := routes.InitRouter()

	app.GET("/", func(ctx *gin.Context) {
		clienWss, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/1d9d0d19d4df45a99d2f4d162a7d830e")
		contractAddress := common.HexToAddress("0xDE2cBDE6E00F529d4DE278d950c1F9E686A2c952")
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}

		logs := make(chan types.Log)
		sub, err := clienWss.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Fatal(err)
		}

		contractAbi, err := abi.JSON(strings.NewReader(string(tptoken.TptokenABI)))
		if err != nil {
			log.Fatal(err)
		}
		logTransferSig := []byte("Transfer(address,address,uint256)")
		LogApprovalSig := []byte("Approval(address,address,uint256)")
		logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
		logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

		for {
			select {
			case err := <-sub.Err():
				log.Fatal("err:", err)
			case vLog := <-logs:
				switch vLog.Topics[0].Hex() {
				case logTransferSigHash.Hex():
					fmt.Printf("Log Name: Transfer\n")

					var transferEvent LogTransfer

					data, err := contractAbi.Unpack("Approval", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("data", data[0])

					transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
					transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
					transferEvent.Tokens = data[0].(*big.Int)

					fmt.Printf("From: %s\n", transferEvent.From.Hex())
					fmt.Printf("To: %s\n", transferEvent.To.Hex())
					fmt.Printf("TPT Tokens wei: %s\n", transferEvent.Tokens.String())
					if transferEvent.To.String() == "0x50710d2de5c8daab8977cea01c0a47bda593272b" {
						fmt.Println("Recive token from user")
					}

					// us := services.NewUsersService(db)
					// senderAdrress := transferEvent.From.String()
					// err, sendUser := us.GetUserByAddress(&senderAdrress)

					// if transferEvent.To.String() == "0x50710d2dE5C8dAAB8977CEa01c0A47BdA593272B" {
					// 	fmt.Println(sendUser)
					// 	walletToken := models.WalletTokens{
					// 		WalletId: sendUser.Wallet.ID,
					// 		Name:     "TPToken",
					// 		Symbol:   "TPT",
					// 		Amount:   transferEvent.Tokens.String(),
					// 	}

					// 	err = db.Create(&walletToken).Error
					// }

				case logApprovalSigHash.Hex():
					fmt.Printf("Log Name: Approval\n")

					var approvalEvent LogApproval

					data, err := contractAbi.Unpack("Approval", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
					approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
					approvalEvent.Tokens = data[0].(*big.Int)

					fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
					fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
					fmt.Printf("TPT Tokens wei: %s\n", approvalEvent.Tokens.String())
				}
			}
		}
	})
	app.Run(":3003")
}
