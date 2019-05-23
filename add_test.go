package test

import (
	"fmt"
	"testing"
	"unsafe"
)

const Name string = "chenyubo"
const Addr string = "beijing"

const (
	a string = "a"
	b string = "b00000000000099999999999ddddddddddddddddddddddmmmmmmmmmmmmmmmmmmmmmmddddddddddd"
	c = unsafe.Sizeof(b)
	d = unsafe.Sizeof(100)
	e = unsafe.Sizeof(1.23)
)

func TestAdd(t *testing.T) {
	sum := Add(1, 9)
	fmt.Println("sum", sum)

	if sum < 100 {
		t.Error("Sum is smaller than 10")
		t.Errorf("Sum is smaller than %d", 10)
	}
}

func BenchmarkAdd(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Add(1, 9)
	}
}

func TestAddr(t *testing.T) {
	s := "abc"
	fmt.Println(&s)

	s, n := "scm", 100
	fmt.Println(&s, n)

	s = "7788"
	fmt.Println(&s)

	{
		s := "9900"
		fmt.Println(&s)

		s = "7yu"
		fmt.Println(&s)
	}

	const Phone string = "18810114921"
	fmt.Println(Name, Addr, Phone)
	fmt.Println(a, b, c, d, e)
}
