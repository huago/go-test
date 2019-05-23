package test

/*
#include <stdlib.h>
#include <stdio.h>
void Hello() {
	printf("Hello Cgo!\n");
}
*/
import "C"
import "fmt"

type IReadWriter interface {
	Read(buf *byte, cb int) int
	Write(buf *byte, cb int) int
}

type B struct {
	name string
}

func NewB(name string) *B {
	fmt.Println("NewB params:", name)
	return &B{name: name}
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func SayHello() {
	C.Hello()
}

func SayHelloWithoutReturn() {
	C.Hello()
}

func Cstr() {
	var p = NewB("chenyubo")
	fmt.Printf("%v\n", p)
}
