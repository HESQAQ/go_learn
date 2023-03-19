package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type InConv func(opt int) int

// 封装方法,计算器耗时: 类似于Java中的aop原理
func timeSpent(inner InConv) InConv {
	return func(opt int) int {
		start := time.Now()
		ret := inner(opt)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 3)
	return op
}

func TestFunc(t *testing.T) {
	ts := timeSpent(slowFunc)
	t.Log(ts(10))
}
