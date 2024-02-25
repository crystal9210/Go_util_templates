package main

import (
	"fmt"
	"time"
)

// タイマーは将来一度だけ何かを実行したい場合に使用します。ティッカーは一定の間隔で繰り返し実行したい場合に使用します。以下は、停止するまで定期的に音を立てるティッカーの例です。

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
