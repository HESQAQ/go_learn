package series

import "fmt"

func GetFibonaicSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func Square(n int) int {
	return n * n
}

// 同一个代码文件允许存在多个 init() 方法, 且按照顺序执行
func init() {
	fmt.Println("init...")
}

func init() {
	fmt.Println("init - 2...")
}
