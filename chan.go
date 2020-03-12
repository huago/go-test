	package test

import (
	"fmt"
	"time"
)

func NoBufferChan() {
	ch := make(chan int)

	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}

		ch <- sum
	}()

	fmt.Println(<-ch)
}

func BufferedChan() string {
	ch := make(chan string, 3)
	capacity := cap(ch)
	length := len(ch)

	fmt.Println("capacity", capacity)
	fmt.Println("length", length)

	go func() {
		ch <- "huagou"
	}()

	go func() {
		ch <- "chenyubo"
	}()

	go func() {
		ch <- "chenhuagou"
	}()

	firstName := <-ch
	secondName := <-ch
	thirdName := <-ch
	fmt.Println("firstname", firstName)
	fmt.Println("secondname", secondName)
	fmt.Println("thirdname", thirdName)

	return firstName
}

func OneDirection() {
	fmt.Println("hhhhhh")
	//var doubleDirection chan int
	doubleDirection := make(chan int)
	//var oneDirectionSend chan<- int
	//var oneDirectionReceive <-chan int
	go func() {
		doubleDirection <- 18
	}()

	fmt.Println("nnnnnnn")
	//oneDirectionSend <- 19
	//content := <-oneDirectionReceive

	num := <-doubleDirection
	fmt.Println(num)
	return

	//fmt.Println(<- oneDirectionReceive)
	//fmt.Println(content)

}

func Channel() {
	//in可以输出
	fmt.Println("in")

	ch := make(chan int, 1)

	ch <- 18

	//输出不了send
	fmt.Println("send")

	num := <-ch

	fmt.Println(num)
}

func Pipe() {
	one := make(chan int)
	two := make(chan int)

	go func() {
		one <- 100
	}()

	go func() {
		v := <-one
		two <- v
	}()

	fmt.Println(<-two)
}

func Select() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	var num int
	var ok bool = true

	for {
		select {
		case num = <-ch:
			fmt.Println(num)
		case <-func() chan int {
			fmt.Println("chan init")
			c := make(chan int)

			return c
		}():
			fmt.Println("luck")
		default:
			fmt.Println("finish")
			ok = false
		}

		if ok != true {
			break
		}
	}

	fmt.Println("end")
}

func DoOrder() {
	for i := 0; i < 10; i++ {
		fmt.Println("ii:", i)
		time.Sleep(time.Second)
	}

	go func() {
		for j := 0; j < 10; j++ {
			fmt.Println("j:", j)
			time.Sleep(time.Second)
		}
	}()

	for h := 0; h < 10; h++ {
		fmt.Println("h:", h)
		time.Sleep(time.Second)
	}
}

func CloseChan() {
	ch := make(chan int, 1000)
	sign := make(chan int, 1)

	for i := 0; i < 1000; i++ {
		ch <- i
	}

	close(ch) //被关闭的channel，可以继续读，但是不能继续写

	go func() {
		var e int
		ok := true

		for {
			select {
			case e, ok = <-ch:
				if !ok {
					fmt.Println("End.")
					break
				}
				fmt.Printf("ch -> %d\n", e)
			}

			if !ok {
				fmt.Println("it's not ok")
				sign <- 0
				break
			}
		}
	}()

	<-sign
}
