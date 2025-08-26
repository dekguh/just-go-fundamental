package helpers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(0)

	got := wallet.Balance()
	expected := 0

	fmt.Printf("wallet balance test address: %p\n", &wallet.balance)

	if got != expected {
		t.Errorf("expected is %d but got %d", expected, got)
	}
}
