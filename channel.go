package test

import (
	"fmt"
	"time"
)

func SelectChannel() {
	ch := make(chan struct{})
	go doTask(ch)

	timeout := time.After(1 * time.Second)

	select {
	case <-ch:
		fmt.Println("task finish")
	case <-timeout:
		fmt.Println("task timeout")
	}
}

func doTask(ch chan struct{}) {
	time.Sleep(2*time.Second)
	s := new(struct{})
	ch <- *s
}
