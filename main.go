package test

import "fmt"

func TypeMethod() {
	var str interface{} = 1234

	if i, ok := str.(int); ok {
		fmt.Println(i)
	} else {
		fmt.Println("str is not a int value")
	}

	switch inst := str.(type) {
	case int:
		fmt.Println("int", inst)
	case string:
		fmt.Println("string", inst)
	}

	fmt.Println(int(str))
}