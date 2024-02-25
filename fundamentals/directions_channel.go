package main

import "fmt"

// pings chan<- string:pingsは送信専用チャネルとしてping関数内で扱われる
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings <-chan string:pingsは受信専用チャネルとしてpong関数内で扱われる
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// ☆チャネルに対し、受信として扱うか送信として扱うかの情報を付与することで関数内での誤用を避け、明確に定義し、可読性を上げることができる→チャネルに関連するバグ、誤用を避けることができる

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
