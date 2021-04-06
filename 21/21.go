package main

import "fmt"

/*标题：求数组的最大区间和
描述信息
输出一个 int 型数组的最大连续子数组（所有元素加和最大）各个元素之和
保证数组中至少有一个正数
例：
输入：{1，2，5，-7，8，-10}
输出：9 (子数组为: {1，2，5，-7，8})*/

func do(arr []int) int {
	var sum int
	var curSum int
	for _, a := range arr {
		curSum += a
		if curSum > sum {
			sum = curSum
		} else if curSum < 0 {
			curSum = 0
		}

	}
	return sum
}

func main() {
	arr := []int{1, 2, 5, -7, 8, -10}
	sum := do(arr)

	fmt.Println("sum: ", sum)
}
