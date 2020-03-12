package test

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice()
	}
}

func TestPushPop(t *testing.T) {
	stack := make([]int, 0, 5)

	var m sync.Mutex
	push := func(x int) error {
		length := len(stack)
		if length == cap(stack) {
			return errors.New("stack is full")
		}

		m.Lock()
		stack = stack[:length+1]
		stack[length] = x
		m.Unlock()

		return nil
	}

	pop := func() (int, error) {
		length := len(stack)
		if length == 0 {
			return 0, errors.New("stack is empty")
		}

		m.Lock()
		x := stack[length-1]
		stack = stack[:length-1]
		m.Unlock()

		return x, nil
	}

	chs1 := make([]chan int, 7)
	chs2 := make([]chan int, 7)

	for i := 0; i < 7; i++ {
		chs1[i] = make(chan int)
		go func(ch chan int, num int) {
			ch <- 1
			fmt.Printf("push %d: %v, %v\n", num, push(num), stack)
		}(chs1[i], i)
	}

	for i := 0; i < 7; i++ {
		chs2[i] = make(chan int)
		go func(ch chan int) {
			ch <- 1
			x, err := pop()
			fmt.Printf("pop: %d, %v, %v\n", x, err, stack)
		}(chs2[i])
	}

	for _, ch := range chs1 {
		<-ch
	}

	for _, ch := range chs2 {
		<-ch
	}
}
