package test

import (
	"fmt"
	"reflect"
)

type MyFunc func (string, string) string

func DoFunc() {
	var f = func() {
		fmt.Println("Hello world")
	}

	fun := reflect.ValueOf(f)
	fun.Call(nil)
}

func myFunc(f MyFunc, str1 string, str2 string) {
	str := f(str1, str2)

	fmt.Println(str)
}

func DoMyFunc() {
	var f MyFunc
	f = func(s string, s2 string) string {
		return s + " " + s2
	}
	myFunc(f, "hello", "world")
}

func AnonymousFuncChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(i int, i2 int) int {
		return i + i2
	}

	fmt.Println("result", (<- c)(1, 2))
}

