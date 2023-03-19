package polymorphsim_test

import (
	"fmt"
	"testing"
)

/**
多态
*/
//类型封装
type Code string

type Programmer interface {
	WriterHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (g *GoProgrammer) WriterHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

func (j *JavaProgrammer) WriterHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

// 此处的入参为接口类型, 所以使用的时候需要传入 指针类型 参数.(&GoProgrammer)
func WritFirstProgrammer(p Programmer) {
	fmt.Printf("Type : %T, %v\n", p, p.WriterHelloWorld())
}

func TestClient(t *testing.T) {
	//p := new(GoProgrammer)
	p := &GoProgrammer{}
	j := new(JavaProgrammer)
	WritFirstProgrammer(p)
	WritFirstProgrammer(j)
}
