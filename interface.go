package test

import (
	"bytes"
	"fmt"
)

type Animal interface {
	PrintInfo()
}

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

type Dog int
type Cat int
type Welcome string

func (w Welcome) Do(k, v interface{}) {
	fmt.Printf("%s, 我叫%s,今年%d岁\n", w, k, v)
}

func (w Welcome) introduceYourself(k, v interface{}) {
	fmt.Printf("%s, 我叫%s, 今年%d岁\n", w, k, v)
}

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func (dog Dog) PrintInfo() {
	fmt.Println("Hi, I am a dog.")
}

func (cat Cat) PrintInfo() {
	fmt.Println("Hi, I am a cat.")
}

func PrintAnimal() {
	var a Animal
	var c Cat
	var d Dog

	a = c
	a.PrintInfo()

	a = d
	a.PrintInfo()
}

func Bytes() {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Hello, ChenYubo")
	fmt.Println(b.String())
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}
