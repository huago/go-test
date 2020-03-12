package test

import (
	"fmt"
	"testing"
)

func TestDoFunc(t *testing.T) {
	DoFunc()
}

func TestDoMyFunc(t *testing.T) {
	DoMyFunc()
}

func TestAnonymousFuncChannel(t *testing.T) {
	AnonymousFuncChannel()
}

func TestHanoi(t *testing.T) {
	var total int
	Hanoi(10, "A", "B", "C", &total)
	fmt.Println("total", total)
}

func TestAge(t *testing.T) {
	age := Age(5)

	fmt.Println("age", age)
}
