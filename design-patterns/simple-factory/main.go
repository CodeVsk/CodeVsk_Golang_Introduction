package main

import "fmt"

type PaymentProcessor interface {
	Pay()
}

type Paypal struct{}

func (p *Paypal) Pay() {
	fmt.Println("Paypal payment.")
}

func PaymentFactory(paymentType string) PaymentProcessor {
	switch paymentType {
	case "paypal":
		return &Paypal{}
	default:
		return nil
	}
}

func main() {
	processor := PaymentFactory("paypal")
	processor.Pay()
}