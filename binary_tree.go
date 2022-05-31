package main

import "fmt"

type Tree struct {
	Num   int
	Left  *Tree
	Right *Tree
}

func main() {
	//init tree
	t7 := &Tree{7, nil, nil}
	t6 := &Tree{6, t7, nil}
	t3 := &Tree{3, t6, nil}

	t5 := &Tree{5, nil, nil}
	t4 := &Tree{4, nil, nil}
	t2 := &Tree{2, t4, t5}

	t1 := &Tree{1, t2, t3}

	fmt.Println(t1)

	fmt.Println("递归计算最小深度", BinaryTreeMinDepth(t1))
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//recurrence
func BinaryTreeMinDepth(root *Tree) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil {
		return 1 + BinaryTreeMinDepth(root.Left)
	}

	if root.Left != nil {
		return 1 + BinaryTreeMinDepth(root.Right)
	}

	return 1 + Min(BinaryTreeMinDepth(root.Left), BinaryTreeMinDepth(root.Right))
}
