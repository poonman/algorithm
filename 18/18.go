package main

import "fmt"

/*标题：求字符串中长度为 k 的最小字典序子序列
描述信息
给定一个字符串，对该字符串进行删除操作，保留 k 个字符且相对位置不变，使字典序最小
例子：
k=2
原始字符串：atdrgatgy
答案：aa
参考答案
方案一： 首先在区间 1, n-k+1 区间内找到最小的字母(位置尽量靠前), 假设其位置为 p1;
第二次就是再区间 p1+1, n-k+2 区间内找最小的字母(位置尽量靠前);
依此类推, 进行 k 次即可 优化方案：补充
评分标准
方案一给出完整思路和代码正确，3.0 分；
能够想到优化解法，降低每次在区间中查找最小字符时间复杂度，3.5 分；
ACM 同学需要思考到优化算法，可得 3.0 分；*/

func getMinimumSeq(str string, k int) string{
	n := len(str)
	if n <= k {
		return str
	}
	bs := []byte(str)

	j := n-k+1
	i := 0

	ret := []byte{}

	for j <= n {
		fmt.Println("i: ", i, " j: ", j)
		min := byte(0xff)
		idx := i
		for i < j {
			if min > bs[i] {
				min = bs[i]
				idx = i
			}
			i++
		}
		ret = append(ret, min)
		i = idx+1
		j++
	}

	return string(ret)
}

func main() {
	str := "atdrgatgy"

	ret := getMinimumSeq(str, 4)

	fmt.Println("ret: ", ret)
}
