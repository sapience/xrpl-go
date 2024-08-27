package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

func main() {

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: "r9cZA1mLK5R5AmHZiRd6CCe83ACaut34Mf",
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: "USD",
			Issuer: "r9cZA1mLK5R5AmHZiRd6CCe83ACaut34Mf",
			Value: "100",
		},
		Destination: "r9cZA1mLK5R5AmHZiRd6CCe83ACaut34Mf",
	}

	fmt.Println(payment)
	fmt.Println(payment.Flatten())
}