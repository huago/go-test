package test

import (
	"fmt"
	"math/rand"
	"time"
)

func Rand100() {
	for i := 0; i < 100000; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(100)
		if num == 100 {
			fmt.Println("hello")
		}
	}

}
