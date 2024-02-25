package main

import "fmt"

// チャネルが(受信を)閉じてもバッファに保持するデータをルーチンに送信できることを確認するコード

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
