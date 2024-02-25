package pool

import (
	"log"
	"sync"

	import_work "go_worker_pool/work"
)

type Work struct {
	ID  int
	Job string
}

type Worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool
	wg            *sync.WaitGroup // Workerが完了した際に通知するため
}

// start worker
func (w *Worker) Start() {
	w.wg.Add(1) // Workerの開始をWaitGroupに通知
	// 各ワーカーが持つゴルーチン
	go func() {
		defer w.wg.Done() // 処理が終了したらDoneを呼び出してWaitGroupのカウンタをデクリメント
		for {
			select {
			case w.WorkerChannel <- w.Channel:
				// ワーカープールに自身のチャネルを送信
				work, ok := <-w.Channel
				if !ok {
					// Channelが閉じられた場合、ループを抜ける
					return
				}
				// ジョブの処理を開始
				log.Printf("Worker %d starting work %d", w.ID, work.ID)
				import_work.DoWork(work.Job, w.ID)
				// ジョブの処理が完了したことをログに出力
				log.Printf("Worker %d finished work %d", w.ID, work.ID)
			case <-w.End:
				log.Printf("Worker %d stopping", w.ID)
				return // Endからシグナルを受け取ったらループを抜ける
			}
		}
	}()
}
