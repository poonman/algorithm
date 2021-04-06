package main

import "fmt"

/*标题：单链表(奇数位升序，偶数位降序)的排序
描述信息
单链表，奇数位升序，偶数位降序，现在要求整体排成全局升序 输入：1->200->10->120->30->8->88->4 输出：1->4->8->10->30->88->120->200
参考答案
思路：链表可以随便拆、组合 先把奇数和偶数拆开，形成两个链表，一个升序和一个降序 1->10->30->88 200->120->8->4 然后将降序的反转，再合并成一个列表

评分标准
3.0 分：能想出 O(N)的思路
3.5 分：代码实现正确，效率比较高，比较整洁可以给 3.5 分；*/

type Node struct {
	Value int
	Next *Node
}

func do(head *Node) *Node{
	p := head
	var p1 *Node = &Node{}
	var p2 *Node = &Node{}
	var pn1 *Node = p1
	var pn2 *Node = p2

	var i int
	for {
		if p == nil {
			pn1.Next = nil
			pn2.Next = nil
			break
		}
		if i%2 == 0 {
			pn1.Next = p
			pn1 = p
		} else {
			pn2.Next = p
			pn2 = p
		}
		p = p.Next
		i++
	}
	//print("p1", p1.Next)
	//
	//print("p2", p2)

	if p2.Next != nil {
		var pr = p2.Next
		p = pr.Next
		var pn *Node
		if p != nil {
			pn = p.Next
		}
		pr.Next = nil

		for {
			fmt.Println("before:", pr, p, pn)
			if p == nil {
				p2 = pr
				break
			}

			p.Next = pr
			pr = p
			p = pn
			if pn != nil {
				pn = pn.Next
			} else {

			}
			fmt.Println("after:", pr, p, pn)

		}
	}

	print("p1", p1)

	print("p2", p2)

	fmt.Println("xxx")
	// 合并p1, p2
	pn1 = p1.Next
	pn2 = p2.Next
	p = &Node{}
	head = p
	for {
		if pn1 == nil {
			p.Next = pn2
			break
		}
		if pn2 == nil {
			p.Next = pn1
			break
		}
		if pn1.Value > pn2.Value {
			p.Next = pn2
			pn2 = pn2.Next
		} else {
			p.Next = pn1
			pn1 = pn1.Next
		}
		p = p.Next
	}
	return head.Next
}

func main() {
	//1->200->10->120->30->8->88->4
	n4 := &Node{4, nil}
	n88 := &Node{88, n4}
	n8 := &Node{8, n88}
	n30 := &Node{30, n8}
	n120 := &Node{120, n30}
	n10 := &Node{10, n120}
	n200 := &Node{200, n10}
	n1 := &Node{1, n200}

	head := do(n1)

	print("head", head)
}

func print(name string, head *Node) {
	p := head
	fmt.Println(name, "...")
	for {
		if p != nil {
			fmt.Println(p.Value)
			p = p.Next
		} else {
			break
		}
	}
	fmt.Println("")
}
