package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amout int64)
}

//--------------------------

type Paypal struct{}

func (p *Paypal) ProcessPayment(amout int64) {
	fmt.Println("Process payment with paypal.")
}

//--------------------------

type PaypalFactory struct {}

func (p *PaypalFactory) CreatePaymentProcessor() PaymentProcessor {
	return &Paypal{}
}

//--------------------------

type PaymentFactory interface {
	CreatePaymentProcessor() PaymentProcessor
}

func ProcessPayment(factory PaymentFactory, amount int64) {
	processor := factory.CreatePaymentProcessor()
	processor.ProcessPayment(amount)
}

func main() {
	ProcessPayment(&PaypalFactory{}, 60)
}