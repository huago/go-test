package test

import (
	"fmt"
	"testing"
)

func TestSomeRand(t *testing.T) {
	sr := SomeRand()

	fmt.Printf("原始数组：")
	for _, v := range sr {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()

	bs := BubbleSort(sr)
	fmt.Printf("冒泡结果：")
	for _, v := range bs {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()

	sr = SomeRand()
	bsf := BubbleSortWithFlag(sr)
	fmt.Printf("Flag结果：")
	for _, v := range bsf {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()

	sr = SomeRand()
	is := InsertSort(sr)
	fmt.Printf("插入结果：")
	for _, v := range is {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()

}
