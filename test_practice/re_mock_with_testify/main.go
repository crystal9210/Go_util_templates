package main

import (
	"fmt"
)

// MessageService handles notifying clients they have
// been charged
type MessageService interface {
	SendChargeNotification(int) error
}

// SMSService is our implementation of MessageService
type SMSService struct{}

// MyService uses the MessageService to notify clients
type MyService struct {
	messageService MessageService
}

// テストでは、SendChargeNotificationメソッドが呼び出されたことを確認し、エラーが発生しないことを確認する必要がある
// 一部の関数では、エラーがない場合には明示的にnilを返すことで、成功したことを明示→errorが返り値の時でもnilは返せる
func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending Production Charge Notification")
	return nil
}

// ChargeCustomer performs the charge to the customer
func (a MyService) ChargeCustomer(value int) error {
	err := a.messageService.SendChargeNotification(value)
	if err != nil {
		return err
	}
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil
}

// A "Production" Example
func main() {
	fmt.Println("Hello World")

	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}
