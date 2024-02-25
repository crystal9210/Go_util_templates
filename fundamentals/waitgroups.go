package main

import (
	"fmt"
	"sync"
	"time"
)

// 【概要】
// sync.WaitGroupは、ゴルーチンのグループの完了を待つための仕組みを提供します。主に以下の3つのメソッドを持ちます：

// Add(delta int): WaitGroupの内部カウンターを増やします。deltaは正の整数で、通常はゴルーチンの数を表します。これにより、待機するゴルーチンの数が増えたことをWaitGroupに通知します。

// Done(): WaitGroupの内部カウンターを減らします。Add()で増やした数だけ、Done()を呼び出す必要があります。これにより、ゴルーチンが完了したことをWaitGroupに通知します。

// Wait(): WaitGroupの内部カウンターが0になるまでブロックし、それがゼロになると解除されます。つまり、全てのDone()が呼び出されるのを待ちます。

// これらのメソッドを組み合わせることで、複数のゴルーチンが完了するのを待つことができます。例えば、以下のような流れになります：

// メインゴルーチン内でWaitGroupを作成します。
// ゴルーチンを起動する前に、Add()を使用してWaitGroupの内部カウンターを増やします。
// 各ゴルーチンが完了したら、その最後にDone()を呼び出します。
// 全てのゴルーチンが完了したいときに、Wait()を呼び出して全てのゴルーチンが終了するのを待ちます。
// これにより、メインゴルーチンが全てのゴルーチンの完了を確認し、次の処理に進むことができます。

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)         // waitGroup wgの内部カウンターを1増やす→最終的に内部カウンターに一致する数だけDoneが呼び出されれば処理終了
		go worker(i, &wg) // ゴルーチンで生成して処理するため非同期
	}

	wg.Wait()
}
