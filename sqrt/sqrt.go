package main

import "fmt"

/*标题：根号2的计算
描述信息
根号2的计算，计算根号2小数点后10000位的值
参考答案
主要考查通过数学方式如何计算根号2

评分标准
候选人能够想出可行的办法并可实际操作 4分*/

func GetSqrtFloat(n, by int) (newBy int, left int, ret int) {
	for i:=9; i >=0; i-- {
		if (2*by*10 + i) * i < n {
			newBy = by*10+i
			ret = i
			left = n-((2*by*10 + i) * i)
			return newBy, left, ret
		}
	}

	return by*10, 0, 0
}

func main() {
	n := 2
	var by, left, ret int
	for i:=0; i<= 9; i++ {
		by, left, ret = GetSqrtFloat(n, by)
		n = left*100
	}

	fmt.Println(ret)
}