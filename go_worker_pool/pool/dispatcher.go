package pool

import (
	"fmt"
	"log"
	"sync"
)

var WorkerChannel = make(chan chan Work)

type Collector struct {
	Work    chan Work
	End     chan bool
	Workers []Worker
	endWg   sync.WaitGroup
}

func StartDispatcher(workerCount int) *Collector {
	var i int
	fmt.Println(i) // 出力: 0
	collector := Collector{
		Work:    make(chan Work),
		End:     make(chan bool),
		Workers: make([]Worker, workerCount),
	}

	for i := 0; i < workerCount; i++ {
		log.Println("Starting worker:", i+1)
		worker := Worker{
			ID:            i + 1,
			WorkerChannel: WorkerChannel,
			Channel:       make(chan Work),
			End:           make(chan bool),
			wg:            &collector.endWg,
		}
		log.Println("worker ID:", worker.ID)
		worker.Start()
		collector.Workers[i] = worker
		log.Println("collector.Workers[i].ID:", collector.Workers[i].ID)
	}

	return &collector
}

func (c *Collector) StopWorkers() {
	for _, worker := range c.Workers {
		worker.End <- true
	}
	c.endWg.Wait() // 全てのワーカーが終了するまで待機
	log.Println("All workers have stopped.")
}
