package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExt(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally")
	//}()
	// 此方式容易导致僵尸进程存在.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
		}
	}()
	fmt.Println("Start...")
	panic(errors.New("something wrong"))
	//os.Exit(-1)
	//fmt.Println("End...")
}
