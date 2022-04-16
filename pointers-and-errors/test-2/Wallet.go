package WalletER

import (
	"fmt"
	"errors"
)

type Bitcoin int

type Wallet struct{
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) Bitcoin {
	// This to add more of bits to your wallet
	w.balance += amount
	return w.balance
}

// We have to check if the incoming amount is greater than the current balance and this will raise an error if so
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	// This to get the balance of your wallet
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}