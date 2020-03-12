package test

import "fmt"

func DoBreak() {
	for i := 0; i < 10; i++ {
	outer:
		for j := 0; j < 10; j++ {
			if j > 6 {
				break outer
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("end.")
}

func DoGoto() {

}

func DoContinue() {

}
