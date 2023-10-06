package itfi

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int) int
	WithDraw(amount int) error
}