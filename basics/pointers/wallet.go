package pointers

// Wallet that you can deposit money into in, and you can check the remain
// balance in it.
type Wallet struct {
	balance int
}

// Deposit save some money into your wallet
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance return the remain balance in the wallet
func (w *Wallet) Balance() int {
	return w.balance
}
