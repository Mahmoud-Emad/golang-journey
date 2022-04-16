package Wallet

import (
	"testing"
)

func TestWallet(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// Refactor and add withdraw
func TestWalletR(t *testing.T){
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if want != got{
			t.Errorf("got %s want %s", got, want)
		}
		
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance:Bitcoin(20)}
		wallet.Withdraw(Bitcoin(5))
		got := wallet.Balance()
		want := Bitcoin(15)
		if want != got{
			t.Errorf("got %s want %s", got, want)
		}

	})
}

// Refactor-2 and add withdraw
func TestWalletRT(t *testing.T){
	assertCorrect := func(t testing.TB, wallet Wallet, outPut Bitcoin){
		t.Helper()
		got := wallet.Balance()
		if got != outPut{
			t.Errorf("got %s output %s", got, outPut)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertCorrect(t, wallet, Bitcoin(10))
		
	})
}