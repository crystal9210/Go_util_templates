package main

import (
	"fmt"
	"time"
)

// 【要点】正確には、プログラムのエントリポイントとなるメインゴルーチンが終了すると、プログラム全体が終了し、その時点で実行中の全てのゴルーチンも強制終了されます。その結果、プログラムが使用していたリソースはオペレーティングシステムによって解放されます。
// →メインルーチンが終了するとそれをエントリポイントとする関連するすべてのゴルーチンは強制的かつ自動的に終了し、それらを制御するプロセスは終了・解放される

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C // timeパッケージのNewTimer関数は内部的にCという名前のチャネルを生成するため、そのチャネルからの通知を受け取るまでメインルーチンの処理をブロックするということ→GPT:タイマーが通知を送信するためのチャネル (C) を生成します
	fmt.Println("Timer 1 fired")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}

// Goプログラムとゴルーチンの終了
// メインゴルーチンの終了とプログラムの終了: Go言語において、メインゴルーチン（main関数）が終了すると、プログラム全体が終了します。これは、メインゴルーチンが終了した時点で、他の未完了のゴルーチンが存在していたとしても、それらは強制的に終了します。

// リソースの解放: プログラムが終了すると、そのプロセスが使用していたメモリやその他のリソースはオペレーティングシステムによって解放されます。これは、ゴルーチンがブロックされた状態であっても同様です。プログラムが終了すれば、ゴルーチンによって消費されていたリソースも解放されます。

// リソースリークの可能性: プログラムが実行中の場合に限り、永遠にブロックされるゴルーチンはリソースリークの原因になり得ます。たとえば、ゴルーチンが終了しないでブロックされ続ける場合、そのゴルーチンが消費するメモリや他のリソースはプログラムが終了するまで解放されません。しかし、プログラム自体が終了すれば、これらのリソースはオペレーティングシステムによって回収されます。

// 結論
// メインゴルーチンが終了すると、Goのプログラムも終了します。この時点で、全てのゴルーチン（ブロックされているものも含む）は強制的に終了し、プログラムによって使用されていたリソースはオペレーティングシステムによって解放されます。

// プログラムがまだ実行中の状態でゴルーチンが永遠にブロックされる場合にのみ、リソースリークの問題が生じる可能性があります。プログラムが終了すれば、リソースリークは発生しません。

// ゴルーチンを保持するプロセスが「永遠に残り続ける」ということはありません。プログラムの終了と共に、関連する全てのプロセスとリソースは終了または解放されます。
