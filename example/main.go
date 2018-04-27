package main

import (
	"time"

	coin "github.com/finejian/go-connect-eth-contract"
)

func main() {
	s := coin.NewConnecter()

	s.WatchTransferAndMint()

	time.Sleep(10 * 60 * 1000)
}
