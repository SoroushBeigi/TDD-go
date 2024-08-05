package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit function", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw function", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}

		err := wallet.Withdraw(Bitcoin(3))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(2))
	})

	t.Run("Withdraw function error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}

		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err, errInsufficientError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("did not get an error when needed")
	}
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("should not have error here")
	}

}
