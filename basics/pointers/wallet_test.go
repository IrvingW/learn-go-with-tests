package pointers

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{0}
	w.Deposit(100)
	got := w.Balance()
	want := Bitcoin(100)
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
