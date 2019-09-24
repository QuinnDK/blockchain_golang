package cli

import (
	"fmt"
	block "myCode/public_blockchain/part7-network/blc"
)

func (cli *Cli) getBalance(address string) {
	bc := block.NewBlockchain()
	defer bc.BD.Close()
	balance := bc.GetBalance(address)
	fmt.Printf("用户:%s的余额为：%d\n", address, balance)
}
