package routine

import (
	"github.com/ysw-jingzhe/go-toolbox/coroutine/sync_queue"
)

// QuantityLimit quantitative restriction
const (
	QuantityLimit = 128
)

type Task struct {
	TID int64
}

var (
	SyncQueue chan *Task
)

func Initialize() {
	sq := sync_queue.NewLimiter(QuantityLimit)
	SyncQueue = make(chan *Task, QuantityLimit)

	go func(read <-chan *Task) {
		for ticket := range read {
			sq.Add()
			go WaitQueue(ticket, sq)
		}
	}(SyncQueue)
}

func WaitQueue(t *Task, sq *sync_queue.Limiter) {
	defer sq.Done()

	//Restrict the same TID task from being executed repeatedly at the same time
	if sq.Open(t.TID) {
		Runner(t)
		sq.Close(t.TID)
	}
}

func PushQueue(t *Task) {
	SyncQueue <- t
}
