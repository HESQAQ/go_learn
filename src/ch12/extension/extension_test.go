package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

// 重载
func (d *Dog) Speak() {
	fmt.Println("Wang Wang Wang...")
}

// 重载
func (d *Dog) SpeakTo(host string) {
	// 方法重载
	d.Speak()
	// 调用父类方法
	d.p.Speak()
	fmt.Println(" ", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("hesq")
}
