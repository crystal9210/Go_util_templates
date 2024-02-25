package main

import "fmt"

// バッファは、チャネルに要素を一時的に格納するためのメモリ領域です。通常、チャネルは送信（<-演算子）と受信（<-演算子）が同時に起こるまで、送信する側と受信する側が同期されます。しかし、バッファ付きのチャネルでは、一定数の要素を保持できるため、送信側がブロックされずに複数の要素を送信できます。
// はい、通常、データ送信時にはチャネルまたは類似の概念を利用します。バッファを利用する場合でも、バッファはチャネルの一部であり、通常バッファ自体はキューとして実装され、チャネルを介してデータが送信されます。バッファがチャネルに含まれる一時的な領域であるため、バッファを通してデータがチャネルに書き込まれ、それが別のゴルーチンで読み取られることになります。つまり、データはチャネルを介して通信され、バッファがその通信のプロセスを補助します。

//例えば、messages := make(chan string, 2)のようにバッファサイズが2のチャネルを作成すると、このチャネルは2つの要素を保持できることになります。
// そして、messages <- "buffered"およびmessages <- "channel"によって、それぞれの文字列がチャネルに送信されます。バッファがいっぱいになるまで待つ必要はないため、送信は即座に行われます。

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// →
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive]:
	// main.main()
	//
	//	/home/crystal9210/Go_util_templates/fundamentals/baffaringchannel.go:19 +0x11a
	//
	// exit status 2
}
