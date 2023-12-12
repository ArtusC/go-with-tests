// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	Balance Bitcoin
}

var ErrorInsufficientFounds = errors.New("cannot withdraw, insufficent founds in your account")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount
}

func (w *Wallet) BalanceMethod() Bitcoin {
	return w.Balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.Balance {
		return ErrorInsufficientFounds
	}

	w.Balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
