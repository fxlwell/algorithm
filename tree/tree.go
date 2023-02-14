package main

import "fmt"

func main() {
	top := InitTreeV1()
	fmt.Println("前序遍历-递归:")
	PreOrderV1(top)
	fmt.Println("前序遍历-非递归:")
	PreOrderV2(top)
	fmt.Println("中序遍历-递归:")
	InOrderV1(top)
	fmt.Println("中序遍历-非递归:")
	InOrderV2(top)
	fmt.Println("后续遍历-递归:")
	LastOrderV1(top)
	fmt.Println("后续遍历-非递归:")
	LastOrderV2(top)

	/*
		fmt.Println("栈测试")
		stack := NewStack()
		for i := 0; i < 10; i++ {
			stack.Push(&Tree{i, nil, nil})
		}
		for {
			v := stack.Pop()
			if v == nil {
				return
			}
			fmt.Println(v.Value, stack.Len())
		}
	*/

}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func InitTreeV1() *Tree {
	/*
				1
			2				3
		4		5		6		7
		    8		9
		10
	*/
	/*	v10 := &Tree{10, nil, nil}
		v8 := &Tree{8, v10, nil}
		v9 := &Tree{9, nil, nil}
		v5 := &Tree{5, v8, v9}
		v6 := &Tree{6, nil, nil}
		v7 := &Tree{7, nil, nil}
		v4 := &Tree{4, nil, nil}
		v3 := &Tree{3, v6, v7}
		v2 := &Tree{2, v4, v5}
		v1 := &Tree{1, v2, v3}
	*/
	v7 := &Tree{7, nil, nil}
	v6 := &Tree{6, nil, nil}
	v5 := &Tree{5, nil, nil}
	v4 := &Tree{4, nil, nil}
	v3 := &Tree{3, v6, v7}
	v2 := &Tree{2, v4, v5}
	v1 := &Tree{1, v2, v3}

	return v1
}

func PreOrderV1(top *Tree) {
	if top == nil {
		return
	}
	fmt.Println(top.Value)
	PreOrderV1(top.Left)
	PreOrderV1(top.Right)
}

func InOrderV1(top *Tree) {
	if top == nil {
		return
	}
	InOrderV1(top.Left)
	fmt.Println(top.Value)
	InOrderV1(top.Right)
}

func LastOrderV1(top *Tree) {
	if top == nil {
		return
	}
	LastOrderV1(top.Left)
	LastOrderV1(top.Right)
	fmt.Println(top.Value)
}

type Stack struct {
	data []*Tree
}

func NewStack() *Stack {
	return &Stack{make([]*Tree, 0, 100)}
}

func (s *Stack) Push(p *Tree) {
	s.data = append(s.data, p)
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Pop() *Tree {
	l := len(s.data)
	if l == 0 {
		return nil
	}

	maxIndex := l - 1

	v := s.data[maxIndex]
	s.data = s.data[:maxIndex]
	return v
}

func PreOrderV2(top *Tree) {
	if top == nil {
		return
	}
	stack := NewStack()
	stack.Push(top)

	for stack.Len() > 0 {
		v := stack.Pop()
		fmt.Println(v.Value)

		if v.Right != nil {
			stack.Push(v.Right)
		}

		if v.Left != nil {
			stack.Push(v.Left)
		}
	}
}

func InOrderV2(top *Tree) {
	if top == nil {
		return
	}
	stack := NewStack()

	root := top

	for stack.Len() > 0 || root != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		v := stack.Pop()
		fmt.Println(v.Value)
		if v.Right != nil {
			root = v.Right
		}
	}
}

func LastOrderV2(top *Tree) {
	if top == nil {
		return
	}
	stack := NewStack()
	stack.Push(top)

	values := []int{}

	for stack.Len() > 0 {
		v := stack.Pop()
		values = append(values, v.Value)

		if v.Left != nil {
			stack.Push(v.Left)
		}

		if v.Right != nil {
			stack.Push(v.Right)
		}
	}

	for i := len(values) - 1; i >= 0; i-- {
		fmt.Println(values[i])
	}
}
