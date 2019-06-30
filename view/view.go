package view

import (
	"fmt"
	"strings"
)

func BalanceMessage(balance float64) string {
	if balance < 0 {
		return fmt.Sprintf("Balance: -€%.2f", balance)
	} else {
		return fmt.Sprintf("Balance: €%.2f", balance)
	}
}

func LogMessage(entries []string) string {
	return strings.Join(entries, "\n")
}
