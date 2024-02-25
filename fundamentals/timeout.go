package main

import (
	"fmt"
	"time"
)

// 【概要】selectを使って複数のチャネルから値を受信し、それぞれのゴルーチンが処理を完了するまで待機します。また、time.Afterを使用して、指定した時間内に受信が行われない場合のタイムアウト処理を実装しています。

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1" // 外部のルーチンとの非同期処理における通信を実現(模倣)
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2" // // 外部のルーチンとの非同期処理における通信を実現(模倣)
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
