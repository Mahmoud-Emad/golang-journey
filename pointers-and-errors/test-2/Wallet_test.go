package WalletER
import "testing"

// func TestWallet(t *testing.T) {
// 	// We want Withdraw to return an error if you try to take out more than you have and the balance should stay the same.
// 	// We then check an error has returned by failing the test if it is nil
// 	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
// 		t.Helper()
// 		got := wallet.Balance()
// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}
// 	t.Run("withdraw insufficient funds", func(t *testing.T) {
// 		startingBalance := Bitcoin(20)
// 		wallet := Wallet{startingBalance}
// 		err := wallet.Withdraw(Bitcoin(100))
// 		assertBalance(t, wallet, startingBalance)
// 		if err == nil {
// 			t.Errorf("wanted error but didn't get one")
// 		}
// 	})
	
// }


// // Refactor
// // Let's make a quick test helper for our error check to improve the test's readability

// func TestWalletR(t *testing.T){
// 	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
// 		t.Helper()
// 		got := wallet.Balance()
// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}
// 	assertError := func(t testing.TB, err error){
// 		t.Helper()
// 		if err == nil {
// 			t.Error("wanted an error but didn't get one")
// 		}
// 	}
// 	t.Run("withdraw insufficient funds", func(t *testing.T){
// 		startingBalance := Bitcoin(20)
// 		wallet := Wallet{startingBalance}
// 		err := wallet.Withdraw(Bitcoin(100))
// 		assertError(t, err)
// 		assertBalance(t, wallet, startingBalance)
// 	})
// }

// // Assuming that the error ultimately gets returned to the user, 
// // let's update our test to assert on some kind of error message rather than just the existence of an error.
// func TestWalletRR(t *testing.T){
// 	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
// 		t.Helper()
// 		got := wallet.Balance()
// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}
// 	assertError := func(t testing.TB, got error, want string) {
// 		t.Helper()
// 		if got == nil {
// 			t.Fatal("didn't get an error but wanted one")
// 		}
// 		if got.Error() != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	}
// 	t.Run("withdraw insufficient funds", func(t *testing.T) {
// 		startingBalance := Bitcoin(20)
// 		wallet := Wallet{startingBalance}
// 		err := wallet.Withdraw(Bitcoin(100))
	
// 		assertError(t, err, "cannot withdraw, insufficient funds")
// 		assertBalance(t, wallet, startingBalance)
// 	})
// }

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}