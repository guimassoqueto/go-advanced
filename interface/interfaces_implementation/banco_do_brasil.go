package itfi

import "fmt"

type BancoDoBrasil struct {
	balance int
}

func NewBancoDoBrasil() *BancoDoBrasil{
	return &BancoDoBrasil{
		balance: 0,
	}
}

func (bb * BancoDoBrasil) GetBalance() int {
	return bb.balance
}

func (bb * BancoDoBrasil) Deposit(amount int) int {
	if amount <= 0 {
		return bb.balance
	}
	bb.balance += amount
	return bb.balance
}

func (bb * BancoDoBrasil) WithDraw(amount int) error {
	if amount < 0 {
		return fmt.Errorf("cannot withdraw negative amount")
	}

	if bb.balance < amount {
		return fmt.Errorf("you have %d, trying to withdraw %d", bb.balance, amount)
	}
	
	bb.balance -= amount
	return nil
}