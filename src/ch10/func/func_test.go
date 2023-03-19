package func_test

import (
	"fmt"
	"testing"
)

func Clear() {
	fmt.Println("Clear resoutce......")
}

func TestDefer(t *testing.T) {
	// 延迟执行,类似于Java中的finally
	defer Clear()
	fmt.Println("Start......")
	//panic("error")
}
