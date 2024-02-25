package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bufferedChan := make(chan string, 2) // バッファサイズ2のチャネル
	quitChan := make(chan bool)          // 終了シグナル用のチャネル

	go func() {
		for {
			fmt.Print("Enter data (or 'exit' to quit): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input) // 改行文字をトリム
			if input == "exit" {
				quitChan <- true // 終了シグナルを送信
				return
			}
			select {
			case bufferedChan <- input:
				fmt.Println("Data sent to the channel")
			default:
				fmt.Println("Buffer is full! Discarding data.")
			}
		}
	}()

loop:
	for {
		select {
		case <-quitChan:
			fmt.Println("Exiting program.")
			break loop // 終了シグナルが受信された場合、ループを抜ける
		case data := <-bufferedChan:
			fmt.Println("Received data from channel:", data)
			fmt.Print("Save data? (y/n): ")
			decision, _ := reader.ReadString('\n')
			decision = strings.TrimSpace(decision)

			switch decision {
			case "y":
				fmt.Println("Data stored")
				// storeVar = data // 必要に応じて使用
			case "n":
				fmt.Println("Data discarded")
				// dumpVar = data // 必要に応じて使用
			default:
				fmt.Println("Invalid input. Discarding data.")
			}

			// バッファの状態を表示
			time.Sleep(1 * time.Second)
			if len(bufferedChan) == 0 {
				fmt.Println("Buffer is empty.")
			} else if len(bufferedChan) == cap(bufferedChan) {
				fmt.Println("Buffer is full.")
			} else {
				fmt.Println("Buffer has", len(bufferedChan), "items.")
			}
		}
	}
}
