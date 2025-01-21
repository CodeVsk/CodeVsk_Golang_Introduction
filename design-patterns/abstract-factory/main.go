package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount int64)
}

type NotificationService interface {
	SendNotification(message string)
}

type BillingFactory interface {
	CreatePaymentProcessor() PaymentProcessor
	CreateNotificationService() NotificationService
}

type PaypalProcessor struct{}

func (p *PaypalProcessor) ProcessPayment(amout int64) {
	fmt.Println("Process payment with Paypal")
}

type PaypalNotification struct{}

func (p *PaypalNotification) SendNotification(message string) {
	fmt.Println("Send notification from Paypal")
}

type PaypalFactory struct{}

func (p *PaypalFactory) CreatePaymentProcessor() PaymentProcessor {
	return &PaypalProcessor{}
}

func (p *PaypalFactory) CreateNotificationService() NotificationService {
	return	&PaypalNotification{}
}

func ProcessPayment(factory BillingFactory, amount int64, message string) {
	processor := factory.CreatePaymentProcessor()
	processor.ProcessPayment(amount)

	notification := factory.CreateNotificationService()
	notification.SendNotification(message)
}

func main() {
	ProcessPayment(&PaypalFactory{}, 50, "capybara")
}