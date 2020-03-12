package test

import "fmt"

var (
	Sm []*BTree
)

type BTree struct {
	data  int
	level int
	left  *BTree
	right *BTree
}

func InitBTree() *BTree {
	var root = &BTree{data: 6}

	root.left = &BTree{data: 4}
	root.left.left = &BTree{data: 3}
	root.left.right = &BTree{data: 5}

	root.right = &BTree{data: 9}
	root.right.left = &BTree{data: 7}
	root.right.right = &BTree{data: 8}

	return root
}

func GetSlice(sm *[]*BTree, root *BTree, level int) {
	if root == nil {
		return
	}

	level = level + 1
	root.level = level
	*sm = append(*sm, root)

	if root.left != nil {
		GetSlice(sm, root.left, level)
	}

	if root.right != nil {
		GetSlice(sm, root.right, level)
	}
}

func GetRightSlice(sm *[]*BTree, root *BTree, level int) {
	if root == nil {
		return
	}

	level = level + 1
	root.level = level
	*sm = append(*sm, root)

	if root.right != nil {
		GetRightSlice(sm, root.right, level)
	}

	if root.left != nil {
		GetRightSlice(sm, root.left, level)
	}
}

func DoTree() {
	bTree := InitBTree()

	var sm = []*BTree{}
	GetRightSlice(&sm, bTree, -1)

	for _, v := range sm {

		fmt.Println(v.level, v.data)
	}

	fmt.Println("----------------------")

	unique := Unique1(sm)

	for _, v := range unique {
		fmt.Println(v.level, v.data)
	}
}

func Unique1(sm []*BTree) []*BTree {
	var nsm = []*BTree{sm[0]}

	for i := 1; i < len(sm); i++ {
		var repeat bool
		for j := 0; j < len(nsm); j++ {
			if sm[i].level == nsm[j].level {
				repeat = true
				break
			}
		}

		if !repeat {
			nsm = append(nsm, sm[i])
		}
	}

	return nsm
}
