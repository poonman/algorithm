package main

import "fmt"

/*
标题：输出二叉树左视角能看到的节点
描述信息
给定一颗二叉树：
                1
            2       3
         4   5    6   7
                        8
从左边看，输出能看到的 1，2，4，8 这四个节点，顺序无所谓。
参考答案
递归或者非递归都可以

评分标准
2.5 分及以下：代码有问题，比如 8 这个节点输出不了
3.0 分：代码正确
3.5 分：递归和非递归都可以正确写出，思路很快
*/


type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Result struct {
	arr []int
}

func loop(n *Node, ret *Result, deep int)  {
	if deep >= len(ret.arr) {
		ret.arr = append(ret.arr, n.Value)
	}

	if n.Left != nil {
		loop(n.Left, ret, deep+1)
	}
	if n.Right != nil {
		loop(n.Right, ret, deep+1)
	}
}

func main() {
	n8 := &Node{
		Value: 8,
		Left:  nil,
		Right: nil,
	}
	n7 := &Node{
		Value: 7,
		Left:  nil,
		Right: n8,
	}
	n6 := &Node{
		Value: 6,
		Left:  nil,
		Right: nil,
	}
	n3 := &Node{
		Value: 3,
		Left:  n6,
		Right: n7,
	}

	n4 := &Node{
		Value: 4,
		Left:  nil,
		Right: nil,
	}

	n5 := &Node{
		Value: 5,
		Left:  nil,
		Right: nil,
	}

	n2 := &Node{
		Value: 2,
		Left:  n4,
		Right: n5,
	}

	n1 := &Node{
		Value: 1,
		Left:  n2,
		Right: n3,
	}

	ret := &Result{
		arr: []int{},
	}

	loop(n1, ret, 0)

	fmt.Println("vals: ", ret.arr)
}