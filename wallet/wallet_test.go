package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit function", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw function", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}
		wallet.Withdraw(Bitcoin(3))
		got := wallet.Balance()
		want := Bitcoin(2)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
