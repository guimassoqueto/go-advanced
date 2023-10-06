package itfi

import "fmt"

type Itau struct {
	balance int
}

func NewItau() *Itau{
	return &Itau{
		balance: 0,
	}
}

func (bb * Itau) GetBalance() int {
	return bb.balance
}

func (bb * Itau) Deposit(amount int) int {
	if amount <= 0 {
		return bb.balance
	}
	bb.balance += amount
	return bb.balance
}

func (bb * Itau) WithDraw(amount int) error {
	if amount < 0 {
		return fmt.Errorf("cannot withdraw negative amount")
	}

	if bb.balance < amount {
		return fmt.Errorf("you have %d, trying to withdraw %d", bb.balance, amount)
	}
	
	bb.balance -= amount
	return nil
}