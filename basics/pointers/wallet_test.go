package pointers

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{0}
	w.Deposit(100)
	got := w.Balance()
	want := 100
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
