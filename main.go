package main

import (
	"fmt"
)

func main() {
	account := banking.Account{Owner: "DongHyeon", Balance: 30000}
	fmt.Println(account)
}
