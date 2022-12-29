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

	usdt "blockchain/contracts/USDT"
	"blockchain/pkg/models"
	"blockchain/pkg/routes"
	"blockchain/pkg/services"

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
	db          *gorm.DB = models.ConnectDataBase()
	ctx         context.Context
	url         = "https://mainnet.infura.io/v3/1d9d0d19d4df45a99d2f4d162a7d830e"
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

	app.GET("/ws", func(ctx *gin.Context) {
		fmt.Println("start subscription")
		clienWss, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/1d9d0d19d4df45a99d2f4d162a7d830e")
		contractAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}

		logs := make(chan types.Log)
		sub, err := clienWss.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Fatal(err)
		}

		contractAbi, err := abi.JSON(strings.NewReader(string(usdt.StoreABI)))
		if err != nil {
			log.Fatal(err)
		}

		logTransferSig := []byte("Transfer(address,address,uint256)")
		LogApprovalSig := []byte("Approval(address,address,uint256)")
		logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
		logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)
		count := 0
		for {
			select {
			case err := <-sub.Err():
				log.Fatal("err:", err)
			case vLog := <-logs:
				count += 1
				switch vLog.Topics[0].Hex() {
				case logTransferSigHash.Hex():
					fmt.Printf("Log Name: Transfer\n")

					var transferEvent LogTransfer

					data, err := contractAbi.Unpack("Transfer", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
					transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
					transferEvent.Tokens = data[0].(*big.Int)

					// fmt.Printf("From: %s\n", transferEvent.From.Hex())
					// fmt.Printf("To: %s\n", transferEvent.To.Hex())
					// fmt.Printf("TPT Tokens wei: %s\n", transferEvent.Tokens.String())

					if transferEvent.To.String() == "0x50710d2dE5C8dAAB8977CEa01c0A47BdA593272B" {
						us := services.NewUsersService(db)
						senderAdrress := strings.ToLower(transferEvent.From.String())
						err, sendUser := us.GetUserByAddress(&senderAdrress)
						if err != nil {
							fmt.Println(err)
						} else {
							walletToken := models.WalletTokens{
								WalletId: sendUser.Wallet.ID,
								Name:     "TPToken",
								Symbol:   "TPT",
								Amount:   uint(transferEvent.Tokens.Uint64()),
							}

							err = db.Create(&walletToken).Error
						}
					}

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

					// fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
					// fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
					// fmt.Printf("TPT Tokens wei: %s\n", approvalEvent.Tokens.String())
				}
				fmt.Println(count)
			}
		}
	})

	app.Run(":3003")
}
