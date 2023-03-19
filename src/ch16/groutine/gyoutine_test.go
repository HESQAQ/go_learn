package groutine

import (
	"fmt"
	"testing"
	"time"
)

// 值传递, 不会被共享
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 10)
}

// 此处结果全部为 10, 存在线程不安全问题, 共享变量 i
func TestGroutine_1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 10)
}
