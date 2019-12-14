package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	judgeSize := func(size int) {
		if queue.Len() != size ||
			(size == 0 && queue.Empty() != true) {
			t.Errorf("Except queue size is %d, but got %d", size, queue.Len())
		}
	}

	judgeSize(0)
	queue.Push(0)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	judgeSize(4)
	if v, err := queue.Front(); v != 0 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 0, nil, v, err)
	}
	judgeSize(4)
	if v, err := queue.Front(); v != 0 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 0, nil, v, err)
	}
	judgeSize(4)
	if v, err := queue.Pop(); v != 0 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 0, nil, v, err)
	}
	judgeSize(3)
	if v, err := queue.Pop(); v != 1 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 1, nil, v, err)
	}
	judgeSize(2)
	if v, err := queue.Pop(); v != 2 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 2, nil, v, err)
	}
	judgeSize(1)
	if v, err := queue.Front(); v != 3 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 3, nil, v, err)
	}
	judgeSize(1)
	if v, err := queue.Pop(); v != 3 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 3, nil, v, err)
	}
	judgeSize(0)
	if _, err := queue.Pop(); err == nil {
		t.Errorf("Except err, but got nil")
	}
	if _, err := queue.Front(); err == nil {
		t.Errorf("Except err, but got nil")
	}
}
