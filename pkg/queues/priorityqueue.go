package queues

import (
	"container/heap"
	"fmt"
)

// Copied from https://pkg.go.dev/container/heap and sprinkled some generics on top.

type Item[T any] struct {
	value    *T
	priority int
	index    int
}

var _ heap.Interface = (*PriorityQueue[any])(nil)

type PriorityQueue[T any] []*Item[T]

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the LOWEST priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) SetPriority(value *T, priority int) {
	for _, item := range *pq {
		if item.value == value {
			item.priority = priority
			heap.Fix(pq, item.index)

			return
		}
	}

	// Could not find the item...
	var stringified string
	if value != nil {
		stringified = fmt.Sprint(*value)
	} else {
		stringified = fmt.Sprint(value)
	}
	panic(fmt.Sprintf("item not in queue: %s", stringified))
}

func (pq *PriorityQueue[T]) Init() {
	heap.Init(pq)
}

func (pq *PriorityQueue[T]) PopT() *T {
	return heap.Pop(pq).(*Item[T]).value
}

func (pq *PriorityQueue[T]) PushT(value *T, priority int) {
	newItem := &Item[T]{
		value,
		priority,
		0,
	}

	pq.Push(newItem)
}
