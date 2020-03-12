package test

import (
	"fmt"
	"reflect"
)

type MyFunc func(string, string) string

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

	fmt.Println("result", (<-c)(1, 2))
}

func movep(a string, b string) {
	fmt.Println(a, "->", b)
}

func Hanoi(n int, a string, b string, c string, total *int) {
	if n == 1 {
		movep(a, c)
		*total += 1
	} else {
		Hanoi(n-1, a, c, b, total)
		movep(a, c)
		*total += 1
		Hanoi(n-1, b, a, c, total)
	}
}

func Age(n int) int {
	if n == 1 {
		return 10
	}

	return Age(n-1) + 2
}
