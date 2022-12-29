package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

type TransactionsServiceImpl struct {
	ctx context.Context
}

func NewTransactionsService(ctx context.Context) TransactionsService {
	return &TransactionsServiceImpl{
		ctx: ctx,
	}
}

type TransferParams struct {
	R      string `json:"recipient" xml:"recipient" form:"recipient"`
	S      string `json:"sender" xml:"sender" form:"sender"`
	Amount uint   `json:"amount" xml:"amount" form:"amount"`
}

var (
	ctx         context.Context
	url         = "https://goerli.infura.io/v3/1d9d0d19d4df45a99d2f4d162a7d830e"
	client, err = ethclient.DialContext(ctx, url)
)

func (r *TransactionsServiceImpl) GetBalance(address *string) (error, *big.Int) {
	account := common.HexToAddress(*address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return err, nil
	}

	return err, balance
}

func (r *TransactionsServiceImpl) Transfer(transferParams *TransferParams) (error, *types.Transaction) {
	recipientAddress := common.HexToAddress(transferParams.R)
	privateKey, err := crypto.HexToECDSA(string(transferParams.S))
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA"), nil
	}
	SenderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), SenderAddress)
	if err != nil {
		return err, nil
	}

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err, nil
	}

	tokenAddress := common.HexToAddress("0xDE2cBDE6E00F529d4DE278d950c1F9E686A2c952")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(recipientAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))
	// amount := new(big.Int)
	amount := big.NewInt(int64(transferParams.Amount))
	// amount.SetString("100000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &recipientAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal("err:", err)
		return err, nil
	}
	fmt.Println(gasLimit)
	ChainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err, nil
	}

	transaction := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(ChainID), privateKey)
	if err != nil {
		return err, nil
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err, nil
	}

	// fmt.Printf("transaction sent: %s", signedTx.Hash().Hex())

	return err, signedTx
}

func (r *TransactionsServiceImpl) QueryTransactions() (error, *types.Block) {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	return err, block
}

type TransactionsService interface {
	Transfer(*TransferParams) (error, *types.Transaction)
	QueryTransactions() (error, *types.Block)
	GetBalance(*string) (error, *big.Int)
}
