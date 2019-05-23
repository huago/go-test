package test

import (
	"fmt"
	"runtime"
	"sync"
)

func Wg() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	sum := 0

	for i := 0; i < 100000000; i++ {
		sum += i
	}

	fmt.Println(index, sum)
	wg.Done()
}
