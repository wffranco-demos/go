package adapter

import "fmt"

type Payment interface {
	Pay()
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type Cash struct{}

func (c *Cash) Pay() {
	fmt.Println("Paying with cash")
}

type BankPayment struct {
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	Account     int
}

func (b *BankPayment) Pay(account int) {
	fmt.Println("Paying with bank account", account)
}

func (b *BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.Account)
}
