package main

import (
	"fmt"
	itfi "itf/interfaces_implementation"
	itf "itf/interfaces"
)

func main() {
	banks := []itf.IBankAccount{
		itfi.NewBancoDoBrasil(),
		itfi.NewItau(),
	}

	for _, bank := range banks {
		bank.Deposit(200)
		fmt.Printf("CurrentBalance: %d\n", bank.GetBalance())
	}
}