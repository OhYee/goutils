package queue

import (
	"container/list"
	"fmt"
)

// Queue queue struct
type Queue struct {
	queue *list.List
}

// ErrQueueEmpty want using Top() or Pop() on a empty queue.
var ErrQueueEmpty = fmt.Errorf("Queue error: queue is empty, no value to get")

// NewQueue declare a queue
func NewQueue() *Queue {
	return &Queue{list.New()}
}

// Push a value to the queue
func (s *Queue) Push(value interface{}) {
	s.queue.PushBack(value)
}

// Pop a value from the queue
func (s *Queue) Pop() (value interface{}, err error) {
	e := s.queue.Front()
	if e != nil {
		value = e.Value
		s.queue.Remove(e)
	} else {
		err = ErrQueueEmpty
	}
	return
}

// Top return the top value of the queue and do not pop it
func (s *Queue) Front() (value interface{}, err error) {
	e := s.queue.Front()
	if e != nil {
		value = e.Value
	} else {
		err = ErrQueueEmpty
	}
	return
}

// Len return the legnth of the queue
func (s *Queue) Len() int {
	return s.queue.Len()
}

// Empty return the queue is empty
func (s *Queue) Empty() bool {
	return s.queue.Len() == 0
}
