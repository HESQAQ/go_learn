package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e0 := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回对象指针
	e2.Age = 40
	e2.Name = "Rose"
	e2.Id = "2"
	t.Log(e0)
	t.Log(e1)
	t.Log(e1.Age)
	t.Log(e2)
	t.Logf("e is %T", e0)
	t.Logf("e2 is %T", e2)
}

// 传递值, 内存会拷贝对象数据数据
func (e Employee) String() string {
	//New Object Address is c00006e970
	fmt.Printf("New Object Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s - Name:%s - Age:%d", e.Id, e.Name, e.Age)
}

// 传递对象引用
//func (e *Employee) String() string {
// New Object Address is c00006e940
//fmt.Printf("New Object Address is %x\n", unsafe.Pointer(&e.Name))
//return fmt.Sprintf("ID:%s - Name:%s - Age:%d", e.Id, e.Name, e.Age)
//}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	//Old Object Address is c00006e940
	fmt.Printf("Old Object Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
