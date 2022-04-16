package Wallet

import (
	"fmt"
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

func (w *Wallet) Withdraw(amount Bitcoin) Bitcoin {
	w.balance -= amount
	return w.balance
}

func (w *Wallet) Balance() Bitcoin {
	// This to get the balance of your wallet
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}