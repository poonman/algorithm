package main

import "fmt"

/*
题目详情
标题：1-n数字字典序第k大
描述信息
给你一个数字n(n < 1e9), 再给你一个数字k(k < n), 要求你找到1, 2, 3, ... n按照字典序排序后, 第k大的数字;
如, n = 15, k = 7;
那1 ~ 15按照字典序排序为: 1, 10, 11, 12, 13, 14, 15, 2, 3, 4, 5, 6, 7, 8, 9;
则答案为15;

参考答案
利用字典树的思想; 我们假设有这么一棵树, 每个节点都要10个儿子, 10条到儿子的边分别对应数据0~9; 那么我们在这棵树上,
对边按照0~9的顺序进行DFS, 当走到第k个节点时, 该节点对应的数字既为我们的第k大字典序数字;
*/

type Node struct {
	Depth int
	Parent *Node
	Child [10]*Node
	Num int
}

func InitHead(n int) (head *Node) {
	head = &Node{
		Depth:  0,
		Parent: nil,
		Child:  [10]*Node{

		},
		Num: 0,
	}

	for i:=1; i<10; i++ {
		node := &Node{
			Depth:  1,
			Parent: head,
			Child:  [10]*Node{},
			Num: i,
		}

		head.Child[i] = node

		InitNode(node, n)
	}

	return head
}

func InitNode(p *Node, n int) {
	base := p.Num * 10
	for i:=0; i< 10; i++ {
		num := base+i
		if num > n {
			break
		}
		p.Child[i] = &Node{
			Parent: p,
			Child:  [10]*Node{},
			Num:    base+i,
		}

		InitNode(p.Child[i], n)
	}
}

func DFS(node *Node, cnt *int, k int) int {
	for _, nextNode :=  range node.Child {
		if nextNode == nil {
			continue
		}
		*cnt++
		if *cnt == k {
			return nextNode.Num
		}
		ret := DFS(nextNode, cnt, k)
		if ret != 0 {
			return ret
		}
	}

	return 0
}

func main() {
	head := InitHead(15)

	cnt := 0
	num := DFS(head, &cnt, 8)
	fmt.Println("num:", num)
}
