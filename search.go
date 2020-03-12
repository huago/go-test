package test

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func BinSearch() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	low := 0
	high := len(arr) - 1

	target := 6

	var pos int

	for low <= high {
		mid := (low + high) / 2

		if target == arr[mid] {
			pos = mid
			break
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("pos", pos)
}

func Reverse() {
	current := new(Node)
	header := current

	for i := 0; i < 10; i++ {
		node := &Node{
			Data: i,
			Next: nil,
		}

		current.Next = node
		current = current.Next
	}

	fmt.Println("反转前...")
	tmp := header
	for tmp.Next != nil {
		fmt.Printf("%v\t", tmp.Next.Data)
		tmp = tmp.Next
	}
	fmt.Println()

	// 开始反转
	curr := header.Next

	header.Next = nil

	left := new(Node)
	for curr != nil {
		next := curr.Next

		curr.Next = left
		left = curr
		curr = next
	}

	header.Next = left

	last := header
	fmt.Println("反转后...")
	for last.Next != nil {
		fmt.Printf("%v\t", last.Next.Data)

		last = last.Next
	}
	fmt.Println()
}
