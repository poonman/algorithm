package main

import "fmt"

/*题目详情
标题：低空间复杂度的数组循环右移
描述信息
将一个长为 n 的数组循环右移 k 位，要求时间复杂度 O(n)，空间复杂度 O(1)
参考答案
方案一： 用一个临时变量，直接调整各元素到对应位置，空间复杂度 O(1)，时间复杂度 O(n)
方案二： 设数组长度为 n, 设 reverse(i, j)表示翻转数组 i 到 j(不包含 j)部分的元素; 则 reverse(0, n); reverse(0, k); reverse(k, n);

评分标准
空间复杂度未达到 O(1)不得分；
方案一和方案二任选一种，代码实现正确，得分 3 分；
两种方案实现不快，代码质量较高，可得 3.5 分；*/

func reverse(arr []int) {
	n := len(arr)
	i := 0
	j := n-1
	for i<j {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
		i++
		j--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2
	reverse(arr[:k])
	reverse(arr[k:])
	reverse(arr[:])

	fmt.Println("arr: ", arr)
}
