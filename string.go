package test

import (
	"fmt"
	"reflect"
	"unsafe"
)

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func TransferString() {
	s := "hello world"
	pp("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)

	pp("string to []byte, bs: %x\n", &bs)
	pp("[]byte to string, s2: %x\n", &s2)

	//ts := "hello world"
	tbs := []byte("hello world")
	tbss := toString(tbs)

	fmt.Printf("tbs=%x\n", &tbs)
	fmt.Printf("tbss=%x\n", &tbss)
}

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

