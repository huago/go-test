package test

import "fmt"

func Scan() {
	var a interface{}
	a = int64(2)
	b, ok := a.(int64)
	fmt.Printf("b=%v, ok=%v\n", b, ok)

}
