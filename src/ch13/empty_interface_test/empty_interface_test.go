package empty_interface_test

import (
	"fmt"
	"testing"
)

// if 判断
func DoSomethingIf(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer......", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("string......", s)
		return
	}
	fmt.Println("Unknow Type")
}

// switch 判断
func DoSomethingSwith(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer...", v)
	case string:
		fmt.Println("string...", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomethingIf(10)
	DoSomethingIf("str")
	DoSomethingIf(3.14)

	DoSomethingSwith(10)
	DoSomethingSwith("str")
	DoSomethingSwith(3.14)
}
