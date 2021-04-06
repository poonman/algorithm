package main

import "fmt"

/*题目详情
标题：给定数组求数组一个区间，该区间满足区间内最小数乘上区间内数字和结果最大
描述信息
给定一个正整数数列 a, 对于其每个区间, 我们都可以计算一个 X 值;
X 值的定义如下: 对于任意区间, 其 X 值等于区间内最小的那个数乘上区间内所有数和;
现在需要你找出数列 a 的所有区间中, X 值最大的那个区间;
如数列 a 为: 3 1 6 4 5 2;
则 X 值最大的区间为 6, 4, 5 , X = 4 * (6+4+5) = 60;
*/
/*
参考答案
单调栈的经典问题
假设我们维护了两个数组, L, R;
L[i]和 R[i]分别表示 a[i]左边和右边, 比 a[i]小的数的位置;
那么[L[i]+1, R[i]-1]就是以 a[i]为最小那个数构成的区间;
那么我们取所有 a[i] * sum([L[i]+1, R[i]-1]) 中的最大值即可;
*/
/*
于是现在的主要问题在于快速求出 L 和 R 两个数组;
具体方法可以维护单调栈;
以求 L[i]为例子, 大致的伪代码如下:
for i := 0; i < n; i++ {
  for stack.top.num > a[i] {
    stack.pop()
  }
  L[i] = stack.top.index
  stack.push(i, a[i])
}
提示：1）枚举每个数，找到该数左右最大值；2）具体实现提示可以考虑 stack；
评分标准
题目主要应用在应届生或刚毕业同学；
题目解答仅有一种解答，参考题目答案，如果采用暴力解法不得分；其他解法如果复杂度为 O(n^2)不得分，其他更优复杂度可视情况给分；
如果思路正确，给 2 分；
代码书写正确，能够灵活应用 stack 等，给 3 分；
ACM 同学最多得 3 分，非 ACM 同学可以参考题目解答时间和代码规范性给 3.5 分；*/


type Node struct {
	Index int
	Value int
}

type Stack struct {
	list []Node
	p int
}

func NewStack(l int) *Stack {
	s := &Stack{
		list: make([]Node, l, l),
		p:-1,
	}
	return s
}

func (s *Stack) Push(i, a int) {
	if s.p+1 >= len(s.list) {
		return
	}
	s.p++
	s.list[s.p].Index = i
	s.list[s.p].Value = a
}

func (s *Stack) Pop() (Node, bool) {
	if s.p<0{
		return Node{
			Index: 0,
			Value: 0,
		}, false
	}
	n := s.list[s.p]
	s.p--
	return n, true
}

func (s *Stack) Top() (Node, bool) {
	if s.p < 0 || s.p >= len(s.list) {
		return Node{
			Index: 0,
			Value: 0,
		}, false
	}

	return s.list[s.p], true
}

func fn(arr []int) (int, int, int, int){
	L := make([]int, len(arr))
	R := make([]int, len(arr))
	stack1 := NewStack(len(arr))
	for i, a := range arr {
		var n Node
		var ok bool
		for {
			if n, ok = stack1.Top(); !ok {
				break
			} else {
				if n.Value > a {
					_, ok = stack1.Pop()
					if !ok {
						break
					}
				} else {
					break
				}
			}
		}
		n, ok = stack1.Top()
		if !ok {
			L[i] = -1
		} else {
			L[i] = n.Index
		}

		stack1.Push(i, a)

		fmt.Println("stack1: ", stack1, "L: ", L)

	}

	stack2 := NewStack(len(arr))
	for i:=len(arr)-1; i>= 0; i-- {
		a := arr[i]
		var n Node
		var ok bool
		for {
			if n, ok = stack2.Top(); !ok {
				break
			} else {
				if n.Value > a {
					_, ok = stack2.Pop()
					if !ok {
						break
					}
				} else {
					break
				}
			}
		}
		n, ok = stack2.Top()
		if !ok {
			R[i] = len(arr)
		} else {
			R[i] = n.Index
		}

		stack2.Push(i, a)
	}

	var max int
	var maxIndex int
	for i, a := range arr {
		sum := 0


		for j:=L[i]+1; j<=R[i]-1; j++ {
			sum += arr[j]

		}
		sum *= a
		fmt.Println("sum: ", i, a, sum)

		if sum > max {
			max = sum
			maxIndex = i
		}
	}

	return maxIndex, L[maxIndex], R[maxIndex], max
}

func main() {
	arr := []int{3, 2, 8, 6, 4}
	maxIndex, l, r, max := fn(arr)
	fmt.Printf("i:%d, left:%d, right:%d, max:%d\n", maxIndex, l, r, max)
}
