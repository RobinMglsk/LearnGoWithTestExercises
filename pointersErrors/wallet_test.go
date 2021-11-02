package wallet

import "testing"

func TestWallet(t *testing.T){

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Desposit", func(t *testing.T){
		wallet := Wallet{}
	
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})
	
	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{Bitcoin(20)}
	
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error bit didn't get one")
		}
	})

}