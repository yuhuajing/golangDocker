package scanevent

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client *ethclient.Client
)

const (
	ethServer = "https://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/"
)

func init() {
	client = getConn(ethServer)
}

func getConn(str string) *ethclient.Client {
	nclient, err := ethclient.Dial(str)
	//https://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/
	if err != nil {
		log.Fatal(err)
	}
	return nclient
}

var TransferTopic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
var startBlockHeight = 16544976
var latestblockNum = 16550289
var address = "0xdD69DB25F6D620A7baD3023c5d32761D353D3De9"

func ScanEvent() {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlockHeight)),
		ToBlock:   big.NewInt(int64(latestblockNum)),
		Addresses: []common.Address{common.HexToAddress(address)},
		Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(logs))
	fmt.Println(logs)
}
