package test

import "fmt"

type Myname struct {
	Name string
	Sex  string
	Age  int
}

func PrintPlus() {
	myname := new(Myname)
	myname.Name = "chenyubo"
	myname.Age = 12
	myname.Sex = "nan"

	fmt.Printf("%v\n", myname)
	fmt.Printf("%+v\n", myname)
	fmt.Printf("%##v\n", myname)
}
