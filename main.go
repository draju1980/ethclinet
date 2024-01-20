package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection with the Ethereum network")
	_ = client

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Printing the latest block number
	fmt.Printf("Latest block number in Ethereum Mainnet: %d\n", blockNumber)

	// Getting the ethereum account from the user
	fmt.Println("Enter Your Ethereum Account: ")
	var ethaccount string
	fmt.Scanln(&ethaccount)

	// Converting ethereum account to address
	account := common.HexToAddress(ethaccount)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// converting wei to ether (string to float)
	accBalance, err := strconv.ParseFloat(balance.String(), 64)
	accBalance = accBalance / 1000000000000000000
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Printing the balance
	fmt.Printf("\nYour Ethereum Account %s Balance is : %f ETH\n", ethaccount, accBalance)

}
