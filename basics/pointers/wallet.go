package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represent bitcon
// inherit from int
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
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds represent an error
var ErrInsufficientFunds = errors.New("Insufficient funds error")

// Withdraw will check if your wallet has sufficient money
// if so, it will withdraw amount money, and return nil
// if not, it will return an error
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
