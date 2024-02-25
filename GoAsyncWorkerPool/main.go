package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"go_async_worker_pool/pool"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
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
	collector := pool.StartDispatcher(workerCount) // WaitGroupを渡す

	for i := 0; i < jobCount; i++ {
		collector.JobWg.Add(1)
		go func(jobID int) {
			collector.Work <- pool.Work{ID: jobID, Job: "Sample Job"}
		}(i + 1)
	}
	// すべてのジョブが処理されるのを待機
	collector.JobWg.Wait()
	collector.End <- true
	log.Println("main routine is Waiting for all workers to stop...")
	<-collector.Finish
	// ゴルーチンの数を出力
	log.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())

	fmt.Println("Congratulations!! you have successfully learned how to properly generate, allocate and release channels and goroutines!")
	fmt.Println("You have successfully completed this program! Good job!!")

}
