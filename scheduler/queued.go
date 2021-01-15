package scheduler

import (
	"concurrent-crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	q.workerChan <- worker
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	panic("implement me")
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- q.requestChan:
				requestQ = append(requestQ, r)
			case w := <- q.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}