package adapter

import "testing"

func TestPayment(t *testing.T) {
	cash := &Cash{}
	ProcessPayment(cash)

	// bank := &BankPayment{}
	// ProcessPayment(bank) // error: BankPayment implement Pay method with different signature
	bank := &BankPaymentAdapter{
		Account: 123456,
	}
	ProcessPayment(bank)
}
