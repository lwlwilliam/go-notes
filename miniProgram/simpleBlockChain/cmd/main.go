package main

import "github.com/lwlwilliam/Golang/miniProgram/simpleBlockChain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 EOS to Jack")
	bc.Print()
}
