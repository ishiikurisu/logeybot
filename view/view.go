package view

import (
	"fmt"
)

func BalanceMessage(balance float64) string {
	if balance < 0 {
		return fmt.Sprintf("Balance: -€%.2f", balance)
	} else {
		return fmt.Sprintf("Balance: €%.2f", balance)
	}
}
