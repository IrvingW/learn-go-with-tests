package pointers

import "fmt"

// Bitcoin represent bitcon
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet that you can deposit money into in, and you can check the remain
// balance in it.
type Wallet struct {
	balance Bitcoin
}

// Deposit save some money into your wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance return the remain balance in the wallet
func (w Wallet) Balance() Bitcoin {
	return w.balance
}
