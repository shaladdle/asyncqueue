package asyncqueue

import (
	"container/list"
	"sync"
)

type Queue interface {
	Push(interface{})
	Pop() interface{}
}

type queue struct {
	*sync.Mutex
	*sync.Cond
	*list.List
}

func NewQueue() Queue {
	mut := &sync.Mutex{}
	return &queue{
		Mutex: mut,
		Cond:  sync.NewCond(mut),
		List:  list.New(),
	}
}

func (q *queue) Push(item interface{}) {
	q.Lock()
	q.PushBack(item)
	q.Unlock()

	q.Signal()
}

func (q *queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	for q.Len() == 0 {
		q.Wait()
	}

	return q.Remove(q.Front())
}
