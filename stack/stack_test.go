package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	judgeSize := func(size int) {
		if stack.Len() != size ||
			(size == 0 && stack.Empty() != true) {
			t.Errorf("Except stack size is %d, but got %d", size, stack.Len())
		}
	}

	judgeSize(0)
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	judgeSize(4)
	if v, err := stack.Top(); v != 3 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 3, nil, v, err)
	}
	judgeSize(4)
	if v, err := stack.Top(); v != 3 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 3, nil, v, err)
	}
	judgeSize(4)
	if v, err := stack.Pop(); v != 3 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 3, nil, v, err)
	}
	judgeSize(3)
	if v, err := stack.Pop(); v != 2 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 2, nil, v, err)
	}
	judgeSize(2)
	if v, err := stack.Pop(); v != 1 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 1, nil, v, err)
	}
	judgeSize(1)
	if v, err := stack.Top(); v != 0 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 0, nil, v, err)
	}
	judgeSize(1)
	if v, err := stack.Pop(); v != 0 || err != nil {
		t.Errorf("Except %d %v, but got %d %v", 0, nil, v, err)
	}
	judgeSize(0)
	if _, err := stack.Pop(); err == nil {
		t.Errorf("Except err, but got nil")
	}
	if _, err := stack.Top(); err == nil {
		t.Errorf("Except err, but got nil")
	}
}
