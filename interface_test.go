package test

import (
	"testing"
)

func TestPrintAnimal(t *testing.T) {
	PrintAnimal()
}

func TestBytes(t *testing.T) {
	Bytes()
}

func TestEach(t *testing.T) {
	var w Welcome
	w = "大家好"
	m := make(map[interface{}]interface{})
	m["chenyubo"] = 12
	m["chenhongyuan"] = 18
	Each(m, w)
}
