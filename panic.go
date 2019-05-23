package test

import (
	"fmt"
	"os"
)

func HandlePanic() {
	defer func() {
		for {
			if err := recover(); err != nil {
				fmt.Println("ERR", err)
			} else {
				fmt.Println("No error")
				os.Exit(0)
			}
		}
	}()

	defer func() {
		panic("you are dead")
	}()

	panic("i am dead")
}

func catch() {
	fmt.Println("catch:", recover())
}

func DelayPanic() {
	defer catch()
	defer fmt.Println("recover", recover())
	panic("i am dead")
}
