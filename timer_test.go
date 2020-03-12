package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimes(t *testing.T) {
	Times()
}

func Test_Duration(t *testing.T) {
	//duration()

	timer := time.NewTimer(time.Duration(2) * time.Second)
	ticker := time.NewTicker(time.Duration(2) * time.Second)

	var wg sync.WaitGroup

	wg.Add(2)
	go func(t *time.Timer) {
		defer wg.Done()

		for c := range t.C {
			fmt.Printf("timer=%v\tc=%v\n", time.Now().Format("2006-01-02 15:04:05"), c)
		}
	}(timer)

	go func(t *time.Ticker) {
		defer wg.Done()

		for c := range t.C {
			fmt.Printf("ticker=%v\tc=%v\n", time.Now().Format("2006-01-02 15:04:05"), c)
		}
	}(ticker)

	wg.Wait()
}
