package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit test", func(t *testing.T) {
		w := Wallet{0}
		w.Deposit(100)
		want := Bitcoin(100)
		assertBalance(t, &w, want)
	})

	t.Run("withdraw test", func(t *testing.T) {
		w := Wallet{100}
		w.Withdraw(100)
		assertBalance(t, &w, Bitcoin(0))
	})

	t.Run("withdraw with insufficient balance", func(t *testing.T) {
		w := Wallet{0}
		err := w.Withdraw(10)
		if err == nil {
			t.Fatal("err is nil")
		}
		if err != ErrInsufficientFunds {
			t.Errorf("got %s but want %s", err, ErrInsufficientFunds)
		}

	})
}

func assertBalance(t *testing.T, wallet *Wallet, want Bitcoin) {
	t.Helper()
	if wallet.Balance() != want {
		t.Errorf("got %s but want %s", wallet.Balance(), want)
	}
}
