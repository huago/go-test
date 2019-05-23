package test

import (
	"fmt"
	"time"
)

func Times() {
	ch11 := make(chan int, 1000)
	sign := make(chan byte, 1)

	for i := 0; i < 1000; i++ {
		fmt.Printf("set %d\n", i)
		ch11 <- i
	}

	go func() {
		var e int
		ok := true

		var timer *time.Timer

		for {
			select {
			case e = <-ch11:
				fmt.Printf("ch11 -> %d\n", e)
			case <-func() <-chan time.Time {
				if timer == nil {
					fmt.Println("new timer")
					timer = time.NewTimer(1000 * time.Millisecond)
				} else {
					fmt.Println("timer reset")
					timer.Reset(1000 * time.Millisecond)
				}
				fmt.Println("return")
				return timer.C
			}():
				fmt.Println("Timeout.")
				ok = false
				break
			case <- time.After(1 * time.Millisecond):
				fmt.Println("Dead.")
				ok = false
				break

			}

			if !ok {
				sign <- 0
				break
			}
		}

	}()

	<-sign
}
