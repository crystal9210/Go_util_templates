// アプリケーションを動かすドライバーの役割
// そもそもドライバーとは？→spec.txt参照

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"go_worker_pool/pool"
)

// const WORKER_COUNT = 5
// const JOB_COUNT = 100

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <WORKER_COUNT> <JOB_COUNT>") // *この形式でコマンドを打ってほしいということを出力してる
		return
	}

	workerCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid worker count: %s", os.Args[1])
	}

	jobCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid job count: %s", os.Args[2])
	}

	var wg sync.WaitGroup
	// var endWg sync.WaitGroup
	// log.Println(workerCount)
	// endWg.Add(workerCount)
	log.Println("Starting application...")
	collector := pool.StartDispatcher(workerCount) // WaitGroupを渡す

	for i := 0; i < jobCount; i++ {
		wg.Add(1)
		go func(jobID int) {
			defer wg.Done()
			collector.Work <- pool.Work{ID: jobID, Job: "Sample Job"}
		}(i + 1)
	}
	wg.Wait() // すべてのゴルーチンるのを待つ
	log.Println("All jobs were completed!")
	// 終了シグナルを送信...→ここ以降が問題
	collector.StopWorkers()

	// ワーカーの終了を待機するための追加の同期メカニズムが必要（例: WaitGroupを使用）ただ、めんどくなったので略

	// collector.Wait() // すべてのワーカーが終了するまで待つ
	log.Println("Shutting down application was completed.")
	// To Shut down application was completed.：は不自然→不自然なのはtoが選択肢、つまり未来の意味を含意するからwasとの相性が悪い
}
