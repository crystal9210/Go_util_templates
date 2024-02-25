package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done) // mainルーチン→goルーチン(worker)、とチャネルを渡している

	<-done // doneチャネル型の値をmainルーチンが受信→main関数のルーチンはworkerゴールーチンが終了するまでブロックされる
}
