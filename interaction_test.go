package test

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	Seed(100)
	fmt.Println("Random:", Random())
}

func TestSayHello(t *testing.T) {
	SayHello()
}

func TestSayHelloWithoutReturn(t *testing.T) {
	SayHelloWithoutReturn()
}

func TestCstr(t *testing.T) {
	Cstr()
}
