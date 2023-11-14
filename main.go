// package main

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"math/big"
// 	"time"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
// 	transferEth()
// }

// func transferEth() {
// 	client, err := ethclient.Dial("http://localhost:3000")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	priKey, _ := crypto.HexToECDSA("bebb5b73e288c580a6cee5070929ab3ff8985422d7a0bc45938faae5332e2e2f")
// 	publicKey := priKey.Public()
// 	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
// 	fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
// 	fmt.Println(fromaddress)
// 	bal, err := client.BalanceAt(context.Background(), common.HexToAddress(fromaddress), nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(bal)
// 	nonce, _ := client.PendingNonceAt(context.Background(), common.HexToAddress(fromaddress))
// 	value := big.NewInt(1000000000000000000) // in wei (1 eth)
// 	gasLimit := uint64(21000)
// 	//gasPrice := big.NewInt(30000000000) // in wei (30 gwei)
// 	gasPrice, _ := client.SuggestGasPrice(context.Background())

// 	toAddress := common.HexToAddress("0x60A6E5aF0525523A617CF6c1F85353FBA0408A7b")
// 	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

// 	chainId, _ := client.ChainID(context.Background())
// 	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainId), priKey)
// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
// 	time.Sleep(5 * time.Second)
// 	fmt.Println()
// 	bal, _ = client.BalanceAt(context.Background(), common.HexToAddress(fromaddress), nil)
// 	fmt.Println(bal)
// }

package main

import (
	"net/http"

	"fmt"
)

func main() {
	http.HandleFunc("/tz", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Docker Form Golang!")
}
