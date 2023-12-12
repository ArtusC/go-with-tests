//go:build unit
// +build unit

package wallet_test

import (
	"testing"

	pae "github.com/ArtusC/go-with-tests/pointers-and-errors"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := pae.Wallet{}
		wallet.Deposit(pae.Bitcoin(10))
		assertBalance(t, wallet, pae.Bitcoin(10))

	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := pae.Wallet{Balance: pae.Bitcoin(10)}
		err := wallet.Withdraw(pae.Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, pae.Bitcoin(0))
	})

	t.Run("withdraw isufficient founds", func(t *testing.T) {
		startingBalance := pae.Bitcoin(20)
		wallet := pae.Wallet{Balance: startingBalance}
		err := wallet.Withdraw(pae.Bitcoin(100))

		assertError(t, err, pae.ErrorInsufficientFounds)
		assertBalance(t, wallet, startingBalance)

	})
}

func assertBalance(t testing.TB, wallet pae.Wallet, want pae.Bitcoin) {
	t.Helper()
	got := wallet.BalanceMethod()

	if got != want {
		t.Errorf("\ngot %s want %s", got, want)
	}

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
