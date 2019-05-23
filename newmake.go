package test

import (
	"fmt"
	"sync"
)

type worker struct {
	lock sync.Mutex
	name string
	age  int
}

func NewInt() *int {
	return new(int)
}

func NewInt8() *int8 {
	return new(int8)
}

func NewInt32() *int32 {
	return new(int32)
}

func PrintInt() {
	var num *int
	num = NewInt()
	*num = 100
	fmt.Println("num", num)
	fmt.Println("*num", *num)
}

func WorkerLock() {
	w := new(worker)
	ww := worker{}
	ww.lock.Lock()
	ww.lock.Unlock()

	w.lock.Lock()
	w.name = "chenyubo"
	w.age = 24
	w.lock.Unlock()

	fmt.Println(w)
}

//make只用于slice,map和chan，并且返回类型本身
//new用于一般的内存分配，返回类型指针
