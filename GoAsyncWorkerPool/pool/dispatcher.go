package pool

import (
	import_work "go_async_worker_pool/work"
	"log"
	"runtime"
	"sync"
	"time"
)

var WorkerChannel = make(chan chan Work)

type Collector struct {
	Work    chan Work
	End     chan bool
	Finish  chan bool
	Workers []Worker
	Wg      sync.WaitGroup
	JobWg   sync.WaitGroup // すべてのジョブの処理を追跡するためのWaitGroup
}

type Work struct {
	ID  int
	Job string
}

type Worker struct {
	ID            int
	WorkerChannel chan chan Work // used to communicate between dispatcher and workers
	Channel       chan Work
	End           chan bool
}

func StartDispatcher(workerCount int) *Collector {
	input := make(chan Work)
	end := make(chan bool)
	finish := make(chan bool)
	collector := &Collector{Work: input, End: end, Finish: finish, JobWg: sync.WaitGroup{}}

	for i := 0; i < workerCount; i++ {
		log.Println("starting worker: ", i+1)
		worker := Worker{
			ID:            i + 1,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}
		collector.Wg.Add(1)
		worker.Start(collector)
		collector.Workers = append(collector.Workers, worker)
	}

	// start collector
	go func() {
		for {
			select {
			case <-end:
				log.Println("collector received the ending signal and then collector sending end signal to workers...") // ここまでできてる
				// すべてのワーカーゴルーチンに終了シグナルを送信
				for i := 0; i < len(collector.Workers); i++ {
					log.Println("now stopping worker(id:[%d])...", collector.Workers[i].ID)
					collector.Workers[i].End <- true
				}
				// ここで5秒間待機
				time.Sleep(time.Millisecond * 50)
				// ゴルーチンの数を出力
				log.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())
				collector.Wg.Wait() // 全てのワーカーが停止するまで待機
				close(WorkerChannel)
				log.Println("WorkerChannel was closed!")
				finish <- true // ここをcollectorインスタンスの中でチャネルをインスタンス化せず送信しようとするとおそらく実態がないためデッドロック
				log.Println("collecor.Finish <- true was success!")
				return
			case work := <-input:
				worker := <-WorkerChannel // wait for available channel
				worker <- work            // dispatch work to worker
			}
		}
	}()

	return collector
}

// start worker
func (w *Worker) Start(c *Collector) {
	go func() {
		for {
			select {
			case w.WorkerChannel <- w.Channel: // when the worker is available place channel in queue
			case work := <-w.Channel: // worker has received job
				// ジョブの処理を開始
				log.Printf("Worker %d starting work %d", w.ID, work.ID)
				import_work.DoWork(work.Job, w.ID)
				// ジョブの処理が完了したことをログに出力
				log.Printf("Worker %d finished work %d", w.ID, work.ID)
				c.JobWg.Done()

			case <-w.End:
				// close(w.Channel) // Workerのジョブチャネルをクローズ
				log.Print("worker[%d]'s Channel was closed!", w.ID)
				c.Wg.Done()
				return
			}
		}
	}()
}
