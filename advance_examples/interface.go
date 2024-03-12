package main

import "fmt"

type BankAccount interface {
	GetBalance() int // 100 = 1 dollar
	Deposit(amount int)
	Witdraw(amount int) error
}

func main() {
	fmt.Println("Hello, playground")

}
